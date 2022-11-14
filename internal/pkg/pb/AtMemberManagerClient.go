// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/AmazingTalker/at-member-manager/pkg/pb"
)

// AtMemberManagerClient is an autogenerated mock type for the AtMemberManagerClient type
type AtMemberManagerClient struct {
	mock.Mock
}

// CreateMember provides a mock function with given fields: ctx, in, opts
func (_m *AtMemberManagerClient) CreateMember(ctx context.Context, in *pb.CreateMemberReq, opts ...grpc.CallOption) (*pb.CreateMemberRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.CreateMemberRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateMemberReq, ...grpc.CallOption) *pb.CreateMemberRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CreateMemberRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateMemberReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMember provides a mock function with given fields: ctx, in, opts
func (_m *AtMemberManagerClient) DeleteMember(ctx context.Context, in *pb.DeleteMemberReq, opts ...grpc.CallOption) (*pb.DeleteMemberRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.DeleteMemberRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteMemberReq, ...grpc.CallOption) *pb.DeleteMemberRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteMemberRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteMemberReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields: ctx, in, opts
func (_m *AtMemberManagerClient) Health(ctx context.Context, in *pb.HealthReq, opts ...grpc.CallOption) (*pb.HealthRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.HealthRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.HealthReq, ...grpc.CallOption) *pb.HealthRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.HealthRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.HealthReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMembers provides a mock function with given fields: ctx, in, opts
func (_m *AtMemberManagerClient) ListMembers(ctx context.Context, in *pb.ListMembersReq, opts ...grpc.CallOption) (*pb.ListMembersRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.ListMembersRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListMembersReq, ...grpc.CallOption) *pb.ListMembersRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ListMembersRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.ListMembersReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMember provides a mock function with given fields: ctx, in, opts
func (_m *AtMemberManagerClient) UpdateMember(ctx context.Context, in *pb.UpdateMemberReq, opts ...grpc.CallOption) (*pb.UpdateMemberRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.UpdateMemberRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateMemberReq, ...grpc.CallOption) *pb.UpdateMemberRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.UpdateMemberRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateMemberReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAtMemberManagerClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewAtMemberManagerClient creates a new instance of AtMemberManagerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAtMemberManagerClient(t mockConstructorTestingTNewAtMemberManagerClient) *AtMemberManagerClient {
	mock := &AtMemberManagerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
