// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/AmazingTalker/go-amazing/pkg/pb"

	testing "testing"
)

// GoAmazingRPC is an autogenerated mock type for the GoAmazingRPC type
type GoAmazingRPC struct {
	mock.Mock
}

// Config provides a mock function with given fields: _a0, _a1
func (_m *GoAmazingRPC) Config(_a0 context.Context, _a1 *pb.ConfigReq) (*pb.ConfigRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.ConfigRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ConfigReq) *pb.ConfigRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ConfigRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.ConfigReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRecord provides a mock function with given fields: _a0, _a1
func (_m *GoAmazingRPC) CreateRecord(_a0 context.Context, _a1 *pb.CreateRecordReq) (*pb.CreateRecordRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.CreateRecordRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateRecordReq) *pb.CreateRecordRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CreateRecordRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateRecordReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecord provides a mock function with given fields: _a0, _a1
func (_m *GoAmazingRPC) GetRecord(_a0 context.Context, _a1 *pb.GetRecordReq) (*pb.GetRecordRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetRecordRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetRecordReq) *pb.GetRecordRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetRecordRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetRecordReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields: _a0, _a1
func (_m *GoAmazingRPC) Health(_a0 context.Context, _a1 *pb.HealthReq) (*pb.HealthRes, error) {
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

// ListRecord provides a mock function with given fields: _a0, _a1
func (_m *GoAmazingRPC) ListRecord(_a0 context.Context, _a1 *pb.ListRecordReq) (*pb.ListRecordRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.ListRecordRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListRecordReq) *pb.ListRecordRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ListRecordRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.ListRecordReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGoAmazingRPC creates a new instance of GoAmazingRPC. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewGoAmazingRPC(t testing.TB) *GoAmazingRPC {
	mock := &GoAmazingRPC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
