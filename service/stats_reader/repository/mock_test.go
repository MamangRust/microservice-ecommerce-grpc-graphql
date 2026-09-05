package repository

import "context"

// mockRepository implements Repository for unit tests.
type mockRepository struct {
	monthlyTotalRevenueFn  func(ctx context.Context, year, month int, merchantID int32) ([]MonthlyRevenue, error)
	yearlyTotalRevenueFn   func(ctx context.Context, year int, merchantID int32) ([]YearlyRevenue, error)
	monthlyOrderStatsFn    func(ctx context.Context, year int, merchantID int32) ([]MonthlyOrder, error)
	yearlyOrderStatsFn     func(ctx context.Context, year int, merchantID int32) ([]YearlyOrder, error)
	monthlyTotalPricingFn  func(ctx context.Context, year, month int, filterField string, filterValue int32) ([]MonthlyRevenue, error)
	yearlyTotalPricingFn   func(ctx context.Context, year int, filterField string, filterValue int32) ([]YearlyRevenue, error)
	monthlyCategoryStatsFn func(ctx context.Context, year int, filterField string, filterValue int32) ([]MonthlyCategory, error)
	yearlyCategoryStatsFn  func(ctx context.Context, year int, filterField string, filterValue int32) ([]YearlyCategory, error)
	monthlyAmountFn        func(ctx context.Context, year, month int, status string, merchantID int32) ([]MonthlyAmount, error)
	yearlyAmountFn         func(ctx context.Context, year int, status string, merchantID int32) ([]YearlyAmount, error)
	monthlyMethodFn        func(ctx context.Context, year, month int, status string, merchantID int32) ([]MonthlyMethod, error)
	yearlyMethodFn         func(ctx context.Context, year int, status string, merchantID int32) ([]YearlyMethod, error)
}

func (m *mockRepository) GetMonthlyTotalRevenue(ctx context.Context, year, month int, merchantID int32) ([]MonthlyRevenue, error) {
	if m.monthlyTotalRevenueFn != nil {
		return m.monthlyTotalRevenueFn(ctx, year, month, merchantID)
	}
	return nil, nil
}

func (m *mockRepository) GetYearlyTotalRevenue(ctx context.Context, year int, merchantID int32) ([]YearlyRevenue, error) {
	if m.yearlyTotalRevenueFn != nil {
		return m.yearlyTotalRevenueFn(ctx, year, merchantID)
	}
	return nil, nil
}

func (m *mockRepository) GetMonthlyOrderStats(ctx context.Context, year int, merchantID int32) ([]MonthlyOrder, error) {
	if m.monthlyOrderStatsFn != nil {
		return m.monthlyOrderStatsFn(ctx, year, merchantID)
	}
	return nil, nil
}

func (m *mockRepository) GetYearlyOrderStats(ctx context.Context, year int, merchantID int32) ([]YearlyOrder, error) {
	if m.yearlyOrderStatsFn != nil {
		return m.yearlyOrderStatsFn(ctx, year, merchantID)
	}
	return nil, nil
}

func (m *mockRepository) GetMonthlyTotalPricing(ctx context.Context, year, month int, filterField string, filterValue int32) ([]MonthlyRevenue, error) {
	if m.monthlyTotalPricingFn != nil {
		return m.monthlyTotalPricingFn(ctx, year, month, filterField, filterValue)
	}
	return nil, nil
}

func (m *mockRepository) GetYearlyTotalPricing(ctx context.Context, year int, filterField string, filterValue int32) ([]YearlyRevenue, error) {
	if m.yearlyTotalPricingFn != nil {
		return m.yearlyTotalPricingFn(ctx, year, filterField, filterValue)
	}
	return nil, nil
}

func (m *mockRepository) GetMonthlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]MonthlyCategory, error) {
	if m.monthlyCategoryStatsFn != nil {
		return m.monthlyCategoryStatsFn(ctx, year, filterField, filterValue)
	}
	return nil, nil
}

func (m *mockRepository) GetYearlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]YearlyCategory, error) {
	if m.yearlyCategoryStatsFn != nil {
		return m.yearlyCategoryStatsFn(ctx, year, filterField, filterValue)
	}
	return nil, nil
}

func (m *mockRepository) GetMonthlyAmount(ctx context.Context, year, month int, status string, merchantID int32) ([]MonthlyAmount, error) {
	if m.monthlyAmountFn != nil {
		return m.monthlyAmountFn(ctx, year, month, status, merchantID)
	}
	return nil, nil
}

func (m *mockRepository) GetYearlyAmount(ctx context.Context, year int, status string, merchantID int32) ([]YearlyAmount, error) {
	if m.yearlyAmountFn != nil {
		return m.yearlyAmountFn(ctx, year, status, merchantID)
	}
	return nil, nil
}

func (m *mockRepository) GetMonthlyMethod(ctx context.Context, year, month int, status string, merchantID int32) ([]MonthlyMethod, error) {
	if m.monthlyMethodFn != nil {
		return m.monthlyMethodFn(ctx, year, month, status, merchantID)
	}
	return nil, nil
}

func (m *mockRepository) GetYearlyMethod(ctx context.Context, year int, status string, merchantID int32) ([]YearlyMethod, error) {
	if m.yearlyMethodFn != nil {
		return m.yearlyMethodFn(ctx, year, status, merchantID)
	}
	return nil, nil
}

var _ Repository = (*mockRepository)(nil)
