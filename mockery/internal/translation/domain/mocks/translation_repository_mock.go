// Code generated by mockery v2.42.3. DO NOT EDIT.

package domain

import (
	context "context"

	domain "github.com/dangquyitt/tech-explorer/mockery/internal/translation/domain"
	mock "github.com/stretchr/testify/mock"
)

// TranslationRepositoryMock is an autogenerated mock type for the TranslationRepository type
type TranslationRepositoryMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, translation
func (_m *TranslationRepositoryMock) Create(ctx context.Context, translation *domain.Translation) error {
	ret := _m.Called(ctx, translation)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Translation) error); ok {
		r0 = rf(ctx, translation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindHistories provides a mock function with given fields: ctx
func (_m *TranslationRepositoryMock) FindHistories(ctx context.Context) ([]domain.Translation, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindHistories")
	}

	var r0 []domain.Translation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Translation, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Translation); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Translation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTranslationRepositoryMock creates a new instance of TranslationRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTranslationRepositoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *TranslationRepositoryMock {
	mock := &TranslationRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}