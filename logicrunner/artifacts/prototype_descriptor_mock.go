package artifacts

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
)

// PrototypeDescriptorMock implements PrototypeDescriptor
type PrototypeDescriptorMock struct {
	t minimock.Tester

	funcCode          func() (rp1 *insolar.Reference)
	inspectFuncCode   func()
	afterCodeCounter  uint64
	beforeCodeCounter uint64
	CodeMock          mPrototypeDescriptorMockCode

	funcHeadRef          func() (rp1 *insolar.Reference)
	inspectFuncHeadRef   func()
	afterHeadRefCounter  uint64
	beforeHeadRefCounter uint64
	HeadRefMock          mPrototypeDescriptorMockHeadRef

	funcStateID          func() (ip1 *insolar.ID)
	inspectFuncStateID   func()
	afterStateIDCounter  uint64
	beforeStateIDCounter uint64
	StateIDMock          mPrototypeDescriptorMockStateID
}

// NewPrototypeDescriptorMock returns a mock for PrototypeDescriptor
func NewPrototypeDescriptorMock(t minimock.Tester) *PrototypeDescriptorMock {
	m := &PrototypeDescriptorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CodeMock = mPrototypeDescriptorMockCode{mock: m}

	m.HeadRefMock = mPrototypeDescriptorMockHeadRef{mock: m}

	m.StateIDMock = mPrototypeDescriptorMockStateID{mock: m}

	return m
}

type mPrototypeDescriptorMockCode struct {
	mock               *PrototypeDescriptorMock
	defaultExpectation *PrototypeDescriptorMockCodeExpectation
	expectations       []*PrototypeDescriptorMockCodeExpectation
}

// PrototypeDescriptorMockCodeExpectation specifies expectation struct of the PrototypeDescriptor.Code
type PrototypeDescriptorMockCodeExpectation struct {
	mock *PrototypeDescriptorMock

	results *PrototypeDescriptorMockCodeResults
	Counter uint64
}

// PrototypeDescriptorMockCodeResults contains results of the PrototypeDescriptor.Code
type PrototypeDescriptorMockCodeResults struct {
	rp1 *insolar.Reference
}

// Expect sets up expected params for PrototypeDescriptor.Code
func (mmCode *mPrototypeDescriptorMockCode) Expect() *mPrototypeDescriptorMockCode {
	if mmCode.mock.funcCode != nil {
		mmCode.mock.t.Fatalf("PrototypeDescriptorMock.Code mock is already set by Set")
	}

	if mmCode.defaultExpectation == nil {
		mmCode.defaultExpectation = &PrototypeDescriptorMockCodeExpectation{}
	}

	return mmCode
}

// Inspect accepts an inspector function that has same arguments as the PrototypeDescriptor.Code
func (mmCode *mPrototypeDescriptorMockCode) Inspect(f func()) *mPrototypeDescriptorMockCode {
	if mmCode.mock.inspectFuncCode != nil {
		mmCode.mock.t.Fatalf("Inspect function is already set for PrototypeDescriptorMock.Code")
	}

	mmCode.mock.inspectFuncCode = f

	return mmCode
}

// Return sets up results that will be returned by PrototypeDescriptor.Code
func (mmCode *mPrototypeDescriptorMockCode) Return(rp1 *insolar.Reference) *PrototypeDescriptorMock {
	if mmCode.mock.funcCode != nil {
		mmCode.mock.t.Fatalf("PrototypeDescriptorMock.Code mock is already set by Set")
	}

	if mmCode.defaultExpectation == nil {
		mmCode.defaultExpectation = &PrototypeDescriptorMockCodeExpectation{mock: mmCode.mock}
	}
	mmCode.defaultExpectation.results = &PrototypeDescriptorMockCodeResults{rp1}
	return mmCode.mock
}

//Set uses given function f to mock the PrototypeDescriptor.Code method
func (mmCode *mPrototypeDescriptorMockCode) Set(f func() (rp1 *insolar.Reference)) *PrototypeDescriptorMock {
	if mmCode.defaultExpectation != nil {
		mmCode.mock.t.Fatalf("Default expectation is already set for the PrototypeDescriptor.Code method")
	}

	if len(mmCode.expectations) > 0 {
		mmCode.mock.t.Fatalf("Some expectations are already set for the PrototypeDescriptor.Code method")
	}

	mmCode.mock.funcCode = f
	return mmCode.mock
}

// Code implements PrototypeDescriptor
func (mmCode *PrototypeDescriptorMock) Code() (rp1 *insolar.Reference) {
	mm_atomic.AddUint64(&mmCode.beforeCodeCounter, 1)
	defer mm_atomic.AddUint64(&mmCode.afterCodeCounter, 1)

	if mmCode.inspectFuncCode != nil {
		mmCode.inspectFuncCode()
	}

	if mmCode.CodeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCode.CodeMock.defaultExpectation.Counter, 1)

		results := mmCode.CodeMock.defaultExpectation.results
		if results == nil {
			mmCode.t.Fatal("No results are set for the PrototypeDescriptorMock.Code")
		}
		return (*results).rp1
	}
	if mmCode.funcCode != nil {
		return mmCode.funcCode()
	}
	mmCode.t.Fatalf("Unexpected call to PrototypeDescriptorMock.Code.")
	return
}

// CodeAfterCounter returns a count of finished PrototypeDescriptorMock.Code invocations
func (mmCode *PrototypeDescriptorMock) CodeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCode.afterCodeCounter)
}

// CodeBeforeCounter returns a count of PrototypeDescriptorMock.Code invocations
func (mmCode *PrototypeDescriptorMock) CodeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCode.beforeCodeCounter)
}

// MinimockCodeDone returns true if the count of the Code invocations corresponds
// the number of defined expectations
func (m *PrototypeDescriptorMock) MinimockCodeDone() bool {
	for _, e := range m.CodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCodeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCode != nil && mm_atomic.LoadUint64(&m.afterCodeCounter) < 1 {
		return false
	}
	return true
}

// MinimockCodeInspect logs each unmet expectation
func (m *PrototypeDescriptorMock) MinimockCodeInspect() {
	for _, e := range m.CodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to PrototypeDescriptorMock.Code")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCodeCounter) < 1 {
		m.t.Error("Expected call to PrototypeDescriptorMock.Code")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCode != nil && mm_atomic.LoadUint64(&m.afterCodeCounter) < 1 {
		m.t.Error("Expected call to PrototypeDescriptorMock.Code")
	}
}

type mPrototypeDescriptorMockHeadRef struct {
	mock               *PrototypeDescriptorMock
	defaultExpectation *PrototypeDescriptorMockHeadRefExpectation
	expectations       []*PrototypeDescriptorMockHeadRefExpectation
}

// PrototypeDescriptorMockHeadRefExpectation specifies expectation struct of the PrototypeDescriptor.HeadRef
type PrototypeDescriptorMockHeadRefExpectation struct {
	mock *PrototypeDescriptorMock

	results *PrototypeDescriptorMockHeadRefResults
	Counter uint64
}

// PrototypeDescriptorMockHeadRefResults contains results of the PrototypeDescriptor.HeadRef
type PrototypeDescriptorMockHeadRefResults struct {
	rp1 *insolar.Reference
}

// Expect sets up expected params for PrototypeDescriptor.HeadRef
func (mmHeadRef *mPrototypeDescriptorMockHeadRef) Expect() *mPrototypeDescriptorMockHeadRef {
	if mmHeadRef.mock.funcHeadRef != nil {
		mmHeadRef.mock.t.Fatalf("PrototypeDescriptorMock.HeadRef mock is already set by Set")
	}

	if mmHeadRef.defaultExpectation == nil {
		mmHeadRef.defaultExpectation = &PrototypeDescriptorMockHeadRefExpectation{}
	}

	return mmHeadRef
}

// Inspect accepts an inspector function that has same arguments as the PrototypeDescriptor.HeadRef
func (mmHeadRef *mPrototypeDescriptorMockHeadRef) Inspect(f func()) *mPrototypeDescriptorMockHeadRef {
	if mmHeadRef.mock.inspectFuncHeadRef != nil {
		mmHeadRef.mock.t.Fatalf("Inspect function is already set for PrototypeDescriptorMock.HeadRef")
	}

	mmHeadRef.mock.inspectFuncHeadRef = f

	return mmHeadRef
}

// Return sets up results that will be returned by PrototypeDescriptor.HeadRef
func (mmHeadRef *mPrototypeDescriptorMockHeadRef) Return(rp1 *insolar.Reference) *PrototypeDescriptorMock {
	if mmHeadRef.mock.funcHeadRef != nil {
		mmHeadRef.mock.t.Fatalf("PrototypeDescriptorMock.HeadRef mock is already set by Set")
	}

	if mmHeadRef.defaultExpectation == nil {
		mmHeadRef.defaultExpectation = &PrototypeDescriptorMockHeadRefExpectation{mock: mmHeadRef.mock}
	}
	mmHeadRef.defaultExpectation.results = &PrototypeDescriptorMockHeadRefResults{rp1}
	return mmHeadRef.mock
}

//Set uses given function f to mock the PrototypeDescriptor.HeadRef method
func (mmHeadRef *mPrototypeDescriptorMockHeadRef) Set(f func() (rp1 *insolar.Reference)) *PrototypeDescriptorMock {
	if mmHeadRef.defaultExpectation != nil {
		mmHeadRef.mock.t.Fatalf("Default expectation is already set for the PrototypeDescriptor.HeadRef method")
	}

	if len(mmHeadRef.expectations) > 0 {
		mmHeadRef.mock.t.Fatalf("Some expectations are already set for the PrototypeDescriptor.HeadRef method")
	}

	mmHeadRef.mock.funcHeadRef = f
	return mmHeadRef.mock
}

// HeadRef implements PrototypeDescriptor
func (mmHeadRef *PrototypeDescriptorMock) HeadRef() (rp1 *insolar.Reference) {
	mm_atomic.AddUint64(&mmHeadRef.beforeHeadRefCounter, 1)
	defer mm_atomic.AddUint64(&mmHeadRef.afterHeadRefCounter, 1)

	if mmHeadRef.inspectFuncHeadRef != nil {
		mmHeadRef.inspectFuncHeadRef()
	}

	if mmHeadRef.HeadRefMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmHeadRef.HeadRefMock.defaultExpectation.Counter, 1)

		results := mmHeadRef.HeadRefMock.defaultExpectation.results
		if results == nil {
			mmHeadRef.t.Fatal("No results are set for the PrototypeDescriptorMock.HeadRef")
		}
		return (*results).rp1
	}
	if mmHeadRef.funcHeadRef != nil {
		return mmHeadRef.funcHeadRef()
	}
	mmHeadRef.t.Fatalf("Unexpected call to PrototypeDescriptorMock.HeadRef.")
	return
}

// HeadRefAfterCounter returns a count of finished PrototypeDescriptorMock.HeadRef invocations
func (mmHeadRef *PrototypeDescriptorMock) HeadRefAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHeadRef.afterHeadRefCounter)
}

// HeadRefBeforeCounter returns a count of PrototypeDescriptorMock.HeadRef invocations
func (mmHeadRef *PrototypeDescriptorMock) HeadRefBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHeadRef.beforeHeadRefCounter)
}

// MinimockHeadRefDone returns true if the count of the HeadRef invocations corresponds
// the number of defined expectations
func (m *PrototypeDescriptorMock) MinimockHeadRefDone() bool {
	for _, e := range m.HeadRefMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.HeadRefMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterHeadRefCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHeadRef != nil && mm_atomic.LoadUint64(&m.afterHeadRefCounter) < 1 {
		return false
	}
	return true
}

// MinimockHeadRefInspect logs each unmet expectation
func (m *PrototypeDescriptorMock) MinimockHeadRefInspect() {
	for _, e := range m.HeadRefMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to PrototypeDescriptorMock.HeadRef")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.HeadRefMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterHeadRefCounter) < 1 {
		m.t.Error("Expected call to PrototypeDescriptorMock.HeadRef")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHeadRef != nil && mm_atomic.LoadUint64(&m.afterHeadRefCounter) < 1 {
		m.t.Error("Expected call to PrototypeDescriptorMock.HeadRef")
	}
}

type mPrototypeDescriptorMockStateID struct {
	mock               *PrototypeDescriptorMock
	defaultExpectation *PrototypeDescriptorMockStateIDExpectation
	expectations       []*PrototypeDescriptorMockStateIDExpectation
}

// PrototypeDescriptorMockStateIDExpectation specifies expectation struct of the PrototypeDescriptor.StateID
type PrototypeDescriptorMockStateIDExpectation struct {
	mock *PrototypeDescriptorMock

	results *PrototypeDescriptorMockStateIDResults
	Counter uint64
}

// PrototypeDescriptorMockStateIDResults contains results of the PrototypeDescriptor.StateID
type PrototypeDescriptorMockStateIDResults struct {
	ip1 *insolar.ID
}

// Expect sets up expected params for PrototypeDescriptor.StateID
func (mmStateID *mPrototypeDescriptorMockStateID) Expect() *mPrototypeDescriptorMockStateID {
	if mmStateID.mock.funcStateID != nil {
		mmStateID.mock.t.Fatalf("PrototypeDescriptorMock.StateID mock is already set by Set")
	}

	if mmStateID.defaultExpectation == nil {
		mmStateID.defaultExpectation = &PrototypeDescriptorMockStateIDExpectation{}
	}

	return mmStateID
}

// Inspect accepts an inspector function that has same arguments as the PrototypeDescriptor.StateID
func (mmStateID *mPrototypeDescriptorMockStateID) Inspect(f func()) *mPrototypeDescriptorMockStateID {
	if mmStateID.mock.inspectFuncStateID != nil {
		mmStateID.mock.t.Fatalf("Inspect function is already set for PrototypeDescriptorMock.StateID")
	}

	mmStateID.mock.inspectFuncStateID = f

	return mmStateID
}

// Return sets up results that will be returned by PrototypeDescriptor.StateID
func (mmStateID *mPrototypeDescriptorMockStateID) Return(ip1 *insolar.ID) *PrototypeDescriptorMock {
	if mmStateID.mock.funcStateID != nil {
		mmStateID.mock.t.Fatalf("PrototypeDescriptorMock.StateID mock is already set by Set")
	}

	if mmStateID.defaultExpectation == nil {
		mmStateID.defaultExpectation = &PrototypeDescriptorMockStateIDExpectation{mock: mmStateID.mock}
	}
	mmStateID.defaultExpectation.results = &PrototypeDescriptorMockStateIDResults{ip1}
	return mmStateID.mock
}

//Set uses given function f to mock the PrototypeDescriptor.StateID method
func (mmStateID *mPrototypeDescriptorMockStateID) Set(f func() (ip1 *insolar.ID)) *PrototypeDescriptorMock {
	if mmStateID.defaultExpectation != nil {
		mmStateID.mock.t.Fatalf("Default expectation is already set for the PrototypeDescriptor.StateID method")
	}

	if len(mmStateID.expectations) > 0 {
		mmStateID.mock.t.Fatalf("Some expectations are already set for the PrototypeDescriptor.StateID method")
	}

	mmStateID.mock.funcStateID = f
	return mmStateID.mock
}

// StateID implements PrototypeDescriptor
func (mmStateID *PrototypeDescriptorMock) StateID() (ip1 *insolar.ID) {
	mm_atomic.AddUint64(&mmStateID.beforeStateIDCounter, 1)
	defer mm_atomic.AddUint64(&mmStateID.afterStateIDCounter, 1)

	if mmStateID.inspectFuncStateID != nil {
		mmStateID.inspectFuncStateID()
	}

	if mmStateID.StateIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmStateID.StateIDMock.defaultExpectation.Counter, 1)

		results := mmStateID.StateIDMock.defaultExpectation.results
		if results == nil {
			mmStateID.t.Fatal("No results are set for the PrototypeDescriptorMock.StateID")
		}
		return (*results).ip1
	}
	if mmStateID.funcStateID != nil {
		return mmStateID.funcStateID()
	}
	mmStateID.t.Fatalf("Unexpected call to PrototypeDescriptorMock.StateID.")
	return
}

// StateIDAfterCounter returns a count of finished PrototypeDescriptorMock.StateID invocations
func (mmStateID *PrototypeDescriptorMock) StateIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStateID.afterStateIDCounter)
}

// StateIDBeforeCounter returns a count of PrototypeDescriptorMock.StateID invocations
func (mmStateID *PrototypeDescriptorMock) StateIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmStateID.beforeStateIDCounter)
}

// MinimockStateIDDone returns true if the count of the StateID invocations corresponds
// the number of defined expectations
func (m *PrototypeDescriptorMock) MinimockStateIDDone() bool {
	for _, e := range m.StateIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StateIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStateIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStateID != nil && mm_atomic.LoadUint64(&m.afterStateIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockStateIDInspect logs each unmet expectation
func (m *PrototypeDescriptorMock) MinimockStateIDInspect() {
	for _, e := range m.StateIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to PrototypeDescriptorMock.StateID")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.StateIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterStateIDCounter) < 1 {
		m.t.Error("Expected call to PrototypeDescriptorMock.StateID")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcStateID != nil && mm_atomic.LoadUint64(&m.afterStateIDCounter) < 1 {
		m.t.Error("Expected call to PrototypeDescriptorMock.StateID")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PrototypeDescriptorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCodeInspect()

		m.MinimockHeadRefInspect()

		m.MinimockStateIDInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PrototypeDescriptorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *PrototypeDescriptorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCodeDone() &&
		m.MinimockHeadRefDone() &&
		m.MinimockStateIDDone()
}
