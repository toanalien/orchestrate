// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	database "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/database"
	store "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store"
	models "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/store/models"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockBuilder) Build(ctx context.Context, name string, configuration interface{}) (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor, func(*grpc.Server), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", ctx, name, configuration)
	ret0, _ := ret[0].(grpc.UnaryServerInterceptor)
	ret1, _ := ret[1].(grpc.StreamServerInterceptor)
	ret2, _ := ret[2].(func(*grpc.Server))
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Build indicates an expected call of Build.
func (mr *MockBuilderMockRecorder) Build(ctx, name, configuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build), ctx, name, configuration)
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

// Connect mocks base method.
func (m *MockStore) Connect(ctx context.Context, conf interface{}) (store.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx, conf)
	ret0, _ := ret[0].(store.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockStoreMockRecorder) Connect(ctx, conf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockStore)(nil).Connect), ctx, conf)
}

// MockAgents is a mock of Agents interface.
type MockAgents struct {
	ctrl     *gomock.Controller
	recorder *MockAgentsMockRecorder
}

// MockAgentsMockRecorder is the mock recorder for MockAgents.
type MockAgentsMockRecorder struct {
	mock *MockAgents
}

// NewMockAgents creates a new mock instance.
func NewMockAgents(ctrl *gomock.Controller) *MockAgents {
	mock := &MockAgents{ctrl: ctrl}
	mock.recorder = &MockAgentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgents) EXPECT() *MockAgentsMockRecorder {
	return m.recorder
}

// Schedule mocks base method.
func (m *MockAgents) Schedule() store.ScheduleAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schedule")
	ret0, _ := ret[0].(store.ScheduleAgent)
	return ret0
}

// Schedule indicates an expected call of Schedule.
func (mr *MockAgentsMockRecorder) Schedule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockAgents)(nil).Schedule))
}

// Job mocks base method.
func (m *MockAgents) Job() store.JobAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Job")
	ret0, _ := ret[0].(store.JobAgent)
	return ret0
}

// Job indicates an expected call of Job.
func (mr *MockAgentsMockRecorder) Job() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Job", reflect.TypeOf((*MockAgents)(nil).Job))
}

// Log mocks base method.
func (m *MockAgents) Log() store.LogAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log")
	ret0, _ := ret[0].(store.LogAgent)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockAgentsMockRecorder) Log() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockAgents)(nil).Log))
}

// Transaction mocks base method.
func (m *MockAgents) Transaction() store.TransactionAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction")
	ret0, _ := ret[0].(store.TransactionAgent)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockAgentsMockRecorder) Transaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockAgents)(nil).Transaction))
}

// TransactionRequest mocks base method.
func (m *MockAgents) TransactionRequest() store.TransactionRequestAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionRequest")
	ret0, _ := ret[0].(store.TransactionRequestAgent)
	return ret0
}

// TransactionRequest indicates an expected call of TransactionRequest.
func (mr *MockAgentsMockRecorder) TransactionRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionRequest", reflect.TypeOf((*MockAgents)(nil).TransactionRequest))
}

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockDB) Begin() (database.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(database.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockDBMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockDB)(nil).Begin))
}

// Schedule mocks base method.
func (m *MockDB) Schedule() store.ScheduleAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schedule")
	ret0, _ := ret[0].(store.ScheduleAgent)
	return ret0
}

// Schedule indicates an expected call of Schedule.
func (mr *MockDBMockRecorder) Schedule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockDB)(nil).Schedule))
}

// Job mocks base method.
func (m *MockDB) Job() store.JobAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Job")
	ret0, _ := ret[0].(store.JobAgent)
	return ret0
}

// Job indicates an expected call of Job.
func (mr *MockDBMockRecorder) Job() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Job", reflect.TypeOf((*MockDB)(nil).Job))
}

// Log mocks base method.
func (m *MockDB) Log() store.LogAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log")
	ret0, _ := ret[0].(store.LogAgent)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockDBMockRecorder) Log() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockDB)(nil).Log))
}

// Transaction mocks base method.
func (m *MockDB) Transaction() store.TransactionAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction")
	ret0, _ := ret[0].(store.TransactionAgent)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockDBMockRecorder) Transaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockDB)(nil).Transaction))
}

// TransactionRequest mocks base method.
func (m *MockDB) TransactionRequest() store.TransactionRequestAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionRequest")
	ret0, _ := ret[0].(store.TransactionRequestAgent)
	return ret0
}

// TransactionRequest indicates an expected call of TransactionRequest.
func (mr *MockDBMockRecorder) TransactionRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionRequest", reflect.TypeOf((*MockDB)(nil).TransactionRequest))
}

// MockTx is a mock of Tx interface.
type MockTx struct {
	ctrl     *gomock.Controller
	recorder *MockTxMockRecorder
}

// MockTxMockRecorder is the mock recorder for MockTx.
type MockTxMockRecorder struct {
	mock *MockTx
}

// NewMockTx creates a new mock instance.
func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{ctrl: ctrl}
	mock.recorder = &MockTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTx) EXPECT() *MockTxMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockTx) Begin() (database.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(database.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockTxMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockTx)(nil).Begin))
}

// Commit mocks base method.
func (m *MockTx) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTxMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTx)(nil).Commit))
}

// Rollback mocks base method.
func (m *MockTx) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTxMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTx)(nil).Rollback))
}

// Close mocks base method.
func (m *MockTx) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTxMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTx)(nil).Close))
}

// Schedule mocks base method.
func (m *MockTx) Schedule() store.ScheduleAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schedule")
	ret0, _ := ret[0].(store.ScheduleAgent)
	return ret0
}

// Schedule indicates an expected call of Schedule.
func (mr *MockTxMockRecorder) Schedule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockTx)(nil).Schedule))
}

// Job mocks base method.
func (m *MockTx) Job() store.JobAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Job")
	ret0, _ := ret[0].(store.JobAgent)
	return ret0
}

// Job indicates an expected call of Job.
func (mr *MockTxMockRecorder) Job() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Job", reflect.TypeOf((*MockTx)(nil).Job))
}

// Log mocks base method.
func (m *MockTx) Log() store.LogAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log")
	ret0, _ := ret[0].(store.LogAgent)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockTxMockRecorder) Log() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockTx)(nil).Log))
}

// Transaction mocks base method.
func (m *MockTx) Transaction() store.TransactionAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction")
	ret0, _ := ret[0].(store.TransactionAgent)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockTxMockRecorder) Transaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockTx)(nil).Transaction))
}

// TransactionRequest mocks base method.
func (m *MockTx) TransactionRequest() store.TransactionRequestAgent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionRequest")
	ret0, _ := ret[0].(store.TransactionRequestAgent)
	return ret0
}

// TransactionRequest indicates an expected call of TransactionRequest.
func (mr *MockTxMockRecorder) TransactionRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionRequest", reflect.TypeOf((*MockTx)(nil).TransactionRequest))
}

// MockTransactionRequestAgent is a mock of TransactionRequestAgent interface.
type MockTransactionRequestAgent struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRequestAgentMockRecorder
}

// MockTransactionRequestAgentMockRecorder is the mock recorder for MockTransactionRequestAgent.
type MockTransactionRequestAgentMockRecorder struct {
	mock *MockTransactionRequestAgent
}

// NewMockTransactionRequestAgent creates a new mock instance.
func NewMockTransactionRequestAgent(ctrl *gomock.Controller) *MockTransactionRequestAgent {
	mock := &MockTransactionRequestAgent{ctrl: ctrl}
	mock.recorder = &MockTransactionRequestAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRequestAgent) EXPECT() *MockTransactionRequestAgentMockRecorder {
	return m.recorder
}

// SelectOrInsert mocks base method.
func (m *MockTransactionRequestAgent) SelectOrInsert(ctx context.Context, txRequest *models.TransactionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectOrInsert", ctx, txRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectOrInsert indicates an expected call of SelectOrInsert.
func (mr *MockTransactionRequestAgentMockRecorder) SelectOrInsert(ctx, txRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOrInsert", reflect.TypeOf((*MockTransactionRequestAgent)(nil).SelectOrInsert), ctx, txRequest)
}

// FindOneByIdempotencyKey mocks base method.
func (m *MockTransactionRequestAgent) FindOneByIdempotencyKey(ctx context.Context, idempotencyKey string) (*models.TransactionRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByIdempotencyKey", ctx, idempotencyKey)
	ret0, _ := ret[0].(*models.TransactionRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByIdempotencyKey indicates an expected call of FindOneByIdempotencyKey.
func (mr *MockTransactionRequestAgentMockRecorder) FindOneByIdempotencyKey(ctx, idempotencyKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByIdempotencyKey", reflect.TypeOf((*MockTransactionRequestAgent)(nil).FindOneByIdempotencyKey), ctx, idempotencyKey)
}

// MockScheduleAgent is a mock of ScheduleAgent interface.
type MockScheduleAgent struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleAgentMockRecorder
}

// MockScheduleAgentMockRecorder is the mock recorder for MockScheduleAgent.
type MockScheduleAgentMockRecorder struct {
	mock *MockScheduleAgent
}

// NewMockScheduleAgent creates a new mock instance.
func NewMockScheduleAgent(ctrl *gomock.Controller) *MockScheduleAgent {
	mock := &MockScheduleAgent{ctrl: ctrl}
	mock.recorder = &MockScheduleAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduleAgent) EXPECT() *MockScheduleAgentMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockScheduleAgent) Insert(ctx context.Context, schedule *models.Schedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, schedule)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockScheduleAgentMockRecorder) Insert(ctx, schedule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockScheduleAgent)(nil).Insert), ctx, schedule)
}

// FindOneByID mocks base method.
func (m *MockScheduleAgent) FindOneByID(ctx context.Context, ID int) (*models.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByID", ctx, ID)
	ret0, _ := ret[0].(*models.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByID indicates an expected call of FindOneByID.
func (mr *MockScheduleAgentMockRecorder) FindOneByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByID", reflect.TypeOf((*MockScheduleAgent)(nil).FindOneByID), ctx, ID)
}

// FindOneByUUID mocks base method.
func (m *MockScheduleAgent) FindOneByUUID(ctx context.Context, uuid, tenantID string) (*models.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUUID", ctx, uuid, tenantID)
	ret0, _ := ret[0].(*models.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUUID indicates an expected call of FindOneByUUID.
func (mr *MockScheduleAgentMockRecorder) FindOneByUUID(ctx, uuid, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUUID", reflect.TypeOf((*MockScheduleAgent)(nil).FindOneByUUID), ctx, uuid, tenantID)
}

// FindAll mocks base method.
func (m *MockScheduleAgent) FindAll(ctx context.Context, tenantID string) ([]*models.Schedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, tenantID)
	ret0, _ := ret[0].([]*models.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockScheduleAgentMockRecorder) FindAll(ctx, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockScheduleAgent)(nil).FindAll), ctx, tenantID)
}

// MockJobAgent is a mock of JobAgent interface.
type MockJobAgent struct {
	ctrl     *gomock.Controller
	recorder *MockJobAgentMockRecorder
}

// MockJobAgentMockRecorder is the mock recorder for MockJobAgent.
type MockJobAgentMockRecorder struct {
	mock *MockJobAgent
}

// NewMockJobAgent creates a new mock instance.
func NewMockJobAgent(ctrl *gomock.Controller) *MockJobAgent {
	mock := &MockJobAgent{ctrl: ctrl}
	mock.recorder = &MockJobAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobAgent) EXPECT() *MockJobAgentMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockJobAgent) Insert(ctx context.Context, job *models.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockJobAgentMockRecorder) Insert(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockJobAgent)(nil).Insert), ctx, job)
}

// Update mocks base method.
func (m *MockJobAgent) Update(ctx context.Context, job *models.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockJobAgentMockRecorder) Update(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockJobAgent)(nil).Update), ctx, job)
}

// FindOneByUUID mocks base method.
func (m *MockJobAgent) FindOneByUUID(ctx context.Context, uuid, tenantID string) (*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUUID", ctx, uuid, tenantID)
	ret0, _ := ret[0].(*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUUID indicates an expected call of FindOneByUUID.
func (mr *MockJobAgentMockRecorder) FindOneByUUID(ctx, uuid, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUUID", reflect.TypeOf((*MockJobAgent)(nil).FindOneByUUID), ctx, uuid, tenantID)
}

// Search mocks base method.
func (m *MockJobAgent) Search(ctx context.Context, filters map[string]string, tenantID string) ([]*models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, filters, tenantID)
	ret0, _ := ret[0].([]*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockJobAgentMockRecorder) Search(ctx, filters, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockJobAgent)(nil).Search), ctx, filters, tenantID)
}

// MockLogAgent is a mock of LogAgent interface.
type MockLogAgent struct {
	ctrl     *gomock.Controller
	recorder *MockLogAgentMockRecorder
}

// MockLogAgentMockRecorder is the mock recorder for MockLogAgent.
type MockLogAgentMockRecorder struct {
	mock *MockLogAgent
}

// NewMockLogAgent creates a new mock instance.
func NewMockLogAgent(ctrl *gomock.Controller) *MockLogAgent {
	mock := &MockLogAgent{ctrl: ctrl}
	mock.recorder = &MockLogAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogAgent) EXPECT() *MockLogAgentMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockLogAgent) Insert(ctx context.Context, log *models.Log) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, log)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockLogAgentMockRecorder) Insert(ctx, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockLogAgent)(nil).Insert), ctx, log)
}

// MockTransactionAgent is a mock of TransactionAgent interface.
type MockTransactionAgent struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionAgentMockRecorder
}

// MockTransactionAgentMockRecorder is the mock recorder for MockTransactionAgent.
type MockTransactionAgentMockRecorder struct {
	mock *MockTransactionAgent
}

// NewMockTransactionAgent creates a new mock instance.
func NewMockTransactionAgent(ctrl *gomock.Controller) *MockTransactionAgent {
	mock := &MockTransactionAgent{ctrl: ctrl}
	mock.recorder = &MockTransactionAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionAgent) EXPECT() *MockTransactionAgentMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockTransactionAgent) Insert(ctx context.Context, tx *models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockTransactionAgentMockRecorder) Insert(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTransactionAgent)(nil).Insert), ctx, tx)
}

// Update mocks base method.
func (m *MockTransactionAgent) Update(ctx context.Context, tx *models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTransactionAgentMockRecorder) Update(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTransactionAgent)(nil).Update), ctx, tx)
}