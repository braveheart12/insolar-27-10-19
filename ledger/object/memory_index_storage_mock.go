package object

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
)

// MemoryIndexStorageMock implements MemoryIndexStorage
type MemoryIndexStorageMock struct {
	t minimock.Tester

	funcForID          func(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) (i1 record.Index, err error)
	inspectFuncForID   func(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID)
	afterForIDCounter  uint64
	beforeForIDCounter uint64
	ForIDMock          mMemoryIndexStorageMockForID

	funcForPulse          func(ctx context.Context, pn insolar.PulseNumber) (ia1 []record.Index, err error)
	inspectFuncForPulse   func(ctx context.Context, pn insolar.PulseNumber)
	afterForPulseCounter  uint64
	beforeForPulseCounter uint64
	ForPulseMock          mMemoryIndexStorageMockForPulse

	funcSet          func(ctx context.Context, pn insolar.PulseNumber, index record.Index)
	inspectFuncSet   func(ctx context.Context, pn insolar.PulseNumber, index record.Index)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mMemoryIndexStorageMockSet

	funcSetIfNone          func(ctx context.Context, pn insolar.PulseNumber, index record.Index)
	inspectFuncSetIfNone   func(ctx context.Context, pn insolar.PulseNumber, index record.Index)
	afterSetIfNoneCounter  uint64
	beforeSetIfNoneCounter uint64
	SetIfNoneMock          mMemoryIndexStorageMockSetIfNone
}

// NewMemoryIndexStorageMock returns a mock for MemoryIndexStorage
func NewMemoryIndexStorageMock(t minimock.Tester) *MemoryIndexStorageMock {
	m := &MemoryIndexStorageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ForIDMock = mMemoryIndexStorageMockForID{mock: m}
	m.ForIDMock.callArgs = []*MemoryIndexStorageMockForIDParams{}

	m.ForPulseMock = mMemoryIndexStorageMockForPulse{mock: m}
	m.ForPulseMock.callArgs = []*MemoryIndexStorageMockForPulseParams{}

	m.SetMock = mMemoryIndexStorageMockSet{mock: m}
	m.SetMock.callArgs = []*MemoryIndexStorageMockSetParams{}

	m.SetIfNoneMock = mMemoryIndexStorageMockSetIfNone{mock: m}
	m.SetIfNoneMock.callArgs = []*MemoryIndexStorageMockSetIfNoneParams{}

	return m
}

type mMemoryIndexStorageMockForID struct {
	mock               *MemoryIndexStorageMock
	defaultExpectation *MemoryIndexStorageMockForIDExpectation
	expectations       []*MemoryIndexStorageMockForIDExpectation

	callArgs []*MemoryIndexStorageMockForIDParams
	mutex    sync.RWMutex
}

// MemoryIndexStorageMockForIDExpectation specifies expectation struct of the MemoryIndexStorage.ForID
type MemoryIndexStorageMockForIDExpectation struct {
	mock    *MemoryIndexStorageMock
	params  *MemoryIndexStorageMockForIDParams
	results *MemoryIndexStorageMockForIDResults
	Counter uint64
}

// MemoryIndexStorageMockForIDParams contains parameters of the MemoryIndexStorage.ForID
type MemoryIndexStorageMockForIDParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	objID insolar.ID
}

// MemoryIndexStorageMockForIDResults contains results of the MemoryIndexStorage.ForID
type MemoryIndexStorageMockForIDResults struct {
	i1  record.Index
	err error
}

// Expect sets up expected params for MemoryIndexStorage.ForID
func (mmForID *mMemoryIndexStorageMockForID) Expect(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) *mMemoryIndexStorageMockForID {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("MemoryIndexStorageMock.ForID mock is already set by Set")
	}

	if mmForID.defaultExpectation == nil {
		mmForID.defaultExpectation = &MemoryIndexStorageMockForIDExpectation{}
	}

	mmForID.defaultExpectation.params = &MemoryIndexStorageMockForIDParams{ctx, pn, objID}
	for _, e := range mmForID.expectations {
		if minimock.Equal(e.params, mmForID.defaultExpectation.params) {
			mmForID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmForID.defaultExpectation.params)
		}
	}

	return mmForID
}

// Inspect accepts an inspector function that has same arguments as the MemoryIndexStorage.ForID
func (mmForID *mMemoryIndexStorageMockForID) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID)) *mMemoryIndexStorageMockForID {
	if mmForID.mock.inspectFuncForID != nil {
		mmForID.mock.t.Fatalf("Inspect function is already set for MemoryIndexStorageMock.ForID")
	}

	mmForID.mock.inspectFuncForID = f

	return mmForID
}

// Return sets up results that will be returned by MemoryIndexStorage.ForID
func (mmForID *mMemoryIndexStorageMockForID) Return(i1 record.Index, err error) *MemoryIndexStorageMock {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("MemoryIndexStorageMock.ForID mock is already set by Set")
	}

	if mmForID.defaultExpectation == nil {
		mmForID.defaultExpectation = &MemoryIndexStorageMockForIDExpectation{mock: mmForID.mock}
	}
	mmForID.defaultExpectation.results = &MemoryIndexStorageMockForIDResults{i1, err}
	return mmForID.mock
}

//Set uses given function f to mock the MemoryIndexStorage.ForID method
func (mmForID *mMemoryIndexStorageMockForID) Set(f func(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) (i1 record.Index, err error)) *MemoryIndexStorageMock {
	if mmForID.defaultExpectation != nil {
		mmForID.mock.t.Fatalf("Default expectation is already set for the MemoryIndexStorage.ForID method")
	}

	if len(mmForID.expectations) > 0 {
		mmForID.mock.t.Fatalf("Some expectations are already set for the MemoryIndexStorage.ForID method")
	}

	mmForID.mock.funcForID = f
	return mmForID.mock
}

// When sets expectation for the MemoryIndexStorage.ForID which will trigger the result defined by the following
// Then helper
func (mmForID *mMemoryIndexStorageMockForID) When(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) *MemoryIndexStorageMockForIDExpectation {
	if mmForID.mock.funcForID != nil {
		mmForID.mock.t.Fatalf("MemoryIndexStorageMock.ForID mock is already set by Set")
	}

	expectation := &MemoryIndexStorageMockForIDExpectation{
		mock:   mmForID.mock,
		params: &MemoryIndexStorageMockForIDParams{ctx, pn, objID},
	}
	mmForID.expectations = append(mmForID.expectations, expectation)
	return expectation
}

// Then sets up MemoryIndexStorage.ForID return parameters for the expectation previously defined by the When method
func (e *MemoryIndexStorageMockForIDExpectation) Then(i1 record.Index, err error) *MemoryIndexStorageMock {
	e.results = &MemoryIndexStorageMockForIDResults{i1, err}
	return e.mock
}

// ForID implements MemoryIndexStorage
func (mmForID *MemoryIndexStorageMock) ForID(ctx context.Context, pn insolar.PulseNumber, objID insolar.ID) (i1 record.Index, err error) {
	mm_atomic.AddUint64(&mmForID.beforeForIDCounter, 1)
	defer mm_atomic.AddUint64(&mmForID.afterForIDCounter, 1)

	if mmForID.inspectFuncForID != nil {
		mmForID.inspectFuncForID(ctx, pn, objID)
	}

	params := &MemoryIndexStorageMockForIDParams{ctx, pn, objID}

	// Record call args
	mmForID.ForIDMock.mutex.Lock()
	mmForID.ForIDMock.callArgs = append(mmForID.ForIDMock.callArgs, params)
	mmForID.ForIDMock.mutex.Unlock()

	for _, e := range mmForID.ForIDMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmForID.ForIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmForID.ForIDMock.defaultExpectation.Counter, 1)
		want := mmForID.ForIDMock.defaultExpectation.params
		got := MemoryIndexStorageMockForIDParams{ctx, pn, objID}
		if want != nil && !minimock.Equal(*want, got) {
			mmForID.t.Errorf("MemoryIndexStorageMock.ForID got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmForID.ForIDMock.defaultExpectation.results
		if results == nil {
			mmForID.t.Fatal("No results are set for the MemoryIndexStorageMock.ForID")
		}
		return (*results).i1, (*results).err
	}
	if mmForID.funcForID != nil {
		return mmForID.funcForID(ctx, pn, objID)
	}
	mmForID.t.Fatalf("Unexpected call to MemoryIndexStorageMock.ForID. %v %v %v", ctx, pn, objID)
	return
}

// ForIDAfterCounter returns a count of finished MemoryIndexStorageMock.ForID invocations
func (mmForID *MemoryIndexStorageMock) ForIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForID.afterForIDCounter)
}

// ForIDBeforeCounter returns a count of MemoryIndexStorageMock.ForID invocations
func (mmForID *MemoryIndexStorageMock) ForIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForID.beforeForIDCounter)
}

// Calls returns a list of arguments used in each call to MemoryIndexStorageMock.ForID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmForID *mMemoryIndexStorageMockForID) Calls() []*MemoryIndexStorageMockForIDParams {
	mmForID.mutex.RLock()

	argCopy := make([]*MemoryIndexStorageMockForIDParams, len(mmForID.callArgs))
	copy(argCopy, mmForID.callArgs)

	mmForID.mutex.RUnlock()

	return argCopy
}

// MinimockForIDDone returns true if the count of the ForID invocations corresponds
// the number of defined expectations
func (m *MemoryIndexStorageMock) MinimockForIDDone() bool {
	for _, e := range m.ForIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForID != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockForIDInspect logs each unmet expectation
func (m *MemoryIndexStorageMock) MinimockForIDInspect() {
	for _, e := range m.ForIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MemoryIndexStorageMock.ForID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		if m.ForIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MemoryIndexStorageMock.ForID")
		} else {
			m.t.Errorf("Expected call to MemoryIndexStorageMock.ForID with params: %#v", *m.ForIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForID != nil && mm_atomic.LoadUint64(&m.afterForIDCounter) < 1 {
		m.t.Error("Expected call to MemoryIndexStorageMock.ForID")
	}
}

type mMemoryIndexStorageMockForPulse struct {
	mock               *MemoryIndexStorageMock
	defaultExpectation *MemoryIndexStorageMockForPulseExpectation
	expectations       []*MemoryIndexStorageMockForPulseExpectation

	callArgs []*MemoryIndexStorageMockForPulseParams
	mutex    sync.RWMutex
}

// MemoryIndexStorageMockForPulseExpectation specifies expectation struct of the MemoryIndexStorage.ForPulse
type MemoryIndexStorageMockForPulseExpectation struct {
	mock    *MemoryIndexStorageMock
	params  *MemoryIndexStorageMockForPulseParams
	results *MemoryIndexStorageMockForPulseResults
	Counter uint64
}

// MemoryIndexStorageMockForPulseParams contains parameters of the MemoryIndexStorage.ForPulse
type MemoryIndexStorageMockForPulseParams struct {
	ctx context.Context
	pn  insolar.PulseNumber
}

// MemoryIndexStorageMockForPulseResults contains results of the MemoryIndexStorage.ForPulse
type MemoryIndexStorageMockForPulseResults struct {
	ia1 []record.Index
	err error
}

// Expect sets up expected params for MemoryIndexStorage.ForPulse
func (mmForPulse *mMemoryIndexStorageMockForPulse) Expect(ctx context.Context, pn insolar.PulseNumber) *mMemoryIndexStorageMockForPulse {
	if mmForPulse.mock.funcForPulse != nil {
		mmForPulse.mock.t.Fatalf("MemoryIndexStorageMock.ForPulse mock is already set by Set")
	}

	if mmForPulse.defaultExpectation == nil {
		mmForPulse.defaultExpectation = &MemoryIndexStorageMockForPulseExpectation{}
	}

	mmForPulse.defaultExpectation.params = &MemoryIndexStorageMockForPulseParams{ctx, pn}
	for _, e := range mmForPulse.expectations {
		if minimock.Equal(e.params, mmForPulse.defaultExpectation.params) {
			mmForPulse.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmForPulse.defaultExpectation.params)
		}
	}

	return mmForPulse
}

// Inspect accepts an inspector function that has same arguments as the MemoryIndexStorage.ForPulse
func (mmForPulse *mMemoryIndexStorageMockForPulse) Inspect(f func(ctx context.Context, pn insolar.PulseNumber)) *mMemoryIndexStorageMockForPulse {
	if mmForPulse.mock.inspectFuncForPulse != nil {
		mmForPulse.mock.t.Fatalf("Inspect function is already set for MemoryIndexStorageMock.ForPulse")
	}

	mmForPulse.mock.inspectFuncForPulse = f

	return mmForPulse
}

// Return sets up results that will be returned by MemoryIndexStorage.ForPulse
func (mmForPulse *mMemoryIndexStorageMockForPulse) Return(ia1 []record.Index, err error) *MemoryIndexStorageMock {
	if mmForPulse.mock.funcForPulse != nil {
		mmForPulse.mock.t.Fatalf("MemoryIndexStorageMock.ForPulse mock is already set by Set")
	}

	if mmForPulse.defaultExpectation == nil {
		mmForPulse.defaultExpectation = &MemoryIndexStorageMockForPulseExpectation{mock: mmForPulse.mock}
	}
	mmForPulse.defaultExpectation.results = &MemoryIndexStorageMockForPulseResults{ia1, err}
	return mmForPulse.mock
}

//Set uses given function f to mock the MemoryIndexStorage.ForPulse method
func (mmForPulse *mMemoryIndexStorageMockForPulse) Set(f func(ctx context.Context, pn insolar.PulseNumber) (ia1 []record.Index, err error)) *MemoryIndexStorageMock {
	if mmForPulse.defaultExpectation != nil {
		mmForPulse.mock.t.Fatalf("Default expectation is already set for the MemoryIndexStorage.ForPulse method")
	}

	if len(mmForPulse.expectations) > 0 {
		mmForPulse.mock.t.Fatalf("Some expectations are already set for the MemoryIndexStorage.ForPulse method")
	}

	mmForPulse.mock.funcForPulse = f
	return mmForPulse.mock
}

// When sets expectation for the MemoryIndexStorage.ForPulse which will trigger the result defined by the following
// Then helper
func (mmForPulse *mMemoryIndexStorageMockForPulse) When(ctx context.Context, pn insolar.PulseNumber) *MemoryIndexStorageMockForPulseExpectation {
	if mmForPulse.mock.funcForPulse != nil {
		mmForPulse.mock.t.Fatalf("MemoryIndexStorageMock.ForPulse mock is already set by Set")
	}

	expectation := &MemoryIndexStorageMockForPulseExpectation{
		mock:   mmForPulse.mock,
		params: &MemoryIndexStorageMockForPulseParams{ctx, pn},
	}
	mmForPulse.expectations = append(mmForPulse.expectations, expectation)
	return expectation
}

// Then sets up MemoryIndexStorage.ForPulse return parameters for the expectation previously defined by the When method
func (e *MemoryIndexStorageMockForPulseExpectation) Then(ia1 []record.Index, err error) *MemoryIndexStorageMock {
	e.results = &MemoryIndexStorageMockForPulseResults{ia1, err}
	return e.mock
}

// ForPulse implements MemoryIndexStorage
func (mmForPulse *MemoryIndexStorageMock) ForPulse(ctx context.Context, pn insolar.PulseNumber) (ia1 []record.Index, err error) {
	mm_atomic.AddUint64(&mmForPulse.beforeForPulseCounter, 1)
	defer mm_atomic.AddUint64(&mmForPulse.afterForPulseCounter, 1)

	if mmForPulse.inspectFuncForPulse != nil {
		mmForPulse.inspectFuncForPulse(ctx, pn)
	}

	params := &MemoryIndexStorageMockForPulseParams{ctx, pn}

	// Record call args
	mmForPulse.ForPulseMock.mutex.Lock()
	mmForPulse.ForPulseMock.callArgs = append(mmForPulse.ForPulseMock.callArgs, params)
	mmForPulse.ForPulseMock.mutex.Unlock()

	for _, e := range mmForPulse.ForPulseMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ia1, e.results.err
		}
	}

	if mmForPulse.ForPulseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmForPulse.ForPulseMock.defaultExpectation.Counter, 1)
		want := mmForPulse.ForPulseMock.defaultExpectation.params
		got := MemoryIndexStorageMockForPulseParams{ctx, pn}
		if want != nil && !minimock.Equal(*want, got) {
			mmForPulse.t.Errorf("MemoryIndexStorageMock.ForPulse got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmForPulse.ForPulseMock.defaultExpectation.results
		if results == nil {
			mmForPulse.t.Fatal("No results are set for the MemoryIndexStorageMock.ForPulse")
		}
		return (*results).ia1, (*results).err
	}
	if mmForPulse.funcForPulse != nil {
		return mmForPulse.funcForPulse(ctx, pn)
	}
	mmForPulse.t.Fatalf("Unexpected call to MemoryIndexStorageMock.ForPulse. %v %v", ctx, pn)
	return
}

// ForPulseAfterCounter returns a count of finished MemoryIndexStorageMock.ForPulse invocations
func (mmForPulse *MemoryIndexStorageMock) ForPulseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForPulse.afterForPulseCounter)
}

// ForPulseBeforeCounter returns a count of MemoryIndexStorageMock.ForPulse invocations
func (mmForPulse *MemoryIndexStorageMock) ForPulseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForPulse.beforeForPulseCounter)
}

// Calls returns a list of arguments used in each call to MemoryIndexStorageMock.ForPulse.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmForPulse *mMemoryIndexStorageMockForPulse) Calls() []*MemoryIndexStorageMockForPulseParams {
	mmForPulse.mutex.RLock()

	argCopy := make([]*MemoryIndexStorageMockForPulseParams, len(mmForPulse.callArgs))
	copy(argCopy, mmForPulse.callArgs)

	mmForPulse.mutex.RUnlock()

	return argCopy
}

// MinimockForPulseDone returns true if the count of the ForPulse invocations corresponds
// the number of defined expectations
func (m *MemoryIndexStorageMock) MinimockForPulseDone() bool {
	for _, e := range m.ForPulseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForPulseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForPulseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForPulse != nil && mm_atomic.LoadUint64(&m.afterForPulseCounter) < 1 {
		return false
	}
	return true
}

// MinimockForPulseInspect logs each unmet expectation
func (m *MemoryIndexStorageMock) MinimockForPulseInspect() {
	for _, e := range m.ForPulseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MemoryIndexStorageMock.ForPulse with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForPulseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForPulseCounter) < 1 {
		if m.ForPulseMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MemoryIndexStorageMock.ForPulse")
		} else {
			m.t.Errorf("Expected call to MemoryIndexStorageMock.ForPulse with params: %#v", *m.ForPulseMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForPulse != nil && mm_atomic.LoadUint64(&m.afterForPulseCounter) < 1 {
		m.t.Error("Expected call to MemoryIndexStorageMock.ForPulse")
	}
}

type mMemoryIndexStorageMockSet struct {
	mock               *MemoryIndexStorageMock
	defaultExpectation *MemoryIndexStorageMockSetExpectation
	expectations       []*MemoryIndexStorageMockSetExpectation

	callArgs []*MemoryIndexStorageMockSetParams
	mutex    sync.RWMutex
}

// MemoryIndexStorageMockSetExpectation specifies expectation struct of the MemoryIndexStorage.Set
type MemoryIndexStorageMockSetExpectation struct {
	mock   *MemoryIndexStorageMock
	params *MemoryIndexStorageMockSetParams

	Counter uint64
}

// MemoryIndexStorageMockSetParams contains parameters of the MemoryIndexStorage.Set
type MemoryIndexStorageMockSetParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	index record.Index
}

// Expect sets up expected params for MemoryIndexStorage.Set
func (mmSet *mMemoryIndexStorageMockSet) Expect(ctx context.Context, pn insolar.PulseNumber, index record.Index) *mMemoryIndexStorageMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("MemoryIndexStorageMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &MemoryIndexStorageMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &MemoryIndexStorageMockSetParams{ctx, pn, index}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the MemoryIndexStorage.Set
func (mmSet *mMemoryIndexStorageMockSet) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, index record.Index)) *mMemoryIndexStorageMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for MemoryIndexStorageMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by MemoryIndexStorage.Set
func (mmSet *mMemoryIndexStorageMockSet) Return() *MemoryIndexStorageMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("MemoryIndexStorageMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &MemoryIndexStorageMockSetExpectation{mock: mmSet.mock}
	}

	return mmSet.mock
}

//Set uses given function f to mock the MemoryIndexStorage.Set method
func (mmSet *mMemoryIndexStorageMockSet) Set(f func(ctx context.Context, pn insolar.PulseNumber, index record.Index)) *MemoryIndexStorageMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the MemoryIndexStorage.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the MemoryIndexStorage.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// Set implements MemoryIndexStorage
func (mmSet *MemoryIndexStorageMock) Set(ctx context.Context, pn insolar.PulseNumber, index record.Index) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(ctx, pn, index)
	}

	params := &MemoryIndexStorageMockSetParams{ctx, pn, index}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		want := mmSet.SetMock.defaultExpectation.params
		got := MemoryIndexStorageMockSetParams{ctx, pn, index}
		if want != nil && !minimock.Equal(*want, got) {
			mmSet.t.Errorf("MemoryIndexStorageMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmSet.funcSet != nil {
		mmSet.funcSet(ctx, pn, index)
		return
	}
	mmSet.t.Fatalf("Unexpected call to MemoryIndexStorageMock.Set. %v %v %v", ctx, pn, index)

}

// SetAfterCounter returns a count of finished MemoryIndexStorageMock.Set invocations
func (mmSet *MemoryIndexStorageMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of MemoryIndexStorageMock.Set invocations
func (mmSet *MemoryIndexStorageMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to MemoryIndexStorageMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mMemoryIndexStorageMockSet) Calls() []*MemoryIndexStorageMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*MemoryIndexStorageMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *MemoryIndexStorageMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *MemoryIndexStorageMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MemoryIndexStorageMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MemoryIndexStorageMock.Set")
		} else {
			m.t.Errorf("Expected call to MemoryIndexStorageMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to MemoryIndexStorageMock.Set")
	}
}

type mMemoryIndexStorageMockSetIfNone struct {
	mock               *MemoryIndexStorageMock
	defaultExpectation *MemoryIndexStorageMockSetIfNoneExpectation
	expectations       []*MemoryIndexStorageMockSetIfNoneExpectation

	callArgs []*MemoryIndexStorageMockSetIfNoneParams
	mutex    sync.RWMutex
}

// MemoryIndexStorageMockSetIfNoneExpectation specifies expectation struct of the MemoryIndexStorage.SetIfNone
type MemoryIndexStorageMockSetIfNoneExpectation struct {
	mock   *MemoryIndexStorageMock
	params *MemoryIndexStorageMockSetIfNoneParams

	Counter uint64
}

// MemoryIndexStorageMockSetIfNoneParams contains parameters of the MemoryIndexStorage.SetIfNone
type MemoryIndexStorageMockSetIfNoneParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	index record.Index
}

// Expect sets up expected params for MemoryIndexStorage.SetIfNone
func (mmSetIfNone *mMemoryIndexStorageMockSetIfNone) Expect(ctx context.Context, pn insolar.PulseNumber, index record.Index) *mMemoryIndexStorageMockSetIfNone {
	if mmSetIfNone.mock.funcSetIfNone != nil {
		mmSetIfNone.mock.t.Fatalf("MemoryIndexStorageMock.SetIfNone mock is already set by Set")
	}

	if mmSetIfNone.defaultExpectation == nil {
		mmSetIfNone.defaultExpectation = &MemoryIndexStorageMockSetIfNoneExpectation{}
	}

	mmSetIfNone.defaultExpectation.params = &MemoryIndexStorageMockSetIfNoneParams{ctx, pn, index}
	for _, e := range mmSetIfNone.expectations {
		if minimock.Equal(e.params, mmSetIfNone.defaultExpectation.params) {
			mmSetIfNone.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSetIfNone.defaultExpectation.params)
		}
	}

	return mmSetIfNone
}

// Inspect accepts an inspector function that has same arguments as the MemoryIndexStorage.SetIfNone
func (mmSetIfNone *mMemoryIndexStorageMockSetIfNone) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, index record.Index)) *mMemoryIndexStorageMockSetIfNone {
	if mmSetIfNone.mock.inspectFuncSetIfNone != nil {
		mmSetIfNone.mock.t.Fatalf("Inspect function is already set for MemoryIndexStorageMock.SetIfNone")
	}

	mmSetIfNone.mock.inspectFuncSetIfNone = f

	return mmSetIfNone
}

// Return sets up results that will be returned by MemoryIndexStorage.SetIfNone
func (mmSetIfNone *mMemoryIndexStorageMockSetIfNone) Return() *MemoryIndexStorageMock {
	if mmSetIfNone.mock.funcSetIfNone != nil {
		mmSetIfNone.mock.t.Fatalf("MemoryIndexStorageMock.SetIfNone mock is already set by Set")
	}

	if mmSetIfNone.defaultExpectation == nil {
		mmSetIfNone.defaultExpectation = &MemoryIndexStorageMockSetIfNoneExpectation{mock: mmSetIfNone.mock}
	}

	return mmSetIfNone.mock
}

//Set uses given function f to mock the MemoryIndexStorage.SetIfNone method
func (mmSetIfNone *mMemoryIndexStorageMockSetIfNone) Set(f func(ctx context.Context, pn insolar.PulseNumber, index record.Index)) *MemoryIndexStorageMock {
	if mmSetIfNone.defaultExpectation != nil {
		mmSetIfNone.mock.t.Fatalf("Default expectation is already set for the MemoryIndexStorage.SetIfNone method")
	}

	if len(mmSetIfNone.expectations) > 0 {
		mmSetIfNone.mock.t.Fatalf("Some expectations are already set for the MemoryIndexStorage.SetIfNone method")
	}

	mmSetIfNone.mock.funcSetIfNone = f
	return mmSetIfNone.mock
}

// SetIfNone implements MemoryIndexStorage
func (mmSetIfNone *MemoryIndexStorageMock) SetIfNone(ctx context.Context, pn insolar.PulseNumber, index record.Index) {
	mm_atomic.AddUint64(&mmSetIfNone.beforeSetIfNoneCounter, 1)
	defer mm_atomic.AddUint64(&mmSetIfNone.afterSetIfNoneCounter, 1)

	if mmSetIfNone.inspectFuncSetIfNone != nil {
		mmSetIfNone.inspectFuncSetIfNone(ctx, pn, index)
	}

	params := &MemoryIndexStorageMockSetIfNoneParams{ctx, pn, index}

	// Record call args
	mmSetIfNone.SetIfNoneMock.mutex.Lock()
	mmSetIfNone.SetIfNoneMock.callArgs = append(mmSetIfNone.SetIfNoneMock.callArgs, params)
	mmSetIfNone.SetIfNoneMock.mutex.Unlock()

	for _, e := range mmSetIfNone.SetIfNoneMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmSetIfNone.SetIfNoneMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSetIfNone.SetIfNoneMock.defaultExpectation.Counter, 1)
		want := mmSetIfNone.SetIfNoneMock.defaultExpectation.params
		got := MemoryIndexStorageMockSetIfNoneParams{ctx, pn, index}
		if want != nil && !minimock.Equal(*want, got) {
			mmSetIfNone.t.Errorf("MemoryIndexStorageMock.SetIfNone got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmSetIfNone.funcSetIfNone != nil {
		mmSetIfNone.funcSetIfNone(ctx, pn, index)
		return
	}
	mmSetIfNone.t.Fatalf("Unexpected call to MemoryIndexStorageMock.SetIfNone. %v %v %v", ctx, pn, index)

}

// SetIfNoneAfterCounter returns a count of finished MemoryIndexStorageMock.SetIfNone invocations
func (mmSetIfNone *MemoryIndexStorageMock) SetIfNoneAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSetIfNone.afterSetIfNoneCounter)
}

// SetIfNoneBeforeCounter returns a count of MemoryIndexStorageMock.SetIfNone invocations
func (mmSetIfNone *MemoryIndexStorageMock) SetIfNoneBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSetIfNone.beforeSetIfNoneCounter)
}

// Calls returns a list of arguments used in each call to MemoryIndexStorageMock.SetIfNone.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSetIfNone *mMemoryIndexStorageMockSetIfNone) Calls() []*MemoryIndexStorageMockSetIfNoneParams {
	mmSetIfNone.mutex.RLock()

	argCopy := make([]*MemoryIndexStorageMockSetIfNoneParams, len(mmSetIfNone.callArgs))
	copy(argCopy, mmSetIfNone.callArgs)

	mmSetIfNone.mutex.RUnlock()

	return argCopy
}

// MinimockSetIfNoneDone returns true if the count of the SetIfNone invocations corresponds
// the number of defined expectations
func (m *MemoryIndexStorageMock) MinimockSetIfNoneDone() bool {
	for _, e := range m.SetIfNoneMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetIfNoneMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetIfNoneCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSetIfNone != nil && mm_atomic.LoadUint64(&m.afterSetIfNoneCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetIfNoneInspect logs each unmet expectation
func (m *MemoryIndexStorageMock) MinimockSetIfNoneInspect() {
	for _, e := range m.SetIfNoneMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MemoryIndexStorageMock.SetIfNone with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetIfNoneMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetIfNoneCounter) < 1 {
		if m.SetIfNoneMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MemoryIndexStorageMock.SetIfNone")
		} else {
			m.t.Errorf("Expected call to MemoryIndexStorageMock.SetIfNone with params: %#v", *m.SetIfNoneMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSetIfNone != nil && mm_atomic.LoadUint64(&m.afterSetIfNoneCounter) < 1 {
		m.t.Error("Expected call to MemoryIndexStorageMock.SetIfNone")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MemoryIndexStorageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockForIDInspect()

		m.MinimockForPulseInspect()

		m.MinimockSetInspect()

		m.MinimockSetIfNoneInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MemoryIndexStorageMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MemoryIndexStorageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockForIDDone() &&
		m.MinimockForPulseDone() &&
		m.MinimockSetDone() &&
		m.MinimockSetIfNoneDone()
}
