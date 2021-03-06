// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/flaki-service/pkg/flaki (interfaces: Sentry)

// Package mock is a generated GoMock package.
package mock

import (
	raven_go "github.com/getsentry/raven-go"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Sentry is a mock of Sentry interface
type Sentry struct {
	ctrl     *gomock.Controller
	recorder *SentryMockRecorder
}

// SentryMockRecorder is the mock recorder for Sentry
type SentryMockRecorder struct {
	mock *Sentry
}

// NewSentry creates a new mock instance
func NewSentry(ctrl *gomock.Controller) *Sentry {
	mock := &Sentry{ctrl: ctrl}
	mock.recorder = &SentryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Sentry) EXPECT() *SentryMockRecorder {
	return m.recorder
}

// CaptureError mocks base method
func (m *Sentry) CaptureError(arg0 error, arg1 map[string]string, arg2 ...raven_go.Interface) string {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CaptureError", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// CaptureError indicates an expected call of CaptureError
func (mr *SentryMockRecorder) CaptureError(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureError", reflect.TypeOf((*Sentry)(nil).CaptureError), varargs...)
}
