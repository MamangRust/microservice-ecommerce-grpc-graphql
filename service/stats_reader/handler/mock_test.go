package handler

import (
	"context"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
)

// mockRepo implements repository.Repository for handler unit tests.
type mockRepo struct {
	monthlyTotalRevenueFn  func(ctx context.Context, year, month int, merchantID int32) ([]repository.MonthlyRevenue, error)
	yearlyTotalRevenueFn   func(ctx context.Context, year int, merchantID int32) ([]repository.YearlyRevenue, error)
	monthlyOrderStatsFn    func(ctx context.Context, year int, merchantID int32) ([]repository.MonthlyOrder, error)
	yearlyOrderStatsFn     func(ctx context.Context, year int, merchantID int32) ([]repository.YearlyOrder, error)
	monthlyTotalPricingFn  func(ctx context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error)
	yearlyTotalPricingFn   func(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyRevenue, error)
	monthlyCategoryStatsFn func(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.MonthlyCategory, error)
	yearlyCategoryStatsFn  func(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyCategory, error)
	monthlyAmountFn        func(ctx context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyAmount, error)
	yearlyAmountFn         func(ctx context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error)
	monthlyMethodFn        func(ctx context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyMethod, error)
	yearlyMethodFn         func(ctx context.Context, year int, status string, merchantID int32) ([]repository.YearlyMethod, error)
}

func (m *mockRepo) GetMonthlyTotalRevenue(ctx context.Context, year, month int, merchantID int32) ([]repository.MonthlyRevenue, error) {
	if m.monthlyTotalRevenueFn != nil {
		return m.monthlyTotalRevenueFn(ctx, year, month, merchantID)
	}
	return nil, nil
}

func (m *mockRepo) GetYearlyTotalRevenue(ctx context.Context, year int, merchantID int32) ([]repository.YearlyRevenue, error) {
	if m.yearlyTotalRevenueFn != nil {
		return m.yearlyTotalRevenueFn(ctx, year, merchantID)
	}
	return nil, nil
}

func (m *mockRepo) GetMonthlyOrderStats(ctx context.Context, year int, merchantID int32) ([]repository.MonthlyOrder, error) {
	if m.monthlyOrderStatsFn != nil {
		return m.monthlyOrderStatsFn(ctx, year, merchantID)
	}
	return nil, nil
}

func (m *mockRepo) GetYearlyOrderStats(ctx context.Context, year int, merchantID int32) ([]repository.YearlyOrder, error) {
	if m.yearlyOrderStatsFn != nil {
		return m.yearlyOrderStatsFn(ctx, year, merchantID)
	}
	return nil, nil
}

func (m *mockRepo) GetMonthlyTotalPricing(ctx context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error) {
	if m.monthlyTotalPricingFn != nil {
		return m.monthlyTotalPricingFn(ctx, year, month, filterField, filterValue)
	}
	return nil, nil
}

func (m *mockRepo) GetYearlyTotalPricing(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyRevenue, error) {
	if m.yearlyTotalPricingFn != nil {
		return m.yearlyTotalPricingFn(ctx, year, filterField, filterValue)
	}
	return nil, nil
}

func (m *mockRepo) GetMonthlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.MonthlyCategory, error) {
	if m.monthlyCategoryStatsFn != nil {
		return m.monthlyCategoryStatsFn(ctx, year, filterField, filterValue)
	}
	return nil, nil
}

func (m *mockRepo) GetYearlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyCategory, error) {
	if m.yearlyCategoryStatsFn != nil {
		return m.yearlyCategoryStatsFn(ctx, year, filterField, filterValue)
	}
	return nil, nil
}

func (m *mockRepo) GetMonthlyAmount(ctx context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyAmount, error) {
	if m.monthlyAmountFn != nil {
		return m.monthlyAmountFn(ctx, year, month, status, merchantID)
	}
	return nil, nil
}

func (m *mockRepo) GetYearlyAmount(ctx context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error) {
	if m.yearlyAmountFn != nil {
		return m.yearlyAmountFn(ctx, year, status, merchantID)
	}
	return nil, nil
}

func (m *mockRepo) GetMonthlyMethod(ctx context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyMethod, error) {
	if m.monthlyMethodFn != nil {
		return m.monthlyMethodFn(ctx, year, month, status, merchantID)
	}
	return nil, nil
}

func (m *mockRepo) GetYearlyMethod(ctx context.Context, year int, status string, merchantID int32) ([]repository.YearlyMethod, error) {
	if m.yearlyMethodFn != nil {
		return m.yearlyMethodFn(ctx, year, status, merchantID)
	}
	return nil, nil
}

var _ repository.Repository = (*mockRepo)(nil)
