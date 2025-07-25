// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"context"

	"github.com/opencloud-eu/opencloud/protogen/gen/opencloud/services/thumbnails/v0"
	mock "github.com/stretchr/testify/mock"
	"go-micro.dev/v4/client"
)

// NewThumbnailService creates a new instance of ThumbnailService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewThumbnailService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ThumbnailService {
	mock := &ThumbnailService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// ThumbnailService is an autogenerated mock type for the ThumbnailService type
type ThumbnailService struct {
	mock.Mock
}

type ThumbnailService_Expecter struct {
	mock *mock.Mock
}

func (_m *ThumbnailService) EXPECT() *ThumbnailService_Expecter {
	return &ThumbnailService_Expecter{mock: &_m.Mock}
}

// GetThumbnail provides a mock function for the type ThumbnailService
func (_mock *ThumbnailService) GetThumbnail(ctx context.Context, in *v0.GetThumbnailRequest, opts ...client.CallOption) (*v0.GetThumbnailResponse, error) {
	var tmpRet mock.Arguments
	if len(opts) > 0 {
		tmpRet = _mock.Called(ctx, in, opts)
	} else {
		tmpRet = _mock.Called(ctx, in)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for GetThumbnail")
	}

	var r0 *v0.GetThumbnailResponse
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, *v0.GetThumbnailRequest, ...client.CallOption) (*v0.GetThumbnailResponse, error)); ok {
		return returnFunc(ctx, in, opts...)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, *v0.GetThumbnailRequest, ...client.CallOption) *v0.GetThumbnailResponse); ok {
		r0 = returnFunc(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v0.GetThumbnailResponse)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, *v0.GetThumbnailRequest, ...client.CallOption) error); ok {
		r1 = returnFunc(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// ThumbnailService_GetThumbnail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetThumbnail'
type ThumbnailService_GetThumbnail_Call struct {
	*mock.Call
}

// GetThumbnail is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v0.GetThumbnailRequest
//   - opts ...client.CallOption
func (_e *ThumbnailService_Expecter) GetThumbnail(ctx interface{}, in interface{}, opts ...interface{}) *ThumbnailService_GetThumbnail_Call {
	return &ThumbnailService_GetThumbnail_Call{Call: _e.mock.On("GetThumbnail",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ThumbnailService_GetThumbnail_Call) Run(run func(ctx context.Context, in *v0.GetThumbnailRequest, opts ...client.CallOption)) *ThumbnailService_GetThumbnail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 *v0.GetThumbnailRequest
		if args[1] != nil {
			arg1 = args[1].(*v0.GetThumbnailRequest)
		}
		var arg2 []client.CallOption
		var variadicArgs []client.CallOption
		if len(args) > 2 {
			variadicArgs = args[2].([]client.CallOption)
		}
		arg2 = variadicArgs
		run(
			arg0,
			arg1,
			arg2...,
		)
	})
	return _c
}

func (_c *ThumbnailService_GetThumbnail_Call) Return(getThumbnailResponse *v0.GetThumbnailResponse, err error) *ThumbnailService_GetThumbnail_Call {
	_c.Call.Return(getThumbnailResponse, err)
	return _c
}

func (_c *ThumbnailService_GetThumbnail_Call) RunAndReturn(run func(ctx context.Context, in *v0.GetThumbnailRequest, opts ...client.CallOption) (*v0.GetThumbnailResponse, error)) *ThumbnailService_GetThumbnail_Call {
	_c.Call.Return(run)
	return _c
}
