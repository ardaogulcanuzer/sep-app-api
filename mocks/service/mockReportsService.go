// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/berkaymuratt/sep-app-api/services (interfaces: ReportsServiceI)

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	dtos "github.com/berkaymuratt/sep-app-api/dtos"
	models "github.com/berkaymuratt/sep-app-api/models"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockReportsServiceI is a mock of ReportsServiceI interface.
type MockReportsServiceI struct {
	ctrl     *gomock.Controller
	recorder *MockReportsServiceIMockRecorder
}

// MockReportsServiceIMockRecorder is the mock recorder for MockReportsServiceI.
type MockReportsServiceIMockRecorder struct {
	mock *MockReportsServiceI
}

// NewMockReportsServiceI creates a new mock instance.
func NewMockReportsServiceI(ctrl *gomock.Controller) *MockReportsServiceI {
	mock := &MockReportsServiceI{ctrl: ctrl}
	mock.recorder = &MockReportsServiceIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReportsServiceI) EXPECT() *MockReportsServiceIMockRecorder {
	return m.recorder
}

// CreateReportByAppointment mocks base method.
func (m *MockReportsServiceI) CreateReportByAppointment(arg0 *models.Appointment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReportByAppointment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReportByAppointment indicates an expected call of CreateReportByAppointment.
func (mr *MockReportsServiceIMockRecorder) CreateReportByAppointment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReportByAppointment", reflect.TypeOf((*MockReportsServiceI)(nil).CreateReportByAppointment), arg0)
}

// DeleteReport mocks base method.
func (m *MockReportsServiceI) DeleteReport(arg0 primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReport", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReport indicates an expected call of DeleteReport.
func (mr *MockReportsServiceIMockRecorder) DeleteReport(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReport", reflect.TypeOf((*MockReportsServiceI)(nil).DeleteReport), arg0)
}

// GetReportById mocks base method.
func (m *MockReportsServiceI) GetReportById(arg0 primitive.ObjectID) (*dtos.ReportDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportById", arg0)
	ret0, _ := ret[0].(*dtos.ReportDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportById indicates an expected call of GetReportById.
func (mr *MockReportsServiceIMockRecorder) GetReportById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportById", reflect.TypeOf((*MockReportsServiceI)(nil).GetReportById), arg0)
}

// GetReports mocks base method.
func (m *MockReportsServiceI) GetReports(arg0, arg1 *primitive.ObjectID) ([]*dtos.ReportDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReports", arg0, arg1)
	ret0, _ := ret[0].([]*dtos.ReportDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReports indicates an expected call of GetReports.
func (mr *MockReportsServiceIMockRecorder) GetReports(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReports", reflect.TypeOf((*MockReportsServiceI)(nil).GetReports), arg0, arg1)
}
