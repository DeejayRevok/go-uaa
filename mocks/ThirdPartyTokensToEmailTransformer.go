// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	thirdParty "go-iam/src/domain/auth/thirdParty"

	mock "github.com/stretchr/testify/mock"
)

// ThirdPartyTokensToEmailTransformer is an autogenerated mock type for the ThirdPartyTokensToEmailTransformer type
type ThirdPartyTokensToEmailTransformer struct {
	mock.Mock
}

// Transform provides a mock function with given fields: tokens
func (_m *ThirdPartyTokensToEmailTransformer) Transform(tokens *thirdParty.ThirdPartyTokens) (string, error) {
	ret := _m.Called(tokens)

	var r0 string
	if rf, ok := ret.Get(0).(func(*thirdParty.ThirdPartyTokens) string); ok {
		r0 = rf(tokens)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*thirdParty.ThirdPartyTokens) error); ok {
		r1 = rf(tokens)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewThirdPartyTokensToEmailTransformer interface {
	mock.TestingT
	Cleanup(func())
}

// NewThirdPartyTokensToEmailTransformer creates a new instance of ThirdPartyTokensToEmailTransformer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewThirdPartyTokensToEmailTransformer(t mockConstructorTestingTNewThirdPartyTokensToEmailTransformer) *ThirdPartyTokensToEmailTransformer {
	mock := &ThirdPartyTokensToEmailTransformer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
