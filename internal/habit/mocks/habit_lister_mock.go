// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package mocks

//go:generate minimock -i habits/internal/habit.habitLister -o habit_lister_mock.go -n HabitListerMock -p mocks

import (
	"context"
	mm_habit "habits/internal/habit"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// HabitListerMock implements mm_habit.habitLister
type HabitListerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcFindAll          func(ctx context.Context) (ha1 []mm_habit.Habit, err error)
	funcFindAllOrigin    string
	inspectFuncFindAll   func(ctx context.Context)
	afterFindAllCounter  uint64
	beforeFindAllCounter uint64
	FindAllMock          mHabitListerMockFindAll
}

// NewHabitListerMock returns a mock for mm_habit.habitLister
func NewHabitListerMock(t minimock.Tester) *HabitListerMock {
	m := &HabitListerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FindAllMock = mHabitListerMockFindAll{mock: m}
	m.FindAllMock.callArgs = []*HabitListerMockFindAllParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mHabitListerMockFindAll struct {
	optional           bool
	mock               *HabitListerMock
	defaultExpectation *HabitListerMockFindAllExpectation
	expectations       []*HabitListerMockFindAllExpectation

	callArgs []*HabitListerMockFindAllParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// HabitListerMockFindAllExpectation specifies expectation struct of the habitLister.FindAll
type HabitListerMockFindAllExpectation struct {
	mock               *HabitListerMock
	params             *HabitListerMockFindAllParams
	paramPtrs          *HabitListerMockFindAllParamPtrs
	expectationOrigins HabitListerMockFindAllExpectationOrigins
	results            *HabitListerMockFindAllResults
	returnOrigin       string
	Counter            uint64
}

// HabitListerMockFindAllParams contains parameters of the habitLister.FindAll
type HabitListerMockFindAllParams struct {
	ctx context.Context
}

// HabitListerMockFindAllParamPtrs contains pointers to parameters of the habitLister.FindAll
type HabitListerMockFindAllParamPtrs struct {
	ctx *context.Context
}

// HabitListerMockFindAllResults contains results of the habitLister.FindAll
type HabitListerMockFindAllResults struct {
	ha1 []mm_habit.Habit
	err error
}

// HabitListerMockFindAllOrigins contains origins of expectations of the habitLister.FindAll
type HabitListerMockFindAllExpectationOrigins struct {
	origin    string
	originCtx string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmFindAll *mHabitListerMockFindAll) Optional() *mHabitListerMockFindAll {
	mmFindAll.optional = true
	return mmFindAll
}

// Expect sets up expected params for habitLister.FindAll
func (mmFindAll *mHabitListerMockFindAll) Expect(ctx context.Context) *mHabitListerMockFindAll {
	if mmFindAll.mock.funcFindAll != nil {
		mmFindAll.mock.t.Fatalf("HabitListerMock.FindAll mock is already set by Set")
	}

	if mmFindAll.defaultExpectation == nil {
		mmFindAll.defaultExpectation = &HabitListerMockFindAllExpectation{}
	}

	if mmFindAll.defaultExpectation.paramPtrs != nil {
		mmFindAll.mock.t.Fatalf("HabitListerMock.FindAll mock is already set by ExpectParams functions")
	}

	mmFindAll.defaultExpectation.params = &HabitListerMockFindAllParams{ctx}
	mmFindAll.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmFindAll.expectations {
		if minimock.Equal(e.params, mmFindAll.defaultExpectation.params) {
			mmFindAll.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFindAll.defaultExpectation.params)
		}
	}

	return mmFindAll
}

// ExpectCtxParam1 sets up expected param ctx for habitLister.FindAll
func (mmFindAll *mHabitListerMockFindAll) ExpectCtxParam1(ctx context.Context) *mHabitListerMockFindAll {
	if mmFindAll.mock.funcFindAll != nil {
		mmFindAll.mock.t.Fatalf("HabitListerMock.FindAll mock is already set by Set")
	}

	if mmFindAll.defaultExpectation == nil {
		mmFindAll.defaultExpectation = &HabitListerMockFindAllExpectation{}
	}

	if mmFindAll.defaultExpectation.params != nil {
		mmFindAll.mock.t.Fatalf("HabitListerMock.FindAll mock is already set by Expect")
	}

	if mmFindAll.defaultExpectation.paramPtrs == nil {
		mmFindAll.defaultExpectation.paramPtrs = &HabitListerMockFindAllParamPtrs{}
	}
	mmFindAll.defaultExpectation.paramPtrs.ctx = &ctx
	mmFindAll.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmFindAll
}

// Inspect accepts an inspector function that has same arguments as the habitLister.FindAll
func (mmFindAll *mHabitListerMockFindAll) Inspect(f func(ctx context.Context)) *mHabitListerMockFindAll {
	if mmFindAll.mock.inspectFuncFindAll != nil {
		mmFindAll.mock.t.Fatalf("Inspect function is already set for HabitListerMock.FindAll")
	}

	mmFindAll.mock.inspectFuncFindAll = f

	return mmFindAll
}

// Return sets up results that will be returned by habitLister.FindAll
func (mmFindAll *mHabitListerMockFindAll) Return(ha1 []mm_habit.Habit, err error) *HabitListerMock {
	if mmFindAll.mock.funcFindAll != nil {
		mmFindAll.mock.t.Fatalf("HabitListerMock.FindAll mock is already set by Set")
	}

	if mmFindAll.defaultExpectation == nil {
		mmFindAll.defaultExpectation = &HabitListerMockFindAllExpectation{mock: mmFindAll.mock}
	}
	mmFindAll.defaultExpectation.results = &HabitListerMockFindAllResults{ha1, err}
	mmFindAll.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmFindAll.mock
}

// Set uses given function f to mock the habitLister.FindAll method
func (mmFindAll *mHabitListerMockFindAll) Set(f func(ctx context.Context) (ha1 []mm_habit.Habit, err error)) *HabitListerMock {
	if mmFindAll.defaultExpectation != nil {
		mmFindAll.mock.t.Fatalf("Default expectation is already set for the habitLister.FindAll method")
	}

	if len(mmFindAll.expectations) > 0 {
		mmFindAll.mock.t.Fatalf("Some expectations are already set for the habitLister.FindAll method")
	}

	mmFindAll.mock.funcFindAll = f
	mmFindAll.mock.funcFindAllOrigin = minimock.CallerInfo(1)
	return mmFindAll.mock
}

// When sets expectation for the habitLister.FindAll which will trigger the result defined by the following
// Then helper
func (mmFindAll *mHabitListerMockFindAll) When(ctx context.Context) *HabitListerMockFindAllExpectation {
	if mmFindAll.mock.funcFindAll != nil {
		mmFindAll.mock.t.Fatalf("HabitListerMock.FindAll mock is already set by Set")
	}

	expectation := &HabitListerMockFindAllExpectation{
		mock:               mmFindAll.mock,
		params:             &HabitListerMockFindAllParams{ctx},
		expectationOrigins: HabitListerMockFindAllExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmFindAll.expectations = append(mmFindAll.expectations, expectation)
	return expectation
}

// Then sets up habitLister.FindAll return parameters for the expectation previously defined by the When method
func (e *HabitListerMockFindAllExpectation) Then(ha1 []mm_habit.Habit, err error) *HabitListerMock {
	e.results = &HabitListerMockFindAllResults{ha1, err}
	return e.mock
}

// Times sets number of times habitLister.FindAll should be invoked
func (mmFindAll *mHabitListerMockFindAll) Times(n uint64) *mHabitListerMockFindAll {
	if n == 0 {
		mmFindAll.mock.t.Fatalf("Times of HabitListerMock.FindAll mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmFindAll.expectedInvocations, n)
	mmFindAll.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmFindAll
}

func (mmFindAll *mHabitListerMockFindAll) invocationsDone() bool {
	if len(mmFindAll.expectations) == 0 && mmFindAll.defaultExpectation == nil && mmFindAll.mock.funcFindAll == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmFindAll.mock.afterFindAllCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmFindAll.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// FindAll implements mm_habit.habitLister
func (mmFindAll *HabitListerMock) FindAll(ctx context.Context) (ha1 []mm_habit.Habit, err error) {
	mm_atomic.AddUint64(&mmFindAll.beforeFindAllCounter, 1)
	defer mm_atomic.AddUint64(&mmFindAll.afterFindAllCounter, 1)

	mmFindAll.t.Helper()

	if mmFindAll.inspectFuncFindAll != nil {
		mmFindAll.inspectFuncFindAll(ctx)
	}

	mm_params := HabitListerMockFindAllParams{ctx}

	// Record call args
	mmFindAll.FindAllMock.mutex.Lock()
	mmFindAll.FindAllMock.callArgs = append(mmFindAll.FindAllMock.callArgs, &mm_params)
	mmFindAll.FindAllMock.mutex.Unlock()

	for _, e := range mmFindAll.FindAllMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ha1, e.results.err
		}
	}

	if mmFindAll.FindAllMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFindAll.FindAllMock.defaultExpectation.Counter, 1)
		mm_want := mmFindAll.FindAllMock.defaultExpectation.params
		mm_want_ptrs := mmFindAll.FindAllMock.defaultExpectation.paramPtrs

		mm_got := HabitListerMockFindAllParams{ctx}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmFindAll.t.Errorf("HabitListerMock.FindAll got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmFindAll.FindAllMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFindAll.t.Errorf("HabitListerMock.FindAll got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmFindAll.FindAllMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFindAll.FindAllMock.defaultExpectation.results
		if mm_results == nil {
			mmFindAll.t.Fatal("No results are set for the HabitListerMock.FindAll")
		}
		return (*mm_results).ha1, (*mm_results).err
	}
	if mmFindAll.funcFindAll != nil {
		return mmFindAll.funcFindAll(ctx)
	}
	mmFindAll.t.Fatalf("Unexpected call to HabitListerMock.FindAll. %v", ctx)
	return
}

// FindAllAfterCounter returns a count of finished HabitListerMock.FindAll invocations
func (mmFindAll *HabitListerMock) FindAllAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindAll.afterFindAllCounter)
}

// FindAllBeforeCounter returns a count of HabitListerMock.FindAll invocations
func (mmFindAll *HabitListerMock) FindAllBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindAll.beforeFindAllCounter)
}

// Calls returns a list of arguments used in each call to HabitListerMock.FindAll.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFindAll *mHabitListerMockFindAll) Calls() []*HabitListerMockFindAllParams {
	mmFindAll.mutex.RLock()

	argCopy := make([]*HabitListerMockFindAllParams, len(mmFindAll.callArgs))
	copy(argCopy, mmFindAll.callArgs)

	mmFindAll.mutex.RUnlock()

	return argCopy
}

// MinimockFindAllDone returns true if the count of the FindAll invocations corresponds
// the number of defined expectations
func (m *HabitListerMock) MinimockFindAllDone() bool {
	if m.FindAllMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.FindAllMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.FindAllMock.invocationsDone()
}

// MinimockFindAllInspect logs each unmet expectation
func (m *HabitListerMock) MinimockFindAllInspect() {
	for _, e := range m.FindAllMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to HabitListerMock.FindAll at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterFindAllCounter := mm_atomic.LoadUint64(&m.afterFindAllCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.FindAllMock.defaultExpectation != nil && afterFindAllCounter < 1 {
		if m.FindAllMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to HabitListerMock.FindAll at\n%s", m.FindAllMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to HabitListerMock.FindAll at\n%s with params: %#v", m.FindAllMock.defaultExpectation.expectationOrigins.origin, *m.FindAllMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindAll != nil && afterFindAllCounter < 1 {
		m.t.Errorf("Expected call to HabitListerMock.FindAll at\n%s", m.funcFindAllOrigin)
	}

	if !m.FindAllMock.invocationsDone() && afterFindAllCounter > 0 {
		m.t.Errorf("Expected %d calls to HabitListerMock.FindAll at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.FindAllMock.expectedInvocations), m.FindAllMock.expectedInvocationsOrigin, afterFindAllCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *HabitListerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockFindAllInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *HabitListerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *HabitListerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockFindAllDone()
}