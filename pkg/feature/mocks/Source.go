// Code generated by mockery v2.32.2. DO NOT EDIT.

package mocks

import (
	context "context"

	feature "github.com/Mellanox/nic-feature-discovery/pkg/feature"
	mock "github.com/stretchr/testify/mock"
)

// Source is an autogenerated mock type for the Source type
type Source struct {
	mock.Mock
}

type Source_Expecter struct {
	mock *mock.Mock
}

func (_m *Source) EXPECT() *Source_Expecter {
	return &Source_Expecter{mock: &_m.Mock}
}

// Discover provides a mock function with given fields: ctx
func (_m *Source) Discover(ctx context.Context) ([]feature.Feature, error) {
	ret := _m.Called(ctx)

	var r0 []feature.Feature
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]feature.Feature, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []feature.Feature); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]feature.Feature)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Source_Discover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Discover'
type Source_Discover_Call struct {
	*mock.Call
}

// Discover is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Source_Expecter) Discover(ctx interface{}) *Source_Discover_Call {
	return &Source_Discover_Call{Call: _e.mock.On("Discover", ctx)}
}

func (_c *Source_Discover_Call) Run(run func(ctx context.Context)) *Source_Discover_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Source_Discover_Call) Return(_a0 []feature.Feature, _a1 error) *Source_Discover_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Source_Discover_Call) RunAndReturn(run func(context.Context) ([]feature.Feature, error)) *Source_Discover_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *Source) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Source_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Source_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Source_Expecter) Name() *Source_Name_Call {
	return &Source_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Source_Name_Call) Run(run func()) *Source_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Source_Name_Call) Return(_a0 string) *Source_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Source_Name_Call) RunAndReturn(run func() string) *Source_Name_Call {
	_c.Call.Return(run)
	return _c
}

// NewSource creates a new instance of Source. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSource(t interface {
	mock.TestingT
	Cleanup(func())
}) *Source {
	mock := &Source{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}