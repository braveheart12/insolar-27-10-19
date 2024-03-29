package executor

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
)

// FilamentCleanerMock implements FilamentCleaner
type FilamentCleanerMock struct {
	t minimock.Tester

	funcClearAllExcept          func(ids []insolar.ID)
	inspectFuncClearAllExcept   func(ids []insolar.ID)
	afterClearAllExceptCounter  uint64
	beforeClearAllExceptCounter uint64
	ClearAllExceptMock          mFilamentCleanerMockClearAllExcept

	funcClearIfLonger          func(limit int)
	inspectFuncClearIfLonger   func(limit int)
	afterClearIfLongerCounter  uint64
	beforeClearIfLongerCounter uint64
	ClearIfLongerMock          mFilamentCleanerMockClearIfLonger
}

// NewFilamentCleanerMock returns a mock for FilamentCleaner
func NewFilamentCleanerMock(t minimock.Tester) *FilamentCleanerMock {
	m := &FilamentCleanerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ClearAllExceptMock = mFilamentCleanerMockClearAllExcept{mock: m}
	m.ClearAllExceptMock.callArgs = []*FilamentCleanerMockClearAllExceptParams{}

	m.ClearIfLongerMock = mFilamentCleanerMockClearIfLonger{mock: m}
	m.ClearIfLongerMock.callArgs = []*FilamentCleanerMockClearIfLongerParams{}

	return m
}

type mFilamentCleanerMockClearAllExcept struct {
	mock               *FilamentCleanerMock
	defaultExpectation *FilamentCleanerMockClearAllExceptExpectation
	expectations       []*FilamentCleanerMockClearAllExceptExpectation

	callArgs []*FilamentCleanerMockClearAllExceptParams
	mutex    sync.RWMutex
}

// FilamentCleanerMockClearAllExceptExpectation specifies expectation struct of the FilamentCleaner.ClearAllExcept
type FilamentCleanerMockClearAllExceptExpectation struct {
	mock   *FilamentCleanerMock
	params *FilamentCleanerMockClearAllExceptParams

	Counter uint64
}

// FilamentCleanerMockClearAllExceptParams contains parameters of the FilamentCleaner.ClearAllExcept
type FilamentCleanerMockClearAllExceptParams struct {
	ids []insolar.ID
}

// Expect sets up expected params for FilamentCleaner.ClearAllExcept
func (mmClearAllExcept *mFilamentCleanerMockClearAllExcept) Expect(ids []insolar.ID) *mFilamentCleanerMockClearAllExcept {
	if mmClearAllExcept.mock.funcClearAllExcept != nil {
		mmClearAllExcept.mock.t.Fatalf("FilamentCleanerMock.ClearAllExcept mock is already set by Set")
	}

	if mmClearAllExcept.defaultExpectation == nil {
		mmClearAllExcept.defaultExpectation = &FilamentCleanerMockClearAllExceptExpectation{}
	}

	mmClearAllExcept.defaultExpectation.params = &FilamentCleanerMockClearAllExceptParams{ids}
	for _, e := range mmClearAllExcept.expectations {
		if minimock.Equal(e.params, mmClearAllExcept.defaultExpectation.params) {
			mmClearAllExcept.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmClearAllExcept.defaultExpectation.params)
		}
	}

	return mmClearAllExcept
}

// Inspect accepts an inspector function that has same arguments as the FilamentCleaner.ClearAllExcept
func (mmClearAllExcept *mFilamentCleanerMockClearAllExcept) Inspect(f func(ids []insolar.ID)) *mFilamentCleanerMockClearAllExcept {
	if mmClearAllExcept.mock.inspectFuncClearAllExcept != nil {
		mmClearAllExcept.mock.t.Fatalf("Inspect function is already set for FilamentCleanerMock.ClearAllExcept")
	}

	mmClearAllExcept.mock.inspectFuncClearAllExcept = f

	return mmClearAllExcept
}

// Return sets up results that will be returned by FilamentCleaner.ClearAllExcept
func (mmClearAllExcept *mFilamentCleanerMockClearAllExcept) Return() *FilamentCleanerMock {
	if mmClearAllExcept.mock.funcClearAllExcept != nil {
		mmClearAllExcept.mock.t.Fatalf("FilamentCleanerMock.ClearAllExcept mock is already set by Set")
	}

	if mmClearAllExcept.defaultExpectation == nil {
		mmClearAllExcept.defaultExpectation = &FilamentCleanerMockClearAllExceptExpectation{mock: mmClearAllExcept.mock}
	}

	return mmClearAllExcept.mock
}

//Set uses given function f to mock the FilamentCleaner.ClearAllExcept method
func (mmClearAllExcept *mFilamentCleanerMockClearAllExcept) Set(f func(ids []insolar.ID)) *FilamentCleanerMock {
	if mmClearAllExcept.defaultExpectation != nil {
		mmClearAllExcept.mock.t.Fatalf("Default expectation is already set for the FilamentCleaner.ClearAllExcept method")
	}

	if len(mmClearAllExcept.expectations) > 0 {
		mmClearAllExcept.mock.t.Fatalf("Some expectations are already set for the FilamentCleaner.ClearAllExcept method")
	}

	mmClearAllExcept.mock.funcClearAllExcept = f
	return mmClearAllExcept.mock
}

// ClearAllExcept implements FilamentCleaner
func (mmClearAllExcept *FilamentCleanerMock) ClearAllExcept(ids []insolar.ID) {
	mm_atomic.AddUint64(&mmClearAllExcept.beforeClearAllExceptCounter, 1)
	defer mm_atomic.AddUint64(&mmClearAllExcept.afterClearAllExceptCounter, 1)

	if mmClearAllExcept.inspectFuncClearAllExcept != nil {
		mmClearAllExcept.inspectFuncClearAllExcept(ids)
	}

	params := &FilamentCleanerMockClearAllExceptParams{ids}

	// Record call args
	mmClearAllExcept.ClearAllExceptMock.mutex.Lock()
	mmClearAllExcept.ClearAllExceptMock.callArgs = append(mmClearAllExcept.ClearAllExceptMock.callArgs, params)
	mmClearAllExcept.ClearAllExceptMock.mutex.Unlock()

	for _, e := range mmClearAllExcept.ClearAllExceptMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmClearAllExcept.ClearAllExceptMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClearAllExcept.ClearAllExceptMock.defaultExpectation.Counter, 1)
		want := mmClearAllExcept.ClearAllExceptMock.defaultExpectation.params
		got := FilamentCleanerMockClearAllExceptParams{ids}
		if want != nil && !minimock.Equal(*want, got) {
			mmClearAllExcept.t.Errorf("FilamentCleanerMock.ClearAllExcept got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmClearAllExcept.funcClearAllExcept != nil {
		mmClearAllExcept.funcClearAllExcept(ids)
		return
	}
	mmClearAllExcept.t.Fatalf("Unexpected call to FilamentCleanerMock.ClearAllExcept. %v", ids)

}

// ClearAllExceptAfterCounter returns a count of finished FilamentCleanerMock.ClearAllExcept invocations
func (mmClearAllExcept *FilamentCleanerMock) ClearAllExceptAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClearAllExcept.afterClearAllExceptCounter)
}

// ClearAllExceptBeforeCounter returns a count of FilamentCleanerMock.ClearAllExcept invocations
func (mmClearAllExcept *FilamentCleanerMock) ClearAllExceptBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClearAllExcept.beforeClearAllExceptCounter)
}

// Calls returns a list of arguments used in each call to FilamentCleanerMock.ClearAllExcept.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmClearAllExcept *mFilamentCleanerMockClearAllExcept) Calls() []*FilamentCleanerMockClearAllExceptParams {
	mmClearAllExcept.mutex.RLock()

	argCopy := make([]*FilamentCleanerMockClearAllExceptParams, len(mmClearAllExcept.callArgs))
	copy(argCopy, mmClearAllExcept.callArgs)

	mmClearAllExcept.mutex.RUnlock()

	return argCopy
}

// MinimockClearAllExceptDone returns true if the count of the ClearAllExcept invocations corresponds
// the number of defined expectations
func (m *FilamentCleanerMock) MinimockClearAllExceptDone() bool {
	for _, e := range m.ClearAllExceptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearAllExceptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterClearAllExceptCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearAllExcept != nil && mm_atomic.LoadUint64(&m.afterClearAllExceptCounter) < 1 {
		return false
	}
	return true
}

// MinimockClearAllExceptInspect logs each unmet expectation
func (m *FilamentCleanerMock) MinimockClearAllExceptInspect() {
	for _, e := range m.ClearAllExceptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FilamentCleanerMock.ClearAllExcept with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearAllExceptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterClearAllExceptCounter) < 1 {
		if m.ClearAllExceptMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FilamentCleanerMock.ClearAllExcept")
		} else {
			m.t.Errorf("Expected call to FilamentCleanerMock.ClearAllExcept with params: %#v", *m.ClearAllExceptMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearAllExcept != nil && mm_atomic.LoadUint64(&m.afterClearAllExceptCounter) < 1 {
		m.t.Error("Expected call to FilamentCleanerMock.ClearAllExcept")
	}
}

type mFilamentCleanerMockClearIfLonger struct {
	mock               *FilamentCleanerMock
	defaultExpectation *FilamentCleanerMockClearIfLongerExpectation
	expectations       []*FilamentCleanerMockClearIfLongerExpectation

	callArgs []*FilamentCleanerMockClearIfLongerParams
	mutex    sync.RWMutex
}

// FilamentCleanerMockClearIfLongerExpectation specifies expectation struct of the FilamentCleaner.ClearIfLonger
type FilamentCleanerMockClearIfLongerExpectation struct {
	mock   *FilamentCleanerMock
	params *FilamentCleanerMockClearIfLongerParams

	Counter uint64
}

// FilamentCleanerMockClearIfLongerParams contains parameters of the FilamentCleaner.ClearIfLonger
type FilamentCleanerMockClearIfLongerParams struct {
	limit int
}

// Expect sets up expected params for FilamentCleaner.ClearIfLonger
func (mmClearIfLonger *mFilamentCleanerMockClearIfLonger) Expect(limit int) *mFilamentCleanerMockClearIfLonger {
	if mmClearIfLonger.mock.funcClearIfLonger != nil {
		mmClearIfLonger.mock.t.Fatalf("FilamentCleanerMock.ClearIfLonger mock is already set by Set")
	}

	if mmClearIfLonger.defaultExpectation == nil {
		mmClearIfLonger.defaultExpectation = &FilamentCleanerMockClearIfLongerExpectation{}
	}

	mmClearIfLonger.defaultExpectation.params = &FilamentCleanerMockClearIfLongerParams{limit}
	for _, e := range mmClearIfLonger.expectations {
		if minimock.Equal(e.params, mmClearIfLonger.defaultExpectation.params) {
			mmClearIfLonger.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmClearIfLonger.defaultExpectation.params)
		}
	}

	return mmClearIfLonger
}

// Inspect accepts an inspector function that has same arguments as the FilamentCleaner.ClearIfLonger
func (mmClearIfLonger *mFilamentCleanerMockClearIfLonger) Inspect(f func(limit int)) *mFilamentCleanerMockClearIfLonger {
	if mmClearIfLonger.mock.inspectFuncClearIfLonger != nil {
		mmClearIfLonger.mock.t.Fatalf("Inspect function is already set for FilamentCleanerMock.ClearIfLonger")
	}

	mmClearIfLonger.mock.inspectFuncClearIfLonger = f

	return mmClearIfLonger
}

// Return sets up results that will be returned by FilamentCleaner.ClearIfLonger
func (mmClearIfLonger *mFilamentCleanerMockClearIfLonger) Return() *FilamentCleanerMock {
	if mmClearIfLonger.mock.funcClearIfLonger != nil {
		mmClearIfLonger.mock.t.Fatalf("FilamentCleanerMock.ClearIfLonger mock is already set by Set")
	}

	if mmClearIfLonger.defaultExpectation == nil {
		mmClearIfLonger.defaultExpectation = &FilamentCleanerMockClearIfLongerExpectation{mock: mmClearIfLonger.mock}
	}

	return mmClearIfLonger.mock
}

//Set uses given function f to mock the FilamentCleaner.ClearIfLonger method
func (mmClearIfLonger *mFilamentCleanerMockClearIfLonger) Set(f func(limit int)) *FilamentCleanerMock {
	if mmClearIfLonger.defaultExpectation != nil {
		mmClearIfLonger.mock.t.Fatalf("Default expectation is already set for the FilamentCleaner.ClearIfLonger method")
	}

	if len(mmClearIfLonger.expectations) > 0 {
		mmClearIfLonger.mock.t.Fatalf("Some expectations are already set for the FilamentCleaner.ClearIfLonger method")
	}

	mmClearIfLonger.mock.funcClearIfLonger = f
	return mmClearIfLonger.mock
}

// ClearIfLonger implements FilamentCleaner
func (mmClearIfLonger *FilamentCleanerMock) ClearIfLonger(limit int) {
	mm_atomic.AddUint64(&mmClearIfLonger.beforeClearIfLongerCounter, 1)
	defer mm_atomic.AddUint64(&mmClearIfLonger.afterClearIfLongerCounter, 1)

	if mmClearIfLonger.inspectFuncClearIfLonger != nil {
		mmClearIfLonger.inspectFuncClearIfLonger(limit)
	}

	params := &FilamentCleanerMockClearIfLongerParams{limit}

	// Record call args
	mmClearIfLonger.ClearIfLongerMock.mutex.Lock()
	mmClearIfLonger.ClearIfLongerMock.callArgs = append(mmClearIfLonger.ClearIfLongerMock.callArgs, params)
	mmClearIfLonger.ClearIfLongerMock.mutex.Unlock()

	for _, e := range mmClearIfLonger.ClearIfLongerMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmClearIfLonger.ClearIfLongerMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClearIfLonger.ClearIfLongerMock.defaultExpectation.Counter, 1)
		want := mmClearIfLonger.ClearIfLongerMock.defaultExpectation.params
		got := FilamentCleanerMockClearIfLongerParams{limit}
		if want != nil && !minimock.Equal(*want, got) {
			mmClearIfLonger.t.Errorf("FilamentCleanerMock.ClearIfLonger got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmClearIfLonger.funcClearIfLonger != nil {
		mmClearIfLonger.funcClearIfLonger(limit)
		return
	}
	mmClearIfLonger.t.Fatalf("Unexpected call to FilamentCleanerMock.ClearIfLonger. %v", limit)

}

// ClearIfLongerAfterCounter returns a count of finished FilamentCleanerMock.ClearIfLonger invocations
func (mmClearIfLonger *FilamentCleanerMock) ClearIfLongerAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClearIfLonger.afterClearIfLongerCounter)
}

// ClearIfLongerBeforeCounter returns a count of FilamentCleanerMock.ClearIfLonger invocations
func (mmClearIfLonger *FilamentCleanerMock) ClearIfLongerBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClearIfLonger.beforeClearIfLongerCounter)
}

// Calls returns a list of arguments used in each call to FilamentCleanerMock.ClearIfLonger.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmClearIfLonger *mFilamentCleanerMockClearIfLonger) Calls() []*FilamentCleanerMockClearIfLongerParams {
	mmClearIfLonger.mutex.RLock()

	argCopy := make([]*FilamentCleanerMockClearIfLongerParams, len(mmClearIfLonger.callArgs))
	copy(argCopy, mmClearIfLonger.callArgs)

	mmClearIfLonger.mutex.RUnlock()

	return argCopy
}

// MinimockClearIfLongerDone returns true if the count of the ClearIfLonger invocations corresponds
// the number of defined expectations
func (m *FilamentCleanerMock) MinimockClearIfLongerDone() bool {
	for _, e := range m.ClearIfLongerMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearIfLongerMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterClearIfLongerCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearIfLonger != nil && mm_atomic.LoadUint64(&m.afterClearIfLongerCounter) < 1 {
		return false
	}
	return true
}

// MinimockClearIfLongerInspect logs each unmet expectation
func (m *FilamentCleanerMock) MinimockClearIfLongerInspect() {
	for _, e := range m.ClearIfLongerMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FilamentCleanerMock.ClearIfLonger with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearIfLongerMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterClearIfLongerCounter) < 1 {
		if m.ClearIfLongerMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FilamentCleanerMock.ClearIfLonger")
		} else {
			m.t.Errorf("Expected call to FilamentCleanerMock.ClearIfLonger with params: %#v", *m.ClearIfLongerMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearIfLonger != nil && mm_atomic.LoadUint64(&m.afterClearIfLongerCounter) < 1 {
		m.t.Error("Expected call to FilamentCleanerMock.ClearIfLonger")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *FilamentCleanerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockClearAllExceptInspect()

		m.MinimockClearIfLongerInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *FilamentCleanerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *FilamentCleanerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockClearAllExceptDone() &&
		m.MinimockClearIfLongerDone()
}
