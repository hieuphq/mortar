// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_server is a generated GoMock package.
package mock_server

import (
	context "context"
	server "github.com/go-masonry/mortar/interfaces/http/server"
	gomock "github.com/golang/mock/gomock"
	runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	grpc "google.golang.org/grpc"
	net "net"
	http "net/http"
	reflect "reflect"
)

// MockWebService is a mock of WebService interface.
type MockWebService struct {
	ctrl     *gomock.Controller
	recorder *MockWebServiceMockRecorder
}

// MockWebServiceMockRecorder is the mock recorder for MockWebService.
type MockWebServiceMockRecorder struct {
	mock *MockWebService
}

// NewMockWebService creates a new mock instance.
func NewMockWebService(ctrl *gomock.Controller) *MockWebService {
	mock := &MockWebService{ctrl: ctrl}
	mock.recorder = &MockWebServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebService) EXPECT() *MockWebServiceMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockWebService) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockWebServiceMockRecorder) Run(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWebService)(nil).Run), ctx)
}

// Stop mocks base method.
func (m *MockWebService) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockWebServiceMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWebService)(nil).Stop), ctx)
}

// Ports mocks base method.
func (m *MockWebService) Ports() []server.ListenInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ports")
	ret0, _ := ret[0].([]server.ListenInfo)
	return ret0
}

// Ports indicates an expected call of Ports.
func (mr *MockWebServiceMockRecorder) Ports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ports", reflect.TypeOf((*MockWebService)(nil).Ports))
}

// MockGRPCWebServiceBuilder is a mock of GRPCWebServiceBuilder interface.
type MockGRPCWebServiceBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockGRPCWebServiceBuilderMockRecorder
}

// MockGRPCWebServiceBuilderMockRecorder is the mock recorder for MockGRPCWebServiceBuilder.
type MockGRPCWebServiceBuilderMockRecorder struct {
	mock *MockGRPCWebServiceBuilder
}

// NewMockGRPCWebServiceBuilder creates a new mock instance.
func NewMockGRPCWebServiceBuilder(ctrl *gomock.Controller) *MockGRPCWebServiceBuilder {
	mock := &MockGRPCWebServiceBuilder{ctrl: ctrl}
	mock.recorder = &MockGRPCWebServiceBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGRPCWebServiceBuilder) EXPECT() *MockGRPCWebServiceBuilderMockRecorder {
	return m.recorder
}

// ListenOn mocks base method.
func (m *MockGRPCWebServiceBuilder) ListenOn(addr string) server.GRPCWebServiceBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenOn", addr)
	ret0, _ := ret[0].(server.GRPCWebServiceBuilder)
	return ret0
}

// ListenOn indicates an expected call of ListenOn.
func (mr *MockGRPCWebServiceBuilderMockRecorder) ListenOn(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenOn", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).ListenOn), addr)
}

// SetCustomGRPCServer mocks base method.
func (m *MockGRPCWebServiceBuilder) SetCustomGRPCServer(customServer *grpc.Server) server.GRPCWebServiceBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCustomGRPCServer", customServer)
	ret0, _ := ret[0].(server.GRPCWebServiceBuilder)
	return ret0
}

// SetCustomGRPCServer indicates an expected call of SetCustomGRPCServer.
func (mr *MockGRPCWebServiceBuilderMockRecorder) SetCustomGRPCServer(customServer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCustomGRPCServer", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).SetCustomGRPCServer), customServer)
}

// SetCustomListener mocks base method.
func (m *MockGRPCWebServiceBuilder) SetCustomListener(listener net.Listener) server.GRPCWebServiceBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCustomListener", listener)
	ret0, _ := ret[0].(server.GRPCWebServiceBuilder)
	return ret0
}

// SetCustomListener indicates an expected call of SetCustomListener.
func (mr *MockGRPCWebServiceBuilderMockRecorder) SetCustomListener(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCustomListener", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).SetCustomListener), listener)
}

// RegisterGRPCAPIs mocks base method.
func (m *MockGRPCWebServiceBuilder) RegisterGRPCAPIs(register ...server.GRPCServerAPI) server.GRPCWebServiceBuilder {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range register {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterGRPCAPIs", varargs...)
	ret0, _ := ret[0].(server.GRPCWebServiceBuilder)
	return ret0
}

// RegisterGRPCAPIs indicates an expected call of RegisterGRPCAPIs.
func (mr *MockGRPCWebServiceBuilderMockRecorder) RegisterGRPCAPIs(register ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCAPIs", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).RegisterGRPCAPIs), register...)
}

// AddGRPCServerOptions mocks base method.
func (m *MockGRPCWebServiceBuilder) AddGRPCServerOptions(options ...grpc.ServerOption) server.GRPCWebServiceBuilder {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddGRPCServerOptions", varargs...)
	ret0, _ := ret[0].(server.GRPCWebServiceBuilder)
	return ret0
}

// AddGRPCServerOptions indicates an expected call of AddGRPCServerOptions.
func (mr *MockGRPCWebServiceBuilderMockRecorder) AddGRPCServerOptions(options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGRPCServerOptions", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).AddGRPCServerOptions), options...)
}

// SetPanicHandler mocks base method.
func (m *MockGRPCWebServiceBuilder) SetPanicHandler(handler func(interface{}) error) server.GRPCWebServiceBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPanicHandler", handler)
	ret0, _ := ret[0].(server.GRPCWebServiceBuilder)
	return ret0
}

// SetPanicHandler indicates an expected call of SetPanicHandler.
func (mr *MockGRPCWebServiceBuilderMockRecorder) SetPanicHandler(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPanicHandler", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).SetPanicHandler), handler)
}

// SetLogger mocks base method.
func (m *MockGRPCWebServiceBuilder) SetLogger(logger func(context.Context, string, ...interface{})) server.GRPCWebServiceBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLogger", logger)
	ret0, _ := ret[0].(server.GRPCWebServiceBuilder)
	return ret0
}

// SetLogger indicates an expected call of SetLogger.
func (mr *MockGRPCWebServiceBuilderMockRecorder) SetLogger(logger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).SetLogger), logger)
}

// AddRESTServerConfiguration mocks base method.
func (m *MockGRPCWebServiceBuilder) AddRESTServerConfiguration() server.RESTBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRESTServerConfiguration")
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// AddRESTServerConfiguration indicates an expected call of AddRESTServerConfiguration.
func (mr *MockGRPCWebServiceBuilderMockRecorder) AddRESTServerConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRESTServerConfiguration", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).AddRESTServerConfiguration))
}

// Build mocks base method.
func (m *MockGRPCWebServiceBuilder) Build() (server.WebService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(server.WebService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockGRPCWebServiceBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockGRPCWebServiceBuilder)(nil).Build))
}

// MockRESTBuilder is a mock of RESTBuilder interface.
type MockRESTBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockRESTBuilderMockRecorder
}

// MockRESTBuilderMockRecorder is the mock recorder for MockRESTBuilder.
type MockRESTBuilderMockRecorder struct {
	mock *MockRESTBuilder
}

// NewMockRESTBuilder creates a new mock instance.
func NewMockRESTBuilder(ctrl *gomock.Controller) *MockRESTBuilder {
	mock := &MockRESTBuilder{ctrl: ctrl}
	mock.recorder = &MockRESTBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRESTBuilder) EXPECT() *MockRESTBuilderMockRecorder {
	return m.recorder
}

// ListenOn mocks base method.
func (m *MockRESTBuilder) ListenOn(addr string) server.RESTBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenOn", addr)
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// ListenOn indicates an expected call of ListenOn.
func (mr *MockRESTBuilderMockRecorder) ListenOn(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenOn", reflect.TypeOf((*MockRESTBuilder)(nil).ListenOn), addr)
}

// SetCustomServer mocks base method.
func (m *MockRESTBuilder) SetCustomServer(customServer *http.Server) server.RESTBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCustomServer", customServer)
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// SetCustomServer indicates an expected call of SetCustomServer.
func (mr *MockRESTBuilderMockRecorder) SetCustomServer(customServer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCustomServer", reflect.TypeOf((*MockRESTBuilder)(nil).SetCustomServer), customServer)
}

// SetCustomListener mocks base method.
func (m *MockRESTBuilder) SetCustomListener(listener net.Listener) server.RESTBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCustomListener", listener)
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// SetCustomListener indicates an expected call of SetCustomListener.
func (mr *MockRESTBuilderMockRecorder) SetCustomListener(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCustomListener", reflect.TypeOf((*MockRESTBuilder)(nil).SetCustomListener), listener)
}

// AddHandler mocks base method.
func (m *MockRESTBuilder) AddHandler(pattern string, handler http.Handler) server.RESTBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHandler", pattern, handler)
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// AddHandler indicates an expected call of AddHandler.
func (mr *MockRESTBuilderMockRecorder) AddHandler(pattern, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*MockRESTBuilder)(nil).AddHandler), pattern, handler)
}

// AddHandlerFunc mocks base method.
func (m *MockRESTBuilder) AddHandlerFunc(pattern string, handlerFunc http.HandlerFunc) server.RESTBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHandlerFunc", pattern, handlerFunc)
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// AddHandlerFunc indicates an expected call of AddHandlerFunc.
func (mr *MockRESTBuilderMockRecorder) AddHandlerFunc(pattern, handlerFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandlerFunc", reflect.TypeOf((*MockRESTBuilder)(nil).AddHandlerFunc), pattern, handlerFunc)
}

// SetCustomGRPCGatewayMux mocks base method.
func (m *MockRESTBuilder) SetCustomGRPCGatewayMux(mux *runtime.ServeMux) server.RESTBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCustomGRPCGatewayMux", mux)
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// SetCustomGRPCGatewayMux indicates an expected call of SetCustomGRPCGatewayMux.
func (mr *MockRESTBuilderMockRecorder) SetCustomGRPCGatewayMux(mux interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCustomGRPCGatewayMux", reflect.TypeOf((*MockRESTBuilder)(nil).SetCustomGRPCGatewayMux), mux)
}

// RegisterGRPCGatewayHandlers mocks base method.
func (m *MockRESTBuilder) RegisterGRPCGatewayHandlers(handlers ...server.GRPCGatewayGeneratedHandlers) server.RESTBuilder {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range handlers {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterGRPCGatewayHandlers", varargs...)
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// RegisterGRPCGatewayHandlers indicates an expected call of RegisterGRPCGatewayHandlers.
func (mr *MockRESTBuilderMockRecorder) RegisterGRPCGatewayHandlers(handlers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayHandlers", reflect.TypeOf((*MockRESTBuilder)(nil).RegisterGRPCGatewayHandlers), handlers...)
}

// AddGRPCGatewayOptions mocks base method.
func (m *MockRESTBuilder) AddGRPCGatewayOptions(options ...runtime.ServeMuxOption) server.RESTBuilder {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddGRPCGatewayOptions", varargs...)
	ret0, _ := ret[0].(server.RESTBuilder)
	return ret0
}

// AddGRPCGatewayOptions indicates an expected call of AddGRPCGatewayOptions.
func (mr *MockRESTBuilderMockRecorder) AddGRPCGatewayOptions(options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGRPCGatewayOptions", reflect.TypeOf((*MockRESTBuilder)(nil).AddGRPCGatewayOptions), options...)
}

// BuildRESTPart mocks base method.
func (m *MockRESTBuilder) BuildRESTPart() server.GRPCWebServiceBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildRESTPart")
	ret0, _ := ret[0].(server.GRPCWebServiceBuilder)
	return ret0
}

// BuildRESTPart indicates an expected call of BuildRESTPart.
func (mr *MockRESTBuilderMockRecorder) BuildRESTPart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildRESTPart", reflect.TypeOf((*MockRESTBuilder)(nil).BuildRESTPart))
}