// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/AmazingTalker/go-amazing/pkg/pb"

	testing "testing"
)

// GoAmazingClient is an autogenerated mock type for the GoAmazingClient type
type GoAmazingClient struct {
	mock.Mock
}

// Config provides a mock function with given fields: ctx, in, opts
func (_m *GoAmazingClient) Config(ctx context.Context, in *pb.ConfigReq, opts ...grpc.CallOption) (*pb.ConfigRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.ConfigRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ConfigReq, ...grpc.CallOption) *pb.ConfigRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ConfigRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.ConfigReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRecord provides a mock function with given fields: ctx, in, opts
func (_m *GoAmazingClient) CreateRecord(ctx context.Context, in *pb.CreateRecordReq, opts ...grpc.CallOption) (*pb.CreateRecordRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.CreateRecordRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateRecordReq, ...grpc.CallOption) *pb.CreateRecordRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CreateRecordRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateRecordReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecord provides a mock function with given fields: ctx, in, opts
func (_m *GoAmazingClient) GetRecord(ctx context.Context, in *pb.GetRecordReq, opts ...grpc.CallOption) (*pb.GetRecordRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.GetRecordRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetRecordReq, ...grpc.CallOption) *pb.GetRecordRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetRecordRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetRecordReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Health provides a mock function with given fields: ctx, in, opts
func (_m *GoAmazingClient) Health(ctx context.Context, in *pb.HealthReq, opts ...grpc.CallOption) (*pb.HealthRes, error) {
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

// ListRecord provides a mock function with given fields: ctx, in, opts
func (_m *GoAmazingClient) ListRecord(ctx context.Context, in *pb.ListRecordReq, opts ...grpc.CallOption) (*pb.ListRecordRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.ListRecordRes
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListRecordReq, ...grpc.CallOption) *pb.ListRecordRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ListRecordRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.ListRecordReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGoAmazingClient creates a new instance of GoAmazingClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewGoAmazingClient(t testing.TB) *GoAmazingClient {
	mock := &GoAmazingClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
