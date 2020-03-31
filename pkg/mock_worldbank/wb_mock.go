// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eloyekunle/world-bank-grpc/pkg/worldbank (interfaces: WorldBankClient,WorldBank_ListRegionsClient)

// Package mock_worldbank is a generated GoMock package.
package mock_worldbank

import (
	context "context"
	worldbank "github.com/eloyekunle/world-bank-grpc/pkg/worldbank"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockWorldBankClient is a mock of WorldBankClient interface
type MockWorldBankClient struct {
	ctrl     *gomock.Controller
	recorder *MockWorldBankClientMockRecorder
}

// MockWorldBankClientMockRecorder is the mock recorder for MockWorldBankClient
type MockWorldBankClientMockRecorder struct {
	mock *MockWorldBankClient
}

// NewMockWorldBankClient creates a new mock instance
func NewMockWorldBankClient(ctrl *gomock.Controller) *MockWorldBankClient {
	mock := &MockWorldBankClient{ctrl: ctrl}
	mock.recorder = &MockWorldBankClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorldBankClient) EXPECT() *MockWorldBankClientMockRecorder {
	return m.recorder
}

// GetCountry mocks base method
func (m *MockWorldBankClient) GetCountry(arg0 context.Context, arg1 *worldbank.CountryID, arg2 ...grpc.CallOption) (*worldbank.Country, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCountry", varargs...)
	ret0, _ := ret[0].(*worldbank.Country)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountry indicates an expected call of GetCountry
func (mr *MockWorldBankClientMockRecorder) GetCountry(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountry", reflect.TypeOf((*MockWorldBankClient)(nil).GetCountry), varargs...)
}

// ListCountries mocks base method
func (m *MockWorldBankClient) ListCountries(arg0 context.Context, arg1 *worldbank.CountryFilter, arg2 ...grpc.CallOption) (worldbank.WorldBank_ListCountriesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCountries", varargs...)
	ret0, _ := ret[0].(worldbank.WorldBank_ListCountriesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCountries indicates an expected call of ListCountries
func (mr *MockWorldBankClientMockRecorder) ListCountries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCountries", reflect.TypeOf((*MockWorldBankClient)(nil).ListCountries), varargs...)
}

// ListIncomeLevels mocks base method
func (m *MockWorldBankClient) ListIncomeLevels(arg0 context.Context, arg1 *worldbank.Void, arg2 ...grpc.CallOption) (worldbank.WorldBank_ListIncomeLevelsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIncomeLevels", varargs...)
	ret0, _ := ret[0].(worldbank.WorldBank_ListIncomeLevelsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIncomeLevels indicates an expected call of ListIncomeLevels
func (mr *MockWorldBankClientMockRecorder) ListIncomeLevels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncomeLevels", reflect.TypeOf((*MockWorldBankClient)(nil).ListIncomeLevels), varargs...)
}

// ListLendingTypes mocks base method
func (m *MockWorldBankClient) ListLendingTypes(arg0 context.Context, arg1 *worldbank.Void, arg2 ...grpc.CallOption) (worldbank.WorldBank_ListLendingTypesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLendingTypes", varargs...)
	ret0, _ := ret[0].(worldbank.WorldBank_ListLendingTypesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLendingTypes indicates an expected call of ListLendingTypes
func (mr *MockWorldBankClientMockRecorder) ListLendingTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLendingTypes", reflect.TypeOf((*MockWorldBankClient)(nil).ListLendingTypes), varargs...)
}

// ListRegions mocks base method
func (m *MockWorldBankClient) ListRegions(arg0 context.Context, arg1 *worldbank.Void, arg2 ...grpc.CallOption) (worldbank.WorldBank_ListRegionsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegions", varargs...)
	ret0, _ := ret[0].(worldbank.WorldBank_ListRegionsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRegions indicates an expected call of ListRegions
func (mr *MockWorldBankClientMockRecorder) ListRegions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegions", reflect.TypeOf((*MockWorldBankClient)(nil).ListRegions), varargs...)
}

// MockWorldBank_ListRegionsClient is a mock of WorldBank_ListRegionsClient interface
type MockWorldBank_ListRegionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockWorldBank_ListRegionsClientMockRecorder
}

// MockWorldBank_ListRegionsClientMockRecorder is the mock recorder for MockWorldBank_ListRegionsClient
type MockWorldBank_ListRegionsClientMockRecorder struct {
	mock *MockWorldBank_ListRegionsClient
}

// NewMockWorldBank_ListRegionsClient creates a new mock instance
func NewMockWorldBank_ListRegionsClient(ctrl *gomock.Controller) *MockWorldBank_ListRegionsClient {
	mock := &MockWorldBank_ListRegionsClient{ctrl: ctrl}
	mock.recorder = &MockWorldBank_ListRegionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorldBank_ListRegionsClient) EXPECT() *MockWorldBank_ListRegionsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockWorldBank_ListRegionsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockWorldBank_ListRegionsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockWorldBank_ListRegionsClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockWorldBank_ListRegionsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockWorldBank_ListRegionsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockWorldBank_ListRegionsClient)(nil).Context))
}

// Header mocks base method
func (m *MockWorldBank_ListRegionsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockWorldBank_ListRegionsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockWorldBank_ListRegionsClient)(nil).Header))
}

// Recv mocks base method
func (m *MockWorldBank_ListRegionsClient) Recv() (*worldbank.Region, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*worldbank.Region)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockWorldBank_ListRegionsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockWorldBank_ListRegionsClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockWorldBank_ListRegionsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockWorldBank_ListRegionsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockWorldBank_ListRegionsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockWorldBank_ListRegionsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockWorldBank_ListRegionsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockWorldBank_ListRegionsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockWorldBank_ListRegionsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockWorldBank_ListRegionsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockWorldBank_ListRegionsClient)(nil).Trailer))
}
