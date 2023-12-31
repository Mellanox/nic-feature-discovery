// Code generated by mockery v2.32.2. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	renameio "github.com/google/renameio/v2"
	mock "github.com/stretchr/testify/mock"
)

// Os is an autogenerated mock type for the Os type
type Os struct {
	mock.Mock
}

type Os_Expecter struct {
	mock *mock.Mock
}

func (_m *Os) EXPECT() *Os_Expecter {
	return &Os_Expecter{mock: &_m.Mock}
}

// Lstat provides a mock function with given fields: name
func (_m *Os) Lstat(name string) (fs.FileInfo, error) {
	ret := _m.Called(name)

	var r0 fs.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (fs.FileInfo, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) fs.FileInfo); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Os_Lstat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Lstat'
type Os_Lstat_Call struct {
	*mock.Call
}

// Lstat is a helper method to define mock.On call
//   - name string
func (_e *Os_Expecter) Lstat(name interface{}) *Os_Lstat_Call {
	return &Os_Lstat_Call{Call: _e.mock.On("Lstat", name)}
}

func (_c *Os_Lstat_Call) Run(run func(name string)) *Os_Lstat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Os_Lstat_Call) Return(_a0 fs.FileInfo, _a1 error) *Os_Lstat_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Os_Lstat_Call) RunAndReturn(run func(string) (fs.FileInfo, error)) *Os_Lstat_Call {
	_c.Call.Return(run)
	return _c
}

// ReadFile provides a mock function with given fields: name
func (_m *Os) ReadFile(name string) ([]byte, error) {
	ret := _m.Called(name)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Os_ReadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadFile'
type Os_ReadFile_Call struct {
	*mock.Call
}

// ReadFile is a helper method to define mock.On call
//   - name string
func (_e *Os_Expecter) ReadFile(name interface{}) *Os_ReadFile_Call {
	return &Os_ReadFile_Call{Call: _e.mock.On("ReadFile", name)}
}

func (_c *Os_ReadFile_Call) Run(run func(name string)) *Os_ReadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Os_ReadFile_Call) Return(_a0 []byte, _a1 error) *Os_ReadFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Os_ReadFile_Call) RunAndReturn(run func(string) ([]byte, error)) *Os_ReadFile_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: name
func (_m *Os) Remove(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Os_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type Os_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - name string
func (_e *Os_Expecter) Remove(name interface{}) *Os_Remove_Call {
	return &Os_Remove_Call{Call: _e.mock.On("Remove", name)}
}

func (_c *Os_Remove_Call) Run(run func(name string)) *Os_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Os_Remove_Call) Return(_a0 error) *Os_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Os_Remove_Call) RunAndReturn(run func(string) error) *Os_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// Stat provides a mock function with given fields: name
func (_m *Os) Stat(name string) (fs.FileInfo, error) {
	ret := _m.Called(name)

	var r0 fs.FileInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (fs.FileInfo, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) fs.FileInfo); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.FileInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Os_Stat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stat'
type Os_Stat_Call struct {
	*mock.Call
}

// Stat is a helper method to define mock.On call
//   - name string
func (_e *Os_Expecter) Stat(name interface{}) *Os_Stat_Call {
	return &Os_Stat_Call{Call: _e.mock.On("Stat", name)}
}

func (_c *Os_Stat_Call) Run(run func(name string)) *Os_Stat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Os_Stat_Call) Return(_a0 fs.FileInfo, _a1 error) *Os_Stat_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Os_Stat_Call) RunAndReturn(run func(string) (fs.FileInfo, error)) *Os_Stat_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFile provides a mock function with given fields: name, data, perm, opts
func (_m *Os) WriteFile(name string, data []byte, perm fs.FileMode, opts ...renameio.Option) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name, data, perm)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte, fs.FileMode, ...renameio.Option) error); ok {
		r0 = rf(name, data, perm, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Os_WriteFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFile'
type Os_WriteFile_Call struct {
	*mock.Call
}

// WriteFile is a helper method to define mock.On call
//   - name string
//   - data []byte
//   - perm fs.FileMode
//   - opts ...renameio.Option
func (_e *Os_Expecter) WriteFile(name interface{}, data interface{}, perm interface{}, opts ...interface{}) *Os_WriteFile_Call {
	return &Os_WriteFile_Call{Call: _e.mock.On("WriteFile",
		append([]interface{}{name, data, perm}, opts...)...)}
}

func (_c *Os_WriteFile_Call) Run(run func(name string, data []byte, perm fs.FileMode, opts ...renameio.Option)) *Os_WriteFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]renameio.Option, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(renameio.Option)
			}
		}
		run(args[0].(string), args[1].([]byte), args[2].(fs.FileMode), variadicArgs...)
	})
	return _c
}

func (_c *Os_WriteFile_Call) Return(_a0 error) *Os_WriteFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Os_WriteFile_Call) RunAndReturn(run func(string, []byte, fs.FileMode, ...renameio.Option) error) *Os_WriteFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewOs creates a new instance of Os. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOs(t interface {
	mock.TestingT
	Cleanup(func())
}) *Os {
	mock := &Os{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
