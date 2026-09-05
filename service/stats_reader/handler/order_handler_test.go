package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

// --- FindMonthlyTotalRevenue ---

func TestFindMonthlyTotalRevenue(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		monthlyTotalRevenueFn: func(_ context.Context, year, month int, merchantID int32) ([]repository.MonthlyRevenue, error) {
			if merchantID != 0 {
				t.Fatalf("expected merchantID=0, got %d", merchantID)
			}
			return []repository.MonthlyRevenue{
				{Year: "2025", Month: "Jan", TotalRevenue: 15000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthlyTotalRevenue(context.Background(), &pb.FindYearMonthTotalRevenue{Year: 2025, Month: 1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != "success" {
		t.Fatalf("expected status success, got %q", resp.Status)
	}
	if resp.Data[0].TotalRevenue != 15000 {
		t.Fatalf("expected TotalRevenue=15000, got %d", resp.Data[0].TotalRevenue)
	}
}

func TestFindMonthlyTotalRevenue_Error(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		monthlyTotalRevenueFn: func(_ context.Context, _ int, _ int, _ int32) ([]repository.MonthlyRevenue, error) {
			return nil, errors.New("clickhouse down")
		},
	}, newNopLogger())

	_, err := h.FindMonthlyTotalRevenue(context.Background(), &pb.FindYearMonthTotalRevenue{Year: 2025, Month: 1})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// --- FindYearlyTotalRevenue ---

func TestFindYearlyTotalRevenue(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		yearlyTotalRevenueFn: func(_ context.Context, year int, merchantID int32) ([]repository.YearlyRevenue, error) {
			return []repository.YearlyRevenue{
				{Year: "2025", TotalRevenue: 100000},
				{Year: "2024", TotalRevenue: 80000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindYearlyTotalRevenue(context.Background(), &pb.FindYearTotalRevenue{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Data) != 2 {
		t.Fatalf("expected 2 items, got %d", len(resp.Data))
	}
}

// --- FindMonthlyRevenue ---

func TestFindMonthlyRevenue(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		monthlyOrderStatsFn: func(_ context.Context, year int, merchantID int32) ([]repository.MonthlyOrder, error) {
			return []repository.MonthlyOrder{
				{Month: "Feb", OrderCount: 50, TotalRevenue: 25000, TotalItemsSold: 120},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthlyRevenue(context.Background(), &pb.FindYearOrder{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].OrderCount != 50 || resp.Data[0].TotalItemsSold != 120 {
		t.Fatalf("unexpected data: %+v", resp.Data[0])
	}
}

// --- FindYearlyRevenue ---

func TestFindYearlyRevenue(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		yearlyOrderStatsFn: func(_ context.Context, year int, merchantID int32) ([]repository.YearlyOrder, error) {
			return []repository.YearlyOrder{
				{Year: "2025", OrderCount: 500, TotalRevenue: 250000, TotalItemsSold: 1200, UniqueProductsSold: 80},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindYearlyRevenue(context.Background(), &pb.FindYearOrder{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].UniqueProductsSold != 80 {
		t.Fatalf("expected UniqueProductsSold=80, got %d", resp.Data[0].UniqueProductsSold)
	}
}

// --- ByMerchant variants ---

func TestFindMonthlyTotalRevenueByMerchant(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		monthlyTotalRevenueFn: func(_ context.Context, year, month int, merchantID int32) ([]repository.MonthlyRevenue, error) {
			if merchantID != 42 {
				t.Fatalf("expected merchantID=42, got %d", merchantID)
			}
			return []repository.MonthlyRevenue{{Year: "2025", Month: "Mar", TotalRevenue: 8000}}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthlyTotalRevenueByMerchant(context.Background(), &pb.FindYearMonthTotalRevenueByMerchant{
		Year:       2025,
		Month:      3,
		MerchantId: 42,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalRevenue != 8000 {
		t.Fatalf("expected TotalRevenue=8000, got %d", resp.Data[0].TotalRevenue)
	}
}

func TestFindYearlyTotalRevenueByMerchant(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		yearlyTotalRevenueFn: func(_ context.Context, year int, merchantID int32) ([]repository.YearlyRevenue, error) {
			if merchantID != 42 {
				t.Fatalf("expected merchantID=42, got %d", merchantID)
			}
			return []repository.YearlyRevenue{{Year: "2025", TotalRevenue: 96000}}, nil
		},
	}, newNopLogger())

	_, err := h.FindYearlyTotalRevenueByMerchant(context.Background(), &pb.FindYearTotalRevenueByMerchant{
		Year:       2025,
		MerchantId: 42,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFindMonthlyRevenueByMerchant(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		monthlyOrderStatsFn: func(_ context.Context, year int, merchantID int32) ([]repository.MonthlyOrder, error) {
			if merchantID != 42 {
				t.Fatalf("expected merchantID=42")
			}
			return []repository.MonthlyOrder{
				{Month: "Apr", OrderCount: 15, TotalRevenue: 7500, TotalItemsSold: 30},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindMonthlyRevenueByMerchant(context.Background(), &pb.FindYearOrderByMerchant{
		Year:       2025,
		MerchantId: 42,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].OrderCount != 15 {
		t.Fatalf("expected OrderCount=15, got %d", resp.Data[0].OrderCount)
	}
}

func TestFindYearlyRevenueByMerchant(t *testing.T) {
	h := NewOrderStatsHandler(&mockRepo{
		yearlyOrderStatsFn: func(_ context.Context, year int, merchantID int32) ([]repository.YearlyOrder, error) {
			if merchantID != 42 {
				t.Fatalf("expected merchantID=42")
			}
			return []repository.YearlyOrder{
				{Year: "2025", OrderCount: 180, TotalRevenue: 90000, TotalItemsSold: 360, UniqueProductsSold: 20},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.FindYearlyRevenueByMerchant(context.Background(), &pb.FindYearOrderByMerchant{
		Year:       2025,
		MerchantId: 42,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].UniqueProductsSold != 20 {
		t.Fatalf("expected UniqueProductsSold=20, got %d", resp.Data[0].UniqueProductsSold)
	}
}
