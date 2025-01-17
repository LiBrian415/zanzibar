// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/zanzibar/examples/example-gateway/build/clients/baz (interfaces: Client)

// Package clientmock is a generated GoMock package.
package clientmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	base "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/baz/base"
	baz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/baz/baz"
	zanzibar "github.com/uber/zanzibar/runtime"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockClient) Call(arg0 context.Context, arg1 map[string]string, arg2 *baz.SimpleService_Call_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Call indicates an expected call of Call.
func (mr *MockClientMockRecorder) Call(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockClient)(nil).Call), arg0, arg1, arg2, arg3)
}

// Compare mocks base method.
func (m *MockClient) Compare(arg0 context.Context, arg1 map[string]string, arg2 *baz.SimpleService_Compare_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, *base.BazResponse, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*base.BazResponse)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Compare indicates an expected call of Compare.
func (mr *MockClientMockRecorder) Compare(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockClient)(nil).Compare), arg0, arg1, arg2, arg3)
}

// DeliberateDiffNoop mocks base method.
func (m *MockClient) DeliberateDiffNoop(arg0 context.Context, arg1 map[string]string, arg2 *zanzibar.TimeoutAndRetryOptions) (context.Context, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliberateDiffNoop", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeliberateDiffNoop indicates an expected call of DeliberateDiffNoop.
func (mr *MockClientMockRecorder) DeliberateDiffNoop(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliberateDiffNoop", reflect.TypeOf((*MockClient)(nil).DeliberateDiffNoop), arg0, arg1, arg2)
}

// EchoBinary mocks base method.
func (m *MockClient) EchoBinary(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoBinary_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, []byte, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoBinary", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoBinary indicates an expected call of EchoBinary.
func (mr *MockClientMockRecorder) EchoBinary(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoBinary", reflect.TypeOf((*MockClient)(nil).EchoBinary), arg0, arg1, arg2, arg3)
}

// EchoBool mocks base method.
func (m *MockClient) EchoBool(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoBool_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, bool, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoBool", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoBool indicates an expected call of EchoBool.
func (mr *MockClientMockRecorder) EchoBool(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoBool", reflect.TypeOf((*MockClient)(nil).EchoBool), arg0, arg1, arg2, arg3)
}

// EchoDouble mocks base method.
func (m *MockClient) EchoDouble(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoDouble_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, float64, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoDouble", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(float64)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoDouble indicates an expected call of EchoDouble.
func (mr *MockClientMockRecorder) EchoDouble(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoDouble", reflect.TypeOf((*MockClient)(nil).EchoDouble), arg0, arg1, arg2, arg3)
}

// EchoEnum mocks base method.
func (m *MockClient) EchoEnum(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoEnum_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, baz.Fruit, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoEnum", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(baz.Fruit)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoEnum indicates an expected call of EchoEnum.
func (mr *MockClientMockRecorder) EchoEnum(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoEnum", reflect.TypeOf((*MockClient)(nil).EchoEnum), arg0, arg1, arg2, arg3)
}

// EchoI16 mocks base method.
func (m *MockClient) EchoI16(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoI16_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, int16, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoI16", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(int16)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoI16 indicates an expected call of EchoI16.
func (mr *MockClientMockRecorder) EchoI16(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI16", reflect.TypeOf((*MockClient)(nil).EchoI16), arg0, arg1, arg2, arg3)
}

// EchoI32 mocks base method.
func (m *MockClient) EchoI32(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoI32_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, int32, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoI32", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoI32 indicates an expected call of EchoI32.
func (mr *MockClientMockRecorder) EchoI32(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI32", reflect.TypeOf((*MockClient)(nil).EchoI32), arg0, arg1, arg2, arg3)
}

// EchoI64 mocks base method.
func (m *MockClient) EchoI64(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoI64_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, int64, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoI64", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoI64 indicates an expected call of EchoI64.
func (mr *MockClientMockRecorder) EchoI64(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI64", reflect.TypeOf((*MockClient)(nil).EchoI64), arg0, arg1, arg2, arg3)
}

// EchoI8 mocks base method.
func (m *MockClient) EchoI8(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoI8_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, int8, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoI8", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(int8)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoI8 indicates an expected call of EchoI8.
func (mr *MockClientMockRecorder) EchoI8(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoI8", reflect.TypeOf((*MockClient)(nil).EchoI8), arg0, arg1, arg2, arg3)
}

// EchoString mocks base method.
func (m *MockClient) EchoString(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoString_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, string, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoString", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoString indicates an expected call of EchoString.
func (mr *MockClientMockRecorder) EchoString(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoString", reflect.TypeOf((*MockClient)(nil).EchoString), arg0, arg1, arg2, arg3)
}

// EchoStringList mocks base method.
func (m *MockClient) EchoStringList(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoStringList_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, []string, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoStringList", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoStringList indicates an expected call of EchoStringList.
func (mr *MockClientMockRecorder) EchoStringList(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStringList", reflect.TypeOf((*MockClient)(nil).EchoStringList), arg0, arg1, arg2, arg3)
}

// EchoStringMap mocks base method.
func (m *MockClient) EchoStringMap(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoStringMap_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, map[string]*base.BazResponse, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoStringMap", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(map[string]*base.BazResponse)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoStringMap indicates an expected call of EchoStringMap.
func (mr *MockClientMockRecorder) EchoStringMap(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStringMap", reflect.TypeOf((*MockClient)(nil).EchoStringMap), arg0, arg1, arg2, arg3)
}

// EchoStringSet mocks base method.
func (m *MockClient) EchoStringSet(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoStringSet_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, map[string]struct{}, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoStringSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(map[string]struct{})
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoStringSet indicates an expected call of EchoStringSet.
func (mr *MockClientMockRecorder) EchoStringSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStringSet", reflect.TypeOf((*MockClient)(nil).EchoStringSet), arg0, arg1, arg2, arg3)
}

// EchoStructList mocks base method.
func (m *MockClient) EchoStructList(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoStructList_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, []*base.BazResponse, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoStructList", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].([]*base.BazResponse)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoStructList indicates an expected call of EchoStructList.
func (mr *MockClientMockRecorder) EchoStructList(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStructList", reflect.TypeOf((*MockClient)(nil).EchoStructList), arg0, arg1, arg2, arg3)
}

// EchoStructSet mocks base method.
func (m *MockClient) EchoStructSet(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoStructSet_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, []*base.BazResponse, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoStructSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].([]*base.BazResponse)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoStructSet indicates an expected call of EchoStructSet.
func (mr *MockClientMockRecorder) EchoStructSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoStructSet", reflect.TypeOf((*MockClient)(nil).EchoStructSet), arg0, arg1, arg2, arg3)
}

// EchoTypedef mocks base method.
func (m *MockClient) EchoTypedef(arg0 context.Context, arg1 map[string]string, arg2 *baz.SecondService_EchoTypedef_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, base.UUID, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoTypedef", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(base.UUID)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EchoTypedef indicates an expected call of EchoTypedef.
func (mr *MockClientMockRecorder) EchoTypedef(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoTypedef", reflect.TypeOf((*MockClient)(nil).EchoTypedef), arg0, arg1, arg2, arg3)
}

// GetProfile mocks base method.
func (m *MockClient) GetProfile(arg0 context.Context, arg1 map[string]string, arg2 *baz.SimpleService_GetProfile_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, *baz.GetProfileResponse, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*baz.GetProfileResponse)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockClientMockRecorder) GetProfile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockClient)(nil).GetProfile), arg0, arg1, arg2, arg3)
}

// HeaderSchema mocks base method.
func (m *MockClient) HeaderSchema(arg0 context.Context, arg1 map[string]string, arg2 *baz.SimpleService_HeaderSchema_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, *baz.HeaderSchema, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderSchema", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*baz.HeaderSchema)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// HeaderSchema indicates an expected call of HeaderSchema.
func (mr *MockClientMockRecorder) HeaderSchema(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderSchema", reflect.TypeOf((*MockClient)(nil).HeaderSchema), arg0, arg1, arg2, arg3)
}

// Ping mocks base method.
func (m *MockClient) Ping(arg0 context.Context, arg1 map[string]string, arg2 *zanzibar.TimeoutAndRetryOptions) (context.Context, *base.BazResponse, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*base.BazResponse)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Ping indicates an expected call of Ping.
func (mr *MockClientMockRecorder) Ping(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockClient)(nil).Ping), arg0, arg1, arg2)
}

// TestUUID mocks base method.
func (m *MockClient) TestUUID(arg0 context.Context, arg1 map[string]string, arg2 *zanzibar.TimeoutAndRetryOptions) (context.Context, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestUUID", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TestUUID indicates an expected call of TestUUID.
func (mr *MockClientMockRecorder) TestUUID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestUUID", reflect.TypeOf((*MockClient)(nil).TestUUID), arg0, arg1, arg2)
}

// Trans mocks base method.
func (m *MockClient) Trans(arg0 context.Context, arg1 map[string]string, arg2 *baz.SimpleService_Trans_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, *base.TransStruct, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trans", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*base.TransStruct)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Trans indicates an expected call of Trans.
func (mr *MockClientMockRecorder) Trans(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trans", reflect.TypeOf((*MockClient)(nil).Trans), arg0, arg1, arg2, arg3)
}

// TransHeaders mocks base method.
func (m *MockClient) TransHeaders(arg0 context.Context, arg1 map[string]string, arg2 *baz.SimpleService_TransHeaders_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, *base.TransHeaders, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransHeaders", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*base.TransHeaders)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// TransHeaders indicates an expected call of TransHeaders.
func (mr *MockClientMockRecorder) TransHeaders(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransHeaders", reflect.TypeOf((*MockClient)(nil).TransHeaders), arg0, arg1, arg2, arg3)
}

// TransHeadersNoReq mocks base method.
func (m *MockClient) TransHeadersNoReq(arg0 context.Context, arg1 map[string]string, arg2 *baz.SimpleService_TransHeadersNoReq_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, *base.TransHeaders, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransHeadersNoReq", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*base.TransHeaders)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// TransHeadersNoReq indicates an expected call of TransHeadersNoReq.
func (mr *MockClientMockRecorder) TransHeadersNoReq(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransHeadersNoReq", reflect.TypeOf((*MockClient)(nil).TransHeadersNoReq), arg0, arg1, arg2, arg3)
}

// TransHeadersType mocks base method.
func (m *MockClient) TransHeadersType(arg0 context.Context, arg1 map[string]string, arg2 *baz.SimpleService_TransHeadersType_Args, arg3 *zanzibar.TimeoutAndRetryOptions) (context.Context, *baz.TransHeaderType, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransHeadersType", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*baz.TransHeaderType)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// TransHeadersType indicates an expected call of TransHeadersType.
func (mr *MockClientMockRecorder) TransHeadersType(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransHeadersType", reflect.TypeOf((*MockClient)(nil).TransHeadersType), arg0, arg1, arg2, arg3)
}

// URLTest mocks base method.
func (m *MockClient) URLTest(arg0 context.Context, arg1 map[string]string, arg2 *zanzibar.TimeoutAndRetryOptions) (context.Context, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URLTest", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// URLTest indicates an expected call of URLTest.
func (mr *MockClientMockRecorder) URLTest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URLTest", reflect.TypeOf((*MockClient)(nil).URLTest), arg0, arg1, arg2)
}
