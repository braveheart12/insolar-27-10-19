package cryptkit

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// SignatureVerifierFactoryMock implements SignatureVerifierFactory
type SignatureVerifierFactoryMock struct {
	t minimock.Tester

	funcGetSignatureVerifierWithPKS          func(pks PublicKeyStore) (s1 SignatureVerifier)
	inspectFuncGetSignatureVerifierWithPKS   func(pks PublicKeyStore)
	afterGetSignatureVerifierWithPKSCounter  uint64
	beforeGetSignatureVerifierWithPKSCounter uint64
	GetSignatureVerifierWithPKSMock          mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS
}

// NewSignatureVerifierFactoryMock returns a mock for SignatureVerifierFactory
func NewSignatureVerifierFactoryMock(t minimock.Tester) *SignatureVerifierFactoryMock {
	m := &SignatureVerifierFactoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetSignatureVerifierWithPKSMock = mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS{mock: m}
	m.GetSignatureVerifierWithPKSMock.callArgs = []*SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams{}

	return m
}

type mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS struct {
	mock               *SignatureVerifierFactoryMock
	defaultExpectation *SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation
	expectations       []*SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation

	callArgs []*SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams
	mutex    sync.RWMutex
}

// SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation specifies expectation struct of the SignatureVerifierFactory.GetSignatureVerifierWithPKS
type SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation struct {
	mock    *SignatureVerifierFactoryMock
	params  *SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams
	results *SignatureVerifierFactoryMockGetSignatureVerifierWithPKSResults
	Counter uint64
}

// SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams contains parameters of the SignatureVerifierFactory.GetSignatureVerifierWithPKS
type SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams struct {
	pks PublicKeyStore
}

// SignatureVerifierFactoryMockGetSignatureVerifierWithPKSResults contains results of the SignatureVerifierFactory.GetSignatureVerifierWithPKS
type SignatureVerifierFactoryMockGetSignatureVerifierWithPKSResults struct {
	s1 SignatureVerifier
}

// Expect sets up expected params for SignatureVerifierFactory.GetSignatureVerifierWithPKS
func (mmGetSignatureVerifierWithPKS *mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS) Expect(pks PublicKeyStore) *mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS {
	if mmGetSignatureVerifierWithPKS.mock.funcGetSignatureVerifierWithPKS != nil {
		mmGetSignatureVerifierWithPKS.mock.t.Fatalf("SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS mock is already set by Set")
	}

	if mmGetSignatureVerifierWithPKS.defaultExpectation == nil {
		mmGetSignatureVerifierWithPKS.defaultExpectation = &SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation{}
	}

	mmGetSignatureVerifierWithPKS.defaultExpectation.params = &SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams{pks}
	for _, e := range mmGetSignatureVerifierWithPKS.expectations {
		if minimock.Equal(e.params, mmGetSignatureVerifierWithPKS.defaultExpectation.params) {
			mmGetSignatureVerifierWithPKS.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetSignatureVerifierWithPKS.defaultExpectation.params)
		}
	}

	return mmGetSignatureVerifierWithPKS
}

// Inspect accepts an inspector function that has same arguments as the SignatureVerifierFactory.GetSignatureVerifierWithPKS
func (mmGetSignatureVerifierWithPKS *mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS) Inspect(f func(pks PublicKeyStore)) *mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS {
	if mmGetSignatureVerifierWithPKS.mock.inspectFuncGetSignatureVerifierWithPKS != nil {
		mmGetSignatureVerifierWithPKS.mock.t.Fatalf("Inspect function is already set for SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS")
	}

	mmGetSignatureVerifierWithPKS.mock.inspectFuncGetSignatureVerifierWithPKS = f

	return mmGetSignatureVerifierWithPKS
}

// Return sets up results that will be returned by SignatureVerifierFactory.GetSignatureVerifierWithPKS
func (mmGetSignatureVerifierWithPKS *mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS) Return(s1 SignatureVerifier) *SignatureVerifierFactoryMock {
	if mmGetSignatureVerifierWithPKS.mock.funcGetSignatureVerifierWithPKS != nil {
		mmGetSignatureVerifierWithPKS.mock.t.Fatalf("SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS mock is already set by Set")
	}

	if mmGetSignatureVerifierWithPKS.defaultExpectation == nil {
		mmGetSignatureVerifierWithPKS.defaultExpectation = &SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation{mock: mmGetSignatureVerifierWithPKS.mock}
	}
	mmGetSignatureVerifierWithPKS.defaultExpectation.results = &SignatureVerifierFactoryMockGetSignatureVerifierWithPKSResults{s1}
	return mmGetSignatureVerifierWithPKS.mock
}

//Set uses given function f to mock the SignatureVerifierFactory.GetSignatureVerifierWithPKS method
func (mmGetSignatureVerifierWithPKS *mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS) Set(f func(pks PublicKeyStore) (s1 SignatureVerifier)) *SignatureVerifierFactoryMock {
	if mmGetSignatureVerifierWithPKS.defaultExpectation != nil {
		mmGetSignatureVerifierWithPKS.mock.t.Fatalf("Default expectation is already set for the SignatureVerifierFactory.GetSignatureVerifierWithPKS method")
	}

	if len(mmGetSignatureVerifierWithPKS.expectations) > 0 {
		mmGetSignatureVerifierWithPKS.mock.t.Fatalf("Some expectations are already set for the SignatureVerifierFactory.GetSignatureVerifierWithPKS method")
	}

	mmGetSignatureVerifierWithPKS.mock.funcGetSignatureVerifierWithPKS = f
	return mmGetSignatureVerifierWithPKS.mock
}

// When sets expectation for the SignatureVerifierFactory.GetSignatureVerifierWithPKS which will trigger the result defined by the following
// Then helper
func (mmGetSignatureVerifierWithPKS *mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS) When(pks PublicKeyStore) *SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation {
	if mmGetSignatureVerifierWithPKS.mock.funcGetSignatureVerifierWithPKS != nil {
		mmGetSignatureVerifierWithPKS.mock.t.Fatalf("SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS mock is already set by Set")
	}

	expectation := &SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation{
		mock:   mmGetSignatureVerifierWithPKS.mock,
		params: &SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams{pks},
	}
	mmGetSignatureVerifierWithPKS.expectations = append(mmGetSignatureVerifierWithPKS.expectations, expectation)
	return expectation
}

// Then sets up SignatureVerifierFactory.GetSignatureVerifierWithPKS return parameters for the expectation previously defined by the When method
func (e *SignatureVerifierFactoryMockGetSignatureVerifierWithPKSExpectation) Then(s1 SignatureVerifier) *SignatureVerifierFactoryMock {
	e.results = &SignatureVerifierFactoryMockGetSignatureVerifierWithPKSResults{s1}
	return e.mock
}

// GetSignatureVerifierWithPKS implements SignatureVerifierFactory
func (mmGetSignatureVerifierWithPKS *SignatureVerifierFactoryMock) GetSignatureVerifierWithPKS(pks PublicKeyStore) (s1 SignatureVerifier) {
	mm_atomic.AddUint64(&mmGetSignatureVerifierWithPKS.beforeGetSignatureVerifierWithPKSCounter, 1)
	defer mm_atomic.AddUint64(&mmGetSignatureVerifierWithPKS.afterGetSignatureVerifierWithPKSCounter, 1)

	if mmGetSignatureVerifierWithPKS.inspectFuncGetSignatureVerifierWithPKS != nil {
		mmGetSignatureVerifierWithPKS.inspectFuncGetSignatureVerifierWithPKS(pks)
	}

	params := &SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams{pks}

	// Record call args
	mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.mutex.Lock()
	mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.callArgs = append(mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.callArgs, params)
	mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.mutex.Unlock()

	for _, e := range mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1
		}
	}

	if mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.defaultExpectation.Counter, 1)
		want := mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.defaultExpectation.params
		got := SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams{pks}
		if want != nil && !minimock.Equal(*want, got) {
			mmGetSignatureVerifierWithPKS.t.Errorf("SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmGetSignatureVerifierWithPKS.GetSignatureVerifierWithPKSMock.defaultExpectation.results
		if results == nil {
			mmGetSignatureVerifierWithPKS.t.Fatal("No results are set for the SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS")
		}
		return (*results).s1
	}
	if mmGetSignatureVerifierWithPKS.funcGetSignatureVerifierWithPKS != nil {
		return mmGetSignatureVerifierWithPKS.funcGetSignatureVerifierWithPKS(pks)
	}
	mmGetSignatureVerifierWithPKS.t.Fatalf("Unexpected call to SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS. %v", pks)
	return
}

// GetSignatureVerifierWithPKSAfterCounter returns a count of finished SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS invocations
func (mmGetSignatureVerifierWithPKS *SignatureVerifierFactoryMock) GetSignatureVerifierWithPKSAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetSignatureVerifierWithPKS.afterGetSignatureVerifierWithPKSCounter)
}

// GetSignatureVerifierWithPKSBeforeCounter returns a count of SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS invocations
func (mmGetSignatureVerifierWithPKS *SignatureVerifierFactoryMock) GetSignatureVerifierWithPKSBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetSignatureVerifierWithPKS.beforeGetSignatureVerifierWithPKSCounter)
}

// Calls returns a list of arguments used in each call to SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetSignatureVerifierWithPKS *mSignatureVerifierFactoryMockGetSignatureVerifierWithPKS) Calls() []*SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams {
	mmGetSignatureVerifierWithPKS.mutex.RLock()

	argCopy := make([]*SignatureVerifierFactoryMockGetSignatureVerifierWithPKSParams, len(mmGetSignatureVerifierWithPKS.callArgs))
	copy(argCopy, mmGetSignatureVerifierWithPKS.callArgs)

	mmGetSignatureVerifierWithPKS.mutex.RUnlock()

	return argCopy
}

// MinimockGetSignatureVerifierWithPKSDone returns true if the count of the GetSignatureVerifierWithPKS invocations corresponds
// the number of defined expectations
func (m *SignatureVerifierFactoryMock) MinimockGetSignatureVerifierWithPKSDone() bool {
	for _, e := range m.GetSignatureVerifierWithPKSMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetSignatureVerifierWithPKSMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetSignatureVerifierWithPKSCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetSignatureVerifierWithPKS != nil && mm_atomic.LoadUint64(&m.afterGetSignatureVerifierWithPKSCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetSignatureVerifierWithPKSInspect logs each unmet expectation
func (m *SignatureVerifierFactoryMock) MinimockGetSignatureVerifierWithPKSInspect() {
	for _, e := range m.GetSignatureVerifierWithPKSMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetSignatureVerifierWithPKSMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetSignatureVerifierWithPKSCounter) < 1 {
		if m.GetSignatureVerifierWithPKSMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS")
		} else {
			m.t.Errorf("Expected call to SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS with params: %#v", *m.GetSignatureVerifierWithPKSMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetSignatureVerifierWithPKS != nil && mm_atomic.LoadUint64(&m.afterGetSignatureVerifierWithPKSCounter) < 1 {
		m.t.Error("Expected call to SignatureVerifierFactoryMock.GetSignatureVerifierWithPKS")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SignatureVerifierFactoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetSignatureVerifierWithPKSInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SignatureVerifierFactoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *SignatureVerifierFactoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetSignatureVerifierWithPKSDone()
}
