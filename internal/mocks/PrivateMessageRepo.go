// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import (
	context "context"

	pm "github.com/bangumi/server/internal/pm"
	mock "github.com/stretchr/testify/mock"
)

// PrivateMessageRepo is an autogenerated mock type for the Repo type
type PrivateMessageRepo struct {
	mock.Mock
}

type PrivateMessageRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *PrivateMessageRepo) EXPECT() *PrivateMessageRepo_Expecter {
	return &PrivateMessageRepo_Expecter{mock: &_m.Mock}
}

// CountByFolder provides a mock function with given fields: ctx, userID, folder
func (_m *PrivateMessageRepo) CountByFolder(ctx context.Context, userID uint32, folder pm.FolderType) (int64, error) {
	ret := _m.Called(ctx, userID, folder)

	if len(ret) == 0 {
		panic("no return value specified for CountByFolder")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, pm.FolderType) (int64, error)); ok {
		return rf(ctx, userID, folder)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, pm.FolderType) int64); ok {
		r0 = rf(ctx, userID, folder)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, pm.FolderType) error); ok {
		r1 = rf(ctx, userID, folder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrivateMessageRepo_CountByFolder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountByFolder'
type PrivateMessageRepo_CountByFolder_Call struct {
	*mock.Call
}

// CountByFolder is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - folder pm.FolderType
func (_e *PrivateMessageRepo_Expecter) CountByFolder(ctx interface{}, userID interface{}, folder interface{}) *PrivateMessageRepo_CountByFolder_Call {
	return &PrivateMessageRepo_CountByFolder_Call{Call: _e.mock.On("CountByFolder", ctx, userID, folder)}
}

func (_c *PrivateMessageRepo_CountByFolder_Call) Run(run func(ctx context.Context, userID uint32, folder pm.FolderType)) *PrivateMessageRepo_CountByFolder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(pm.FolderType))
	})
	return _c
}

func (_c *PrivateMessageRepo_CountByFolder_Call) Return(_a0 int64, _a1 error) *PrivateMessageRepo_CountByFolder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PrivateMessageRepo_CountByFolder_Call) RunAndReturn(run func(context.Context, uint32, pm.FolderType) (int64, error)) *PrivateMessageRepo_CountByFolder_Call {
	_c.Call.Return(run)
	return _c
}

// CountTypes provides a mock function with given fields: ctx, userID
func (_m *PrivateMessageRepo) CountTypes(ctx context.Context, userID uint32) (pm.PrivateMessageTypeCounts, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for CountTypes")
	}

	var r0 pm.PrivateMessageTypeCounts
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (pm.PrivateMessageTypeCounts, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) pm.PrivateMessageTypeCounts); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(pm.PrivateMessageTypeCounts)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrivateMessageRepo_CountTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountTypes'
type PrivateMessageRepo_CountTypes_Call struct {
	*mock.Call
}

// CountTypes is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
func (_e *PrivateMessageRepo_Expecter) CountTypes(ctx interface{}, userID interface{}) *PrivateMessageRepo_CountTypes_Call {
	return &PrivateMessageRepo_CountTypes_Call{Call: _e.mock.On("CountTypes", ctx, userID)}
}

func (_c *PrivateMessageRepo_CountTypes_Call) Run(run func(ctx context.Context, userID uint32)) *PrivateMessageRepo_CountTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *PrivateMessageRepo_CountTypes_Call) Return(_a0 pm.PrivateMessageTypeCounts, _a1 error) *PrivateMessageRepo_CountTypes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PrivateMessageRepo_CountTypes_Call) RunAndReturn(run func(context.Context, uint32) (pm.PrivateMessageTypeCounts, error)) *PrivateMessageRepo_CountTypes_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, senderID, receiverIDs, relatedIDFilter, title, content
func (_m *PrivateMessageRepo) Create(ctx context.Context, senderID uint32, receiverIDs []uint32, relatedIDFilter pm.IDFilter, title string, content string) ([]pm.PrivateMessage, error) {
	ret := _m.Called(ctx, senderID, receiverIDs, relatedIDFilter, title, content)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 []pm.PrivateMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, []uint32, pm.IDFilter, string, string) ([]pm.PrivateMessage, error)); ok {
		return rf(ctx, senderID, receiverIDs, relatedIDFilter, title, content)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, []uint32, pm.IDFilter, string, string) []pm.PrivateMessage); ok {
		r0 = rf(ctx, senderID, receiverIDs, relatedIDFilter, title, content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pm.PrivateMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, []uint32, pm.IDFilter, string, string) error); ok {
		r1 = rf(ctx, senderID, receiverIDs, relatedIDFilter, title, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrivateMessageRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PrivateMessageRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - senderID uint32
//   - receiverIDs []uint32
//   - relatedIDFilter pm.IDFilter
//   - title string
//   - content string
func (_e *PrivateMessageRepo_Expecter) Create(ctx interface{}, senderID interface{}, receiverIDs interface{}, relatedIDFilter interface{}, title interface{}, content interface{}) *PrivateMessageRepo_Create_Call {
	return &PrivateMessageRepo_Create_Call{Call: _e.mock.On("Create", ctx, senderID, receiverIDs, relatedIDFilter, title, content)}
}

func (_c *PrivateMessageRepo_Create_Call) Run(run func(ctx context.Context, senderID uint32, receiverIDs []uint32, relatedIDFilter pm.IDFilter, title string, content string)) *PrivateMessageRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].([]uint32), args[3].(pm.IDFilter), args[4].(string), args[5].(string))
	})
	return _c
}

func (_c *PrivateMessageRepo_Create_Call) Return(_a0 []pm.PrivateMessage, _a1 error) *PrivateMessageRepo_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PrivateMessageRepo_Create_Call) RunAndReturn(run func(context.Context, uint32, []uint32, pm.IDFilter, string, string) ([]pm.PrivateMessage, error)) *PrivateMessageRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, userID, ids
func (_m *PrivateMessageRepo) Delete(ctx context.Context, userID uint32, ids []uint32) error {
	ret := _m.Called(ctx, userID, ids)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, []uint32) error); ok {
		r0 = rf(ctx, userID, ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrivateMessageRepo_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type PrivateMessageRepo_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - ids []uint32
func (_e *PrivateMessageRepo_Expecter) Delete(ctx interface{}, userID interface{}, ids interface{}) *PrivateMessageRepo_Delete_Call {
	return &PrivateMessageRepo_Delete_Call{Call: _e.mock.On("Delete", ctx, userID, ids)}
}

func (_c *PrivateMessageRepo_Delete_Call) Run(run func(ctx context.Context, userID uint32, ids []uint32)) *PrivateMessageRepo_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].([]uint32))
	})
	return _c
}

func (_c *PrivateMessageRepo_Delete_Call) Return(_a0 error) *PrivateMessageRepo_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PrivateMessageRepo_Delete_Call) RunAndReturn(run func(context.Context, uint32, []uint32) error) *PrivateMessageRepo_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, userID, folder, offset, limit
func (_m *PrivateMessageRepo) List(ctx context.Context, userID uint32, folder pm.FolderType, offset int, limit int) ([]pm.PrivateMessageListItem, error) {
	ret := _m.Called(ctx, userID, folder, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []pm.PrivateMessageListItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, pm.FolderType, int, int) ([]pm.PrivateMessageListItem, error)); ok {
		return rf(ctx, userID, folder, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, pm.FolderType, int, int) []pm.PrivateMessageListItem); ok {
		r0 = rf(ctx, userID, folder, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pm.PrivateMessageListItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, pm.FolderType, int, int) error); ok {
		r1 = rf(ctx, userID, folder, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrivateMessageRepo_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type PrivateMessageRepo_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - folder pm.FolderType
//   - offset int
//   - limit int
func (_e *PrivateMessageRepo_Expecter) List(ctx interface{}, userID interface{}, folder interface{}, offset interface{}, limit interface{}) *PrivateMessageRepo_List_Call {
	return &PrivateMessageRepo_List_Call{Call: _e.mock.On("List", ctx, userID, folder, offset, limit)}
}

func (_c *PrivateMessageRepo_List_Call) Run(run func(ctx context.Context, userID uint32, folder pm.FolderType, offset int, limit int)) *PrivateMessageRepo_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(pm.FolderType), args[3].(int), args[4].(int))
	})
	return _c
}

func (_c *PrivateMessageRepo_List_Call) Return(_a0 []pm.PrivateMessageListItem, _a1 error) *PrivateMessageRepo_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PrivateMessageRepo_List_Call) RunAndReturn(run func(context.Context, uint32, pm.FolderType, int, int) ([]pm.PrivateMessageListItem, error)) *PrivateMessageRepo_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListRecentContact provides a mock function with given fields: ctx, userID
func (_m *PrivateMessageRepo) ListRecentContact(ctx context.Context, userID uint32) ([]uint32, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for ListRecentContact")
	}

	var r0 []uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]uint32, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []uint32); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrivateMessageRepo_ListRecentContact_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListRecentContact'
type PrivateMessageRepo_ListRecentContact_Call struct {
	*mock.Call
}

// ListRecentContact is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
func (_e *PrivateMessageRepo_Expecter) ListRecentContact(ctx interface{}, userID interface{}) *PrivateMessageRepo_ListRecentContact_Call {
	return &PrivateMessageRepo_ListRecentContact_Call{Call: _e.mock.On("ListRecentContact", ctx, userID)}
}

func (_c *PrivateMessageRepo_ListRecentContact_Call) Run(run func(ctx context.Context, userID uint32)) *PrivateMessageRepo_ListRecentContact_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *PrivateMessageRepo_ListRecentContact_Call) Return(_a0 []uint32, _a1 error) *PrivateMessageRepo_ListRecentContact_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PrivateMessageRepo_ListRecentContact_Call) RunAndReturn(run func(context.Context, uint32) ([]uint32, error)) *PrivateMessageRepo_ListRecentContact_Call {
	_c.Call.Return(run)
	return _c
}

// ListRelated provides a mock function with given fields: ctx, userID, id
func (_m *PrivateMessageRepo) ListRelated(ctx context.Context, userID uint32, id uint32) ([]pm.PrivateMessage, error) {
	ret := _m.Called(ctx, userID, id)

	if len(ret) == 0 {
		panic("no return value specified for ListRelated")
	}

	var r0 []pm.PrivateMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) ([]pm.PrivateMessage, error)); ok {
		return rf(ctx, userID, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) []pm.PrivateMessage); ok {
		r0 = rf(ctx, userID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pm.PrivateMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint32) error); ok {
		r1 = rf(ctx, userID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrivateMessageRepo_ListRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListRelated'
type PrivateMessageRepo_ListRelated_Call struct {
	*mock.Call
}

// ListRelated is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - id uint32
func (_e *PrivateMessageRepo_Expecter) ListRelated(ctx interface{}, userID interface{}, id interface{}) *PrivateMessageRepo_ListRelated_Call {
	return &PrivateMessageRepo_ListRelated_Call{Call: _e.mock.On("ListRelated", ctx, userID, id)}
}

func (_c *PrivateMessageRepo_ListRelated_Call) Run(run func(ctx context.Context, userID uint32, id uint32)) *PrivateMessageRepo_ListRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *PrivateMessageRepo_ListRelated_Call) Return(_a0 []pm.PrivateMessage, _a1 error) *PrivateMessageRepo_ListRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PrivateMessageRepo_ListRelated_Call) RunAndReturn(run func(context.Context, uint32, uint32) ([]pm.PrivateMessage, error)) *PrivateMessageRepo_ListRelated_Call {
	_c.Call.Return(run)
	return _c
}

// MarkRead provides a mock function with given fields: ctx, userID, relatedID
func (_m *PrivateMessageRepo) MarkRead(ctx context.Context, userID uint32, relatedID uint32) error {
	ret := _m.Called(ctx, userID, relatedID)

	if len(ret) == 0 {
		panic("no return value specified for MarkRead")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) error); ok {
		r0 = rf(ctx, userID, relatedID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrivateMessageRepo_MarkRead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkRead'
type PrivateMessageRepo_MarkRead_Call struct {
	*mock.Call
}

// MarkRead is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - relatedID uint32
func (_e *PrivateMessageRepo_Expecter) MarkRead(ctx interface{}, userID interface{}, relatedID interface{}) *PrivateMessageRepo_MarkRead_Call {
	return &PrivateMessageRepo_MarkRead_Call{Call: _e.mock.On("MarkRead", ctx, userID, relatedID)}
}

func (_c *PrivateMessageRepo_MarkRead_Call) Run(run func(ctx context.Context, userID uint32, relatedID uint32)) *PrivateMessageRepo_MarkRead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *PrivateMessageRepo_MarkRead_Call) Return(_a0 error) *PrivateMessageRepo_MarkRead_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PrivateMessageRepo_MarkRead_Call) RunAndReturn(run func(context.Context, uint32, uint32) error) *PrivateMessageRepo_MarkRead_Call {
	_c.Call.Return(run)
	return _c
}

// NewPrivateMessageRepo creates a new instance of PrivateMessageRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPrivateMessageRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *PrivateMessageRepo {
	mock := &PrivateMessageRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
