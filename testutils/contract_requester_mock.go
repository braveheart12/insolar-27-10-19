package testutils

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	mm_insolar "github.com/insolar/insolar/insolar"
)

// ContractRequesterMock implements insolar.ContractRequester
type ContractRequesterMock struct {
	t minimock.Tester

	funcCall          func(ctx context.Context, ref *mm_insolar.Reference, method string, argsIn []interface{}, pulse mm_insolar.PulseNumber) (r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error)
	inspectFuncCall   func(ctx context.Context, ref *mm_insolar.Reference, method string, argsIn []interface{}, pulse mm_insolar.PulseNumber)
	afterCallCounter  uint64
	beforeCallCounter uint64
	CallMock          mContractRequesterMockCall

	funcSendRequest          func(ctx context.Context, msg mm_insolar.Payload) (r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error)
	inspectFuncSendRequest   func(ctx context.Context, msg mm_insolar.Payload)
	afterSendRequestCounter  uint64
	beforeSendRequestCounter uint64
	SendRequestMock          mContractRequesterMockSendRequest
}

// NewContractRequesterMock returns a mock for insolar.ContractRequester
func NewContractRequesterMock(t minimock.Tester) *ContractRequesterMock {
	m := &ContractRequesterMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CallMock = mContractRequesterMockCall{mock: m}
	m.CallMock.callArgs = []*ContractRequesterMockCallParams{}

	m.SendRequestMock = mContractRequesterMockSendRequest{mock: m}
	m.SendRequestMock.callArgs = []*ContractRequesterMockSendRequestParams{}

	return m
}

type mContractRequesterMockCall struct {
	mock               *ContractRequesterMock
	defaultExpectation *ContractRequesterMockCallExpectation
	expectations       []*ContractRequesterMockCallExpectation

	callArgs []*ContractRequesterMockCallParams
	mutex    sync.RWMutex
}

// ContractRequesterMockCallExpectation specifies expectation struct of the ContractRequester.Call
type ContractRequesterMockCallExpectation struct {
	mock    *ContractRequesterMock
	params  *ContractRequesterMockCallParams
	results *ContractRequesterMockCallResults
	Counter uint64
}

// ContractRequesterMockCallParams contains parameters of the ContractRequester.Call
type ContractRequesterMockCallParams struct {
	ctx    context.Context
	ref    *mm_insolar.Reference
	method string
	argsIn []interface{}
	pulse  mm_insolar.PulseNumber
}

// ContractRequesterMockCallResults contains results of the ContractRequester.Call
type ContractRequesterMockCallResults struct {
	r1  mm_insolar.Reply
	rp1 *mm_insolar.Reference
	err error
}

// Expect sets up expected params for ContractRequester.Call
func (mmCall *mContractRequesterMockCall) Expect(ctx context.Context, ref *mm_insolar.Reference, method string, argsIn []interface{}, pulse mm_insolar.PulseNumber) *mContractRequesterMockCall {
	if mmCall.mock.funcCall != nil {
		mmCall.mock.t.Fatalf("ContractRequesterMock.Call mock is already set by Set")
	}

	if mmCall.defaultExpectation == nil {
		mmCall.defaultExpectation = &ContractRequesterMockCallExpectation{}
	}

	mmCall.defaultExpectation.params = &ContractRequesterMockCallParams{ctx, ref, method, argsIn, pulse}
	for _, e := range mmCall.expectations {
		if minimock.Equal(e.params, mmCall.defaultExpectation.params) {
			mmCall.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCall.defaultExpectation.params)
		}
	}

	return mmCall
}

// Inspect accepts an inspector function that has same arguments as the ContractRequester.Call
func (mmCall *mContractRequesterMockCall) Inspect(f func(ctx context.Context, ref *mm_insolar.Reference, method string, argsIn []interface{}, pulse mm_insolar.PulseNumber)) *mContractRequesterMockCall {
	if mmCall.mock.inspectFuncCall != nil {
		mmCall.mock.t.Fatalf("Inspect function is already set for ContractRequesterMock.Call")
	}

	mmCall.mock.inspectFuncCall = f

	return mmCall
}

// Return sets up results that will be returned by ContractRequester.Call
func (mmCall *mContractRequesterMockCall) Return(r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error) *ContractRequesterMock {
	if mmCall.mock.funcCall != nil {
		mmCall.mock.t.Fatalf("ContractRequesterMock.Call mock is already set by Set")
	}

	if mmCall.defaultExpectation == nil {
		mmCall.defaultExpectation = &ContractRequesterMockCallExpectation{mock: mmCall.mock}
	}
	mmCall.defaultExpectation.results = &ContractRequesterMockCallResults{r1, rp1, err}
	return mmCall.mock
}

//Set uses given function f to mock the ContractRequester.Call method
func (mmCall *mContractRequesterMockCall) Set(f func(ctx context.Context, ref *mm_insolar.Reference, method string, argsIn []interface{}, pulse mm_insolar.PulseNumber) (r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error)) *ContractRequesterMock {
	if mmCall.defaultExpectation != nil {
		mmCall.mock.t.Fatalf("Default expectation is already set for the ContractRequester.Call method")
	}

	if len(mmCall.expectations) > 0 {
		mmCall.mock.t.Fatalf("Some expectations are already set for the ContractRequester.Call method")
	}

	mmCall.mock.funcCall = f
	return mmCall.mock
}

// When sets expectation for the ContractRequester.Call which will trigger the result defined by the following
// Then helper
func (mmCall *mContractRequesterMockCall) When(ctx context.Context, ref *mm_insolar.Reference, method string, argsIn []interface{}, pulse mm_insolar.PulseNumber) *ContractRequesterMockCallExpectation {
	if mmCall.mock.funcCall != nil {
		mmCall.mock.t.Fatalf("ContractRequesterMock.Call mock is already set by Set")
	}

	expectation := &ContractRequesterMockCallExpectation{
		mock:   mmCall.mock,
		params: &ContractRequesterMockCallParams{ctx, ref, method, argsIn, pulse},
	}
	mmCall.expectations = append(mmCall.expectations, expectation)
	return expectation
}

// Then sets up ContractRequester.Call return parameters for the expectation previously defined by the When method
func (e *ContractRequesterMockCallExpectation) Then(r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error) *ContractRequesterMock {
	e.results = &ContractRequesterMockCallResults{r1, rp1, err}
	return e.mock
}

// Call implements insolar.ContractRequester
func (mmCall *ContractRequesterMock) Call(ctx context.Context, ref *mm_insolar.Reference, method string, argsIn []interface{}, pulse mm_insolar.PulseNumber) (r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error) {
	mm_atomic.AddUint64(&mmCall.beforeCallCounter, 1)
	defer mm_atomic.AddUint64(&mmCall.afterCallCounter, 1)

	if mmCall.inspectFuncCall != nil {
		mmCall.inspectFuncCall(ctx, ref, method, argsIn, pulse)
	}

	params := &ContractRequesterMockCallParams{ctx, ref, method, argsIn, pulse}

	// Record call args
	mmCall.CallMock.mutex.Lock()
	mmCall.CallMock.callArgs = append(mmCall.CallMock.callArgs, params)
	mmCall.CallMock.mutex.Unlock()

	for _, e := range mmCall.CallMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.rp1, e.results.err
		}
	}

	if mmCall.CallMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCall.CallMock.defaultExpectation.Counter, 1)
		want := mmCall.CallMock.defaultExpectation.params
		got := ContractRequesterMockCallParams{ctx, ref, method, argsIn, pulse}
		if want != nil && !minimock.Equal(*want, got) {
			mmCall.t.Errorf("ContractRequesterMock.Call got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmCall.CallMock.defaultExpectation.results
		if results == nil {
			mmCall.t.Fatal("No results are set for the ContractRequesterMock.Call")
		}
		return (*results).r1, (*results).rp1, (*results).err
	}
	if mmCall.funcCall != nil {
		return mmCall.funcCall(ctx, ref, method, argsIn, pulse)
	}
	mmCall.t.Fatalf("Unexpected call to ContractRequesterMock.Call. %v %v %v %v %v", ctx, ref, method, argsIn, pulse)
	return
}

// CallAfterCounter returns a count of finished ContractRequesterMock.Call invocations
func (mmCall *ContractRequesterMock) CallAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCall.afterCallCounter)
}

// CallBeforeCounter returns a count of ContractRequesterMock.Call invocations
func (mmCall *ContractRequesterMock) CallBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCall.beforeCallCounter)
}

// Calls returns a list of arguments used in each call to ContractRequesterMock.Call.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCall *mContractRequesterMockCall) Calls() []*ContractRequesterMockCallParams {
	mmCall.mutex.RLock()

	argCopy := make([]*ContractRequesterMockCallParams, len(mmCall.callArgs))
	copy(argCopy, mmCall.callArgs)

	mmCall.mutex.RUnlock()

	return argCopy
}

// MinimockCallDone returns true if the count of the Call invocations corresponds
// the number of defined expectations
func (m *ContractRequesterMock) MinimockCallDone() bool {
	for _, e := range m.CallMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CallMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCallCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCall != nil && mm_atomic.LoadUint64(&m.afterCallCounter) < 1 {
		return false
	}
	return true
}

// MinimockCallInspect logs each unmet expectation
func (m *ContractRequesterMock) MinimockCallInspect() {
	for _, e := range m.CallMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ContractRequesterMock.Call with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CallMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCallCounter) < 1 {
		if m.CallMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ContractRequesterMock.Call")
		} else {
			m.t.Errorf("Expected call to ContractRequesterMock.Call with params: %#v", *m.CallMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCall != nil && mm_atomic.LoadUint64(&m.afterCallCounter) < 1 {
		m.t.Error("Expected call to ContractRequesterMock.Call")
	}
}

type mContractRequesterMockSendRequest struct {
	mock               *ContractRequesterMock
	defaultExpectation *ContractRequesterMockSendRequestExpectation
	expectations       []*ContractRequesterMockSendRequestExpectation

	callArgs []*ContractRequesterMockSendRequestParams
	mutex    sync.RWMutex
}

// ContractRequesterMockSendRequestExpectation specifies expectation struct of the ContractRequester.SendRequest
type ContractRequesterMockSendRequestExpectation struct {
	mock    *ContractRequesterMock
	params  *ContractRequesterMockSendRequestParams
	results *ContractRequesterMockSendRequestResults
	Counter uint64
}

// ContractRequesterMockSendRequestParams contains parameters of the ContractRequester.SendRequest
type ContractRequesterMockSendRequestParams struct {
	ctx context.Context
	msg mm_insolar.Payload
}

// ContractRequesterMockSendRequestResults contains results of the ContractRequester.SendRequest
type ContractRequesterMockSendRequestResults struct {
	r1  mm_insolar.Reply
	rp1 *mm_insolar.Reference
	err error
}

// Expect sets up expected params for ContractRequester.SendRequest
func (mmSendRequest *mContractRequesterMockSendRequest) Expect(ctx context.Context, msg mm_insolar.Payload) *mContractRequesterMockSendRequest {
	if mmSendRequest.mock.funcSendRequest != nil {
		mmSendRequest.mock.t.Fatalf("ContractRequesterMock.SendRequest mock is already set by Set")
	}

	if mmSendRequest.defaultExpectation == nil {
		mmSendRequest.defaultExpectation = &ContractRequesterMockSendRequestExpectation{}
	}

	mmSendRequest.defaultExpectation.params = &ContractRequesterMockSendRequestParams{ctx, msg}
	for _, e := range mmSendRequest.expectations {
		if minimock.Equal(e.params, mmSendRequest.defaultExpectation.params) {
			mmSendRequest.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendRequest.defaultExpectation.params)
		}
	}

	return mmSendRequest
}

// Inspect accepts an inspector function that has same arguments as the ContractRequester.SendRequest
func (mmSendRequest *mContractRequesterMockSendRequest) Inspect(f func(ctx context.Context, msg mm_insolar.Payload)) *mContractRequesterMockSendRequest {
	if mmSendRequest.mock.inspectFuncSendRequest != nil {
		mmSendRequest.mock.t.Fatalf("Inspect function is already set for ContractRequesterMock.SendRequest")
	}

	mmSendRequest.mock.inspectFuncSendRequest = f

	return mmSendRequest
}

// Return sets up results that will be returned by ContractRequester.SendRequest
func (mmSendRequest *mContractRequesterMockSendRequest) Return(r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error) *ContractRequesterMock {
	if mmSendRequest.mock.funcSendRequest != nil {
		mmSendRequest.mock.t.Fatalf("ContractRequesterMock.SendRequest mock is already set by Set")
	}

	if mmSendRequest.defaultExpectation == nil {
		mmSendRequest.defaultExpectation = &ContractRequesterMockSendRequestExpectation{mock: mmSendRequest.mock}
	}
	mmSendRequest.defaultExpectation.results = &ContractRequesterMockSendRequestResults{r1, rp1, err}
	return mmSendRequest.mock
}

//Set uses given function f to mock the ContractRequester.SendRequest method
func (mmSendRequest *mContractRequesterMockSendRequest) Set(f func(ctx context.Context, msg mm_insolar.Payload) (r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error)) *ContractRequesterMock {
	if mmSendRequest.defaultExpectation != nil {
		mmSendRequest.mock.t.Fatalf("Default expectation is already set for the ContractRequester.SendRequest method")
	}

	if len(mmSendRequest.expectations) > 0 {
		mmSendRequest.mock.t.Fatalf("Some expectations are already set for the ContractRequester.SendRequest method")
	}

	mmSendRequest.mock.funcSendRequest = f
	return mmSendRequest.mock
}

// When sets expectation for the ContractRequester.SendRequest which will trigger the result defined by the following
// Then helper
func (mmSendRequest *mContractRequesterMockSendRequest) When(ctx context.Context, msg mm_insolar.Payload) *ContractRequesterMockSendRequestExpectation {
	if mmSendRequest.mock.funcSendRequest != nil {
		mmSendRequest.mock.t.Fatalf("ContractRequesterMock.SendRequest mock is already set by Set")
	}

	expectation := &ContractRequesterMockSendRequestExpectation{
		mock:   mmSendRequest.mock,
		params: &ContractRequesterMockSendRequestParams{ctx, msg},
	}
	mmSendRequest.expectations = append(mmSendRequest.expectations, expectation)
	return expectation
}

// Then sets up ContractRequester.SendRequest return parameters for the expectation previously defined by the When method
func (e *ContractRequesterMockSendRequestExpectation) Then(r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error) *ContractRequesterMock {
	e.results = &ContractRequesterMockSendRequestResults{r1, rp1, err}
	return e.mock
}

// SendRequest implements insolar.ContractRequester
func (mmSendRequest *ContractRequesterMock) SendRequest(ctx context.Context, msg mm_insolar.Payload) (r1 mm_insolar.Reply, rp1 *mm_insolar.Reference, err error) {
	mm_atomic.AddUint64(&mmSendRequest.beforeSendRequestCounter, 1)
	defer mm_atomic.AddUint64(&mmSendRequest.afterSendRequestCounter, 1)

	if mmSendRequest.inspectFuncSendRequest != nil {
		mmSendRequest.inspectFuncSendRequest(ctx, msg)
	}

	params := &ContractRequesterMockSendRequestParams{ctx, msg}

	// Record call args
	mmSendRequest.SendRequestMock.mutex.Lock()
	mmSendRequest.SendRequestMock.callArgs = append(mmSendRequest.SendRequestMock.callArgs, params)
	mmSendRequest.SendRequestMock.mutex.Unlock()

	for _, e := range mmSendRequest.SendRequestMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.rp1, e.results.err
		}
	}

	if mmSendRequest.SendRequestMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendRequest.SendRequestMock.defaultExpectation.Counter, 1)
		want := mmSendRequest.SendRequestMock.defaultExpectation.params
		got := ContractRequesterMockSendRequestParams{ctx, msg}
		if want != nil && !minimock.Equal(*want, got) {
			mmSendRequest.t.Errorf("ContractRequesterMock.SendRequest got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSendRequest.SendRequestMock.defaultExpectation.results
		if results == nil {
			mmSendRequest.t.Fatal("No results are set for the ContractRequesterMock.SendRequest")
		}
		return (*results).r1, (*results).rp1, (*results).err
	}
	if mmSendRequest.funcSendRequest != nil {
		return mmSendRequest.funcSendRequest(ctx, msg)
	}
	mmSendRequest.t.Fatalf("Unexpected call to ContractRequesterMock.SendRequest. %v %v", ctx, msg)
	return
}

// SendRequestAfterCounter returns a count of finished ContractRequesterMock.SendRequest invocations
func (mmSendRequest *ContractRequesterMock) SendRequestAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendRequest.afterSendRequestCounter)
}

// SendRequestBeforeCounter returns a count of ContractRequesterMock.SendRequest invocations
func (mmSendRequest *ContractRequesterMock) SendRequestBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendRequest.beforeSendRequestCounter)
}

// Calls returns a list of arguments used in each call to ContractRequesterMock.SendRequest.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendRequest *mContractRequesterMockSendRequest) Calls() []*ContractRequesterMockSendRequestParams {
	mmSendRequest.mutex.RLock()

	argCopy := make([]*ContractRequesterMockSendRequestParams, len(mmSendRequest.callArgs))
	copy(argCopy, mmSendRequest.callArgs)

	mmSendRequest.mutex.RUnlock()

	return argCopy
}

// MinimockSendRequestDone returns true if the count of the SendRequest invocations corresponds
// the number of defined expectations
func (m *ContractRequesterMock) MinimockSendRequestDone() bool {
	for _, e := range m.SendRequestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendRequestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendRequestCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendRequest != nil && mm_atomic.LoadUint64(&m.afterSendRequestCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendRequestInspect logs each unmet expectation
func (m *ContractRequesterMock) MinimockSendRequestInspect() {
	for _, e := range m.SendRequestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ContractRequesterMock.SendRequest with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendRequestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendRequestCounter) < 1 {
		if m.SendRequestMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ContractRequesterMock.SendRequest")
		} else {
			m.t.Errorf("Expected call to ContractRequesterMock.SendRequest with params: %#v", *m.SendRequestMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendRequest != nil && mm_atomic.LoadUint64(&m.afterSendRequestCounter) < 1 {
		m.t.Error("Expected call to ContractRequesterMock.SendRequest")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ContractRequesterMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCallInspect()

		m.MinimockSendRequestInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ContractRequesterMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ContractRequesterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCallDone() &&
		m.MinimockSendRequestDone()
}
