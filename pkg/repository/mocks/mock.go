// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	model "github.com/go-park-mail-ru/2023_2_Umlaut/model"
	gomock "github.com/golang/mock/gomock"
	pgx "github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(ctx context.Context, user model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), ctx, user)
}

// GetNextUser mocks base method.
func (m *MockUser) GetNextUser(ctx context.Context, user model.User, params model.FilterParams) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextUser", ctx, user, params)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextUser indicates an expected call of GetNextUser.
func (mr *MockUserMockRecorder) GetNextUser(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextUser", reflect.TypeOf((*MockUser)(nil).GetNextUser), ctx, user, params)
}

// GetUser mocks base method.
func (m *MockUser) GetUser(ctx context.Context, mail string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, mail)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserMockRecorder) GetUser(ctx, mail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUser)(nil).GetUser), ctx, mail)
}

// GetUserById mocks base method.
func (m *MockUser) GetUserById(ctx context.Context, id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUser)(nil).GetUserById), ctx, id)
}

// GetUserInvites mocks base method.
func (m *MockUser) GetUserInvites(ctx context.Context, userId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInvites", ctx, userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInvites indicates an expected call of GetUserInvites.
func (mr *MockUserMockRecorder) GetUserInvites(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInvites", reflect.TypeOf((*MockUser)(nil).GetUserInvites), ctx, userId)
}

// ResetLikeCounter mocks base method.
func (m *MockUser) ResetLikeCounter(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetLikeCounter", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetLikeCounter indicates an expected call of ResetLikeCounter.
func (mr *MockUserMockRecorder) ResetLikeCounter(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetLikeCounter", reflect.TypeOf((*MockUser)(nil).ResetLikeCounter), ctx)
}

// ShowCSAT mocks base method.
func (m *MockUser) ShowCSAT(ctx context.Context, userId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowCSAT", ctx, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowCSAT indicates an expected call of ShowCSAT.
func (mr *MockUserMockRecorder) ShowCSAT(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowCSAT", reflect.TypeOf((*MockUser)(nil).ShowCSAT), ctx, userId)
}

// UpdateUser mocks base method.
func (m *MockUser) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), ctx, user)
}

// UpdateUserPassword mocks base method.
func (m *MockUser) UpdateUserPassword(ctx context.Context, user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockUserMockRecorder) UpdateUserPassword(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockUser)(nil).UpdateUserPassword), ctx, user)
}

// MockLike is a mock of Like interface.
type MockLike struct {
	ctrl     *gomock.Controller
	recorder *MockLikeMockRecorder
}

// MockLikeMockRecorder is the mock recorder for MockLike.
type MockLikeMockRecorder struct {
	mock *MockLike
}

// NewMockLike creates a new mock instance.
func NewMockLike(ctrl *gomock.Controller) *MockLike {
	mock := &MockLike{ctrl: ctrl}
	mock.recorder = &MockLikeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLike) EXPECT() *MockLikeMockRecorder {
	return m.recorder
}

// CreateLike mocks base method.
func (m *MockLike) CreateLike(ctx context.Context, like model.Like) (model.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLike", ctx, like)
	ret0, _ := ret[0].(model.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLike indicates an expected call of CreateLike.
func (mr *MockLikeMockRecorder) CreateLike(ctx, like interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike", reflect.TypeOf((*MockLike)(nil).CreateLike), ctx, like)
}

// GetUserLikedToLikes mocks base method.
func (m *MockLike) GetUserLikedToLikes(ctx context.Context, userId int) ([]model.PremiumLike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserLikedToLikes", ctx, userId)
	ret0, _ := ret[0].([]model.PremiumLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserLikedToLikes indicates an expected call of GetUserLikedToLikes.
func (mr *MockLikeMockRecorder) GetUserLikedToLikes(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserLikedToLikes", reflect.TypeOf((*MockLike)(nil).GetUserLikedToLikes), ctx, userId)
}

// IsMutualLike mocks base method.
func (m *MockLike) IsMutualLike(ctx context.Context, like model.Like) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMutualLike", ctx, like)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMutualLike indicates an expected call of IsMutualLike.
func (mr *MockLikeMockRecorder) IsMutualLike(ctx, like interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMutualLike", reflect.TypeOf((*MockLike)(nil).IsMutualLike), ctx, like)
}

// ResetDislike mocks base method.
func (m *MockLike) ResetDislike(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetDislike", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetDislike indicates an expected call of ResetDislike.
func (mr *MockLikeMockRecorder) ResetDislike(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetDislike", reflect.TypeOf((*MockLike)(nil).ResetDislike), ctx)
}

// MockDialog is a mock of Dialog interface.
type MockDialog struct {
	ctrl     *gomock.Controller
	recorder *MockDialogMockRecorder
}

// MockDialogMockRecorder is the mock recorder for MockDialog.
type MockDialogMockRecorder struct {
	mock *MockDialog
}

// NewMockDialog creates a new mock instance.
func NewMockDialog(ctrl *gomock.Controller) *MockDialog {
	mock := &MockDialog{ctrl: ctrl}
	mock.recorder = &MockDialogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialog) EXPECT() *MockDialogMockRecorder {
	return m.recorder
}

// CreateDialog mocks base method.
func (m *MockDialog) CreateDialog(ctx context.Context, dialog model.Dialog) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDialog", ctx, dialog)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDialog indicates an expected call of CreateDialog.
func (mr *MockDialogMockRecorder) CreateDialog(ctx, dialog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDialog", reflect.TypeOf((*MockDialog)(nil).CreateDialog), ctx, dialog)
}

// GetDialogById mocks base method.
func (m *MockDialog) GetDialogById(ctx context.Context, id int) (model.Dialog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialogById", ctx, id)
	ret0, _ := ret[0].(model.Dialog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogById indicates an expected call of GetDialogById.
func (mr *MockDialogMockRecorder) GetDialogById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogById", reflect.TypeOf((*MockDialog)(nil).GetDialogById), ctx, id)
}

// GetDialogs mocks base method.
func (m *MockDialog) GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialogs", ctx, userId)
	ret0, _ := ret[0].([]model.Dialog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogs indicates an expected call of GetDialogs.
func (mr *MockDialogMockRecorder) GetDialogs(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogs", reflect.TypeOf((*MockDialog)(nil).GetDialogs), ctx, userId)
}

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method.
func (m *MockMessage) CreateMessage(ctx context.Context, message model.Message) (model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", ctx, message)
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockMessageMockRecorder) CreateMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockMessage)(nil).CreateMessage), ctx, message)
}

// GetDialogMessages mocks base method.
func (m *MockMessage) GetDialogMessages(ctx context.Context, userId, recipientId int) ([]model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDialogMessages", ctx, userId, recipientId)
	ret0, _ := ret[0].([]model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDialogMessages indicates an expected call of GetDialogMessages.
func (mr *MockMessageMockRecorder) GetDialogMessages(ctx, userId, recipientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDialogMessages", reflect.TypeOf((*MockMessage)(nil).GetDialogMessages), ctx, userId, recipientId)
}

// UpdateMessage mocks base method.
func (m *MockMessage) UpdateMessage(ctx context.Context, message model.Message) (model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMessage", ctx, message)
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMessage indicates an expected call of UpdateMessage.
func (mr *MockMessageMockRecorder) UpdateMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMessage", reflect.TypeOf((*MockMessage)(nil).UpdateMessage), ctx, message)
}

// MockTag is a mock of Tag interface.
type MockTag struct {
	ctrl     *gomock.Controller
	recorder *MockTagMockRecorder
}

// MockTagMockRecorder is the mock recorder for MockTag.
type MockTagMockRecorder struct {
	mock *MockTag
}

// NewMockTag creates a new mock instance.
func NewMockTag(ctrl *gomock.Controller) *MockTag {
	mock := &MockTag{ctrl: ctrl}
	mock.recorder = &MockTagMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTag) EXPECT() *MockTagMockRecorder {
	return m.recorder
}

// GetAllTags mocks base method.
func (m *MockTag) GetAllTags(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTags", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTags indicates an expected call of GetAllTags.
func (mr *MockTagMockRecorder) GetAllTags(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTags", reflect.TypeOf((*MockTag)(nil).GetAllTags), ctx)
}

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// CreateFeedFeedback mocks base method.
func (m *MockAdmin) CreateFeedFeedback(ctx context.Context, rec model.Recommendation) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeedFeedback", ctx, rec)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeedFeedback indicates an expected call of CreateFeedFeedback.
func (mr *MockAdminMockRecorder) CreateFeedFeedback(ctx, rec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeedFeedback", reflect.TypeOf((*MockAdmin)(nil).CreateFeedFeedback), ctx, rec)
}

// CreateFeedback mocks base method.
func (m *MockAdmin) CreateFeedback(ctx context.Context, stat model.Feedback) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeedback", ctx, stat)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeedback indicates an expected call of CreateFeedback.
func (mr *MockAdminMockRecorder) CreateFeedback(ctx, stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeedback", reflect.TypeOf((*MockAdmin)(nil).CreateFeedback), ctx, stat)
}

// CreateRecommendation mocks base method.
func (m *MockAdmin) CreateRecommendation(ctx context.Context, rec model.Recommendation) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecommendation", ctx, rec)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecommendation indicates an expected call of CreateRecommendation.
func (mr *MockAdminMockRecorder) CreateRecommendation(ctx, rec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecommendation", reflect.TypeOf((*MockAdmin)(nil).CreateRecommendation), ctx, rec)
}

// GetAdmin mocks base method.
func (m *MockAdmin) GetAdmin(ctx context.Context, mail string) (model.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", ctx, mail)
	ret0, _ := ret[0].(model.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockAdminMockRecorder) GetAdmin(ctx, mail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockAdmin)(nil).GetAdmin), ctx, mail)
}

// GetFeedbacks mocks base method.
func (m *MockAdmin) GetFeedbacks(ctx context.Context) ([]model.Feedback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeedbacks", ctx)
	ret0, _ := ret[0].([]model.Feedback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeedbacks indicates an expected call of GetFeedbacks.
func (mr *MockAdminMockRecorder) GetFeedbacks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeedbacks", reflect.TypeOf((*MockAdmin)(nil).GetFeedbacks), ctx)
}

// GetRecommendations mocks base method.
func (m *MockAdmin) GetRecommendations(ctx context.Context) ([]model.Recommendation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendations", ctx)
	ret0, _ := ret[0].([]model.Recommendation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendations indicates an expected call of GetRecommendations.
func (mr *MockAdminMockRecorder) GetRecommendations(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendations", reflect.TypeOf((*MockAdmin)(nil).GetRecommendations), ctx)
}

// ShowFeedback mocks base method.
func (m *MockAdmin) ShowFeedback(ctx context.Context, userId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowFeedback", ctx, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowFeedback indicates an expected call of ShowFeedback.
func (mr *MockAdminMockRecorder) ShowFeedback(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowFeedback", reflect.TypeOf((*MockAdmin)(nil).ShowFeedback), ctx, userId)
}

// ShowRecommendation mocks base method.
func (m *MockAdmin) ShowRecommendation(ctx context.Context, userId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowRecommendation", ctx, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowRecommendation indicates an expected call of ShowRecommendation.
func (mr *MockAdminMockRecorder) ShowRecommendation(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowRecommendation", reflect.TypeOf((*MockAdmin)(nil).ShowRecommendation), ctx, userId)
}

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(ctx context.Context, SID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, SID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(ctx, SID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), ctx, SID)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(ctx context.Context, SID string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", ctx, SID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(ctx, SID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), ctx, SID)
}

// SetSession mocks base method.
func (m *MockStore) SetSession(ctx context.Context, SID string, id int, lifetime time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSession", ctx, SID, id, lifetime)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSession indicates an expected call of SetSession.
func (mr *MockStoreMockRecorder) SetSession(ctx, SID, id, lifetime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockStore)(nil).SetSession), ctx, SID, id, lifetime)
}

// MockFileServer is a mock of FileServer interface.
type MockFileServer struct {
	ctrl     *gomock.Controller
	recorder *MockFileServerMockRecorder
}

// MockFileServerMockRecorder is the mock recorder for MockFileServer.
type MockFileServerMockRecorder struct {
	mock *MockFileServer
}

// NewMockFileServer creates a new mock instance.
func NewMockFileServer(ctrl *gomock.Controller) *MockFileServer {
	mock := &MockFileServer{ctrl: ctrl}
	mock.recorder = &MockFileServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileServer) EXPECT() *MockFileServerMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method.
func (m *MockFileServer) CreateBucket(ctx context.Context, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", ctx, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockFileServerMockRecorder) CreateBucket(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockFileServer)(nil).CreateBucket), ctx, bucketName)
}

// DeleteFile mocks base method.
func (m *MockFileServer) DeleteFile(ctx context.Context, bucketName, link string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", ctx, bucketName, link)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockFileServerMockRecorder) DeleteFile(ctx, bucketName, link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileServer)(nil).DeleteFile), ctx, bucketName, link)
}

// UploadFile mocks base method.
func (m *MockFileServer) UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader, size int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", ctx, bucketName, fileName, contentType, file, size)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockFileServerMockRecorder) UploadFile(ctx, bucketName, fileName, contentType, file, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileServer)(nil).UploadFile), ctx, bucketName, fileName, contentType, file, size)
}

// MockComplaint is a mock of Complaint interface.
type MockComplaint struct {
	ctrl     *gomock.Controller
	recorder *MockComplaintMockRecorder
}

// MockComplaintMockRecorder is the mock recorder for MockComplaint.
type MockComplaintMockRecorder struct {
	mock *MockComplaint
}

// NewMockComplaint creates a new mock instance.
func NewMockComplaint(ctrl *gomock.Controller) *MockComplaint {
	mock := &MockComplaint{ctrl: ctrl}
	mock.recorder = &MockComplaintMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplaint) EXPECT() *MockComplaintMockRecorder {
	return m.recorder
}

// AcceptComplaint mocks base method.
func (m *MockComplaint) AcceptComplaint(ctx context.Context, complaintId int) (model.Complaint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptComplaint", ctx, complaintId)
	ret0, _ := ret[0].(model.Complaint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptComplaint indicates an expected call of AcceptComplaint.
func (mr *MockComplaintMockRecorder) AcceptComplaint(ctx, complaintId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptComplaint", reflect.TypeOf((*MockComplaint)(nil).AcceptComplaint), ctx, complaintId)
}

// CreateComplaint mocks base method.
func (m *MockComplaint) CreateComplaint(ctx context.Context, complaint model.Complaint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComplaint", ctx, complaint)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComplaint indicates an expected call of CreateComplaint.
func (mr *MockComplaintMockRecorder) CreateComplaint(ctx, complaint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComplaint", reflect.TypeOf((*MockComplaint)(nil).CreateComplaint), ctx, complaint)
}

// DeleteComplaint mocks base method.
func (m *MockComplaint) DeleteComplaint(ctx context.Context, complaintId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComplaint", ctx, complaintId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComplaint indicates an expected call of DeleteComplaint.
func (mr *MockComplaintMockRecorder) DeleteComplaint(ctx, complaintId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComplaint", reflect.TypeOf((*MockComplaint)(nil).DeleteComplaint), ctx, complaintId)
}

// GetComplaintTypes mocks base method.
func (m *MockComplaint) GetComplaintTypes(ctx context.Context) ([]model.ComplaintType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComplaintTypes", ctx)
	ret0, _ := ret[0].([]model.ComplaintType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComplaintTypes indicates an expected call of GetComplaintTypes.
func (mr *MockComplaintMockRecorder) GetComplaintTypes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComplaintTypes", reflect.TypeOf((*MockComplaint)(nil).GetComplaintTypes), ctx)
}

// GetNextComplaint mocks base method.
func (m *MockComplaint) GetNextComplaint(ctx context.Context) (model.Complaint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextComplaint", ctx)
	ret0, _ := ret[0].(model.Complaint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextComplaint indicates an expected call of GetNextComplaint.
func (mr *MockComplaintMockRecorder) GetNextComplaint(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextComplaint", reflect.TypeOf((*MockComplaint)(nil).GetNextComplaint), ctx)
}

// MockPgxPoolInterface is a mock of PgxPoolInterface interface.
type MockPgxPoolInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPgxPoolInterfaceMockRecorder
}

// MockPgxPoolInterfaceMockRecorder is the mock recorder for MockPgxPoolInterface.
type MockPgxPoolInterfaceMockRecorder struct {
	mock *MockPgxPoolInterface
}

// NewMockPgxPoolInterface creates a new mock instance.
func NewMockPgxPoolInterface(ctrl *gomock.Controller) *MockPgxPoolInterface {
	mock := &MockPgxPoolInterface{ctrl: ctrl}
	mock.recorder = &MockPgxPoolInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPgxPoolInterface) EXPECT() *MockPgxPoolInterfaceMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockPgxPoolInterface) Begin(arg0 context.Context) (pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockPgxPoolInterfaceMockRecorder) Begin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockPgxPoolInterface)(nil).Begin), arg0)
}

// Close mocks base method.
func (m *MockPgxPoolInterface) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockPgxPoolInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPgxPoolInterface)(nil).Close))
}

// Exec mocks base method.
func (m *MockPgxPoolInterface) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range arguments {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockPgxPoolInterfaceMockRecorder) Exec(ctx, sql interface{}, arguments ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, arguments...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockPgxPoolInterface)(nil).Exec), varargs...)
}

// Ping mocks base method.
func (m *MockPgxPoolInterface) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockPgxPoolInterfaceMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockPgxPoolInterface)(nil).Ping), ctx)
}

// Query mocks base method.
func (m *MockPgxPoolInterface) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockPgxPoolInterfaceMockRecorder) Query(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockPgxPoolInterface)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockPgxPoolInterface) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockPgxPoolInterfaceMockRecorder) QueryRow(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockPgxPoolInterface)(nil).QueryRow), varargs...)
}
