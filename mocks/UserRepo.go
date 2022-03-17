// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/model"
)

// UserRepo is an autogenerated mock type for the UserRepo type
type UserRepo struct {
	mock.Mock
}

type UserRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepo) EXPECT() *UserRepo_Expecter {
	return &UserRepo_Expecter{mock: &_m.Mock}
}

// CountCollections provides a mock function with given fields: ctx, userID, subjectType, collectionType, showPrivate
func (_m *UserRepo) CountCollections(ctx context.Context, userID uint32, subjectType uint8, collectionType uint8, showPrivate bool) (int64, error) {
	ret := _m.Called(ctx, userID, subjectType, collectionType, showPrivate)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, uint8, bool) int64); ok {
		r0 = rf(ctx, userID, subjectType, collectionType, showPrivate)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8, uint8, bool) error); ok {
		r1 = rf(ctx, userID, subjectType, collectionType, showPrivate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_CountCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountCollections'
type UserRepo_CountCollections_Call struct {
	*mock.Call
}

// CountCollections is a helper method to define mock.On call
//  - ctx context.Context
//  - userID uint32
//  - subjectType uint8
//  - collectionType uint8
//  - showPrivate bool
func (_e *UserRepo_Expecter) CountCollections(ctx interface{}, userID interface{}, subjectType interface{}, collectionType interface{}, showPrivate interface{}) *UserRepo_CountCollections_Call {
	return &UserRepo_CountCollections_Call{Call: _e.mock.On("CountCollections", ctx, userID, subjectType, collectionType, showPrivate)}
}

func (_c *UserRepo_CountCollections_Call) Run(run func(ctx context.Context, userID uint32, subjectType uint8, collectionType uint8, showPrivate bool)) *UserRepo_CountCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8), args[3].(uint8), args[4].(bool))
	})
	return _c
}

func (_c *UserRepo_CountCollections_Call) Return(_a0 int64, _a1 error) *UserRepo_CountCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByID provides a mock function with given fields: ctx, userID
func (_m *UserRepo) GetByID(ctx context.Context, userID uint32) (model.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.User); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type UserRepo_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//  - ctx context.Context
//  - userID uint32
func (_e *UserRepo_Expecter) GetByID(ctx interface{}, userID interface{}) *UserRepo_GetByID_Call {
	return &UserRepo_GetByID_Call{Call: _e.mock.On("GetByID", ctx, userID)}
}

func (_c *UserRepo_GetByID_Call) Run(run func(ctx context.Context, userID uint32)) *UserRepo_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *UserRepo_GetByID_Call) Return(_a0 model.User, _a1 error) *UserRepo_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, ids
func (_m *UserRepo) GetByIDs(ctx context.Context, ids ...uint32) (map[uint32]model.User, error) {
	_va := make([]interface{}, len(ids))
	for _i := range ids {
		_va[_i] = ids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[uint32]model.User
	if rf, ok := ret.Get(0).(func(context.Context, ...uint32) map[uint32]model.User); ok {
		r0 = rf(ctx, ids...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...uint32) error); ok {
		r1 = rf(ctx, ids...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type UserRepo_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//  - ctx context.Context
//  - ids ...uint32
func (_e *UserRepo_Expecter) GetByIDs(ctx interface{}, ids ...interface{}) *UserRepo_GetByIDs_Call {
	return &UserRepo_GetByIDs_Call{Call: _e.mock.On("GetByIDs",
		append([]interface{}{ctx}, ids...)...)}
}

func (_c *UserRepo_GetByIDs_Call) Run(run func(ctx context.Context, ids ...uint32)) *UserRepo_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]uint32, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(uint32)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *UserRepo_GetByIDs_Call) Return(_a0 map[uint32]model.User, _a1 error) *UserRepo_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByName provides a mock function with given fields: ctx, username
func (_m *UserRepo) GetByName(ctx context.Context, username string) (model.User, error) {
	ret := _m.Called(ctx, username)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) model.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type UserRepo_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//  - ctx context.Context
//  - username string
func (_e *UserRepo_Expecter) GetByName(ctx interface{}, username interface{}) *UserRepo_GetByName_Call {
	return &UserRepo_GetByName_Call{Call: _e.mock.On("GetByName", ctx, username)}
}

func (_c *UserRepo_GetByName_Call) Run(run func(ctx context.Context, username string)) *UserRepo_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepo_GetByName_Call) Return(_a0 model.User, _a1 error) *UserRepo_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListCollections provides a mock function with given fields: ctx, userID, subjectType, collectionType, showPrivate, limit, offset
func (_m *UserRepo) ListCollections(ctx context.Context, userID uint32, subjectType uint8, collectionType uint8, showPrivate bool, limit int, offset int) ([]model.Collection, error) {
	ret := _m.Called(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)

	var r0 []model.Collection
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, uint8, bool, int, int) []model.Collection); ok {
		r0 = rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8, uint8, bool, int, int) error); ok {
		r1 = rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_ListCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListCollections'
type UserRepo_ListCollections_Call struct {
	*mock.Call
}

// ListCollections is a helper method to define mock.On call
//  - ctx context.Context
//  - userID uint32
//  - subjectType uint8
//  - collectionType uint8
//  - showPrivate bool
//  - limit int
//  - offset int
func (_e *UserRepo_Expecter) ListCollections(ctx interface{}, userID interface{}, subjectType interface{}, collectionType interface{}, showPrivate interface{}, limit interface{}, offset interface{}) *UserRepo_ListCollections_Call {
	return &UserRepo_ListCollections_Call{Call: _e.mock.On("ListCollections", ctx, userID, subjectType, collectionType, showPrivate, limit, offset)}
}

func (_c *UserRepo_ListCollections_Call) Run(run func(ctx context.Context, userID uint32, subjectType uint8, collectionType uint8, showPrivate bool, limit int, offset int)) *UserRepo_ListCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8), args[3].(uint8), args[4].(bool), args[5].(int), args[6].(int))
	})
	return _c
}

func (_c *UserRepo_ListCollections_Call) Return(_a0 []model.Collection, _a1 error) *UserRepo_ListCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}