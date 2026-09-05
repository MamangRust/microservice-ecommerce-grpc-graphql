package stats_reader_test

import (
	"context"
	"testing"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/handler"
	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// --- mock repo ---

type mockRepo struct {
	monthlyRevenue    func(ctx context.Context, year, month int, merchantID int32) ([]repository.MonthlyRevenue, error)
	yearlyRevenue     func(ctx context.Context, year int, merchantID int32) ([]repository.YearlyRevenue, error)
	monthlyOrder      func(ctx context.Context, year int, merchantID int32) ([]repository.MonthlyOrder, error)
	yearlyOrder       func(ctx context.Context, year int, merchantID int32) ([]repository.YearlyOrder, error)
	monthlyPricing    func(ctx context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error)
	yearlyPricing     func(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyRevenue, error)
	monthlyCategory   func(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.MonthlyCategory, error)
	yearlyCategory    func(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyCategory, error)
	monthlyAmount     func(ctx context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyAmount, error)
	yearlyAmount      func(ctx context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error)
	monthlyMethod     func(ctx context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyMethod, error)
	yearlyMethod      func(ctx context.Context, year int, status string, merchantID int32) ([]repository.YearlyMethod, error)
}

func (m *mockRepo) GetMonthlyTotalRevenue(ctx context.Context, year, month int, merchantID int32) ([]repository.MonthlyRevenue, error) {
	if m.monthlyRevenue != nil {
		return m.monthlyRevenue(ctx, year, month, merchantID)
	}
	return nil, nil
}
func (m *mockRepo) GetYearlyTotalRevenue(ctx context.Context, year int, merchantID int32) ([]repository.YearlyRevenue, error) {
	if m.yearlyRevenue != nil {
		return m.yearlyRevenue(ctx, year, merchantID)
	}
	return nil, nil
}
func (m *mockRepo) GetMonthlyOrderStats(ctx context.Context, year int, merchantID int32) ([]repository.MonthlyOrder, error) {
	if m.monthlyOrder != nil {
		return m.monthlyOrder(ctx, year, merchantID)
	}
	return nil, nil
}
func (m *mockRepo) GetYearlyOrderStats(ctx context.Context, year int, merchantID int32) ([]repository.YearlyOrder, error) {
	if m.yearlyOrder != nil {
		return m.yearlyOrder(ctx, year, merchantID)
	}
	return nil, nil
}
func (m *mockRepo) GetMonthlyTotalPricing(ctx context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error) {
	if m.monthlyPricing != nil {
		return m.monthlyPricing(ctx, year, month, filterField, filterValue)
	}
	return nil, nil
}
func (m *mockRepo) GetYearlyTotalPricing(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyRevenue, error) {
	if m.yearlyPricing != nil {
		return m.yearlyPricing(ctx, year, filterField, filterValue)
	}
	return nil, nil
}
func (m *mockRepo) GetMonthlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.MonthlyCategory, error) {
	if m.monthlyCategory != nil {
		return m.monthlyCategory(ctx, year, filterField, filterValue)
	}
	return nil, nil
}
func (m *mockRepo) GetYearlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyCategory, error) {
	if m.yearlyCategory != nil {
		return m.yearlyCategory(ctx, year, filterField, filterValue)
	}
	return nil, nil
}
func (m *mockRepo) GetMonthlyAmount(ctx context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyAmount, error) {
	if m.monthlyAmount != nil {
		return m.monthlyAmount(ctx, year, month, status, merchantID)
	}
	return nil, nil
}
func (m *mockRepo) GetYearlyAmount(ctx context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error) {
	if m.yearlyAmount != nil {
		return m.yearlyAmount(ctx, year, status, merchantID)
	}
	return nil, nil
}
func (m *mockRepo) GetMonthlyMethod(ctx context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyMethod, error) {
	if m.monthlyMethod != nil {
		return m.monthlyMethod(ctx, year, month, status, merchantID)
	}
	return nil, nil
}
func (m *mockRepo) GetYearlyMethod(ctx context.Context, year int, status string, merchantID int32) ([]repository.YearlyMethod, error) {
	if m.yearlyMethod != nil {
		return m.yearlyMethod(ctx, year, status, merchantID)
	}
	return nil, nil
}

var _ repository.Repository = (*mockRepo)(nil)

func newTestLogger() logger.LoggerInterface {
	z, _ := zap.NewDevelopment()
	return &logger.Logger{Log: z}
}

// --- OrderStatsHandler ---

func TestOrderStatsHandler_FindMonthlyTotalRevenue(t *testing.T) {
	repo := &mockRepo{
		monthlyRevenue: func(_ context.Context, year, month int, _ int32) ([]repository.MonthlyRevenue, error) {
			return []repository.MonthlyRevenue{
				{Year: "2026", Month: "Mar", TotalRevenue: 80000},
			}, nil
		},
	}
	h := handler.NewOrderStatsHandler(repo, newTestLogger())

	resp, err := h.FindMonthlyTotalRevenue(context.Background(), &pb.FindYearMonthTotalRevenue{Year: 2026, Month: 3})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(80000), resp.Data[0].TotalRevenue)
}

func TestOrderStatsHandler_FindMonthlyTotalRevenue_Error(t *testing.T) {
	repo := &mockRepo{
		monthlyRevenue: func(_ context.Context, _, _ int, _ int32) ([]repository.MonthlyRevenue, error) {
			return nil, assert.AnError
		},
	}
	h := handler.NewOrderStatsHandler(repo, newTestLogger())

	_, err := h.FindMonthlyTotalRevenue(context.Background(), &pb.FindYearMonthTotalRevenue{Year: 2026, Month: 3})
	require.Error(t, err)
}

func TestOrderStatsHandler_FindYearlyTotalRevenue(t *testing.T) {
	repo := &mockRepo{
		yearlyRevenue: func(_ context.Context, year int, _ int32) ([]repository.YearlyRevenue, error) {
			return []repository.YearlyRevenue{
				{Year: "2026", TotalRevenue: 500000},
				{Year: "2025", TotalRevenue: 300000},
			}, nil
		},
	}
	h := handler.NewOrderStatsHandler(repo, newTestLogger())

	resp, err := h.FindYearlyTotalRevenue(context.Background(), &pb.FindYearTotalRevenue{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 2)
}

func TestOrderStatsHandler_FindYearlyTotalRevenue_Empty(t *testing.T) {
	repo := &mockRepo{
		yearlyRevenue: func(_ context.Context, _ int, _ int32) ([]repository.YearlyRevenue, error) {
			return nil, nil
		},
	}
	h := handler.NewOrderStatsHandler(repo, newTestLogger())

	resp, err := h.FindYearlyTotalRevenue(context.Background(), &pb.FindYearTotalRevenue{Year: 2099})
	require.NoError(t, err)
	assert.Empty(t, resp.Data)
}

func TestOrderStatsHandler_FindMonthlyRevenue(t *testing.T) {
	repo := &mockRepo{
		monthlyOrder: func(_ context.Context, year int, _ int32) ([]repository.MonthlyOrder, error) {
			return []repository.MonthlyOrder{
				{Month: "Mar", OrderCount: 10, TotalRevenue: 100000, TotalItemsSold: 25},
			}, nil
		},
	}
	h := handler.NewOrderStatsHandler(repo, newTestLogger())

	resp, err := h.FindMonthlyRevenue(context.Background(), &pb.FindYearOrder{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(10), resp.Data[0].OrderCount)
	assert.Equal(t, int32(100000), resp.Data[0].TotalRevenue)
}

func TestOrderStatsHandler_FindYearlyRevenue(t *testing.T) {
	repo := &mockRepo{
		yearlyOrder: func(_ context.Context, year int, _ int32) ([]repository.YearlyOrder, error) {
			return []repository.YearlyOrder{
				{Year: "2026", OrderCount: 100, TotalRevenue: 500000, TotalItemsSold: 200, UniqueProductsSold: 15},
			}, nil
		},
	}
	h := handler.NewOrderStatsHandler(repo, newTestLogger())

	resp, err := h.FindYearlyRevenue(context.Background(), &pb.FindYearOrder{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(15), resp.Data[0].UniqueProductsSold)
}

func TestOrderStatsHandler_FindMonthlyTotalRevenueByMerchant(t *testing.T) {
	repo := &mockRepo{
		monthlyRevenue: func(_ context.Context, year, month int, merchantID int32) ([]repository.MonthlyRevenue, error) {
			return []repository.MonthlyRevenue{
				{Year: "2026", Month: "Apr", TotalRevenue: 200000},
			}, nil
		},
	}
	h := handler.NewOrderStatsHandler(repo, newTestLogger())

	resp, err := h.FindMonthlyTotalRevenueByMerchant(context.Background(), &pb.FindYearMonthTotalRevenueByMerchant{Year: 2026, Month: 4, MerchantId: 10})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(200000), resp.Data[0].TotalRevenue)
}

// --- CategoryStatsHandler ---

func TestCategoryStatsHandler_FindMonthlyTotalPrices(t *testing.T) {
	repo := &mockRepo{
		monthlyPricing: func(_ context.Context, _, _ int, _ string, _ int32) ([]repository.MonthlyRevenue, error) {
			return []repository.MonthlyRevenue{
				{Year: "2026", Month: "Feb", TotalRevenue: 50000},
			}, nil
		},
	}
	h := handler.NewCategoryStatsHandler(repo, newTestLogger())

	resp, err := h.FindMonthlyTotalPrices(context.Background(), &pb.FindYearMonthTotalPrices{Year: 2026, Month: 2})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(50000), resp.Data[0].TotalRevenue)
}

func TestCategoryStatsHandler_FindMonthlyTotalPrices_Error(t *testing.T) {
	repo := &mockRepo{
		monthlyPricing: func(_ context.Context, _, _ int, _ string, _ int32) ([]repository.MonthlyRevenue, error) {
			return nil, assert.AnError
		},
	}
	h := handler.NewCategoryStatsHandler(repo, newTestLogger())

	_, err := h.FindMonthlyTotalPrices(context.Background(), &pb.FindYearMonthTotalPrices{Year: 2026, Month: 2})
	require.Error(t, err)
}

func TestCategoryStatsHandler_FindYearlyTotalPrices(t *testing.T) {
	repo := &mockRepo{
		yearlyPricing: func(_ context.Context, year int, _ string, _ int32) ([]repository.YearlyRevenue, error) {
			return []repository.YearlyRevenue{
				{Year: "2026", TotalRevenue: 600000},
			}, nil
		},
	}
	h := handler.NewCategoryStatsHandler(repo, newTestLogger())

	resp, err := h.FindYearlyTotalPrices(context.Background(), &pb.FindYearTotalPrices{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(600000), resp.Data[0].TotalRevenue)
}

func TestCategoryStatsHandler_FindMonthPrice(t *testing.T) {
	repo := &mockRepo{
		monthlyCategory: func(_ context.Context, year int, _ string, _ int32) ([]repository.MonthlyCategory, error) {
			return []repository.MonthlyCategory{
				{Month: "Jan", CategoryID: 1, CategoryName: "Electronics", OrderCount: 20, ItemsSold: 50, TotalRevenue: 250000},
			}, nil
		},
	}
	h := handler.NewCategoryStatsHandler(repo, newTestLogger())

	resp, err := h.FindMonthPrice(context.Background(), &pb.FindYearCategory{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, "Electronics", resp.Data[0].CategoryName)
	assert.Equal(t, int32(50), resp.Data[0].ItemsSold)
}

func TestCategoryStatsHandler_FindYearPrice(t *testing.T) {
	repo := &mockRepo{
		yearlyCategory: func(_ context.Context, year int, _ string, _ int32) ([]repository.YearlyCategory, error) {
			return []repository.YearlyCategory{
				{Year: "2026", CategoryID: 1, CategoryName: "Fashion", OrderCount: 100, ItemsSold: 200, TotalRevenue: 1000000, UniqueProductsSold: 30},
			}, nil
		},
	}
	h := handler.NewCategoryStatsHandler(repo, newTestLogger())

	resp, err := h.FindYearPrice(context.Background(), &pb.FindYearCategory{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(30), resp.Data[0].UniqueProductsSold)
}

func TestCategoryStatsHandler_FindMonthlyTotalPricesById(t *testing.T) {
	repo := &mockRepo{
		monthlyPricing: func(_ context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error) {
			assert.Equal(t, "category_id", filterField)
			assert.Equal(t, int32(5), filterValue)
			return []repository.MonthlyRevenue{
				{Year: "2026", Month: "May", TotalRevenue: 30000},
			}, nil
		},
	}
	h := handler.NewCategoryStatsHandler(repo, newTestLogger())

	resp, err := h.FindMonthlyTotalPricesById(context.Background(), &pb.FindYearMonthTotalPriceById{Year: 2026, Month: 5, CategoryId: 5})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
}

func TestCategoryStatsHandler_FindMonthlyTotalPricesByMerchant(t *testing.T) {
	repo := &mockRepo{
		monthlyPricing: func(_ context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error) {
			assert.Equal(t, "merchant_id", filterField)
			assert.Equal(t, int32(10), filterValue)
			return []repository.MonthlyRevenue{
				{Year: "2026", Month: "Jun", TotalRevenue: 70000},
			}, nil
		},
	}
	h := handler.NewCategoryStatsHandler(repo, newTestLogger())

	resp, err := h.FindMonthlyTotalPricesByMerchant(context.Background(), &pb.FindYearMonthTotalPriceByMerchant{Year: 2026, Month: 6, MerchantId: 10})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
}

// --- TransactionStatsHandler ---

func TestTransactionStatsHandler_GetMonthlyAmountSuccess(t *testing.T) {
	repo := &mockRepo{
		monthlyAmount: func(_ context.Context, _, _ int, status string, _ int32) ([]repository.MonthlyAmount, error) {
			assert.Equal(t, "success", status)
			return []repository.MonthlyAmount{
				{Year: "2026", Month: "Jul", TotalCount: 30, TotalAmount: 1500000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetMonthlyAmountSuccess(context.Background(), &pb.MonthAmountTransactionRequest{Year: 2026, Month: 7})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(30), resp.Data[0].TotalSuccess)
	assert.Equal(t, int32(1500000), resp.Data[0].TotalAmount)
}

func TestTransactionStatsHandler_GetMonthlyAmountSuccess_Error(t *testing.T) {
	repo := &mockRepo{
		monthlyAmount: func(_ context.Context, _, _ int, _ string, _ int32) ([]repository.MonthlyAmount, error) {
			return nil, assert.AnError
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	_, err := h.GetMonthlyAmountSuccess(context.Background(), &pb.MonthAmountTransactionRequest{Year: 2026, Month: 7})
	require.Error(t, err)
}

func TestTransactionStatsHandler_GetYearlyAmountSuccess(t *testing.T) {
	repo := &mockRepo{
		yearlyAmount: func(_ context.Context, year int, status string, _ int32) ([]repository.YearlyAmount, error) {
			assert.Equal(t, "success", status)
			return []repository.YearlyAmount{
				{Year: "2026", TotalCount: 300, TotalAmount: 15000000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetYearlyAmountSuccess(context.Background(), &pb.YearAmountTransactionRequest{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(300), resp.Data[0].TotalSuccess)
}

func TestTransactionStatsHandler_GetMonthlyAmountFailed(t *testing.T) {
	repo := &mockRepo{
		monthlyAmount: func(_ context.Context, _, _ int, status string, _ int32) ([]repository.MonthlyAmount, error) {
			assert.Equal(t, "failed", status)
			return []repository.MonthlyAmount{
				{Year: "2026", Month: "Aug", TotalCount: 5, TotalAmount: 25000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetMonthlyAmountFailed(context.Background(), &pb.MonthAmountTransactionRequest{Year: 2026, Month: 8})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(5), resp.Data[0].TotalFailed)
}

func TestTransactionStatsHandler_GetYearlyAmountFailed(t *testing.T) {
	repo := &mockRepo{
		yearlyAmount: func(_ context.Context, year int, status string, _ int32) ([]repository.YearlyAmount, error) {
			assert.Equal(t, "failed", status)
			return []repository.YearlyAmount{
				{Year: "2026", TotalCount: 50, TotalAmount: 250000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetYearlyAmountFailed(context.Background(), &pb.YearAmountTransactionRequest{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(50), resp.Data[0].TotalFailed)
}

func TestTransactionStatsHandler_GetMonthlyTransactionMethodSuccess(t *testing.T) {
	repo := &mockRepo{
		monthlyMethod: func(_ context.Context, _, _ int, status string, _ int32) ([]repository.MonthlyMethod, error) {
			assert.Equal(t, "success", status)
			return []repository.MonthlyMethod{
				{Month: "Sep", PaymentMethod: "credit_card", TotalCount: 20, TotalAmount: 800000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetMonthlyTransactionMethodSuccess(context.Background(), &pb.MonthMethodTransactionRequest{Year: 2026, Month: 9})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, "credit_card", resp.Data[0].PaymentMethod)
	assert.Equal(t, int32(20), resp.Data[0].TotalTransactions)
}

func TestTransactionStatsHandler_GetYearlyTransactionMethodSuccess(t *testing.T) {
	repo := &mockRepo{
		yearlyMethod: func(_ context.Context, year int, status string, _ int32) ([]repository.YearlyMethod, error) {
			assert.Equal(t, "success", status)
			return []repository.YearlyMethod{
				{Year: "2026", PaymentMethod: "bank_transfer", TotalCount: 100, TotalAmount: 5000000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetYearlyTransactionMethodSuccess(context.Background(), &pb.YearMethodTransactionRequest{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, "bank_transfer", resp.Data[0].PaymentMethod)
}

func TestTransactionStatsHandler_GetMonthlyTransactionMethodFailed(t *testing.T) {
	repo := &mockRepo{
		monthlyMethod: func(_ context.Context, _, _ int, status string, _ int32) ([]repository.MonthlyMethod, error) {
			assert.Equal(t, "failed", status)
			return []repository.MonthlyMethod{
				{Month: "Oct", PaymentMethod: "ewallet", TotalCount: 3, TotalAmount: 15000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetMonthlyTransactionMethodFailed(context.Background(), &pb.MonthMethodTransactionRequest{Year: 2026, Month: 10})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
}

func TestTransactionStatsHandler_GetYearlyTransactionMethodFailed(t *testing.T) {
	repo := &mockRepo{
		yearlyMethod: func(_ context.Context, year int, status string, _ int32) ([]repository.YearlyMethod, error) {
			assert.Equal(t, "failed", status)
			return []repository.YearlyMethod{
				{Year: "2026", PaymentMethod: "cod", TotalCount: 10, TotalAmount: 50000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetYearlyTransactionMethodFailed(context.Background(), &pb.YearMethodTransactionRequest{Year: 2026})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
}

// --- ByMerchant transaction tests ---

func TestTransactionStatsHandler_GetMonthlyAmountSuccessByMerchant(t *testing.T) {
	repo := &mockRepo{
		monthlyAmount: func(_ context.Context, _, _ int, status string, merchantID int32) ([]repository.MonthlyAmount, error) {
			assert.Equal(t, "success", status)
			assert.Equal(t, int32(10), merchantID)
			return []repository.MonthlyAmount{
				{Year: "2026", Month: "Nov", TotalCount: 15, TotalAmount: 750000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetMonthlyAmountSuccessByMerchant(context.Background(), &pb.MonthAmountTransactionMerchantRequest{Year: 2026, Month: 11, MerchantId: 10})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
	assert.Equal(t, int32(15), resp.Data[0].TotalSuccess)
}

func TestTransactionStatsHandler_GetYearlyAmountSuccessByMerchant(t *testing.T) {
	repo := &mockRepo{
		yearlyAmount: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error) {
			assert.Equal(t, "success", status)
			assert.Equal(t, int32(20), merchantID)
			return []repository.YearlyAmount{
				{Year: "2026", TotalCount: 150, TotalAmount: 7500000},
			}, nil
		},
	}
	h := handler.NewTransactionStatsHandler(repo, newTestLogger())

	resp, err := h.GetYearlyAmountSuccessByMerchant(context.Background(), &pb.YearAmountTransactionMerchantRequest{Year: 2026, MerchantId: 20})
	require.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	require.Len(t, resp.Data, 1)
}
