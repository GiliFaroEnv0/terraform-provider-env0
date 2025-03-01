// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/env0/terraform-provider-env0/client (interfaces: ApiClientInterface)

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockApiClientInterface is a mock of ApiClientInterface interface.
type MockApiClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockApiClientInterfaceMockRecorder
}

// MockApiClientInterfaceMockRecorder is the mock recorder for MockApiClientInterface.
type MockApiClientInterfaceMockRecorder struct {
	mock *MockApiClientInterface
}

// NewMockApiClientInterface creates a new mock instance.
func NewMockApiClientInterface(ctrl *gomock.Controller) *MockApiClientInterface {
	mock := &MockApiClientInterface{ctrl: ctrl}
	mock.recorder = &MockApiClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiClientInterface) EXPECT() *MockApiClientInterfaceMockRecorder {
	return m.recorder
}

// AssignCloudCredentialsToProject mocks base method.
func (m *MockApiClientInterface) AssignCloudCredentialsToProject(arg0, arg1 string) (CloudCredentialsProjectAssignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignCloudCredentialsToProject", arg0, arg1)
	ret0, _ := ret[0].(CloudCredentialsProjectAssignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignCloudCredentialsToProject indicates an expected call of AssignCloudCredentialsToProject.
func (mr *MockApiClientInterfaceMockRecorder) AssignCloudCredentialsToProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignCloudCredentialsToProject", reflect.TypeOf((*MockApiClientInterface)(nil).AssignCloudCredentialsToProject), arg0, arg1)
}

// AssignTemplateToProject mocks base method.
func (m *MockApiClientInterface) AssignTemplateToProject(arg0 string, arg1 TemplateAssignmentToProjectPayload) (Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignTemplateToProject", arg0, arg1)
	ret0, _ := ret[0].(Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignTemplateToProject indicates an expected call of AssignTemplateToProject.
func (mr *MockApiClientInterfaceMockRecorder) AssignTemplateToProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignTemplateToProject", reflect.TypeOf((*MockApiClientInterface)(nil).AssignTemplateToProject), arg0, arg1)
}

// AwsCredentials mocks base method.
func (m *MockApiClientInterface) AwsCredentials(arg0 string) (ApiKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwsCredentials", arg0)
	ret0, _ := ret[0].(ApiKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AwsCredentials indicates an expected call of AwsCredentials.
func (mr *MockApiClientInterfaceMockRecorder) AwsCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwsCredentials", reflect.TypeOf((*MockApiClientInterface)(nil).AwsCredentials), arg0)
}

// AwsCredentialsCreate mocks base method.
func (m *MockApiClientInterface) AwsCredentialsCreate(arg0 AwsCredentialsCreatePayload) (ApiKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwsCredentialsCreate", arg0)
	ret0, _ := ret[0].(ApiKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AwsCredentialsCreate indicates an expected call of AwsCredentialsCreate.
func (mr *MockApiClientInterfaceMockRecorder) AwsCredentialsCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwsCredentialsCreate", reflect.TypeOf((*MockApiClientInterface)(nil).AwsCredentialsCreate), arg0)
}

// AwsCredentialsDelete mocks base method.
func (m *MockApiClientInterface) AwsCredentialsDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwsCredentialsDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AwsCredentialsDelete indicates an expected call of AwsCredentialsDelete.
func (mr *MockApiClientInterfaceMockRecorder) AwsCredentialsDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwsCredentialsDelete", reflect.TypeOf((*MockApiClientInterface)(nil).AwsCredentialsDelete), arg0)
}

// AwsCredentialsList mocks base method.
func (m *MockApiClientInterface) AwsCredentialsList() ([]ApiKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwsCredentialsList")
	ret0, _ := ret[0].([]ApiKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AwsCredentialsList indicates an expected call of AwsCredentialsList.
func (mr *MockApiClientInterfaceMockRecorder) AwsCredentialsList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwsCredentialsList", reflect.TypeOf((*MockApiClientInterface)(nil).AwsCredentialsList))
}

// CloudCredentialIdsInProject mocks base method.
func (m *MockApiClientInterface) CloudCredentialIdsInProject(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredentialIdsInProject", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudCredentialIdsInProject indicates an expected call of CloudCredentialIdsInProject.
func (mr *MockApiClientInterfaceMockRecorder) CloudCredentialIdsInProject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredentialIdsInProject", reflect.TypeOf((*MockApiClientInterface)(nil).CloudCredentialIdsInProject), arg0)
}

// ConfigurationVariableCreate mocks base method.
func (m *MockApiClientInterface) ConfigurationVariableCreate(arg0 ConfigurationVariableCreateParams) (ConfigurationVariable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigurationVariableCreate", arg0)
	ret0, _ := ret[0].(ConfigurationVariable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigurationVariableCreate indicates an expected call of ConfigurationVariableCreate.
func (mr *MockApiClientInterfaceMockRecorder) ConfigurationVariableCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigurationVariableCreate", reflect.TypeOf((*MockApiClientInterface)(nil).ConfigurationVariableCreate), arg0)
}

// ConfigurationVariableDelete mocks base method.
func (m *MockApiClientInterface) ConfigurationVariableDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigurationVariableDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigurationVariableDelete indicates an expected call of ConfigurationVariableDelete.
func (mr *MockApiClientInterfaceMockRecorder) ConfigurationVariableDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigurationVariableDelete", reflect.TypeOf((*MockApiClientInterface)(nil).ConfigurationVariableDelete), arg0)
}

// ConfigurationVariableUpdate mocks base method.
func (m *MockApiClientInterface) ConfigurationVariableUpdate(arg0 ConfigurationVariableUpdateParams) (ConfigurationVariable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigurationVariableUpdate", arg0)
	ret0, _ := ret[0].(ConfigurationVariable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigurationVariableUpdate indicates an expected call of ConfigurationVariableUpdate.
func (mr *MockApiClientInterfaceMockRecorder) ConfigurationVariableUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigurationVariableUpdate", reflect.TypeOf((*MockApiClientInterface)(nil).ConfigurationVariableUpdate), arg0)
}

// ConfigurationVariablesById mocks base method.
func (m *MockApiClientInterface) ConfigurationVariablesById(arg0 string) (ConfigurationVariable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigurationVariablesById", arg0)
	ret0, _ := ret[0].(ConfigurationVariable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigurationVariables indicates an expected call of ConfigurationVariables.
func (mr *MockApiClientInterfaceMockRecorder) ConfigurationVariables(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigurationVariablesByScope", reflect.TypeOf((*MockApiClientInterface)(nil).ConfigurationVariablesByScope), arg0, arg1)
}

// ConfigurationVariablesById indicates an expected call of ConfigurationVariablesById.
func (mr *MockApiClientInterfaceMockRecorder) ConfigurationVariablesById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigurationVariablesById", reflect.TypeOf((*MockApiClientInterface)(nil).ConfigurationVariablesById), arg0)
}

// ConfigurationVariablesByScope mocks base method.
func (m *MockApiClientInterface) ConfigurationVariablesByScope(arg0 Scope, arg1 string) ([]ConfigurationVariable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigurationVariablesByScope", arg0, arg1)
	ret0, _ := ret[0].([]ConfigurationVariable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigurationVariablesByScope indicates an expected call of ConfigurationVariablesByScope.
func (mr *MockApiClientInterfaceMockRecorder) ConfigurationVariablesByScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigurationVariablesByScope", reflect.TypeOf((*MockApiClientInterface)(nil).ConfigurationVariablesByScope), arg0, arg1)
}

// Environment mocks base method.
func (m *MockApiClientInterface) Environment(arg0 string) (Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Environment", arg0)
	ret0, _ := ret[0].(Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Environment indicates an expected call of Environment.
func (mr *MockApiClientInterfaceMockRecorder) Environment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Environment", reflect.TypeOf((*MockApiClientInterface)(nil).Environment), arg0)
}

// EnvironmentCreate mocks base method.
func (m *MockApiClientInterface) EnvironmentCreate(arg0 EnvironmentCreate) (Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentCreate", arg0)
	ret0, _ := ret[0].(Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvironmentCreate indicates an expected call of EnvironmentCreate.
func (mr *MockApiClientInterfaceMockRecorder) EnvironmentCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentCreate", reflect.TypeOf((*MockApiClientInterface)(nil).EnvironmentCreate), arg0)
}

// EnvironmentDeploy mocks base method.
func (m *MockApiClientInterface) EnvironmentDeploy(arg0 string, arg1 DeployRequest) (EnvironmentDeployResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentDeploy", arg0, arg1)
	ret0, _ := ret[0].(EnvironmentDeployResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvironmentDeploy indicates an expected call of EnvironmentDeploy.
func (mr *MockApiClientInterfaceMockRecorder) EnvironmentDeploy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentDeploy", reflect.TypeOf((*MockApiClientInterface)(nil).EnvironmentDeploy), arg0, arg1)
}

// EnvironmentDestroy mocks base method.
func (m *MockApiClientInterface) EnvironmentDestroy(arg0 string) (Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentDestroy", arg0)
	ret0, _ := ret[0].(Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvironmentDestroy indicates an expected call of EnvironmentDestroy.
func (mr *MockApiClientInterfaceMockRecorder) EnvironmentDestroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentDestroy", reflect.TypeOf((*MockApiClientInterface)(nil).EnvironmentDestroy), arg0)
}

// EnvironmentScheduling mocks base method.
func (m *MockApiClientInterface) EnvironmentScheduling(arg0 string) (EnvironmentScheduling, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentScheduling", arg0)
	ret0, _ := ret[0].(EnvironmentScheduling)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvironmentScheduling indicates an expected call of EnvironmentScheduling.
func (mr *MockApiClientInterfaceMockRecorder) EnvironmentScheduling(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentScheduling", reflect.TypeOf((*MockApiClientInterface)(nil).EnvironmentScheduling), arg0)
}

// EnvironmentSchedulingDelete mocks base method.
func (m *MockApiClientInterface) EnvironmentSchedulingDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentSchedulingDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnvironmentSchedulingDelete indicates an expected call of EnvironmentSchedulingDelete.
func (mr *MockApiClientInterfaceMockRecorder) EnvironmentSchedulingDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentSchedulingDelete", reflect.TypeOf((*MockApiClientInterface)(nil).EnvironmentSchedulingDelete), arg0)
}

// EnvironmentSchedulingUpdate mocks base method.
func (m *MockApiClientInterface) EnvironmentSchedulingUpdate(arg0 string, arg1 EnvironmentScheduling) (EnvironmentScheduling, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentSchedulingUpdate", arg0, arg1)
	ret0, _ := ret[0].(EnvironmentScheduling)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvironmentSchedulingUpdate indicates an expected call of EnvironmentSchedulingUpdate.
func (mr *MockApiClientInterfaceMockRecorder) EnvironmentSchedulingUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentSchedulingUpdate", reflect.TypeOf((*MockApiClientInterface)(nil).EnvironmentSchedulingUpdate), arg0, arg1)
}

// EnvironmentUpdate mocks base method.
func (m *MockApiClientInterface) EnvironmentUpdate(arg0 string, arg1 EnvironmentUpdate) (Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentUpdate", arg0, arg1)
	ret0, _ := ret[0].(Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvironmentUpdate indicates an expected call of EnvironmentUpdate.
func (mr *MockApiClientInterfaceMockRecorder) EnvironmentUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentUpdate", reflect.TypeOf((*MockApiClientInterface)(nil).EnvironmentUpdate), arg0, arg1)
}

// EnvironmentUpdateTTL mocks base method.
func (m *MockApiClientInterface) EnvironmentUpdateTTL(arg0 string, arg1 TTL) (Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnvironmentUpdateTTL", arg0, arg1)
	ret0, _ := ret[0].(Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnvironmentUpdateTTL indicates an expected call of EnvironmentUpdateTTL.
func (mr *MockApiClientInterfaceMockRecorder) EnvironmentUpdateTTL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnvironmentUpdateTTL", reflect.TypeOf((*MockApiClientInterface)(nil).EnvironmentUpdateTTL), arg0, arg1)
}

// Environments mocks base method.
func (m *MockApiClientInterface) Environments() ([]Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Environments")
	ret0, _ := ret[0].([]Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Environments indicates an expected call of Environments.
func (mr *MockApiClientInterfaceMockRecorder) Environments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Environments", reflect.TypeOf((*MockApiClientInterface)(nil).Environments))
}

// Organization mocks base method.
func (m *MockApiClientInterface) Organization() (Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Organization")
	ret0, _ := ret[0].(Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Organization indicates an expected call of Organization.
func (mr *MockApiClientInterfaceMockRecorder) Organization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Organization", reflect.TypeOf((*MockApiClientInterface)(nil).Organization))
}

// Policy mocks base method.
func (m *MockApiClientInterface) Policy(arg0 string) (Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Policy", arg0)
	ret0, _ := ret[0].(Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Policy indicates an expected call of Policy.
func (mr *MockApiClientInterfaceMockRecorder) Policy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Policy", reflect.TypeOf((*MockApiClientInterface)(nil).Policy), arg0)
}

// PolicyUpdate mocks base method.
func (m *MockApiClientInterface) PolicyUpdate(arg0 PolicyUpdatePayload) (Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicyUpdate", arg0)
	ret0, _ := ret[0].(Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PolicyUpdate indicates an expected call of PolicyUpdate.
func (mr *MockApiClientInterfaceMockRecorder) PolicyUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyUpdate", reflect.TypeOf((*MockApiClientInterface)(nil).PolicyUpdate), arg0)
}

// Project mocks base method.
func (m *MockApiClientInterface) Project(arg0 string) (Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Project", arg0)
	ret0, _ := ret[0].(Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Project indicates an expected call of Project.
func (mr *MockApiClientInterfaceMockRecorder) Project(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Project", reflect.TypeOf((*MockApiClientInterface)(nil).Project), arg0)
}

// ProjectCreate mocks base method.
func (m *MockApiClientInterface) ProjectCreate(arg0 ProjectCreatePayload) (Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectCreate", arg0)
	ret0, _ := ret[0].(Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectCreate indicates an expected call of ProjectCreate.
func (mr *MockApiClientInterfaceMockRecorder) ProjectCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectCreate", reflect.TypeOf((*MockApiClientInterface)(nil).ProjectCreate), arg0)
}

// ProjectDelete mocks base method.
func (m *MockApiClientInterface) ProjectDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProjectDelete indicates an expected call of ProjectDelete.
func (mr *MockApiClientInterfaceMockRecorder) ProjectDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectDelete", reflect.TypeOf((*MockApiClientInterface)(nil).ProjectDelete), arg0)
}

// ProjectUpdate mocks base method.
func (m *MockApiClientInterface) ProjectUpdate(arg0 string, arg1 ProjectCreatePayload) (Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectUpdate", arg0, arg1)
	ret0, _ := ret[0].(Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProjectUpdate indicates an expected call of ProjectUpdate.
func (mr *MockApiClientInterfaceMockRecorder) ProjectUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectUpdate", reflect.TypeOf((*MockApiClientInterface)(nil).ProjectUpdate), arg0, arg1)
}

// Projects mocks base method.
func (m *MockApiClientInterface) Projects() ([]Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects")
	ret0, _ := ret[0].([]Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Projects indicates an expected call of Projects.
func (mr *MockApiClientInterfaceMockRecorder) Projects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockApiClientInterface)(nil).Projects))
}

// RemoveCloudCredentialsFromProject mocks base method.
func (m *MockApiClientInterface) RemoveCloudCredentialsFromProject(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCloudCredentialsFromProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCloudCredentialsFromProject indicates an expected call of RemoveCloudCredentialsFromProject.
func (mr *MockApiClientInterfaceMockRecorder) RemoveCloudCredentialsFromProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCloudCredentialsFromProject", reflect.TypeOf((*MockApiClientInterface)(nil).RemoveCloudCredentialsFromProject), arg0, arg1)
}

// RemoveTemplateFromProject mocks base method.
func (m *MockApiClientInterface) RemoveTemplateFromProject(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTemplateFromProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTemplateFromProject indicates an expected call of RemoveTemplateFromProject.
func (mr *MockApiClientInterfaceMockRecorder) RemoveTemplateFromProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTemplateFromProject", reflect.TypeOf((*MockApiClientInterface)(nil).RemoveTemplateFromProject), arg0, arg1)
}

// SshKeyCreate mocks base method.
func (m *MockApiClientInterface) SshKeyCreate(arg0 SshKeyCreatePayload) (SshKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshKeyCreate", arg0)
	ret0, _ := ret[0].(SshKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SshKeyCreate indicates an expected call of SshKeyCreate.
func (mr *MockApiClientInterfaceMockRecorder) SshKeyCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshKeyCreate", reflect.TypeOf((*MockApiClientInterface)(nil).SshKeyCreate), arg0)
}

// SshKeyDelete mocks base method.
func (m *MockApiClientInterface) SshKeyDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshKeyDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SshKeyDelete indicates an expected call of SshKeyDelete.
func (mr *MockApiClientInterfaceMockRecorder) SshKeyDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshKeyDelete", reflect.TypeOf((*MockApiClientInterface)(nil).SshKeyDelete), arg0)
}

// SshKeys mocks base method.
func (m *MockApiClientInterface) SshKeys() ([]SshKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshKeys")
	ret0, _ := ret[0].([]SshKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SshKeys indicates an expected call of SshKeys.
func (mr *MockApiClientInterfaceMockRecorder) SshKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshKeys", reflect.TypeOf((*MockApiClientInterface)(nil).SshKeys))
}

// Team mocks base method.
func (m *MockApiClientInterface) Team(arg0 string) (Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Team", arg0)
	ret0, _ := ret[0].(Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Team indicates an expected call of Team.
func (mr *MockApiClientInterfaceMockRecorder) Team(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Team", reflect.TypeOf((*MockApiClientInterface)(nil).Team), arg0)
}

// TeamCreate mocks base method.
func (m *MockApiClientInterface) TeamCreate(arg0 TeamCreatePayload) (Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamCreate", arg0)
	ret0, _ := ret[0].(Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamCreate indicates an expected call of TeamCreate.
func (mr *MockApiClientInterfaceMockRecorder) TeamCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamCreate", reflect.TypeOf((*MockApiClientInterface)(nil).TeamCreate), arg0)
}

// TeamDelete mocks base method.
func (m *MockApiClientInterface) TeamDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TeamDelete indicates an expected call of TeamDelete.
func (mr *MockApiClientInterfaceMockRecorder) TeamDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamDelete", reflect.TypeOf((*MockApiClientInterface)(nil).TeamDelete), arg0)
}

// TeamProjectAssignmentCreateOrUpdate mocks base method.
func (m *MockApiClientInterface) TeamProjectAssignmentCreateOrUpdate(arg0 TeamProjectAssignmentPayload) (TeamProjectAssignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamProjectAssignmentCreateOrUpdate", arg0)
	ret0, _ := ret[0].(TeamProjectAssignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamProjectAssignmentCreateOrUpdate indicates an expected call of TeamProjectAssignmentCreateOrUpdate.
func (mr *MockApiClientInterfaceMockRecorder) TeamProjectAssignmentCreateOrUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamProjectAssignmentCreateOrUpdate", reflect.TypeOf((*MockApiClientInterface)(nil).TeamProjectAssignmentCreateOrUpdate), arg0)
}

// TeamProjectAssignmentDelete mocks base method.
func (m *MockApiClientInterface) TeamProjectAssignmentDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamProjectAssignmentDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TeamProjectAssignmentDelete indicates an expected call of TeamProjectAssignmentDelete.
func (mr *MockApiClientInterfaceMockRecorder) TeamProjectAssignmentDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamProjectAssignmentDelete", reflect.TypeOf((*MockApiClientInterface)(nil).TeamProjectAssignmentDelete), arg0)
}

// TeamProjectAssignments mocks base method.
func (m *MockApiClientInterface) TeamProjectAssignments(arg0 string) ([]TeamProjectAssignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamProjectAssignments", arg0)
	ret0, _ := ret[0].([]TeamProjectAssignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamProjectAssignments indicates an expected call of TeamProjectAssignments.
func (mr *MockApiClientInterfaceMockRecorder) TeamProjectAssignments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamProjectAssignments", reflect.TypeOf((*MockApiClientInterface)(nil).TeamProjectAssignments), arg0)
}

// TeamUpdate mocks base method.
func (m *MockApiClientInterface) TeamUpdate(arg0 string, arg1 TeamUpdatePayload) (Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamUpdate", arg0, arg1)
	ret0, _ := ret[0].(Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamUpdate indicates an expected call of TeamUpdate.
func (mr *MockApiClientInterfaceMockRecorder) TeamUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamUpdate", reflect.TypeOf((*MockApiClientInterface)(nil).TeamUpdate), arg0, arg1)
}

// Teams mocks base method.
func (m *MockApiClientInterface) Teams() ([]Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teams")
	ret0, _ := ret[0].([]Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teams indicates an expected call of Teams.
func (mr *MockApiClientInterfaceMockRecorder) Teams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teams", reflect.TypeOf((*MockApiClientInterface)(nil).Teams))
}

// Template mocks base method.
func (m *MockApiClientInterface) Template(arg0 string) (Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template", arg0)
	ret0, _ := ret[0].(Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MockApiClientInterfaceMockRecorder) Template(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockApiClientInterface)(nil).Template), arg0)
}

// TemplateCreate mocks base method.
func (m *MockApiClientInterface) TemplateCreate(arg0 TemplateCreatePayload) (Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemplateCreate", arg0)
	ret0, _ := ret[0].(Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TemplateCreate indicates an expected call of TemplateCreate.
func (mr *MockApiClientInterfaceMockRecorder) TemplateCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemplateCreate", reflect.TypeOf((*MockApiClientInterface)(nil).TemplateCreate), arg0)
}

// TemplateDelete mocks base method.
func (m *MockApiClientInterface) TemplateDelete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemplateDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TemplateDelete indicates an expected call of TemplateDelete.
func (mr *MockApiClientInterfaceMockRecorder) TemplateDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemplateDelete", reflect.TypeOf((*MockApiClientInterface)(nil).TemplateDelete), arg0)
}

// TemplateUpdate mocks base method.
func (m *MockApiClientInterface) TemplateUpdate(arg0 string, arg1 TemplateCreatePayload) (Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemplateUpdate", arg0, arg1)
	ret0, _ := ret[0].(Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TemplateUpdate indicates an expected call of TemplateUpdate.
func (mr *MockApiClientInterfaceMockRecorder) TemplateUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemplateUpdate", reflect.TypeOf((*MockApiClientInterface)(nil).TemplateUpdate), arg0, arg1)
}

// Templates mocks base method.
func (m *MockApiClientInterface) Templates() ([]Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Templates")
	ret0, _ := ret[0].([]Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Templates indicates an expected call of Templates.
func (mr *MockApiClientInterfaceMockRecorder) Templates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Templates", reflect.TypeOf((*MockApiClientInterface)(nil).Templates))
}

// WorkflowTrigger mocks base method.
func (m *MockApiClientInterface) WorkflowTrigger(arg0 string) ([]WorkflowTrigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkflowTrigger", arg0)
	ret0, _ := ret[0].([]WorkflowTrigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkflowTrigger indicates an expected call of WorkflowTrigger.
func (mr *MockApiClientInterfaceMockRecorder) WorkflowTrigger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkflowTrigger", reflect.TypeOf((*MockApiClientInterface)(nil).WorkflowTrigger), arg0)
}

// WorkflowTriggerUpsert mocks base method.
func (m *MockApiClientInterface) WorkflowTriggerUpsert(arg0 string, arg1 WorkflowTriggerUpsertPayload) ([]WorkflowTrigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkflowTriggerUpsert", arg0, arg1)
	ret0, _ := ret[0].([]WorkflowTrigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkflowTriggerUpsert indicates an expected call of WorkflowTriggerUpsert.
func (mr *MockApiClientInterfaceMockRecorder) WorkflowTriggerUpsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkflowTriggerUpsert", reflect.TypeOf((*MockApiClientInterface)(nil).WorkflowTriggerUpsert), arg0, arg1)
}

// organizationId mocks base method.
func (m *MockApiClientInterface) organizationId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "organizationId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// organizationId indicates an expected call of organizationId.
func (mr *MockApiClientInterfaceMockRecorder) organizationId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "organizationId", reflect.TypeOf((*MockApiClientInterface)(nil).organizationId))
}
