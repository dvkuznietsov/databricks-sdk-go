// Code generated by mockery v2.43.0. DO NOT EDIT.

package billing

import (
	context "context"

	billing "github.com/databricks/databricks-sdk-go/service/billing"

	mock "github.com/stretchr/testify/mock"
)

// MockUsageDashboardsInterface is an autogenerated mock type for the UsageDashboardsInterface type
type MockUsageDashboardsInterface struct {
	mock.Mock
}

type MockUsageDashboardsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUsageDashboardsInterface) EXPECT() *MockUsageDashboardsInterface_Expecter {
	return &MockUsageDashboardsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockUsageDashboardsInterface) Create(ctx context.Context, request billing.CreateBillingUsageDashboardRequest) (*billing.CreateBillingUsageDashboardResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *billing.CreateBillingUsageDashboardResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, billing.CreateBillingUsageDashboardRequest) (*billing.CreateBillingUsageDashboardResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, billing.CreateBillingUsageDashboardRequest) *billing.CreateBillingUsageDashboardResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.CreateBillingUsageDashboardResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, billing.CreateBillingUsageDashboardRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsageDashboardsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUsageDashboardsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request billing.CreateBillingUsageDashboardRequest
func (_e *MockUsageDashboardsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockUsageDashboardsInterface_Create_Call {
	return &MockUsageDashboardsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockUsageDashboardsInterface_Create_Call) Run(run func(ctx context.Context, request billing.CreateBillingUsageDashboardRequest)) *MockUsageDashboardsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(billing.CreateBillingUsageDashboardRequest))
	})
	return _c
}

func (_c *MockUsageDashboardsInterface_Create_Call) Return(_a0 *billing.CreateBillingUsageDashboardResponse, _a1 error) *MockUsageDashboardsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsageDashboardsInterface_Create_Call) RunAndReturn(run func(context.Context, billing.CreateBillingUsageDashboardRequest) (*billing.CreateBillingUsageDashboardResponse, error)) *MockUsageDashboardsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockUsageDashboardsInterface) Get(ctx context.Context, request billing.GetBillingUsageDashboardRequest) (*billing.GetBillingUsageDashboardResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *billing.GetBillingUsageDashboardResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, billing.GetBillingUsageDashboardRequest) (*billing.GetBillingUsageDashboardResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, billing.GetBillingUsageDashboardRequest) *billing.GetBillingUsageDashboardResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.GetBillingUsageDashboardResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, billing.GetBillingUsageDashboardRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUsageDashboardsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockUsageDashboardsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request billing.GetBillingUsageDashboardRequest
func (_e *MockUsageDashboardsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockUsageDashboardsInterface_Get_Call {
	return &MockUsageDashboardsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockUsageDashboardsInterface_Get_Call) Run(run func(ctx context.Context, request billing.GetBillingUsageDashboardRequest)) *MockUsageDashboardsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(billing.GetBillingUsageDashboardRequest))
	})
	return _c
}

func (_c *MockUsageDashboardsInterface_Get_Call) Return(_a0 *billing.GetBillingUsageDashboardResponse, _a1 error) *MockUsageDashboardsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsageDashboardsInterface_Get_Call) RunAndReturn(run func(context.Context, billing.GetBillingUsageDashboardRequest) (*billing.GetBillingUsageDashboardResponse, error)) *MockUsageDashboardsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockUsageDashboardsInterface) Impl() billing.UsageDashboardsService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 billing.UsageDashboardsService
	if rf, ok := ret.Get(0).(func() billing.UsageDashboardsService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billing.UsageDashboardsService)
		}
	}

	return r0
}

// MockUsageDashboardsInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockUsageDashboardsInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockUsageDashboardsInterface_Expecter) Impl() *MockUsageDashboardsInterface_Impl_Call {
	return &MockUsageDashboardsInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockUsageDashboardsInterface_Impl_Call) Run(run func()) *MockUsageDashboardsInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUsageDashboardsInterface_Impl_Call) Return(_a0 billing.UsageDashboardsService) *MockUsageDashboardsInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsageDashboardsInterface_Impl_Call) RunAndReturn(run func() billing.UsageDashboardsService) *MockUsageDashboardsInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockUsageDashboardsInterface) WithImpl(impl billing.UsageDashboardsService) billing.UsageDashboardsInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 billing.UsageDashboardsInterface
	if rf, ok := ret.Get(0).(func(billing.UsageDashboardsService) billing.UsageDashboardsInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(billing.UsageDashboardsInterface)
		}
	}

	return r0
}

// MockUsageDashboardsInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockUsageDashboardsInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl billing.UsageDashboardsService
func (_e *MockUsageDashboardsInterface_Expecter) WithImpl(impl interface{}) *MockUsageDashboardsInterface_WithImpl_Call {
	return &MockUsageDashboardsInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockUsageDashboardsInterface_WithImpl_Call) Run(run func(impl billing.UsageDashboardsService)) *MockUsageDashboardsInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(billing.UsageDashboardsService))
	})
	return _c
}

func (_c *MockUsageDashboardsInterface_WithImpl_Call) Return(_a0 billing.UsageDashboardsInterface) *MockUsageDashboardsInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsageDashboardsInterface_WithImpl_Call) RunAndReturn(run func(billing.UsageDashboardsService) billing.UsageDashboardsInterface) *MockUsageDashboardsInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUsageDashboardsInterface creates a new instance of MockUsageDashboardsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUsageDashboardsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUsageDashboardsInterface {
	mock := &MockUsageDashboardsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
