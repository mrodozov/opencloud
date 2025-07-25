// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/opencloud-eu/opencloud/pkg/oidc"
	mock "github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

// NewOIDCClient creates a new instance of OIDCClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOIDCClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *OIDCClient {
	mock := &OIDCClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// OIDCClient is an autogenerated mock type for the OIDCClient type
type OIDCClient struct {
	mock.Mock
}

type OIDCClient_Expecter struct {
	mock *mock.Mock
}

func (_m *OIDCClient) EXPECT() *OIDCClient_Expecter {
	return &OIDCClient_Expecter{mock: &_m.Mock}
}

// UserInfo provides a mock function for the type OIDCClient
func (_mock *OIDCClient) UserInfo(ctx context.Context, ts oauth2.TokenSource) (*oidc.UserInfo, error) {
	ret := _mock.Called(ctx, ts)

	if len(ret) == 0 {
		panic("no return value specified for UserInfo")
	}

	var r0 *oidc.UserInfo
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, oauth2.TokenSource) (*oidc.UserInfo, error)); ok {
		return returnFunc(ctx, ts)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, oauth2.TokenSource) *oidc.UserInfo); ok {
		r0 = returnFunc(ctx, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.UserInfo)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, oauth2.TokenSource) error); ok {
		r1 = returnFunc(ctx, ts)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// OIDCClient_UserInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserInfo'
type OIDCClient_UserInfo_Call struct {
	*mock.Call
}

// UserInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - ts oauth2.TokenSource
func (_e *OIDCClient_Expecter) UserInfo(ctx interface{}, ts interface{}) *OIDCClient_UserInfo_Call {
	return &OIDCClient_UserInfo_Call{Call: _e.mock.On("UserInfo", ctx, ts)}
}

func (_c *OIDCClient_UserInfo_Call) Run(run func(ctx context.Context, ts oauth2.TokenSource)) *OIDCClient_UserInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 oauth2.TokenSource
		if args[1] != nil {
			arg1 = args[1].(oauth2.TokenSource)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *OIDCClient_UserInfo_Call) Return(userInfo *oidc.UserInfo, err error) *OIDCClient_UserInfo_Call {
	_c.Call.Return(userInfo, err)
	return _c
}

func (_c *OIDCClient_UserInfo_Call) RunAndReturn(run func(ctx context.Context, ts oauth2.TokenSource) (*oidc.UserInfo, error)) *OIDCClient_UserInfo_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyAccessToken provides a mock function for the type OIDCClient
func (_mock *OIDCClient) VerifyAccessToken(ctx context.Context, token string) (oidc.RegClaimsWithSID, jwt.MapClaims, error) {
	ret := _mock.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyAccessToken")
	}

	var r0 oidc.RegClaimsWithSID
	var r1 jwt.MapClaims
	var r2 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (oidc.RegClaimsWithSID, jwt.MapClaims, error)); ok {
		return returnFunc(ctx, token)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) oidc.RegClaimsWithSID); ok {
		r0 = returnFunc(ctx, token)
	} else {
		r0 = ret.Get(0).(oidc.RegClaimsWithSID)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) jwt.MapClaims); ok {
		r1 = returnFunc(ctx, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(jwt.MapClaims)
		}
	}
	if returnFunc, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = returnFunc(ctx, token)
	} else {
		r2 = ret.Error(2)
	}
	return r0, r1, r2
}

// OIDCClient_VerifyAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyAccessToken'
type OIDCClient_VerifyAccessToken_Call struct {
	*mock.Call
}

// VerifyAccessToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *OIDCClient_Expecter) VerifyAccessToken(ctx interface{}, token interface{}) *OIDCClient_VerifyAccessToken_Call {
	return &OIDCClient_VerifyAccessToken_Call{Call: _e.mock.On("VerifyAccessToken", ctx, token)}
}

func (_c *OIDCClient_VerifyAccessToken_Call) Run(run func(ctx context.Context, token string)) *OIDCClient_VerifyAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *OIDCClient_VerifyAccessToken_Call) Return(regClaimsWithSID oidc.RegClaimsWithSID, mapClaims jwt.MapClaims, err error) *OIDCClient_VerifyAccessToken_Call {
	_c.Call.Return(regClaimsWithSID, mapClaims, err)
	return _c
}

func (_c *OIDCClient_VerifyAccessToken_Call) RunAndReturn(run func(ctx context.Context, token string) (oidc.RegClaimsWithSID, jwt.MapClaims, error)) *OIDCClient_VerifyAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyLogoutToken provides a mock function for the type OIDCClient
func (_mock *OIDCClient) VerifyLogoutToken(ctx context.Context, token string) (*oidc.LogoutToken, error) {
	ret := _mock.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyLogoutToken")
	}

	var r0 *oidc.LogoutToken
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (*oidc.LogoutToken, error)); ok {
		return returnFunc(ctx, token)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) *oidc.LogoutToken); ok {
		r0 = returnFunc(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.LogoutToken)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(ctx, token)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// OIDCClient_VerifyLogoutToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyLogoutToken'
type OIDCClient_VerifyLogoutToken_Call struct {
	*mock.Call
}

// VerifyLogoutToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *OIDCClient_Expecter) VerifyLogoutToken(ctx interface{}, token interface{}) *OIDCClient_VerifyLogoutToken_Call {
	return &OIDCClient_VerifyLogoutToken_Call{Call: _e.mock.On("VerifyLogoutToken", ctx, token)}
}

func (_c *OIDCClient_VerifyLogoutToken_Call) Run(run func(ctx context.Context, token string)) *OIDCClient_VerifyLogoutToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *OIDCClient_VerifyLogoutToken_Call) Return(logoutToken *oidc.LogoutToken, err error) *OIDCClient_VerifyLogoutToken_Call {
	_c.Call.Return(logoutToken, err)
	return _c
}

func (_c *OIDCClient_VerifyLogoutToken_Call) RunAndReturn(run func(ctx context.Context, token string) (*oidc.LogoutToken, error)) *OIDCClient_VerifyLogoutToken_Call {
	_c.Call.Return(run)
	return _c
}
