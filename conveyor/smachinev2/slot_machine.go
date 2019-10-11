///
//    Copyright 2019 Insolar Technologies
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
///

package smachine

import (
	"context"
	"fmt"
	"math"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type SlotMachineConfig struct {
	PollingPeriod   time.Duration
	PollingTruncate time.Duration
	SlotPageSize    uint16
	ScanCountLimit  int
}

func NewSlotMachine(config SlotMachineConfig, signalCallback func(), injectables DependencyRegistry) SlotMachine {
	return SlotMachine{
		config:        config,
		injectables:   injectables,
		slotPool:      NewSlotPool(config.SlotPageSize),
		activeSlots:   NewSlotQueue(ActiveSlots),
		prioritySlots: NewSlotQueue(ActiveSlots),
		workingSlots:  NewSlotQueue(WorkingSlots),
		syncQueue:     NewSlotMachineSync(signalCallback),
	}
}

var _ DependencyRegistry = &SlotMachine{}

type SlotMachine struct {
	config     SlotMachineConfig
	lastSlotID SlotID // atomic
	slotPool   SlotPool

	injectables DependencyRegistry
	sharedData  sync.Map

	scanCount        uint32
	machineStartedAt time.Time
	scanStartedAt    time.Time

	scanWakeUpAt time.Time

	migrationCount uint16

	hotWaitOnly   bool         // true when activeSlots has only slots added by "hot wait" / WaitAny
	activeSlots   SlotQueue    //they are are moved to workingSlots on every full Scan
	prioritySlots SlotQueue    //they are are moved to workingSlots on every partial or full Scan
	pollingSlots  PollingQueue //they are are moved to workingSlots on every full Scan when time has passed
	workingSlots  SlotQueue    //slots are currently in processing

	syncQueue SlotMachineSync
}

func (m *SlotMachine) IsZero() bool {
	return m.syncQueue.IsZero()
}

func (m *SlotMachine) IsEmpty() bool {
	return m.slotPool.IsEmpty()
}

/* -- Methods for dependencies ------------------------------ */

func (m *SlotMachine) FindDependency(id string) interface{} {
	if v, ok := m.sharedData.Load(id); ok {
		return v
	}
	if m.injectables != nil {
		return m.injectables.FindDependency(id)
	}
	return nil
}

func (m *SlotMachine) PutDependency(id string, v interface{}) {
	m.sharedData.Store(id, v)
}

func (m *SlotMachine) TryPutDependency(id string, v interface{}) bool {
	_, loaded := m.sharedData.LoadOrStore(id, v)
	return !loaded
}

/* -- Methods to run state machines ------------------------------ */

func (m *SlotMachine) ScanOnce(worker FixedSlotWorker) (repeatNow bool, nextPollTime time.Time) {
	scanTime := time.Now()
	m.beforeScan(scanTime)

	switch {
	case m.machineStartedAt.IsZero():
		m.machineStartedAt = scanTime
		fallthrough
	case m.workingSlots.IsEmpty():
		m.scanCount++

		m.hotWaitOnly = true
		m.workingSlots.AppendAll(&m.prioritySlots)
		m.workingSlots.AppendAll(&m.activeSlots)
		m.pollingSlots.FilterOut(scanTime, m.workingSlots.AppendAll)
	default:
		// we were interrupted
		m.workingSlots.PrependAll(&m.prioritySlots)
	}

	m.pollingSlots.PrepareFor(scanTime.Add(m.config.PollingPeriod).Truncate(m.config.PollingTruncate))

	repeatNow = m.syncQueue.ProcessUpdates(worker)
	hasUpdates, hasSignal, wasDetached := m.syncQueue.ProcessCallbacks(worker)
	if hasUpdates {
		repeatNow = true
	}

	if !hasSignal && !wasDetached {
		m.scanWorkingSlots(worker)
	}

	repeatNow = repeatNow || !m.hotWaitOnly
	return repeatNow, minTime(m.scanWakeUpAt, m.pollingSlots.GetNearestPollTime())
}

func (m *SlotMachine) beforeScan(scanTime time.Time) {
	if m.machineStartedAt.IsZero() {
		m.machineStartedAt = scanTime
	}
	m.scanStartedAt = scanTime
	m.scanWakeUpAt = time.Time{}
}

const maxLoopCount = 10000

func (m *SlotMachine) scanWorkingSlots(worker FixedSlotWorker) {
	limit := m.config.ScanCountLimit
	if limit <= 0 || limit > maxLoopCount {
		limit = maxLoopCount
	}

	for i := 0; i < limit; {
		currentSlot := m.workingSlots.First()
		if currentSlot == nil {
			return
		}
		loopLimit := 1 + ((limit - i) / m.workingSlots.Count())

		prevStepNo := currentSlot.startWorking(m.scanCount) // its counterpart is in slotPostExecution()
		currentSlot.removeFromQueue()

		if stopNow, loopIncrement := m.executeSlot(currentSlot, prevStepNo, worker, loopLimit); stopNow {
			return
		} else {
			i += loopIncrement
		}
	}
}

func (m *SlotMachine) executeSlot(currentSlot *Slot, prevStepNo uint32, worker FixedSlotWorker, loopLimit int) (hasSignal bool, loopCount int) {

	if dep := currentSlot.dependency; dep != nil && dep.IsReleaseOnWorking() {
		currentSlot.releaseDependency(worker)
	}

	// TODO consider use of sync.Pool for executionContext if they are allocated on heap
	var stateUpdate StateUpdate
	wasDetached := worker.DetachableCall(func(worker DetachableSlotWorker) {
		//defer func() {
		//	err = recoverSlotPanic("slot execution failed", recover(), err)
		//}()

		for ; loopCount < loopLimit; loopCount++ {
			canLoop := false
			canLoop, hasSignal = worker.CanLoopOrHasSignal(loopCount)
			if !canLoop || hasSignal {
				if loopCount == 0 {
					// a very special update type, not to be used anywhere else
					stateUpdate = StateUpdate{updKind: uint8(stateUpdInternalRepeatNow)}
				} else {
					stateUpdate = newStateUpdateTemplate(updCtxExec, 0, stateUpdRepeat).newUint(0)
				}
				return
			}

			var asyncCnt uint16
			var sut StateUpdateType

			ec := executionContext{slotContext: slotContext{s: currentSlot, w: worker}}
			stateUpdate, sut, asyncCnt = ec.executeNextStep()

			if asyncCnt > 0 {
				currentSlot.addAsyncCount(asyncCnt)
			}

			if !sut.ShortLoop(currentSlot, stateUpdate, uint32(loopCount)) {
				return
			}
		}
	})
	loopCount++

	if wasDetached {
		// MUST NOT apply any changes in the current routine, as it is no more safe to update queues
		m.asyncPostSlotExecution(currentSlot, stateUpdate, prevStepNo)
		return true, loopCount
	}

	hasAsync := m.slotPostExecution(currentSlot, stateUpdate, worker, prevStepNo, false)
	if hasAsync && !hasSignal {
		_, hasSignal, wasDetached = m.syncQueue.ProcessCallbacks(worker)
		return hasSignal || wasDetached, loopCount
	}
	return hasSignal, loopCount
}

func (m *SlotMachine) _executeSlotInitByCreator(currentSlot *Slot, worker DetachableSlotWorker) {

	currentSlot.ensureInitializing()

	ec := executionContext{slotContext: slotContext{s: currentSlot, w: worker}}
	stateUpdate, _, asyncCnt := ec.executeNextStep()

	if asyncCnt > 0 {
		currentSlot.addAsyncCount(asyncCnt)
	}

	if !worker.NonDetachableCall(func(worker FixedSlotWorker) {
		m.slotPostExecution(currentSlot, stateUpdate, worker, 0, false)
	}) {
		m.asyncPostSlotExecution(currentSlot, stateUpdate, 0)
	}
}

func (m *SlotMachine) slotPostExecution(slot *Slot, stateUpdate StateUpdate, worker FixedSlotWorker,
	prevStepNo uint32, wasAsync bool) (hasAsync bool) {

	if !stateUpdate.IsZero() {
		slotLink := slot.NewLink() // MUST be taken BEFORE any slot updates
		if !m.applyStateUpdate(slot, stateUpdate, worker) {
			m._flushMissingSlotQueue(slotLink)
			return false
		}
	}

	if slot.canMigrateWorking(prevStepNo, wasAsync) {
		slotLink := slot.NewLink() // MUST be taken BEFORE any slot updates
		if _, isAvailable := m._migrateSlot(slot, worker); !isAvailable {
			m._flushMissingSlotQueue(slotLink)
			return false
		}
	}

	hasAsync = wasAsync || slot.hasAsyncOrBargeIn()
	m.stopSlotWorking(slot, prevStepNo, worker)
	return hasAsync
}

func (m *SlotMachine) queueAsyncCallback(link SlotLink,
	callbackFn func(slot *Slot, worker DetachableSlotWorker) StateUpdate,
	recovered interface{}) {

	if callbackFn == nil && recovered == nil || !m._canCallback(link) {
		return
	}

	m.syncQueue.AddAsyncCallback(link, func(link SlotLink, worker DetachableSlotWorker) (isDone bool) {
		if !m._canCallback(link) {
			return true
		}
		if worker == nil {
			// TODO _handleAsyncDetachmentLimitExceeded
			return true
		}

		slot := link.s
		isStarted, prevStepNo := slot.tryStartWorking()
		if !isStarted {
			return false
		}
		var stateUpdate StateUpdate
		func() {
			defer func() {
				prevErr := recoverSlotPanic("async call panic", recovered, nil)
				recoverSlotPanicAsUpdate(&stateUpdate, "async callback panic", recover(), prevErr)
			}()
			if callbackFn != nil {
				stateUpdate = callbackFn(slot, worker)
			}
		}()

		if worker.NonDetachableCall(func(worker FixedSlotWorker) {
			m.slotPostExecution(slot, stateUpdate, worker, prevStepNo, true)
		}) {
			m.syncQueue.ProcessDetachQueue(link, worker)
		} else {
			m.asyncPostSlotExecution(slot, stateUpdate, prevStepNo)
		}

		return true
	})
}

func (m *SlotMachine) _canCallback(link SlotLink) bool {
	if link.s.machine != m {
		panic("illegal state")
	}
	if link.IsValid() {
		return true
	}
	// TODO log missed async result?
	m._flushMissingSlotQueue(link)
	return false
}

func (m *SlotMachine) _flushMissingSlotQueue(slotLink SlotLink) {
	detached := m.syncQueue.FlushDetachQueue(slotLink)
	if len(detached) > 0 {
		runtime.KeepAlive(detached)
		// TODO m._handleMissedSlotCallbacks(slotLink, nil, detached)
	}
}

func (m *SlotMachine) asyncPostSlotExecution(s *Slot, stateUpdate StateUpdate, prevStepNo uint32) {
	s.asyncCallCount++
	m.syncQueue.AddAsyncUpdate(s.NewLink(), func(link SlotLink, worker FixedSlotWorker) {
		if !link.IsValid() {
			return
		}
		slot := link.s
		slot.asyncCallCount--
		if m.slotPostExecution(slot, stateUpdate, worker, prevStepNo, true) {
			m.syncQueue.AppendSlotDetachQueue(link)
		}
	})
}

func recoverSlotPanic(msg string, recovered interface{}, prev error) error {
	if recovered == nil {
		return prev
	}
	stack := debug.Stack()
	if prev != nil {
		return fmt.Errorf("%s: %v, %v\n%s", msg, recovered, prev, string(stack))
		//return errors.Wrap(prev, fmt.Sprintf("%s: %v", msg, recovered))
	}
	return fmt.Errorf("%s: %v\n%s", msg, recovered, string(stack))
}

func recoverSlotPanicWithStack(msg string, recovered interface{}, prev error) error {
	if recovered == nil {
		return prev
	}
	stack := debug.Stack()
	if prev != nil {
		return fmt.Errorf("%s: %v, %v\n%s", msg, recovered, prev, string(stack))
		//return errors.Wrap(prev, fmt.Sprintf("%s: %v", msg, recovered))
	}
	return fmt.Errorf("%s: %v\n%s", msg, recovered, string(stack))
}

func recoverSlotPanicAsUpdate(update *StateUpdate, msg string, recovered interface{}, prev error) {
	if recovered != nil {
		*update = newPanicStateUpdate(recoverSlotPanicWithStack(msg, recovered, prev))
	} else if prev != nil {
		*update = newPanicStateUpdate(prev)
	}
}

/* -- Methods to migrate slots ------------------------------ */

func (m *SlotMachine) Migrate(cleanupWeak bool, worker FixedSlotWorker) {

	m.migrationCount++
	//if m.ownsAdapters {
	//	m.adapters.migrate(m.containerState, m.migrationCount)
	//}
	m.slotPool.ScanAndCleanup(cleanupWeak, worker, m.recycleSlot, m.migratePage)
	m.syncQueue.CleanupDetachQueues(nil)
}

func (m *SlotMachine) migratePage(slotPage []Slot, worker FixedSlotWorker) (isPageEmptyOrWeak, hasWeakSlots bool) {
	isPageEmptyOrWeak = true
	hasWeakSlots = false
	for i := range slotPage {
		isSlotEmptyOrWeak, isSlotAvailable := m.migrateSlot(&slotPage[i], worker)
		switch {
		case !isSlotEmptyOrWeak:
			isPageEmptyOrWeak = false
		case isSlotAvailable:
			hasWeakSlots = true
		}
	}
	return isPageEmptyOrWeak, hasWeakSlots
}

func (m *SlotMachine) migrateSlot(slot *Slot, w FixedSlotWorker) (isEmptyOrWeak, isAvailable bool) {
	isEmpty, isStarted, prevStepNo := slot.tryStartMigrate()
	if !isStarted {
		return isEmpty, false
	}
	isEmptyOrWeak, isAvailable = m._migrateSlot(slot, w)
	if isAvailable {
		m.stopSlotWorking(slot, prevStepNo, w)
	}
	return isEmptyOrWeak, isAvailable
}

func (m *SlotMachine) _migrateSlot(slot *Slot, worker FixedSlotWorker) (isEmptyOrWeak, isAvailable bool) {

	for m.migrationCount != slot.migrationCount {
		migrateFn := slot.getMigration()
		if migrateFn == nil {
			slot.migrationCount = m.migrationCount
			break
		}

		mc := migrationContext{slotContext{s: slot}}
		stateUpdate := mc.executeMigration(migrateFn)

		slotLink := slot.NewLink() // MUST be taken BEFORE any slot updates

		if !m.applyStateUpdate(slot, stateUpdate, worker) {
			m._flushMissingSlotQueue(slotLink)
			return true, false
		}
		slot.migrationCount++
	}

	return slot.step.Flags&StepWeak != 0, true
}

/* -- Methods to dispose/reuse slots ------------------------------ */

func (m *SlotMachine) Cleanup(worker FixedSlotWorker) {
	m.slotPool.ScanAndCleanup(true, worker, m.recycleSlot, m.verifyPage)
	m.syncQueue.CleanupDetachQueues(nil)
}

func (m *SlotMachine) verifyPage(slotPage []Slot, _ FixedSlotWorker) (isPageEmptyOrWeak, hasWeakSlots bool) {
	isPageEmptyOrWeak = true
	hasWeakSlots = false

	for i := range slotPage {
		slot := &slotPage[i]

		switch {
		case slot.isEmpty():
			continue
		case slot.isWorking():
			break
		case slot.step.Flags&StepWeak != 0:
			hasWeakSlots = true
			continue
		}
		return false, hasWeakSlots
	}
	return isPageEmptyOrWeak, hasWeakSlots
}

func (m *SlotMachine) recycleSlot(slot *Slot, w FixedSlotWorker) {
	m._recycleSlot(slot, w)
}

func (m *SlotMachine) recycleEmptySlot(slot *Slot) {
	m._recycleSlot(slot, nil)
}

func (m *SlotMachine) _recycleSlot(slot *Slot, worker FixedSlotWorker) {

	dep := slot.dependency
	slot.dependency = nil
	if dep != nil {
		dep.ReleaseOnDisposed(func(link SlotLink) {
			if worker == nil {
				panic("illegal state")
			}
			m.activateDependantByLink(link, worker)
		})
	}

	if slot.isQueueHead() {
		s := slot.removeHeadedQueue()
		for s != nil {
			c := s
			s = s._cutNext()
			m._activateDependant(c, worker)
		}
	} else {
		slot.removeFromQueue()
	}
	slot.dispose()

	m.slotPool.RecycleSlot(slot)
}

func (m *SlotMachine) OccupiedSlotCount() int {
	return m.slotPool.Count()
}

func (m *SlotMachine) AllocatedSlotCount() int {
	return m.slotPool.Capacity()
}

/* -- Methods to create/allocate and start new slots ------------------------------ */

func (m *SlotMachine) AddNew(ctx context.Context, parent SlotLink, sm StateMachine) SlotLink {
	if sm == nil {
		panic("illegal value")
	}
	if ctx == nil {
		panic("illegal value")
	}

	newSlot := m.allocateSlot()
	newSlot.parent = parent
	newSlot.ctx = ctx
	link := newSlot.NewLink()

	m.prepareNewSlot(newSlot, nil, nil, sm)
	m.syncQueue.AddAsyncUpdate(link, m._startAddedSlot)
	return link
}

// TODO migrate MUST not take slots at step=0
// TODO slot must apply migrate after apply when step=0, but only outside of detachable
func (m *SlotMachine) prepareNewSlot(slot, creator *Slot, fn CreateFunc, sm StateMachine) {

	defer func() {
		recovered := recover()
		if recovered != nil {
			m.recycleEmptySlot(slot)
			panic(recovered)
		}
	}()

	if fn != nil {
		if sm != nil {
			panic("illegal value")
		}
		cc := constructionContext{s: slot}
		sm = cc.executeCreate(fn)
	}

	if sm == nil {
		m.recycleEmptySlot(slot)
		return
	}

	decl := sm.GetStateMachineDeclaration()
	if decl == nil {
		panic("illegal state")
	}
	slot.declaration = decl

	link := slot.NewLink()
	decl.InjectDependencies(sm, link, m)

	initFn := slot.declaration.GetInitStateFor(sm)
	if initFn == nil {
		panic("illegal value")
	}

	if creator != nil {
		slot.migrationCount = creator.migrationCount
		slot.lastWorkScan = creator.lastWorkScan
	} else {
		slot.migrationCount = m.migrationCount
		slot.lastWorkScan = uint8(m.scanCount)
	}

	slot.step = SlotStep{Transition: func(ctx ExecutionContext) StateUpdate {
		ec := ctx.(*executionContext)
		ic := initializationContext{ec.clone()}
		slot.incStep()
		su := ic.executeInitialization(initFn)
		su.marker = ec.getMarker()
		return su
	}}
}

func (m *SlotMachine) startNewSlot(slot *Slot, worker FixedSlotWorker) {
	slot.ensureInitializing()
	m.stopSlotWorking(slot, 0, worker)
	m.updateSlotQueue(slot, worker, activateSlot)
}

func (m *SlotMachine) startNewSlotByDetachable(slot *Slot, runInit bool, w DetachableSlotWorker) {
	if runInit {
		m._executeSlotInitByCreator(slot, w)
		return
	}

	slot.ensureInitializing()
	if !w.NonDetachableCall(func(worker FixedSlotWorker) {
		m.stopSlotWorking(slot, 0, worker)
		m.updateSlotQueue(slot, worker, activateSlot)
	}) {
		m.syncQueue.AddAsyncUpdate(slot.NewLink(), m._startAddedSlot)
	}
}

func (m *SlotMachine) _startAddedSlot(link SlotLink, worker FixedSlotWorker) {
	if !link.IsValid() {
		panic("unexpected")
	}
	slot := link.s
	slot.ensureInitializing()
	m.stopSlotWorking(slot, 0, worker)
	list := m._updateSlotQueue(slot, false, activateSlot)
	if list != nil {
		panic("unexpected")
	}
}

func (m *SlotMachine) allocateNextSlotID() SlotID {
	for {
		r := atomic.LoadUint32((*uint32)(&m.lastSlotID))
		if r == math.MaxUint32 {
			panic("overflow")
		}
		if atomic.CompareAndSwapUint32((*uint32)(&m.lastSlotID), r, r+1) {
			return SlotID(r + 1)
		}
	}
}

func (m *SlotMachine) allocateSlot() *Slot {
	return m.slotPool.AllocateSlot(m, m.allocateNextSlotID())
}

/* -- Methods to manage processing queues, activate/deactivate slots ------------------------------ */

type slotActivationMode uint8

const (
	deactivateSlot slotActivationMode = iota
	activateSlot
	activateHotWaitSlot
)

func (m *SlotMachine) updateSlotQueue(slot *Slot, w FixedSlotWorker, activation slotActivationMode) {
	if slot.machine != m {
		panic("illegal state")
	}

	s := m._updateSlotQueue(slot, slot.isInQueue(), activation)
	for s != nil {
		c := s
		s = s._cutNext()
		m._activateDependant(c, w)
	}
}

func (m *SlotMachine) _updateSlotQueue(slot *Slot, inplaceUpdate bool, activation slotActivationMode) *Slot {
	if !slot.isQueueHead() {
		if inplaceUpdate {
			switch activation {
			case activateSlot:
				switch slot.QueueType() {
				case ActiveSlots, WorkingSlots:
					return nil
				}
			case activateHotWaitSlot:
				if slot.QueueType() == ActiveSlots {
					return nil
				}
			}
			slot.removeFromQueue()
		} else {
			slot.ensureNotInQueue()
		}

		if activation == deactivateSlot {
			return nil
		}
		m._activateSlot(slot, activation)
		return nil
	}

	if slot.QueueType() != ActivationOfSlot {
		panic("illegal state")
	}

	if activation == deactivateSlot {
		if !inplaceUpdate {
			slot.ensureNotInQueue()
		}
		return nil
	}

	nextDep := slot.removeHeadedQueue()
	m._activateSlot(slot, activation)
	return nextDep
}

func (m *SlotMachine) _activateSlot(slot *Slot, mode slotActivationMode) {
	switch {
	case mode == activateHotWaitSlot:
		m._addSlotToActiveQueue(slot)
	case slot.isLastScan(m.scanCount):
		m.hotWaitOnly = false
		m._addSlotToActiveQueue(slot)
	default:
		m._addSlotToWorkingQueue(slot)
	}
}

func (m *SlotMachine) _addSlotToActiveQueue(slot *Slot) {
	if slot.isPriority() {
		m.prioritySlots.AddLast(slot)
	} else {
		m.activeSlots.AddLast(slot)
	}
}

func (m *SlotMachine) _addSlotToWorkingQueue(slot *Slot) {
	if slot.isPriority() {
		m.workingSlots.AddFirst(slot)
	} else {
		m.workingSlots.AddLast(slot)
	}
}

/* -------------------------------- */

func (m *SlotMachine) applyStateUpdate(slot *Slot, stateUpdate StateUpdate, w FixedSlotWorker) bool {
	if slot.machine != m {
		panic("illegal state")
	}

	var err error
	isAvailable := false
	isPanic := false

	func() {
		defer func() {
			isPanic = true
			err = recoverSlotPanic("apply state update panic", recover(), err)
		}()
		isAvailable, err = typeOfStateUpdate(stateUpdate).Apply(slot, stateUpdate, w)
	}()

	if err == nil {
		return isAvailable
	}

	return m.handleSlotUpdateError(slot, w, isPanic, false, err)
}

func (m *SlotMachine) handleSlotUpdateError(slot *Slot, worker FixedSlotWorker, isPanic bool, isAsync bool, err error) bool {

	eh := slot.getErrorHandler()
	if eh != nil {
		fc := failureContext{isPanic: isPanic, isAsync: isAsync, err: err}
		err = fc.executeFailure(eh)
	}
	if err != nil {
		// TODO log error m._handleStateUpdateError(slot, stateUpdate, w, err)
		fmt.Printf("SLOT ERROR: slot=%v err=%v\n", slot.GetSlotID(), err)
		runtime.KeepAlive(err)
	}

	m.recycleSlot(slot, worker)
	return false
}

/* -------------------------------- */

func (m *SlotMachine) createBargeIn(link StepLink, applyFn BargeInApplyFunc) BargeInParamFunc {

	link.s.bargeInCount++
	return func(param interface{}) bool {
		if !link.IsValid() {
			return false
		}
		m.queueAsyncCallback(link.SlotLink, func(slot *Slot, worker DetachableSlotWorker) StateUpdate {
			_, atExactStep := link.isValidAndAtExactStep()
			bc := bargingInContext{slotContext{s: slot}, param, atExactStep}
			return bc.executeBargeIn(applyFn)
		}, nil)
		return true
	}
}

/* -------------------------------- */

func minTime(t1, t2 time.Time) time.Time {
	if t1.IsZero() {
		return t2
	}
	if t2.IsZero() || t1.Before(t2) {
		return t1
	}
	return t2
}

func (m *SlotMachine) toRelativeTime(t time.Time) uint32 {
	if m.scanStartedAt.IsZero() {
		panic("illegal state")
	}
	if t.IsZero() {
		return 0
	}

	d := t.Sub(m.scanStartedAt)
	if d <= time.Microsecond {
		return 1
	}
	d = 1 + d/time.Microsecond
	if d > math.MaxUint32 {
		panic("illegal value")
	}
	return uint32(d)
}

func (m *SlotMachine) fromRelativeTime(rel uint32) time.Time {
	switch rel {
	case 0:
		return time.Time{}
	case 1:
		return m.scanStartedAt
	}
	return m.scanStartedAt.Add(time.Duration(rel-1) * time.Microsecond)
}

/* -------------------------------- */

func (m *SlotMachine) wakeupOnDeactivationOf(slot *Slot, waitOn SlotLink, worker FixedSlotWorker) {
	if waitOn.s == slot || !waitOn.IsValid() {
		// don't wait for self
		// don't wait for an expired slot
		m.updateSlotQueue(slot, worker, activateSlot)
		return
	}

	wakeupLink := slot.NewLink()
	m.syncQueue.AddAsyncCallback(waitOn, func(waitOn SlotLink, worker DetachableSlotWorker) bool {
		if !wakeupLink.IsValid() {
			return true
		}
		if waitOn.isValidAndBusy() {
			// add this back
			return false
		}

		if !worker.NonDetachableCall(wakeupLink.s.activateSlot) {
			m.syncQueue.AddAsyncUpdate(wakeupLink, SlotLink.activateSlot)
		}

		return true
	})
}

func (m *SlotMachine) useSlotAsShared(link SharedDataLink, accessFn SharedDataFunc, worker DetachableSlotWorker) SharedAccessReport {
	isValid, isBusy := link.link.getIsValidAndBusy()

	if !isValid {
		return SharedSlotAbsent
	}

	isLocal := link.link.s.machine == m
	if isBusy {
		if isLocal {
			return SharedSlotLocalBusy
		}
		return SharedSlotRemoteBusy
	}

	if isLocal {
		if m._useLocalSlotAsShared(link, accessFn, worker) {
			return SharedSlotLocalAvailable
		}
		return SharedSlotLocalBusy
	}

	panic("not implemented")
}

func (m *SlotMachine) stopSlotWorking(slot *Slot, prevStepNo uint32, worker FixedSlotWorker) {
	dep := slot.dependency
	newStepNo := slot.stopWorking()
	if dep == nil || newStepNo == prevStepNo || newStepNo <= 1 {
		return
	}
	if prevStepNo == 0 {
		panic("illegal state")
	}

	if !dep.IsReleaseOnStepping() {
		return
	}

	slot.dependency = nil
	dep.Release(func(link SlotLink) {
		m.activateDependantByLink(link, worker)
	})
}

func (m *SlotMachine) _useLocalSlotAsShared(link SharedDataLink, accessFn SharedDataFunc, worker DetachableSlotWorker) bool {
	slot := link.link.s
	isStarted, _ := slot.tryStartWorking()
	if !isStarted {
		return false
	}

	defer slot.stopWorking()
	accessFn(link.data)

	_, hasSignal := m.syncQueue.ProcessSlotCallbacksByDetachable(link.link, worker)
	if hasSignal || !link.wakeup || slot.QueueType().IsActiveOrPolling() {
		return true
	}

	if !worker.NonDetachableCall(slot.activateSlot) {
		m.syncQueue.AddAsyncUpdate(link.link, SlotLink.activateSlot)
	}
	return true
}

func (m *SlotMachine) activateDependantByLink(link SlotLink, worker FixedSlotWorker) {
	switch {
	case !link.IsValid():
		return
	case m == link.s.machine:
		link.s.activateSlot(worker)
	case worker.OuterCall(link.s.machine, link.s.activateSlot):
		return
	default:
		link.s.machine.syncQueue.AddAsyncUpdate(link, SlotLink.activateSlot)
	}
}

func (m *SlotMachine) _activateDependant(s *Slot, worker FixedSlotWorker) {
	switch {
	case m == s.machine:
		s.activateSlot(worker)
	case worker.OuterCall(s.machine, s.activateSlot):
		return
	default:
		s.machine.syncQueue.AddAsyncUpdate(s.NewLink(), SlotLink.activateSlot)
	}
}