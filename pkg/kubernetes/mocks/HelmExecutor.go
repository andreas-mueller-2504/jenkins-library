// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// HelmExecutor is an autogenerated mock type for the HelmExecutor type
type HelmExecutor struct {
	mock.Mock
}

type HelmExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *HelmExecutor) EXPECT() *HelmExecutor_Expecter {
	return &HelmExecutor_Expecter{mock: &_m.Mock}
}

// RunHelmDependency provides a mock function with given fields:
func (_m *HelmExecutor) RunHelmDependency() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RunHelmDependency")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HelmExecutor_RunHelmDependency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunHelmDependency'
type HelmExecutor_RunHelmDependency_Call struct {
	*mock.Call
}

// RunHelmDependency is a helper method to define mock.On call
func (_e *HelmExecutor_Expecter) RunHelmDependency() *HelmExecutor_RunHelmDependency_Call {
	return &HelmExecutor_RunHelmDependency_Call{Call: _e.mock.On("RunHelmDependency")}
}

func (_c *HelmExecutor_RunHelmDependency_Call) Run(run func()) *HelmExecutor_RunHelmDependency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HelmExecutor_RunHelmDependency_Call) Return(_a0 error) *HelmExecutor_RunHelmDependency_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HelmExecutor_RunHelmDependency_Call) RunAndReturn(run func() error) *HelmExecutor_RunHelmDependency_Call {
	_c.Call.Return(run)
	return _c
}

// RunHelmInstall provides a mock function with given fields:
func (_m *HelmExecutor) RunHelmInstall() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RunHelmInstall")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HelmExecutor_RunHelmInstall_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunHelmInstall'
type HelmExecutor_RunHelmInstall_Call struct {
	*mock.Call
}

// RunHelmInstall is a helper method to define mock.On call
func (_e *HelmExecutor_Expecter) RunHelmInstall() *HelmExecutor_RunHelmInstall_Call {
	return &HelmExecutor_RunHelmInstall_Call{Call: _e.mock.On("RunHelmInstall")}
}

func (_c *HelmExecutor_RunHelmInstall_Call) Run(run func()) *HelmExecutor_RunHelmInstall_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HelmExecutor_RunHelmInstall_Call) Return(_a0 error) *HelmExecutor_RunHelmInstall_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HelmExecutor_RunHelmInstall_Call) RunAndReturn(run func() error) *HelmExecutor_RunHelmInstall_Call {
	_c.Call.Return(run)
	return _c
}

// RunHelmLint provides a mock function with given fields:
func (_m *HelmExecutor) RunHelmLint() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RunHelmLint")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HelmExecutor_RunHelmLint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunHelmLint'
type HelmExecutor_RunHelmLint_Call struct {
	*mock.Call
}

// RunHelmLint is a helper method to define mock.On call
func (_e *HelmExecutor_Expecter) RunHelmLint() *HelmExecutor_RunHelmLint_Call {
	return &HelmExecutor_RunHelmLint_Call{Call: _e.mock.On("RunHelmLint")}
}

func (_c *HelmExecutor_RunHelmLint_Call) Run(run func()) *HelmExecutor_RunHelmLint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HelmExecutor_RunHelmLint_Call) Return(_a0 error) *HelmExecutor_RunHelmLint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HelmExecutor_RunHelmLint_Call) RunAndReturn(run func() error) *HelmExecutor_RunHelmLint_Call {
	_c.Call.Return(run)
	return _c
}

// RunHelmPublish provides a mock function with given fields:
func (_m *HelmExecutor) RunHelmPublish() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RunHelmPublish")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HelmExecutor_RunHelmPublish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunHelmPublish'
type HelmExecutor_RunHelmPublish_Call struct {
	*mock.Call
}

// RunHelmPublish is a helper method to define mock.On call
func (_e *HelmExecutor_Expecter) RunHelmPublish() *HelmExecutor_RunHelmPublish_Call {
	return &HelmExecutor_RunHelmPublish_Call{Call: _e.mock.On("RunHelmPublish")}
}

func (_c *HelmExecutor_RunHelmPublish_Call) Run(run func()) *HelmExecutor_RunHelmPublish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HelmExecutor_RunHelmPublish_Call) Return(_a0 string, _a1 error) *HelmExecutor_RunHelmPublish_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HelmExecutor_RunHelmPublish_Call) RunAndReturn(run func() (string, error)) *HelmExecutor_RunHelmPublish_Call {
	_c.Call.Return(run)
	return _c
}

// RunHelmTest provides a mock function with given fields:
func (_m *HelmExecutor) RunHelmTest() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RunHelmTest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HelmExecutor_RunHelmTest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunHelmTest'
type HelmExecutor_RunHelmTest_Call struct {
	*mock.Call
}

// RunHelmTest is a helper method to define mock.On call
func (_e *HelmExecutor_Expecter) RunHelmTest() *HelmExecutor_RunHelmTest_Call {
	return &HelmExecutor_RunHelmTest_Call{Call: _e.mock.On("RunHelmTest")}
}

func (_c *HelmExecutor_RunHelmTest_Call) Run(run func()) *HelmExecutor_RunHelmTest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HelmExecutor_RunHelmTest_Call) Return(_a0 error) *HelmExecutor_RunHelmTest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HelmExecutor_RunHelmTest_Call) RunAndReturn(run func() error) *HelmExecutor_RunHelmTest_Call {
	_c.Call.Return(run)
	return _c
}

// RunHelmUninstall provides a mock function with given fields:
func (_m *HelmExecutor) RunHelmUninstall() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RunHelmUninstall")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HelmExecutor_RunHelmUninstall_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunHelmUninstall'
type HelmExecutor_RunHelmUninstall_Call struct {
	*mock.Call
}

// RunHelmUninstall is a helper method to define mock.On call
func (_e *HelmExecutor_Expecter) RunHelmUninstall() *HelmExecutor_RunHelmUninstall_Call {
	return &HelmExecutor_RunHelmUninstall_Call{Call: _e.mock.On("RunHelmUninstall")}
}

func (_c *HelmExecutor_RunHelmUninstall_Call) Run(run func()) *HelmExecutor_RunHelmUninstall_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HelmExecutor_RunHelmUninstall_Call) Return(_a0 error) *HelmExecutor_RunHelmUninstall_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HelmExecutor_RunHelmUninstall_Call) RunAndReturn(run func() error) *HelmExecutor_RunHelmUninstall_Call {
	_c.Call.Return(run)
	return _c
}

// RunHelmUpgrade provides a mock function with given fields:
func (_m *HelmExecutor) RunHelmUpgrade() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RunHelmUpgrade")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HelmExecutor_RunHelmUpgrade_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunHelmUpgrade'
type HelmExecutor_RunHelmUpgrade_Call struct {
	*mock.Call
}

// RunHelmUpgrade is a helper method to define mock.On call
func (_e *HelmExecutor_Expecter) RunHelmUpgrade() *HelmExecutor_RunHelmUpgrade_Call {
	return &HelmExecutor_RunHelmUpgrade_Call{Call: _e.mock.On("RunHelmUpgrade")}
}

func (_c *HelmExecutor_RunHelmUpgrade_Call) Run(run func()) *HelmExecutor_RunHelmUpgrade_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HelmExecutor_RunHelmUpgrade_Call) Return(_a0 error) *HelmExecutor_RunHelmUpgrade_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HelmExecutor_RunHelmUpgrade_Call) RunAndReturn(run func() error) *HelmExecutor_RunHelmUpgrade_Call {
	_c.Call.Return(run)
	return _c
}

// NewHelmExecutor creates a new instance of HelmExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHelmExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *HelmExecutor {
	mock := &HelmExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
