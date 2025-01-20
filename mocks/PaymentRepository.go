// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/igorscandido/go-transactions-api/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// PaymentRepository is an autogenerated mock type for the PaymentRepository type
type PaymentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, payment
func (_m *PaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	ret := _m.Called(ctx, payment)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Payment) error); ok {
		r0 = rf(ctx, payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *PaymentRepository) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Payment, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Payment); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPaymentRepository creates a new instance of PaymentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPaymentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PaymentRepository {
	mock := &PaymentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
