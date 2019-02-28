package jet

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "JetStorage" can be found in github.com/insolar/insolar/ledger/storage/jet
*/
import (
	context "context"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	core "github.com/insolar/insolar/core"

	testify_assert "github.com/stretchr/testify/assert"
)

//JetStorageMock implements github.com/insolar/insolar/ledger/storage/jet.JetStorage
type JetStorageMock struct {
	t minimock.Tester

	AddJetsFunc       func(p context.Context, p1 ...core.RecordID) (r error)
	AddJetsCounter    uint64
	AddJetsPreCounter uint64
	AddJetsMock       mJetStorageMockAddJets

	CloneJetTreeFunc       func(p context.Context, p1 core.PulseNumber, p2 core.PulseNumber) (r *Tree)
	CloneJetTreeCounter    uint64
	CloneJetTreePreCounter uint64
	CloneJetTreeMock       mJetStorageMockCloneJetTree

	DeleteJetTreeFunc       func(p context.Context, p1 core.PulseNumber)
	DeleteJetTreeCounter    uint64
	DeleteJetTreePreCounter uint64
	DeleteJetTreeMock       mJetStorageMockDeleteJetTree

	FindJetFunc       func(p context.Context, p1 core.PulseNumber, p2 core.RecordID) (r *core.RecordID, r1 bool)
	FindJetCounter    uint64
	FindJetPreCounter uint64
	FindJetMock       mJetStorageMockFindJet

	GetJetsFunc       func(p context.Context) (r IDSet, r1 error)
	GetJetsCounter    uint64
	GetJetsPreCounter uint64
	GetJetsMock       mJetStorageMockGetJets

	SplitJetTreeFunc       func(p context.Context, p1 core.PulseNumber, p2 core.RecordID) (r *core.RecordID, r1 *core.RecordID, r2 error)
	SplitJetTreeCounter    uint64
	SplitJetTreePreCounter uint64
	SplitJetTreeMock       mJetStorageMockSplitJetTree

	UpdateJetTreeFunc       func(p context.Context, p1 core.PulseNumber, p2 bool, p3 ...core.RecordID)
	UpdateJetTreeCounter    uint64
	UpdateJetTreePreCounter uint64
	UpdateJetTreeMock       mJetStorageMockUpdateJetTree
}

//NewJetStorageMock returns a mock for github.com/insolar/insolar/ledger/storage/jet.JetStorage
func NewJetStorageMock(t minimock.Tester) *JetStorageMock {
	m := &JetStorageMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddJetsMock = mJetStorageMockAddJets{mock: m}
	m.CloneJetTreeMock = mJetStorageMockCloneJetTree{mock: m}
	m.DeleteJetTreeMock = mJetStorageMockDeleteJetTree{mock: m}
	m.FindJetMock = mJetStorageMockFindJet{mock: m}
	m.GetJetsMock = mJetStorageMockGetJets{mock: m}
	m.SplitJetTreeMock = mJetStorageMockSplitJetTree{mock: m}
	m.UpdateJetTreeMock = mJetStorageMockUpdateJetTree{mock: m}

	return m
}

type mJetStorageMockAddJets struct {
	mock              *JetStorageMock
	mainExpectation   *JetStorageMockAddJetsExpectation
	expectationSeries []*JetStorageMockAddJetsExpectation
}

type JetStorageMockAddJetsExpectation struct {
	input  *JetStorageMockAddJetsInput
	result *JetStorageMockAddJetsResult
}

type JetStorageMockAddJetsInput struct {
	p  context.Context
	p1 []core.RecordID
}

type JetStorageMockAddJetsResult struct {
	r error
}

//Expect specifies that invocation of JetStorage.AddJets is expected from 1 to Infinity times
func (m *mJetStorageMockAddJets) Expect(p context.Context, p1 ...core.RecordID) *mJetStorageMockAddJets {
	m.mock.AddJetsFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockAddJetsExpectation{}
	}
	m.mainExpectation.input = &JetStorageMockAddJetsInput{p, p1}
	return m
}

//Return specifies results of invocation of JetStorage.AddJets
func (m *mJetStorageMockAddJets) Return(r error) *JetStorageMock {
	m.mock.AddJetsFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockAddJetsExpectation{}
	}
	m.mainExpectation.result = &JetStorageMockAddJetsResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of JetStorage.AddJets is expected once
func (m *mJetStorageMockAddJets) ExpectOnce(p context.Context, p1 ...core.RecordID) *JetStorageMockAddJetsExpectation {
	m.mock.AddJetsFunc = nil
	m.mainExpectation = nil

	expectation := &JetStorageMockAddJetsExpectation{}
	expectation.input = &JetStorageMockAddJetsInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetStorageMockAddJetsExpectation) Return(r error) {
	e.result = &JetStorageMockAddJetsResult{r}
}

//Set uses given function f as a mock of JetStorage.AddJets method
func (m *mJetStorageMockAddJets) Set(f func(p context.Context, p1 ...core.RecordID) (r error)) *JetStorageMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.AddJetsFunc = f
	return m.mock
}

//AddJets implements github.com/insolar/insolar/ledger/storage/jet.JetStorage interface
func (m *JetStorageMock) AddJets(p context.Context, p1 ...core.RecordID) (r error) {
	counter := atomic.AddUint64(&m.AddJetsPreCounter, 1)
	defer atomic.AddUint64(&m.AddJetsCounter, 1)

	if len(m.AddJetsMock.expectationSeries) > 0 {
		if counter > uint64(len(m.AddJetsMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetStorageMock.AddJets. %v %v", p, p1)
			return
		}

		input := m.AddJetsMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetStorageMockAddJetsInput{p, p1}, "JetStorage.AddJets got unexpected parameters")

		result := m.AddJetsMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.AddJets")
			return
		}

		r = result.r

		return
	}

	if m.AddJetsMock.mainExpectation != nil {

		input := m.AddJetsMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetStorageMockAddJetsInput{p, p1}, "JetStorage.AddJets got unexpected parameters")
		}

		result := m.AddJetsMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.AddJets")
		}

		r = result.r

		return
	}

	if m.AddJetsFunc == nil {
		m.t.Fatalf("Unexpected call to JetStorageMock.AddJets. %v %v", p, p1)
		return
	}

	return m.AddJetsFunc(p, p1...)
}

//AddJetsMinimockCounter returns a count of JetStorageMock.AddJetsFunc invocations
func (m *JetStorageMock) AddJetsMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.AddJetsCounter)
}

//AddJetsMinimockPreCounter returns the value of JetStorageMock.AddJets invocations
func (m *JetStorageMock) AddJetsMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.AddJetsPreCounter)
}

//AddJetsFinished returns true if mock invocations count is ok
func (m *JetStorageMock) AddJetsFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.AddJetsMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.AddJetsCounter) == uint64(len(m.AddJetsMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.AddJetsMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.AddJetsCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.AddJetsFunc != nil {
		return atomic.LoadUint64(&m.AddJetsCounter) > 0
	}

	return true
}

type mJetStorageMockCloneJetTree struct {
	mock              *JetStorageMock
	mainExpectation   *JetStorageMockCloneJetTreeExpectation
	expectationSeries []*JetStorageMockCloneJetTreeExpectation
}

type JetStorageMockCloneJetTreeExpectation struct {
	input  *JetStorageMockCloneJetTreeInput
	result *JetStorageMockCloneJetTreeResult
}

type JetStorageMockCloneJetTreeInput struct {
	p  context.Context
	p1 core.PulseNumber
	p2 core.PulseNumber
}

type JetStorageMockCloneJetTreeResult struct {
	r *Tree
}

//Expect specifies that invocation of JetStorage.CloneJetTree is expected from 1 to Infinity times
func (m *mJetStorageMockCloneJetTree) Expect(p context.Context, p1 core.PulseNumber, p2 core.PulseNumber) *mJetStorageMockCloneJetTree {
	m.mock.CloneJetTreeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockCloneJetTreeExpectation{}
	}
	m.mainExpectation.input = &JetStorageMockCloneJetTreeInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetStorage.CloneJetTree
func (m *mJetStorageMockCloneJetTree) Return(r *Tree) *JetStorageMock {
	m.mock.CloneJetTreeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockCloneJetTreeExpectation{}
	}
	m.mainExpectation.result = &JetStorageMockCloneJetTreeResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of JetStorage.CloneJetTree is expected once
func (m *mJetStorageMockCloneJetTree) ExpectOnce(p context.Context, p1 core.PulseNumber, p2 core.PulseNumber) *JetStorageMockCloneJetTreeExpectation {
	m.mock.CloneJetTreeFunc = nil
	m.mainExpectation = nil

	expectation := &JetStorageMockCloneJetTreeExpectation{}
	expectation.input = &JetStorageMockCloneJetTreeInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetStorageMockCloneJetTreeExpectation) Return(r *Tree) {
	e.result = &JetStorageMockCloneJetTreeResult{r}
}

//Set uses given function f as a mock of JetStorage.CloneJetTree method
func (m *mJetStorageMockCloneJetTree) Set(f func(p context.Context, p1 core.PulseNumber, p2 core.PulseNumber) (r *Tree)) *JetStorageMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.CloneJetTreeFunc = f
	return m.mock
}

//CloneJetTree implements github.com/insolar/insolar/ledger/storage/jet.JetStorage interface
func (m *JetStorageMock) CloneJetTree(p context.Context, p1 core.PulseNumber, p2 core.PulseNumber) (r *Tree) {
	counter := atomic.AddUint64(&m.CloneJetTreePreCounter, 1)
	defer atomic.AddUint64(&m.CloneJetTreeCounter, 1)

	if len(m.CloneJetTreeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.CloneJetTreeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetStorageMock.CloneJetTree. %v %v %v", p, p1, p2)
			return
		}

		input := m.CloneJetTreeMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetStorageMockCloneJetTreeInput{p, p1, p2}, "JetStorage.CloneJetTree got unexpected parameters")

		result := m.CloneJetTreeMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.CloneJetTree")
			return
		}

		r = result.r

		return
	}

	if m.CloneJetTreeMock.mainExpectation != nil {

		input := m.CloneJetTreeMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetStorageMockCloneJetTreeInput{p, p1, p2}, "JetStorage.CloneJetTree got unexpected parameters")
		}

		result := m.CloneJetTreeMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.CloneJetTree")
		}

		r = result.r

		return
	}

	if m.CloneJetTreeFunc == nil {
		m.t.Fatalf("Unexpected call to JetStorageMock.CloneJetTree. %v %v %v", p, p1, p2)
		return
	}

	return m.CloneJetTreeFunc(p, p1, p2)
}

//CloneJetTreeMinimockCounter returns a count of JetStorageMock.CloneJetTreeFunc invocations
func (m *JetStorageMock) CloneJetTreeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CloneJetTreeCounter)
}

//CloneJetTreeMinimockPreCounter returns the value of JetStorageMock.CloneJetTree invocations
func (m *JetStorageMock) CloneJetTreeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.CloneJetTreePreCounter)
}

//CloneJetTreeFinished returns true if mock invocations count is ok
func (m *JetStorageMock) CloneJetTreeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.CloneJetTreeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.CloneJetTreeCounter) == uint64(len(m.CloneJetTreeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.CloneJetTreeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.CloneJetTreeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.CloneJetTreeFunc != nil {
		return atomic.LoadUint64(&m.CloneJetTreeCounter) > 0
	}

	return true
}

type mJetStorageMockDeleteJetTree struct {
	mock              *JetStorageMock
	mainExpectation   *JetStorageMockDeleteJetTreeExpectation
	expectationSeries []*JetStorageMockDeleteJetTreeExpectation
}

type JetStorageMockDeleteJetTreeExpectation struct {
	input *JetStorageMockDeleteJetTreeInput
}

type JetStorageMockDeleteJetTreeInput struct {
	p  context.Context
	p1 core.PulseNumber
}

//Expect specifies that invocation of JetStorage.DeleteJetTree is expected from 1 to Infinity times
func (m *mJetStorageMockDeleteJetTree) Expect(p context.Context, p1 core.PulseNumber) *mJetStorageMockDeleteJetTree {
	m.mock.DeleteJetTreeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockDeleteJetTreeExpectation{}
	}
	m.mainExpectation.input = &JetStorageMockDeleteJetTreeInput{p, p1}
	return m
}

//Return specifies results of invocation of JetStorage.DeleteJetTree
func (m *mJetStorageMockDeleteJetTree) Return() *JetStorageMock {
	m.mock.DeleteJetTreeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockDeleteJetTreeExpectation{}
	}

	return m.mock
}

//ExpectOnce specifies that invocation of JetStorage.DeleteJetTree is expected once
func (m *mJetStorageMockDeleteJetTree) ExpectOnce(p context.Context, p1 core.PulseNumber) *JetStorageMockDeleteJetTreeExpectation {
	m.mock.DeleteJetTreeFunc = nil
	m.mainExpectation = nil

	expectation := &JetStorageMockDeleteJetTreeExpectation{}
	expectation.input = &JetStorageMockDeleteJetTreeInput{p, p1}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

//Set uses given function f as a mock of JetStorage.DeleteJetTree method
func (m *mJetStorageMockDeleteJetTree) Set(f func(p context.Context, p1 core.PulseNumber)) *JetStorageMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.DeleteJetTreeFunc = f
	return m.mock
}

//DeleteJetTree implements github.com/insolar/insolar/ledger/storage/jet.JetStorage interface
func (m *JetStorageMock) DeleteJetTree(p context.Context, p1 core.PulseNumber) {
	counter := atomic.AddUint64(&m.DeleteJetTreePreCounter, 1)
	defer atomic.AddUint64(&m.DeleteJetTreeCounter, 1)

	if len(m.DeleteJetTreeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.DeleteJetTreeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetStorageMock.DeleteJetTree. %v %v", p, p1)
			return
		}

		input := m.DeleteJetTreeMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetStorageMockDeleteJetTreeInput{p, p1}, "JetStorage.DeleteJetTree got unexpected parameters")

		return
	}

	if m.DeleteJetTreeMock.mainExpectation != nil {

		input := m.DeleteJetTreeMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetStorageMockDeleteJetTreeInput{p, p1}, "JetStorage.DeleteJetTree got unexpected parameters")
		}

		return
	}

	if m.DeleteJetTreeFunc == nil {
		m.t.Fatalf("Unexpected call to JetStorageMock.DeleteJetTree. %v %v", p, p1)
		return
	}

	m.DeleteJetTreeFunc(p, p1)
}

//DeleteJetTreeMinimockCounter returns a count of JetStorageMock.DeleteJetTreeFunc invocations
func (m *JetStorageMock) DeleteJetTreeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DeleteJetTreeCounter)
}

//DeleteJetTreeMinimockPreCounter returns the value of JetStorageMock.DeleteJetTree invocations
func (m *JetStorageMock) DeleteJetTreeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.DeleteJetTreePreCounter)
}

//DeleteJetTreeFinished returns true if mock invocations count is ok
func (m *JetStorageMock) DeleteJetTreeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.DeleteJetTreeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.DeleteJetTreeCounter) == uint64(len(m.DeleteJetTreeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.DeleteJetTreeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.DeleteJetTreeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.DeleteJetTreeFunc != nil {
		return atomic.LoadUint64(&m.DeleteJetTreeCounter) > 0
	}

	return true
}

type mJetStorageMockFindJet struct {
	mock              *JetStorageMock
	mainExpectation   *JetStorageMockFindJetExpectation
	expectationSeries []*JetStorageMockFindJetExpectation
}

type JetStorageMockFindJetExpectation struct {
	input  *JetStorageMockFindJetInput
	result *JetStorageMockFindJetResult
}

type JetStorageMockFindJetInput struct {
	p  context.Context
	p1 core.PulseNumber
	p2 core.RecordID
}

type JetStorageMockFindJetResult struct {
	r  *core.RecordID
	r1 bool
}

//Expect specifies that invocation of JetStorage.FindJet is expected from 1 to Infinity times
func (m *mJetStorageMockFindJet) Expect(p context.Context, p1 core.PulseNumber, p2 core.RecordID) *mJetStorageMockFindJet {
	m.mock.FindJetFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockFindJetExpectation{}
	}
	m.mainExpectation.input = &JetStorageMockFindJetInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetStorage.FindJet
func (m *mJetStorageMockFindJet) Return(r *core.RecordID, r1 bool) *JetStorageMock {
	m.mock.FindJetFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockFindJetExpectation{}
	}
	m.mainExpectation.result = &JetStorageMockFindJetResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetStorage.FindJet is expected once
func (m *mJetStorageMockFindJet) ExpectOnce(p context.Context, p1 core.PulseNumber, p2 core.RecordID) *JetStorageMockFindJetExpectation {
	m.mock.FindJetFunc = nil
	m.mainExpectation = nil

	expectation := &JetStorageMockFindJetExpectation{}
	expectation.input = &JetStorageMockFindJetInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetStorageMockFindJetExpectation) Return(r *core.RecordID, r1 bool) {
	e.result = &JetStorageMockFindJetResult{r, r1}
}

//Set uses given function f as a mock of JetStorage.FindJet method
func (m *mJetStorageMockFindJet) Set(f func(p context.Context, p1 core.PulseNumber, p2 core.RecordID) (r *core.RecordID, r1 bool)) *JetStorageMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.FindJetFunc = f
	return m.mock
}

//FindJet implements github.com/insolar/insolar/ledger/storage/jet.JetStorage interface
func (m *JetStorageMock) FindJet(p context.Context, p1 core.PulseNumber, p2 core.RecordID) (r *core.RecordID, r1 bool) {
	counter := atomic.AddUint64(&m.FindJetPreCounter, 1)
	defer atomic.AddUint64(&m.FindJetCounter, 1)

	if len(m.FindJetMock.expectationSeries) > 0 {
		if counter > uint64(len(m.FindJetMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetStorageMock.FindJet. %v %v %v", p, p1, p2)
			return
		}

		input := m.FindJetMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetStorageMockFindJetInput{p, p1, p2}, "JetStorage.FindJet got unexpected parameters")

		result := m.FindJetMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.FindJet")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.FindJetMock.mainExpectation != nil {

		input := m.FindJetMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetStorageMockFindJetInput{p, p1, p2}, "JetStorage.FindJet got unexpected parameters")
		}

		result := m.FindJetMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.FindJet")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.FindJetFunc == nil {
		m.t.Fatalf("Unexpected call to JetStorageMock.FindJet. %v %v %v", p, p1, p2)
		return
	}

	return m.FindJetFunc(p, p1, p2)
}

//FindJetMinimockCounter returns a count of JetStorageMock.FindJetFunc invocations
func (m *JetStorageMock) FindJetMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.FindJetCounter)
}

//FindJetMinimockPreCounter returns the value of JetStorageMock.FindJet invocations
func (m *JetStorageMock) FindJetMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.FindJetPreCounter)
}

//FindJetFinished returns true if mock invocations count is ok
func (m *JetStorageMock) FindJetFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.FindJetMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.FindJetCounter) == uint64(len(m.FindJetMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.FindJetMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.FindJetCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.FindJetFunc != nil {
		return atomic.LoadUint64(&m.FindJetCounter) > 0
	}

	return true
}

type mJetStorageMockGetJets struct {
	mock              *JetStorageMock
	mainExpectation   *JetStorageMockGetJetsExpectation
	expectationSeries []*JetStorageMockGetJetsExpectation
}

type JetStorageMockGetJetsExpectation struct {
	input  *JetStorageMockGetJetsInput
	result *JetStorageMockGetJetsResult
}

type JetStorageMockGetJetsInput struct {
	p context.Context
}

type JetStorageMockGetJetsResult struct {
	r  IDSet
	r1 error
}

//Expect specifies that invocation of JetStorage.GetJets is expected from 1 to Infinity times
func (m *mJetStorageMockGetJets) Expect(p context.Context) *mJetStorageMockGetJets {
	m.mock.GetJetsFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockGetJetsExpectation{}
	}
	m.mainExpectation.input = &JetStorageMockGetJetsInput{p}
	return m
}

//Return specifies results of invocation of JetStorage.GetJets
func (m *mJetStorageMockGetJets) Return(r IDSet, r1 error) *JetStorageMock {
	m.mock.GetJetsFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockGetJetsExpectation{}
	}
	m.mainExpectation.result = &JetStorageMockGetJetsResult{r, r1}
	return m.mock
}

//ExpectOnce specifies that invocation of JetStorage.GetJets is expected once
func (m *mJetStorageMockGetJets) ExpectOnce(p context.Context) *JetStorageMockGetJetsExpectation {
	m.mock.GetJetsFunc = nil
	m.mainExpectation = nil

	expectation := &JetStorageMockGetJetsExpectation{}
	expectation.input = &JetStorageMockGetJetsInput{p}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetStorageMockGetJetsExpectation) Return(r IDSet, r1 error) {
	e.result = &JetStorageMockGetJetsResult{r, r1}
}

//Set uses given function f as a mock of JetStorage.GetJets method
func (m *mJetStorageMockGetJets) Set(f func(p context.Context) (r IDSet, r1 error)) *JetStorageMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetJetsFunc = f
	return m.mock
}

//GetJets implements github.com/insolar/insolar/ledger/storage/jet.JetStorage interface
func (m *JetStorageMock) GetJets(p context.Context) (r IDSet, r1 error) {
	counter := atomic.AddUint64(&m.GetJetsPreCounter, 1)
	defer atomic.AddUint64(&m.GetJetsCounter, 1)

	if len(m.GetJetsMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetJetsMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetStorageMock.GetJets. %v", p)
			return
		}

		input := m.GetJetsMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetStorageMockGetJetsInput{p}, "JetStorage.GetJets got unexpected parameters")

		result := m.GetJetsMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.GetJets")
			return
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.GetJetsMock.mainExpectation != nil {

		input := m.GetJetsMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetStorageMockGetJetsInput{p}, "JetStorage.GetJets got unexpected parameters")
		}

		result := m.GetJetsMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.GetJets")
		}

		r = result.r
		r1 = result.r1

		return
	}

	if m.GetJetsFunc == nil {
		m.t.Fatalf("Unexpected call to JetStorageMock.GetJets. %v", p)
		return
	}

	return m.GetJetsFunc(p)
}

//GetJetsMinimockCounter returns a count of JetStorageMock.GetJetsFunc invocations
func (m *JetStorageMock) GetJetsMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetJetsCounter)
}

//GetJetsMinimockPreCounter returns the value of JetStorageMock.GetJets invocations
func (m *JetStorageMock) GetJetsMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetJetsPreCounter)
}

//GetJetsFinished returns true if mock invocations count is ok
func (m *JetStorageMock) GetJetsFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetJetsMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetJetsCounter) == uint64(len(m.GetJetsMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetJetsMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetJetsCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetJetsFunc != nil {
		return atomic.LoadUint64(&m.GetJetsCounter) > 0
	}

	return true
}

type mJetStorageMockSplitJetTree struct {
	mock              *JetStorageMock
	mainExpectation   *JetStorageMockSplitJetTreeExpectation
	expectationSeries []*JetStorageMockSplitJetTreeExpectation
}

type JetStorageMockSplitJetTreeExpectation struct {
	input  *JetStorageMockSplitJetTreeInput
	result *JetStorageMockSplitJetTreeResult
}

type JetStorageMockSplitJetTreeInput struct {
	p  context.Context
	p1 core.PulseNumber
	p2 core.RecordID
}

type JetStorageMockSplitJetTreeResult struct {
	r  *core.RecordID
	r1 *core.RecordID
	r2 error
}

//Expect specifies that invocation of JetStorage.SplitJetTree is expected from 1 to Infinity times
func (m *mJetStorageMockSplitJetTree) Expect(p context.Context, p1 core.PulseNumber, p2 core.RecordID) *mJetStorageMockSplitJetTree {
	m.mock.SplitJetTreeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockSplitJetTreeExpectation{}
	}
	m.mainExpectation.input = &JetStorageMockSplitJetTreeInput{p, p1, p2}
	return m
}

//Return specifies results of invocation of JetStorage.SplitJetTree
func (m *mJetStorageMockSplitJetTree) Return(r *core.RecordID, r1 *core.RecordID, r2 error) *JetStorageMock {
	m.mock.SplitJetTreeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockSplitJetTreeExpectation{}
	}
	m.mainExpectation.result = &JetStorageMockSplitJetTreeResult{r, r1, r2}
	return m.mock
}

//ExpectOnce specifies that invocation of JetStorage.SplitJetTree is expected once
func (m *mJetStorageMockSplitJetTree) ExpectOnce(p context.Context, p1 core.PulseNumber, p2 core.RecordID) *JetStorageMockSplitJetTreeExpectation {
	m.mock.SplitJetTreeFunc = nil
	m.mainExpectation = nil

	expectation := &JetStorageMockSplitJetTreeExpectation{}
	expectation.input = &JetStorageMockSplitJetTreeInput{p, p1, p2}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *JetStorageMockSplitJetTreeExpectation) Return(r *core.RecordID, r1 *core.RecordID, r2 error) {
	e.result = &JetStorageMockSplitJetTreeResult{r, r1, r2}
}

//Set uses given function f as a mock of JetStorage.SplitJetTree method
func (m *mJetStorageMockSplitJetTree) Set(f func(p context.Context, p1 core.PulseNumber, p2 core.RecordID) (r *core.RecordID, r1 *core.RecordID, r2 error)) *JetStorageMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.SplitJetTreeFunc = f
	return m.mock
}

//SplitJetTree implements github.com/insolar/insolar/ledger/storage/jet.JetStorage interface
func (m *JetStorageMock) SplitJetTree(p context.Context, p1 core.PulseNumber, p2 core.RecordID) (r *core.RecordID, r1 *core.RecordID, r2 error) {
	counter := atomic.AddUint64(&m.SplitJetTreePreCounter, 1)
	defer atomic.AddUint64(&m.SplitJetTreeCounter, 1)

	if len(m.SplitJetTreeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.SplitJetTreeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetStorageMock.SplitJetTree. %v %v %v", p, p1, p2)
			return
		}

		input := m.SplitJetTreeMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetStorageMockSplitJetTreeInput{p, p1, p2}, "JetStorage.SplitJetTree got unexpected parameters")

		result := m.SplitJetTreeMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.SplitJetTree")
			return
		}

		r = result.r
		r1 = result.r1
		r2 = result.r2

		return
	}

	if m.SplitJetTreeMock.mainExpectation != nil {

		input := m.SplitJetTreeMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetStorageMockSplitJetTreeInput{p, p1, p2}, "JetStorage.SplitJetTree got unexpected parameters")
		}

		result := m.SplitJetTreeMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the JetStorageMock.SplitJetTree")
		}

		r = result.r
		r1 = result.r1
		r2 = result.r2

		return
	}

	if m.SplitJetTreeFunc == nil {
		m.t.Fatalf("Unexpected call to JetStorageMock.SplitJetTree. %v %v %v", p, p1, p2)
		return
	}

	return m.SplitJetTreeFunc(p, p1, p2)
}

//SplitJetTreeMinimockCounter returns a count of JetStorageMock.SplitJetTreeFunc invocations
func (m *JetStorageMock) SplitJetTreeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SplitJetTreeCounter)
}

//SplitJetTreeMinimockPreCounter returns the value of JetStorageMock.SplitJetTree invocations
func (m *JetStorageMock) SplitJetTreeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.SplitJetTreePreCounter)
}

//SplitJetTreeFinished returns true if mock invocations count is ok
func (m *JetStorageMock) SplitJetTreeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.SplitJetTreeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.SplitJetTreeCounter) == uint64(len(m.SplitJetTreeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.SplitJetTreeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.SplitJetTreeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.SplitJetTreeFunc != nil {
		return atomic.LoadUint64(&m.SplitJetTreeCounter) > 0
	}

	return true
}

type mJetStorageMockUpdateJetTree struct {
	mock              *JetStorageMock
	mainExpectation   *JetStorageMockUpdateJetTreeExpectation
	expectationSeries []*JetStorageMockUpdateJetTreeExpectation
}

type JetStorageMockUpdateJetTreeExpectation struct {
	input *JetStorageMockUpdateJetTreeInput
}

type JetStorageMockUpdateJetTreeInput struct {
	p  context.Context
	p1 core.PulseNumber
	p2 bool
	p3 []core.RecordID
}

//Expect specifies that invocation of JetStorage.UpdateJetTree is expected from 1 to Infinity times
func (m *mJetStorageMockUpdateJetTree) Expect(p context.Context, p1 core.PulseNumber, p2 bool, p3 ...core.RecordID) *mJetStorageMockUpdateJetTree {
	m.mock.UpdateJetTreeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockUpdateJetTreeExpectation{}
	}
	m.mainExpectation.input = &JetStorageMockUpdateJetTreeInput{p, p1, p2, p3}
	return m
}

//Return specifies results of invocation of JetStorage.UpdateJetTree
func (m *mJetStorageMockUpdateJetTree) Return() *JetStorageMock {
	m.mock.UpdateJetTreeFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &JetStorageMockUpdateJetTreeExpectation{}
	}

	return m.mock
}

//ExpectOnce specifies that invocation of JetStorage.UpdateJetTree is expected once
func (m *mJetStorageMockUpdateJetTree) ExpectOnce(p context.Context, p1 core.PulseNumber, p2 bool, p3 ...core.RecordID) *JetStorageMockUpdateJetTreeExpectation {
	m.mock.UpdateJetTreeFunc = nil
	m.mainExpectation = nil

	expectation := &JetStorageMockUpdateJetTreeExpectation{}
	expectation.input = &JetStorageMockUpdateJetTreeInput{p, p1, p2, p3}
	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

//Set uses given function f as a mock of JetStorage.UpdateJetTree method
func (m *mJetStorageMockUpdateJetTree) Set(f func(p context.Context, p1 core.PulseNumber, p2 bool, p3 ...core.RecordID)) *JetStorageMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.UpdateJetTreeFunc = f
	return m.mock
}

//UpdateJetTree implements github.com/insolar/insolar/ledger/storage/jet.JetStorage interface
func (m *JetStorageMock) UpdateJetTree(p context.Context, p1 core.PulseNumber, p2 bool, p3 ...core.RecordID) {
	counter := atomic.AddUint64(&m.UpdateJetTreePreCounter, 1)
	defer atomic.AddUint64(&m.UpdateJetTreeCounter, 1)

	if len(m.UpdateJetTreeMock.expectationSeries) > 0 {
		if counter > uint64(len(m.UpdateJetTreeMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to JetStorageMock.UpdateJetTree. %v %v %v %v", p, p1, p2, p3)
			return
		}

		input := m.UpdateJetTreeMock.expectationSeries[counter-1].input
		testify_assert.Equal(m.t, *input, JetStorageMockUpdateJetTreeInput{p, p1, p2, p3}, "JetStorage.UpdateJetTree got unexpected parameters")

		return
	}

	if m.UpdateJetTreeMock.mainExpectation != nil {

		input := m.UpdateJetTreeMock.mainExpectation.input
		if input != nil {
			testify_assert.Equal(m.t, *input, JetStorageMockUpdateJetTreeInput{p, p1, p2, p3}, "JetStorage.UpdateJetTree got unexpected parameters")
		}

		return
	}

	if m.UpdateJetTreeFunc == nil {
		m.t.Fatalf("Unexpected call to JetStorageMock.UpdateJetTree. %v %v %v %v", p, p1, p2, p3)
		return
	}

	m.UpdateJetTreeFunc(p, p1, p2, p3...)
}

//UpdateJetTreeMinimockCounter returns a count of JetStorageMock.UpdateJetTreeFunc invocations
func (m *JetStorageMock) UpdateJetTreeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.UpdateJetTreeCounter)
}

//UpdateJetTreeMinimockPreCounter returns the value of JetStorageMock.UpdateJetTree invocations
func (m *JetStorageMock) UpdateJetTreeMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.UpdateJetTreePreCounter)
}

//UpdateJetTreeFinished returns true if mock invocations count is ok
func (m *JetStorageMock) UpdateJetTreeFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.UpdateJetTreeMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.UpdateJetTreeCounter) == uint64(len(m.UpdateJetTreeMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.UpdateJetTreeMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.UpdateJetTreeCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.UpdateJetTreeFunc != nil {
		return atomic.LoadUint64(&m.UpdateJetTreeCounter) > 0
	}

	return true
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *JetStorageMock) ValidateCallCounters() {

	if !m.AddJetsFinished() {
		m.t.Fatal("Expected call to JetStorageMock.AddJets")
	}

	if !m.CloneJetTreeFinished() {
		m.t.Fatal("Expected call to JetStorageMock.CloneJetTree")
	}

	if !m.DeleteJetTreeFinished() {
		m.t.Fatal("Expected call to JetStorageMock.DeleteJetTree")
	}

	if !m.FindJetFinished() {
		m.t.Fatal("Expected call to JetStorageMock.FindJet")
	}

	if !m.GetJetsFinished() {
		m.t.Fatal("Expected call to JetStorageMock.GetJets")
	}

	if !m.SplitJetTreeFinished() {
		m.t.Fatal("Expected call to JetStorageMock.SplitJetTree")
	}

	if !m.UpdateJetTreeFinished() {
		m.t.Fatal("Expected call to JetStorageMock.UpdateJetTree")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *JetStorageMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *JetStorageMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *JetStorageMock) MinimockFinish() {

	if !m.AddJetsFinished() {
		m.t.Fatal("Expected call to JetStorageMock.AddJets")
	}

	if !m.CloneJetTreeFinished() {
		m.t.Fatal("Expected call to JetStorageMock.CloneJetTree")
	}

	if !m.DeleteJetTreeFinished() {
		m.t.Fatal("Expected call to JetStorageMock.DeleteJetTree")
	}

	if !m.FindJetFinished() {
		m.t.Fatal("Expected call to JetStorageMock.FindJet")
	}

	if !m.GetJetsFinished() {
		m.t.Fatal("Expected call to JetStorageMock.GetJets")
	}

	if !m.SplitJetTreeFinished() {
		m.t.Fatal("Expected call to JetStorageMock.SplitJetTree")
	}

	if !m.UpdateJetTreeFinished() {
		m.t.Fatal("Expected call to JetStorageMock.UpdateJetTree")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *JetStorageMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *JetStorageMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.AddJetsFinished()
		ok = ok && m.CloneJetTreeFinished()
		ok = ok && m.DeleteJetTreeFinished()
		ok = ok && m.FindJetFinished()
		ok = ok && m.GetJetsFinished()
		ok = ok && m.SplitJetTreeFinished()
		ok = ok && m.UpdateJetTreeFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.AddJetsFinished() {
				m.t.Error("Expected call to JetStorageMock.AddJets")
			}

			if !m.CloneJetTreeFinished() {
				m.t.Error("Expected call to JetStorageMock.CloneJetTree")
			}

			if !m.DeleteJetTreeFinished() {
				m.t.Error("Expected call to JetStorageMock.DeleteJetTree")
			}

			if !m.FindJetFinished() {
				m.t.Error("Expected call to JetStorageMock.FindJet")
			}

			if !m.GetJetsFinished() {
				m.t.Error("Expected call to JetStorageMock.GetJets")
			}

			if !m.SplitJetTreeFinished() {
				m.t.Error("Expected call to JetStorageMock.SplitJetTree")
			}

			if !m.UpdateJetTreeFinished() {
				m.t.Error("Expected call to JetStorageMock.UpdateJetTree")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *JetStorageMock) AllMocksCalled() bool {

	if !m.AddJetsFinished() {
		return false
	}

	if !m.CloneJetTreeFinished() {
		return false
	}

	if !m.DeleteJetTreeFinished() {
		return false
	}

	if !m.FindJetFinished() {
		return false
	}

	if !m.GetJetsFinished() {
		return false
	}

	if !m.SplitJetTreeFinished() {
		return false
	}

	if !m.UpdateJetTreeFinished() {
		return false
	}

	return true
}
