package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"go.uber.org/zap"
)

func newNopLogger() logger.LoggerInterface {
	return &logger.Logger{Log: zap.NewNop()}
}

// --- CategoryStatsService ---

func TestFindMonthlyTotalPrices(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		monthlyTotalPricingFn: func(_ context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error) {
			if filterField != "" || filterValue != 0 {
				t.Fatalf("expected empty filter, got field=%q value=%d", filterField, filterValue)
			}
			if year != 2025 || month != 3 {
				t.Fatalf("expected year=2025 month=3, got year=%d month=%d", year, month)
			}
			return []repository.MonthlyRevenue{
				{Year: "2025", Month: "Mar", TotalRevenue: 5000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthlyTotalPrices(context.Background(), &pb.FindYearMonthTotalPrices{Year: 2025, Month: 3})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "success" {
		t.Fatalf("expected status success, got %q", resp.Status)
	}
	if len(resp.Data) != 1 || resp.Data[0].TotalRevenue != 5000 {
		t.Fatalf("unexpected data: %+v", resp.Data)
	}
}

func TestFindMonthlyTotalPrices_Error(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		monthlyTotalPricingFn: func(_ context.Context, _, _ int, _ string, _ int32) ([]repository.MonthlyRevenue, error) {
			return nil, errors.New("db error")
		},
	}, newNopLogger())

	_, err := h.FindMonthlyTotalPrices(context.Background(), &pb.FindYearMonthTotalPrices{Year: 2025, Month: 1})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestFindYearlyTotalPrices(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		yearlyTotalPricingFn: func(_ context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyRevenue, error) {
			if filterField != "" || filterValue != 0 {
				t.Fatalf("expected empty filter, got field=%q value=%d", filterField, filterValue)
			}
			return []repository.YearlyRevenue{
				{Year: "2025", TotalRevenue: 10000},
				{Year: "2024", TotalRevenue: 8000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindYearlyTotalPrices(context.Background(), &pb.FindYearTotalPrices{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Data) != 2 {
		t.Fatalf("expected 2 items, got %d", len(resp.Data))
	}
}

func TestFindMonthPrice(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		monthlyCategoryStatsFn: func(_ context.Context, year int, filterField string, filterValue int32) ([]repository.MonthlyCategory, error) {
			if filterField != "" || filterValue != 0 {
				t.Fatalf("expected empty filter")
			}
			return []repository.MonthlyCategory{
				{Month: "Jan", CategoryID: 1, CategoryName: "Electronics", OrderCount: 10, ItemsSold: 25, TotalRevenue: 50000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthPrice(context.Background(), &pb.FindYearCategory{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Data) != 1 || resp.Data[0].CategoryName != "Electronics" {
		t.Fatalf("unexpected data: %+v", resp.Data)
	}
}

func TestFindYearPrice(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		yearlyCategoryStatsFn: func(_ context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyCategory, error) {
			return []repository.YearlyCategory{
				{Year: "2025", CategoryID: 2, CategoryName: "Books", OrderCount: 50, ItemsSold: 120, TotalRevenue: 200000, UniqueProductsSold: 30},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindYearPrice(context.Background(), &pb.FindYearCategory{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].UniqueProductsSold != 30 {
		t.Fatalf("expected UniqueProductsSold=30, got %d", resp.Data[0].UniqueProductsSold)
	}
}

// --- CategoryStatsByIdService ---

func TestFindMonthlyTotalPricesById(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		monthlyTotalPricingFn: func(_ context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error) {
			if filterField != "category_id" || filterValue != 5 {
				t.Fatalf("expected category_id=5, got field=%q value=%d", filterField, filterValue)
			}
			return []repository.MonthlyRevenue{{Year: "2025", Month: "Jun", TotalRevenue: 3000}}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthlyTotalPricesById(context.Background(), &pb.FindYearMonthTotalPriceById{
		Year:       2025,
		Month:      6,
		CategoryId: 5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalRevenue != 3000 {
		t.Fatalf("expected TotalRevenue=3000, got %d", resp.Data[0].TotalRevenue)
	}
}

func TestFindYearlyTotalPricesById(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		yearlyTotalPricingFn: func(_ context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyRevenue, error) {
			if filterField != "category_id" || filterValue != 3 {
				t.Fatalf("expected category_id=3, got field=%q value=%d", filterField, filterValue)
			}
			return []repository.YearlyRevenue{{Year: "2025", TotalRevenue: 12000}}, nil
		},
	}, newNopLogger())

	_, err := h.FindYearlyTotalPricesById(context.Background(), &pb.FindYearTotalPriceById{
		Year:       2025,
		CategoryId: 3,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFindMonthPriceById(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		monthlyCategoryStatsFn: func(_ context.Context, year int, filterField string, filterValue int32) ([]repository.MonthlyCategory, error) {
			if filterField != "category_id" {
				t.Fatalf("expected filterField=category_id, got %q", filterField)
			}
			return []repository.MonthlyCategory{
				{Month: "Apr", CategoryID: uint64(filterValue), CategoryName: "Food", OrderCount: 20, ItemsSold: 40, TotalRevenue: 8000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthPriceById(context.Background(), &pb.FindYearCategoryById{
		Year:       2025,
		CategoryId: 7,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].CategoryId != 7 {
		t.Fatalf("expected CategoryId=7, got %d", resp.Data[0].CategoryId)
	}
}

func TestFindYearPriceById(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		yearlyCategoryStatsFn: func(_ context.Context, _ int, filterField string, filterValue int32) ([]repository.YearlyCategory, error) {
			if filterField != "category_id" {
				t.Fatalf("expected filterField=category_id, got %q", filterField)
			}
			return []repository.YearlyCategory{
				{Year: "2025", CategoryID: uint64(filterValue), CategoryName: "Food", OrderCount: 100, ItemsSold: 200, TotalRevenue: 40000, UniqueProductsSold: 15},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindYearPriceById(context.Background(), &pb.FindYearCategoryById{
		Year:       2025,
		CategoryId: 7,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].CategoryName != "Food" {
		t.Fatalf("expected CategoryName=Food, got %q", resp.Data[0].CategoryName)
	}
}

// --- CategoryStatsByMerchantService ---

func TestFindMonthlyTotalPricesByMerchant(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		monthlyTotalPricingFn: func(_ context.Context, year, month int, filterField string, filterValue int32) ([]repository.MonthlyRevenue, error) {
			if filterField != "merchant_id" || filterValue != 10 {
				t.Fatalf("expected merchant_id=10, got field=%q value=%d", filterField, filterValue)
			}
			return []repository.MonthlyRevenue{{Year: "2025", Month: "May", TotalRevenue: 7000}}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthlyTotalPricesByMerchant(context.Background(), &pb.FindYearMonthTotalPriceByMerchant{
		Year:       2025,
		Month:      5,
		MerchantId: 10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalRevenue != 7000 {
		t.Fatalf("expected TotalRevenue=7000, got %d", resp.Data[0].TotalRevenue)
	}
}

func TestFindYearlyTotalPricesByMerchant(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		yearlyTotalPricingFn: func(_ context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyRevenue, error) {
			if filterField != "merchant_id" || filterValue != 10 {
				t.Fatalf("expected merchant_id=10")
			}
			return []repository.YearlyRevenue{{Year: "2025", TotalRevenue: 50000}}, nil
		},
	}, newNopLogger())

	resp, err := h.FindYearlyTotalPricesByMerchant(context.Background(), &pb.FindYearTotalPriceByMerchant{
		Year:       2025,
		MerchantId: 10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalRevenue != 50000 {
		t.Fatalf("expected TotalRevenue=50000, got %d", resp.Data[0].TotalRevenue)
	}
}

func TestFindMonthPriceByMerchant(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		monthlyCategoryStatsFn: func(_ context.Context, year int, filterField string, filterValue int32) ([]repository.MonthlyCategory, error) {
			if filterField != "merchant_id" || filterValue != 10 {
				t.Fatalf("expected merchant_id=10")
			}
			return []repository.MonthlyCategory{
				{Month: "Jul", CategoryID: 1, CategoryName: "Toys", OrderCount: 5, ItemsSold: 10, TotalRevenue: 2000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthPriceByMerchant(context.Background(), &pb.FindYearCategoryByMerchant{
		Year:       2025,
		MerchantId: 10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].CategoryName != "Toys" {
		t.Fatalf("expected CategoryName=Toys, got %q", resp.Data[0].CategoryName)
	}
}

func TestFindYearPriceByMerchant(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		yearlyCategoryStatsFn: func(_ context.Context, year int, filterField string, filterValue int32) ([]repository.YearlyCategory, error) {
			if filterField != "merchant_id" || filterValue != 10 {
				t.Fatalf("expected merchant_id=10")
			}
			return []repository.YearlyCategory{
				{Year: "2025", CategoryID: 1, CategoryName: "Toys", OrderCount: 60, ItemsSold: 120, TotalRevenue: 24000, UniqueProductsSold: 8},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindYearPriceByMerchant(context.Background(), &pb.FindYearCategoryByMerchant{
		Year:       2025,
		MerchantId: 10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].UniqueProductsSold != 8 {
		t.Fatalf("expected UniqueProductsSold=8, got %d", resp.Data[0].UniqueProductsSold)
	}
}

// --- Nil data handling ---

func TestFindMonthlyTotalPrices_EmptyData(t *testing.T) {
	h := NewCategoryStatsHandler(&mockRepo{
		monthlyTotalPricingFn: func(_ context.Context, _ int, _ int, _ string, _ int32) ([]repository.MonthlyRevenue, error) {
			return nil, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthlyTotalPrices(context.Background(), &pb.FindYearMonthTotalPrices{Year: 2025, Month: 1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Data) != 0 {
		t.Fatalf("expected empty data, got %d items", len(resp.Data))
	}
}
