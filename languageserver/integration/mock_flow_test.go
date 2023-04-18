// Code generated by mockery v1.0.0. DO NOT EDIT.

package integration

import (
	cadence "github.com/onflow/cadence"
	flow "github.com/onflow/flow-go-sdk"

	flowkit "github.com/onflow/flow-cli/flowkit"

	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// mockFlowClient is an autogenerated mock type for the flowClient type
type mockFlowClient struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields:
func (_m *mockFlowClient) CreateAccount() (*clientAccount, error) {
	ret := _m.Called()

	var r0 *clientAccount
	if rf, ok := ret.Get(0).(func() *clientAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeployContract provides a mock function with given fields: address, name, location
func (_m *mockFlowClient) DeployContract(address flow.Address, name string, location *url.URL) error {
	ret := _m.Called(address, name, location)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Address, string, *url.URL) error); ok {
		r0 = rf(address, name, location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecuteScript provides a mock function with given fields: location, args
func (_m *mockFlowClient) ExecuteScript(location *url.URL, args []cadence.Value) (cadence.Value, error) {
	ret := _m.Called(location, args)

	var r0 cadence.Value
	if rf, ok := ret.Get(0).(func(*url.URL, []cadence.Value) cadence.Value); ok {
		r0 = rf(location, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cadence.Value)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*url.URL, []cadence.Value) error); ok {
		r1 = rf(location, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: address
func (_m *mockFlowClient) GetAccount(address flow.Address) (*flow.Account, error) {
	ret := _m.Called(address)

	var r0 *flow.Account
	if rf, ok := ret.Get(0).(func(flow.Address) *flow.Account); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveClientAccount provides a mock function with given fields:
func (_m *mockFlowClient) GetActiveClientAccount() *clientAccount {
	ret := _m.Called()

	var r0 *clientAccount
	if rf, ok := ret.Get(0).(func() *clientAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientAccount)
		}
	}

	return r0
}

// GetClientAccount provides a mock function with given fields: name
func (_m *mockFlowClient) GetClientAccount(name string) *clientAccount {
	ret := _m.Called(name)

	var r0 *clientAccount
	if rf, ok := ret.Get(0).(func(string) *clientAccount); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clientAccount)
		}
	}

	return r0
}

// GetClientAccounts provides a mock function with given fields:
func (_m *mockFlowClient) GetClientAccounts() []*clientAccount {
	ret := _m.Called()

	var r0 []*clientAccount
	if rf, ok := ret.Get(0).(func() []*clientAccount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*clientAccount)
		}
	}

	return r0
}

// GetCodeByName provides a mock function with given fields: name
func (_m *mockFlowClient) GetCodeByName(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: configPath, numberOfAccounts
func (_m *mockFlowClient) Initialize(configPath string, numberOfAccounts int) error {
	ret := _m.Called(configPath, numberOfAccounts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(configPath, numberOfAccounts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reload provides a mock function with given fields:
func (_m *mockFlowClient) Reload() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendTransaction provides a mock function with given fields: authorizers, location, args
func (_m *mockFlowClient) SendTransaction(authorizers []flow.Address, location *url.URL, args []cadence.Value) (*flow.Transaction, *flow.TransactionResult, error) {
	ret := _m.Called(authorizers, location, args)

	var r0 *flow.Transaction
	if rf, ok := ret.Get(0).(func([]flow.Address, *url.URL, []cadence.Value) *flow.Transaction); ok {
		r0 = rf(authorizers, location, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Transaction)
		}
	}

	var r1 *flow.TransactionResult
	if rf, ok := ret.Get(1).(func([]flow.Address, *url.URL, []cadence.Value) *flow.TransactionResult); ok {
		r1 = rf(authorizers, location, args)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*flow.TransactionResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]flow.Address, *url.URL, []cadence.Value) error); ok {
		r2 = rf(authorizers, location, args)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetActiveClientAccount provides a mock function with given fields: name
func (_m *mockFlowClient) SetActiveClientAccount(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// getConfigPath provides a mock function with given fields:
func (_m *mockFlowClient) getConfigPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getState provides a mock function with given fields:
func (_m *mockFlowClient) getState() *flowkit.State {
	ret := _m.Called()

	var r0 *flowkit.State
	if rf, ok := ret.Get(0).(func() *flowkit.State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flowkit.State)
		}
	}

	return r0
}
