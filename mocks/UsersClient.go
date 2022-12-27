// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1beta1 "github.com/input-stream/cli/build/stack/inputstream/v1beta1"
)

// UsersClient is an autogenerated mock type for the UsersClient type
type UsersClient struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: ctx, in, opts
func (_m *UsersClient) GetUser(ctx context.Context, in *v1beta1.GetUserRequest, opts ...grpc.CallOption) (*v1beta1.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta1.User
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.GetUserRequest, ...grpc.CallOption) *v1beta1.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.GetUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx, in, opts
func (_m *UsersClient) ListUsers(ctx context.Context, in *v1beta1.ListUsersRequest, opts ...grpc.CallOption) (*v1beta1.ListUsersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta1.ListUsersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.ListUsersRequest, ...grpc.CallOption) *v1beta1.ListUsersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.ListUsersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.ListUsersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, in, opts
func (_m *UsersClient) UpdateUser(ctx context.Context, in *v1beta1.UpdateUserRequest, opts ...grpc.CallOption) (*v1beta1.User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta1.User
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.UpdateUserRequest, ...grpc.CallOption) *v1beta1.User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.UpdateUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUsersClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersClient creates a new instance of UsersClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersClient(t mockConstructorTestingTNewUsersClient) *UsersClient {
	mock := &UsersClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}