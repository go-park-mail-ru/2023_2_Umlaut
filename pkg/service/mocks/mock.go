// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	model "github.com/go-park-mail-ru/2023_2_Umlaut/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// DeleteCookie mocks base method.
func (m *MockAuthorization) DeleteCookie(ctx context.Context, session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCookie", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCookie indicates an expected call of DeleteCookie.
func (mr *MockAuthorizationMockRecorder) DeleteCookie(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCookie", reflect.TypeOf((*MockAuthorization)(nil).DeleteCookie), ctx, session)
}

// GenerateCookie mocks base method.
func (m *MockAuthorization) GenerateCookie(ctx context.Context, id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCookie", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCookie indicates an expected call of GenerateCookie.
func (mr *MockAuthorizationMockRecorder) GenerateCookie(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCookie", reflect.TypeOf((*MockAuthorization)(nil).GenerateCookie), ctx, id)
}

// GetUser mocks base method.
func (m *MockAuthorization) GetUser(mail, password string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", mail, password)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAuthorizationMockRecorder) GetUser(mail, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthorization)(nil).GetUser), mail, password)
}

// MockFeed is a mock of Feed interface.
type MockFeed struct {
	ctrl     *gomock.Controller
	recorder *MockFeedMockRecorder
}

// MockFeedMockRecorder is the mock recorder for MockFeed.
type MockFeedMockRecorder struct {
	mock *MockFeed
}

// NewMockFeed creates a new mock instance.
func NewMockFeed(ctrl *gomock.Controller) *MockFeed {
	mock := &MockFeed{ctrl: ctrl}
	mock.recorder = &MockFeedMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeed) EXPECT() *MockFeedMockRecorder {
	return m.recorder
}

// GetNextUser mocks base method.
func (m *MockFeed) GetNextUser(ctx context.Context, session string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextUser", ctx, session)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextUser indicates an expected call of GetNextUser.
func (mr *MockFeedMockRecorder) GetNextUser(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextUser", reflect.TypeOf((*MockFeed)(nil).GetNextUser), ctx, session)
}
