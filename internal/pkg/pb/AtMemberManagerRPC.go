// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/AmazingTalker/at-member-manager/pkg/pb"
)

// AtMemberManagerRPC is an autogenerated mock type for the AtMemberManagerRPC type
type AtMemberManagerRPC struct {
	mock.Mock
}

// CreateMember provides a mock function with given fields: _a0, _a1
func (_m *AtMemberManagerRPC) CreateMember(_a0 context.Context, _a1 *pb.CreateMemberReq) (*pb.CreateMemberRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.CreateMemberRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateMemberReq) *pb.CreateMemberRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CreateMemberRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateMemberReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMember provides a mock function with given fields: _a0, _a1
func (_m *AtMemberManagerRPC) DeleteMember(_a0 context.Context, _a1 *pb.DeleteMemberReq) (*pb.DeleteMemberRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.DeleteMemberRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteMemberReq) *pb.DeleteMemberRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteMemberRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteMemberReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields: _a0, _a1
func (_m *AtMemberManagerRPC) Health(_a0 context.Context, _a1 *pb.HealthReq) (*pb.HealthRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.HealthRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.HealthReq) *pb.HealthRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.HealthRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.HealthReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListMembers provides a mock function with given fields: _a0, _a1
func (_m *AtMemberManagerRPC) ListMembers(_a0 context.Context, _a1 *pb.ListMembersReq) (*pb.ListMembersRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.ListMembersRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListMembersReq) *pb.ListMembersRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ListMembersRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.ListMembersReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMember provides a mock function with given fields: _a0, _a1
func (_m *AtMemberManagerRPC) UpdateMember(_a0 context.Context, _a1 *pb.UpdateMemberReq) (*pb.UpdateMemberRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.UpdateMemberRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateMemberReq) *pb.UpdateMemberRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.UpdateMemberRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateMemberReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAtMemberManagerRPC interface {
	mock.TestingT
	Cleanup(func())
}

// NewAtMemberManagerRPC creates a new instance of AtMemberManagerRPC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAtMemberManagerRPC(t mockConstructorTestingTNewAtMemberManagerRPC) *AtMemberManagerRPC {
	mock := &AtMemberManagerRPC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
