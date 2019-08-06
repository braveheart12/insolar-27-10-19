package testutils

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"crypto"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	mm_insolar "github.com/insolar/insolar/insolar"
)

// CryptographyServiceMock implements insolar.CryptographyService
type CryptographyServiceMock struct {
	t minimock.Tester

	funcGetPublicKey          func() (p1 crypto.PublicKey, err error)
	inspectFuncGetPublicKey   func()
	afterGetPublicKeyCounter  uint64
	beforeGetPublicKeyCounter uint64
	GetPublicKeyMock          mCryptographyServiceMockGetPublicKey

	funcSign          func(ba1 []byte) (sp1 *mm_insolar.Signature, err error)
	inspectFuncSign   func(ba1 []byte)
	afterSignCounter  uint64
	beforeSignCounter uint64
	SignMock          mCryptographyServiceMockSign

	funcVerify          func(p1 crypto.PublicKey, s1 mm_insolar.Signature, ba1 []byte) (b1 bool)
	inspectFuncVerify   func(p1 crypto.PublicKey, s1 mm_insolar.Signature, ba1 []byte)
	afterVerifyCounter  uint64
	beforeVerifyCounter uint64
	VerifyMock          mCryptographyServiceMockVerify
}

// NewCryptographyServiceMock returns a mock for insolar.CryptographyService
func NewCryptographyServiceMock(t minimock.Tester) *CryptographyServiceMock {
	m := &CryptographyServiceMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetPublicKeyMock = mCryptographyServiceMockGetPublicKey{mock: m}

	m.SignMock = mCryptographyServiceMockSign{mock: m}
	m.SignMock.callArgs = []*CryptographyServiceMockSignParams{}

	m.VerifyMock = mCryptographyServiceMockVerify{mock: m}
	m.VerifyMock.callArgs = []*CryptographyServiceMockVerifyParams{}

	return m
}

type mCryptographyServiceMockGetPublicKey struct {
	mock               *CryptographyServiceMock
	defaultExpectation *CryptographyServiceMockGetPublicKeyExpectation
	expectations       []*CryptographyServiceMockGetPublicKeyExpectation
}

// CryptographyServiceMockGetPublicKeyExpectation specifies expectation struct of the CryptographyService.GetPublicKey
type CryptographyServiceMockGetPublicKeyExpectation struct {
	mock *CryptographyServiceMock

	results *CryptographyServiceMockGetPublicKeyResults
	Counter uint64
}

// CryptographyServiceMockGetPublicKeyResults contains results of the CryptographyService.GetPublicKey
type CryptographyServiceMockGetPublicKeyResults struct {
	p1  crypto.PublicKey
	err error
}

// Expect sets up expected params for CryptographyService.GetPublicKey
func (mmGetPublicKey *mCryptographyServiceMockGetPublicKey) Expect() *mCryptographyServiceMockGetPublicKey {
	if mmGetPublicKey.mock.funcGetPublicKey != nil {
		mmGetPublicKey.mock.t.Fatalf("CryptographyServiceMock.GetPublicKey mock is already set by Set")
	}

	if mmGetPublicKey.defaultExpectation == nil {
		mmGetPublicKey.defaultExpectation = &CryptographyServiceMockGetPublicKeyExpectation{}
	}

	return mmGetPublicKey
}

// Inspect accepts an inspector function that has same arguments as the CryptographyService.GetPublicKey
func (mmGetPublicKey *mCryptographyServiceMockGetPublicKey) Inspect(f func()) *mCryptographyServiceMockGetPublicKey {
	if mmGetPublicKey.mock.inspectFuncGetPublicKey != nil {
		mmGetPublicKey.mock.t.Fatalf("Inspect function is already set for CryptographyServiceMock.GetPublicKey")
	}

	mmGetPublicKey.mock.inspectFuncGetPublicKey = f

	return mmGetPublicKey
}

// Return sets up results that will be returned by CryptographyService.GetPublicKey
func (mmGetPublicKey *mCryptographyServiceMockGetPublicKey) Return(p1 crypto.PublicKey, err error) *CryptographyServiceMock {
	if mmGetPublicKey.mock.funcGetPublicKey != nil {
		mmGetPublicKey.mock.t.Fatalf("CryptographyServiceMock.GetPublicKey mock is already set by Set")
	}

	if mmGetPublicKey.defaultExpectation == nil {
		mmGetPublicKey.defaultExpectation = &CryptographyServiceMockGetPublicKeyExpectation{mock: mmGetPublicKey.mock}
	}
	mmGetPublicKey.defaultExpectation.results = &CryptographyServiceMockGetPublicKeyResults{p1, err}
	return mmGetPublicKey.mock
}

//Set uses given function f to mock the CryptographyService.GetPublicKey method
func (mmGetPublicKey *mCryptographyServiceMockGetPublicKey) Set(f func() (p1 crypto.PublicKey, err error)) *CryptographyServiceMock {
	if mmGetPublicKey.defaultExpectation != nil {
		mmGetPublicKey.mock.t.Fatalf("Default expectation is already set for the CryptographyService.GetPublicKey method")
	}

	if len(mmGetPublicKey.expectations) > 0 {
		mmGetPublicKey.mock.t.Fatalf("Some expectations are already set for the CryptographyService.GetPublicKey method")
	}

	mmGetPublicKey.mock.funcGetPublicKey = f
	return mmGetPublicKey.mock
}

// GetPublicKey implements insolar.CryptographyService
func (mmGetPublicKey *CryptographyServiceMock) GetPublicKey() (p1 crypto.PublicKey, err error) {
	mm_atomic.AddUint64(&mmGetPublicKey.beforeGetPublicKeyCounter, 1)
	defer mm_atomic.AddUint64(&mmGetPublicKey.afterGetPublicKeyCounter, 1)

	if mmGetPublicKey.inspectFuncGetPublicKey != nil {
		mmGetPublicKey.inspectFuncGetPublicKey()
	}

	if mmGetPublicKey.GetPublicKeyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetPublicKey.GetPublicKeyMock.defaultExpectation.Counter, 1)

		results := mmGetPublicKey.GetPublicKeyMock.defaultExpectation.results
		if results == nil {
			mmGetPublicKey.t.Fatal("No results are set for the CryptographyServiceMock.GetPublicKey")
		}
		return (*results).p1, (*results).err
	}
	if mmGetPublicKey.funcGetPublicKey != nil {
		return mmGetPublicKey.funcGetPublicKey()
	}
	mmGetPublicKey.t.Fatalf("Unexpected call to CryptographyServiceMock.GetPublicKey.")
	return
}

// GetPublicKeyAfterCounter returns a count of finished CryptographyServiceMock.GetPublicKey invocations
func (mmGetPublicKey *CryptographyServiceMock) GetPublicKeyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPublicKey.afterGetPublicKeyCounter)
}

// GetPublicKeyBeforeCounter returns a count of CryptographyServiceMock.GetPublicKey invocations
func (mmGetPublicKey *CryptographyServiceMock) GetPublicKeyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPublicKey.beforeGetPublicKeyCounter)
}

// MinimockGetPublicKeyDone returns true if the count of the GetPublicKey invocations corresponds
// the number of defined expectations
func (m *CryptographyServiceMock) MinimockGetPublicKeyDone() bool {
	for _, e := range m.GetPublicKeyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetPublicKeyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetPublicKeyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPublicKey != nil && mm_atomic.LoadUint64(&m.afterGetPublicKeyCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetPublicKeyInspect logs each unmet expectation
func (m *CryptographyServiceMock) MinimockGetPublicKeyInspect() {
	for _, e := range m.GetPublicKeyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CryptographyServiceMock.GetPublicKey")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetPublicKeyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetPublicKeyCounter) < 1 {
		m.t.Error("Expected call to CryptographyServiceMock.GetPublicKey")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPublicKey != nil && mm_atomic.LoadUint64(&m.afterGetPublicKeyCounter) < 1 {
		m.t.Error("Expected call to CryptographyServiceMock.GetPublicKey")
	}
}

type mCryptographyServiceMockSign struct {
	mock               *CryptographyServiceMock
	defaultExpectation *CryptographyServiceMockSignExpectation
	expectations       []*CryptographyServiceMockSignExpectation

	callArgs []*CryptographyServiceMockSignParams
	mutex    sync.RWMutex
}

// CryptographyServiceMockSignExpectation specifies expectation struct of the CryptographyService.Sign
type CryptographyServiceMockSignExpectation struct {
	mock    *CryptographyServiceMock
	params  *CryptographyServiceMockSignParams
	results *CryptographyServiceMockSignResults
	Counter uint64
}

// CryptographyServiceMockSignParams contains parameters of the CryptographyService.Sign
type CryptographyServiceMockSignParams struct {
	ba1 []byte
}

// CryptographyServiceMockSignResults contains results of the CryptographyService.Sign
type CryptographyServiceMockSignResults struct {
	sp1 *mm_insolar.Signature
	err error
}

// Expect sets up expected params for CryptographyService.Sign
func (mmSign *mCryptographyServiceMockSign) Expect(ba1 []byte) *mCryptographyServiceMockSign {
	if mmSign.mock.funcSign != nil {
		mmSign.mock.t.Fatalf("CryptographyServiceMock.Sign mock is already set by Set")
	}

	if mmSign.defaultExpectation == nil {
		mmSign.defaultExpectation = &CryptographyServiceMockSignExpectation{}
	}

	mmSign.defaultExpectation.params = &CryptographyServiceMockSignParams{ba1}
	for _, e := range mmSign.expectations {
		if minimock.Equal(e.params, mmSign.defaultExpectation.params) {
			mmSign.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSign.defaultExpectation.params)
		}
	}

	return mmSign
}

// Inspect accepts an inspector function that has same arguments as the CryptographyService.Sign
func (mmSign *mCryptographyServiceMockSign) Inspect(f func(ba1 []byte)) *mCryptographyServiceMockSign {
	if mmSign.mock.inspectFuncSign != nil {
		mmSign.mock.t.Fatalf("Inspect function is already set for CryptographyServiceMock.Sign")
	}

	mmSign.mock.inspectFuncSign = f

	return mmSign
}

// Return sets up results that will be returned by CryptographyService.Sign
func (mmSign *mCryptographyServiceMockSign) Return(sp1 *mm_insolar.Signature, err error) *CryptographyServiceMock {
	if mmSign.mock.funcSign != nil {
		mmSign.mock.t.Fatalf("CryptographyServiceMock.Sign mock is already set by Set")
	}

	if mmSign.defaultExpectation == nil {
		mmSign.defaultExpectation = &CryptographyServiceMockSignExpectation{mock: mmSign.mock}
	}
	mmSign.defaultExpectation.results = &CryptographyServiceMockSignResults{sp1, err}
	return mmSign.mock
}

//Set uses given function f to mock the CryptographyService.Sign method
func (mmSign *mCryptographyServiceMockSign) Set(f func(ba1 []byte) (sp1 *mm_insolar.Signature, err error)) *CryptographyServiceMock {
	if mmSign.defaultExpectation != nil {
		mmSign.mock.t.Fatalf("Default expectation is already set for the CryptographyService.Sign method")
	}

	if len(mmSign.expectations) > 0 {
		mmSign.mock.t.Fatalf("Some expectations are already set for the CryptographyService.Sign method")
	}

	mmSign.mock.funcSign = f
	return mmSign.mock
}

// When sets expectation for the CryptographyService.Sign which will trigger the result defined by the following
// Then helper
func (mmSign *mCryptographyServiceMockSign) When(ba1 []byte) *CryptographyServiceMockSignExpectation {
	if mmSign.mock.funcSign != nil {
		mmSign.mock.t.Fatalf("CryptographyServiceMock.Sign mock is already set by Set")
	}

	expectation := &CryptographyServiceMockSignExpectation{
		mock:   mmSign.mock,
		params: &CryptographyServiceMockSignParams{ba1},
	}
	mmSign.expectations = append(mmSign.expectations, expectation)
	return expectation
}

// Then sets up CryptographyService.Sign return parameters for the expectation previously defined by the When method
func (e *CryptographyServiceMockSignExpectation) Then(sp1 *mm_insolar.Signature, err error) *CryptographyServiceMock {
	e.results = &CryptographyServiceMockSignResults{sp1, err}
	return e.mock
}

// Sign implements insolar.CryptographyService
func (mmSign *CryptographyServiceMock) Sign(ba1 []byte) (sp1 *mm_insolar.Signature, err error) {
	mm_atomic.AddUint64(&mmSign.beforeSignCounter, 1)
	defer mm_atomic.AddUint64(&mmSign.afterSignCounter, 1)

	if mmSign.inspectFuncSign != nil {
		mmSign.inspectFuncSign(ba1)
	}

	params := &CryptographyServiceMockSignParams{ba1}

	// Record call args
	mmSign.SignMock.mutex.Lock()
	mmSign.SignMock.callArgs = append(mmSign.SignMock.callArgs, params)
	mmSign.SignMock.mutex.Unlock()

	for _, e := range mmSign.SignMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sp1, e.results.err
		}
	}

	if mmSign.SignMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSign.SignMock.defaultExpectation.Counter, 1)
		want := mmSign.SignMock.defaultExpectation.params
		got := CryptographyServiceMockSignParams{ba1}
		if want != nil && !minimock.Equal(*want, got) {
			mmSign.t.Errorf("CryptographyServiceMock.Sign got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSign.SignMock.defaultExpectation.results
		if results == nil {
			mmSign.t.Fatal("No results are set for the CryptographyServiceMock.Sign")
		}
		return (*results).sp1, (*results).err
	}
	if mmSign.funcSign != nil {
		return mmSign.funcSign(ba1)
	}
	mmSign.t.Fatalf("Unexpected call to CryptographyServiceMock.Sign. %v", ba1)
	return
}

// SignAfterCounter returns a count of finished CryptographyServiceMock.Sign invocations
func (mmSign *CryptographyServiceMock) SignAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSign.afterSignCounter)
}

// SignBeforeCounter returns a count of CryptographyServiceMock.Sign invocations
func (mmSign *CryptographyServiceMock) SignBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSign.beforeSignCounter)
}

// Calls returns a list of arguments used in each call to CryptographyServiceMock.Sign.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSign *mCryptographyServiceMockSign) Calls() []*CryptographyServiceMockSignParams {
	mmSign.mutex.RLock()

	argCopy := make([]*CryptographyServiceMockSignParams, len(mmSign.callArgs))
	copy(argCopy, mmSign.callArgs)

	mmSign.mutex.RUnlock()

	return argCopy
}

// MinimockSignDone returns true if the count of the Sign invocations corresponds
// the number of defined expectations
func (m *CryptographyServiceMock) MinimockSignDone() bool {
	for _, e := range m.SignMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SignMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSignCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSign != nil && mm_atomic.LoadUint64(&m.afterSignCounter) < 1 {
		return false
	}
	return true
}

// MinimockSignInspect logs each unmet expectation
func (m *CryptographyServiceMock) MinimockSignInspect() {
	for _, e := range m.SignMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CryptographyServiceMock.Sign with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SignMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSignCounter) < 1 {
		if m.SignMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CryptographyServiceMock.Sign")
		} else {
			m.t.Errorf("Expected call to CryptographyServiceMock.Sign with params: %#v", *m.SignMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSign != nil && mm_atomic.LoadUint64(&m.afterSignCounter) < 1 {
		m.t.Error("Expected call to CryptographyServiceMock.Sign")
	}
}

type mCryptographyServiceMockVerify struct {
	mock               *CryptographyServiceMock
	defaultExpectation *CryptographyServiceMockVerifyExpectation
	expectations       []*CryptographyServiceMockVerifyExpectation

	callArgs []*CryptographyServiceMockVerifyParams
	mutex    sync.RWMutex
}

// CryptographyServiceMockVerifyExpectation specifies expectation struct of the CryptographyService.Verify
type CryptographyServiceMockVerifyExpectation struct {
	mock    *CryptographyServiceMock
	params  *CryptographyServiceMockVerifyParams
	results *CryptographyServiceMockVerifyResults
	Counter uint64
}

// CryptographyServiceMockVerifyParams contains parameters of the CryptographyService.Verify
type CryptographyServiceMockVerifyParams struct {
	p1  crypto.PublicKey
	s1  mm_insolar.Signature
	ba1 []byte
}

// CryptographyServiceMockVerifyResults contains results of the CryptographyService.Verify
type CryptographyServiceMockVerifyResults struct {
	b1 bool
}

// Expect sets up expected params for CryptographyService.Verify
func (mmVerify *mCryptographyServiceMockVerify) Expect(p1 crypto.PublicKey, s1 mm_insolar.Signature, ba1 []byte) *mCryptographyServiceMockVerify {
	if mmVerify.mock.funcVerify != nil {
		mmVerify.mock.t.Fatalf("CryptographyServiceMock.Verify mock is already set by Set")
	}

	if mmVerify.defaultExpectation == nil {
		mmVerify.defaultExpectation = &CryptographyServiceMockVerifyExpectation{}
	}

	mmVerify.defaultExpectation.params = &CryptographyServiceMockVerifyParams{p1, s1, ba1}
	for _, e := range mmVerify.expectations {
		if minimock.Equal(e.params, mmVerify.defaultExpectation.params) {
			mmVerify.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmVerify.defaultExpectation.params)
		}
	}

	return mmVerify
}

// Inspect accepts an inspector function that has same arguments as the CryptographyService.Verify
func (mmVerify *mCryptographyServiceMockVerify) Inspect(f func(p1 crypto.PublicKey, s1 mm_insolar.Signature, ba1 []byte)) *mCryptographyServiceMockVerify {
	if mmVerify.mock.inspectFuncVerify != nil {
		mmVerify.mock.t.Fatalf("Inspect function is already set for CryptographyServiceMock.Verify")
	}

	mmVerify.mock.inspectFuncVerify = f

	return mmVerify
}

// Return sets up results that will be returned by CryptographyService.Verify
func (mmVerify *mCryptographyServiceMockVerify) Return(b1 bool) *CryptographyServiceMock {
	if mmVerify.mock.funcVerify != nil {
		mmVerify.mock.t.Fatalf("CryptographyServiceMock.Verify mock is already set by Set")
	}

	if mmVerify.defaultExpectation == nil {
		mmVerify.defaultExpectation = &CryptographyServiceMockVerifyExpectation{mock: mmVerify.mock}
	}
	mmVerify.defaultExpectation.results = &CryptographyServiceMockVerifyResults{b1}
	return mmVerify.mock
}

//Set uses given function f to mock the CryptographyService.Verify method
func (mmVerify *mCryptographyServiceMockVerify) Set(f func(p1 crypto.PublicKey, s1 mm_insolar.Signature, ba1 []byte) (b1 bool)) *CryptographyServiceMock {
	if mmVerify.defaultExpectation != nil {
		mmVerify.mock.t.Fatalf("Default expectation is already set for the CryptographyService.Verify method")
	}

	if len(mmVerify.expectations) > 0 {
		mmVerify.mock.t.Fatalf("Some expectations are already set for the CryptographyService.Verify method")
	}

	mmVerify.mock.funcVerify = f
	return mmVerify.mock
}

// When sets expectation for the CryptographyService.Verify which will trigger the result defined by the following
// Then helper
func (mmVerify *mCryptographyServiceMockVerify) When(p1 crypto.PublicKey, s1 mm_insolar.Signature, ba1 []byte) *CryptographyServiceMockVerifyExpectation {
	if mmVerify.mock.funcVerify != nil {
		mmVerify.mock.t.Fatalf("CryptographyServiceMock.Verify mock is already set by Set")
	}

	expectation := &CryptographyServiceMockVerifyExpectation{
		mock:   mmVerify.mock,
		params: &CryptographyServiceMockVerifyParams{p1, s1, ba1},
	}
	mmVerify.expectations = append(mmVerify.expectations, expectation)
	return expectation
}

// Then sets up CryptographyService.Verify return parameters for the expectation previously defined by the When method
func (e *CryptographyServiceMockVerifyExpectation) Then(b1 bool) *CryptographyServiceMock {
	e.results = &CryptographyServiceMockVerifyResults{b1}
	return e.mock
}

// Verify implements insolar.CryptographyService
func (mmVerify *CryptographyServiceMock) Verify(p1 crypto.PublicKey, s1 mm_insolar.Signature, ba1 []byte) (b1 bool) {
	mm_atomic.AddUint64(&mmVerify.beforeVerifyCounter, 1)
	defer mm_atomic.AddUint64(&mmVerify.afterVerifyCounter, 1)

	if mmVerify.inspectFuncVerify != nil {
		mmVerify.inspectFuncVerify(p1, s1, ba1)
	}

	params := &CryptographyServiceMockVerifyParams{p1, s1, ba1}

	// Record call args
	mmVerify.VerifyMock.mutex.Lock()
	mmVerify.VerifyMock.callArgs = append(mmVerify.VerifyMock.callArgs, params)
	mmVerify.VerifyMock.mutex.Unlock()

	for _, e := range mmVerify.VerifyMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1
		}
	}

	if mmVerify.VerifyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmVerify.VerifyMock.defaultExpectation.Counter, 1)
		want := mmVerify.VerifyMock.defaultExpectation.params
		got := CryptographyServiceMockVerifyParams{p1, s1, ba1}
		if want != nil && !minimock.Equal(*want, got) {
			mmVerify.t.Errorf("CryptographyServiceMock.Verify got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmVerify.VerifyMock.defaultExpectation.results
		if results == nil {
			mmVerify.t.Fatal("No results are set for the CryptographyServiceMock.Verify")
		}
		return (*results).b1
	}
	if mmVerify.funcVerify != nil {
		return mmVerify.funcVerify(p1, s1, ba1)
	}
	mmVerify.t.Fatalf("Unexpected call to CryptographyServiceMock.Verify. %v %v %v", p1, s1, ba1)
	return
}

// VerifyAfterCounter returns a count of finished CryptographyServiceMock.Verify invocations
func (mmVerify *CryptographyServiceMock) VerifyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmVerify.afterVerifyCounter)
}

// VerifyBeforeCounter returns a count of CryptographyServiceMock.Verify invocations
func (mmVerify *CryptographyServiceMock) VerifyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmVerify.beforeVerifyCounter)
}

// Calls returns a list of arguments used in each call to CryptographyServiceMock.Verify.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmVerify *mCryptographyServiceMockVerify) Calls() []*CryptographyServiceMockVerifyParams {
	mmVerify.mutex.RLock()

	argCopy := make([]*CryptographyServiceMockVerifyParams, len(mmVerify.callArgs))
	copy(argCopy, mmVerify.callArgs)

	mmVerify.mutex.RUnlock()

	return argCopy
}

// MinimockVerifyDone returns true if the count of the Verify invocations corresponds
// the number of defined expectations
func (m *CryptographyServiceMock) MinimockVerifyDone() bool {
	for _, e := range m.VerifyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.VerifyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterVerifyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcVerify != nil && mm_atomic.LoadUint64(&m.afterVerifyCounter) < 1 {
		return false
	}
	return true
}

// MinimockVerifyInspect logs each unmet expectation
func (m *CryptographyServiceMock) MinimockVerifyInspect() {
	for _, e := range m.VerifyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CryptographyServiceMock.Verify with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.VerifyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterVerifyCounter) < 1 {
		if m.VerifyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CryptographyServiceMock.Verify")
		} else {
			m.t.Errorf("Expected call to CryptographyServiceMock.Verify with params: %#v", *m.VerifyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcVerify != nil && mm_atomic.LoadUint64(&m.afterVerifyCounter) < 1 {
		m.t.Error("Expected call to CryptographyServiceMock.Verify")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CryptographyServiceMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetPublicKeyInspect()

		m.MinimockSignInspect()

		m.MinimockVerifyInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CryptographyServiceMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CryptographyServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetPublicKeyDone() &&
		m.MinimockSignDone() &&
		m.MinimockVerifyDone()
}
