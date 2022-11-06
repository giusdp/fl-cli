// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	openapi "github.com/funlessdev/fl-client-sdk-go"
	mock "github.com/stretchr/testify/mock"

	os "os"
)

// FnHandler is an autogenerated mock type for the FnHandler type
type FnHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, fnName, fnNamespace, code
func (_m *FnHandler) Create(ctx context.Context, fnName string, fnNamespace string, code *os.File) (openapi.FunctionCreationSuccess, error) {
	ret := _m.Called(ctx, fnName, fnNamespace, code)

	var r0 openapi.FunctionCreationSuccess
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *os.File) openapi.FunctionCreationSuccess); ok {
		r0 = rf(ctx, fnName, fnNamespace, code)
	} else {
		r0 = ret.Get(0).(openapi.FunctionCreationSuccess)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *os.File) error); ok {
		r1 = rf(ctx, fnName, fnNamespace, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, fnName, fnNamespace
func (_m *FnHandler) Delete(ctx context.Context, fnName string, fnNamespace string) (openapi.FunctionDeletionSuccess, error) {
	ret := _m.Called(ctx, fnName, fnNamespace)

	var r0 openapi.FunctionDeletionSuccess
	if rf, ok := ret.Get(0).(func(context.Context, string, string) openapi.FunctionDeletionSuccess); ok {
		r0 = rf(ctx, fnName, fnNamespace)
	} else {
		r0 = ret.Get(0).(openapi.FunctionDeletionSuccess)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, fnName, fnNamespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Invoke provides a mock function with given fields: ctx, fnName, fnNamespace, fnArgs
func (_m *FnHandler) Invoke(ctx context.Context, fnName string, fnNamespace string, fnArgs map[string]interface{}) (openapi.FunctionInvocationSuccess, error) {
	ret := _m.Called(ctx, fnName, fnNamespace, fnArgs)

	var r0 openapi.FunctionInvocationSuccess
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]interface{}) openapi.FunctionInvocationSuccess); ok {
		r0 = rf(ctx, fnName, fnNamespace, fnArgs)
	} else {
		r0 = ret.Get(0).(openapi.FunctionInvocationSuccess)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, map[string]interface{}) error); ok {
		r1 = rf(ctx, fnName, fnNamespace, fnArgs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFnHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewFnHandler creates a new instance of FnHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFnHandler(t mockConstructorTestingTNewFnHandler) *FnHandler {
	mock := &FnHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
