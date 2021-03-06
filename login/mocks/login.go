// Automatically generated by MockGen. DO NOT EDIT!
// Source: login/interfaces/login.go

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/topfreegames/extensions/pg/interfaces"
	oauth2 "golang.org/x/oauth2"
)

// Mock of Login interface
type MockLogin struct {
	ctrl     *gomock.Controller
	recorder *_MockLoginRecorder
}

// Recorder for MockLogin (not exported)
type _MockLoginRecorder struct {
	mock *MockLogin
}

func NewMockLogin(ctrl *gomock.Controller) *MockLogin {
	mock := &MockLogin{ctrl: ctrl}
	mock.recorder = &_MockLoginRecorder{mock}
	return mock
}

func (_m *MockLogin) EXPECT() *_MockLoginRecorder {
	return _m.recorder
}

func (_m *MockLogin) Setup() {
	_m.ctrl.Call(_m, "Setup")
}

func (_mr *_MockLoginRecorder) Setup() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Setup")
}

func (_m *MockLogin) GenerateLoginURL(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GenerateLoginURL", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLoginRecorder) GenerateLoginURL(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateLoginURL", arg0)
}

func (_m *MockLogin) GetAccessToken(_param0, _param1 string) (*oauth2.Token, error) {
	ret := _m.ctrl.Call(_m, "GetAccessToken", _param0, _param1)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLoginRecorder) GetAccessToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccessToken", arg0, arg1)
}

func (_m *MockLogin) Authenticate(_param0 *oauth2.Token, _param1 interfaces.DB) (string, int, error) {
	ret := _m.ctrl.Call(_m, "Authenticate", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockLoginRecorder) Authenticate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Authenticate", arg0, arg1)
}
