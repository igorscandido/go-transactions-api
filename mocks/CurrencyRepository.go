// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CurrencyRepository is an autogenerated mock type for the CurrencyRepository type
type CurrencyRepository struct {
	mock.Mock
}

// GetConversionRateForCurrency provides a mock function with given fields: ctx, baseCurrency, destCurrency
func (_m *CurrencyRepository) GetConversionRateForCurrency(ctx context.Context, baseCurrency string, destCurrency string) (*float64, error) {
	ret := _m.Called(ctx, baseCurrency, destCurrency)

	if len(ret) == 0 {
		panic("no return value specified for GetConversionRateForCurrency")
	}

	var r0 *float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*float64, error)); ok {
		return rf(ctx, baseCurrency, destCurrency)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *float64); ok {
		r0 = rf(ctx, baseCurrency, destCurrency)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*float64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, baseCurrency, destCurrency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCurrencyRepository creates a new instance of CurrencyRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCurrencyRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CurrencyRepository {
	mock := &CurrencyRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
