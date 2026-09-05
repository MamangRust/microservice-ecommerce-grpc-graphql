package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

// --- TransactionStatsService ---

func TestGetMonthlyAmountSuccess(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyAmountFn: func(_ context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyAmount, error) {
			if status != "success" {
				t.Fatalf("expected status=success, got %q", status)
			}
			if merchantID != 0 {
				t.Fatalf("expected merchantID=0, got %d", merchantID)
			}
			return []repository.MonthlyAmount{
				{Year: "2025", Month: "Jan", TotalCount: 100, TotalAmount: 500000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetMonthlyAmountSuccess(context.Background(), &pb.MonthAmountTransactionRequest{Year: 2025, Month: 1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalSuccess != 100 || resp.Data[0].TotalAmount != 500000 {
		t.Fatalf("unexpected data: %+v", resp.Data[0])
	}
}

func TestGetMonthlyAmountSuccess_Error(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyAmountFn: func(_ context.Context, _ int, _ int, _ string, _ int32) ([]repository.MonthlyAmount, error) {
			return nil, errors.New("clickhouse error")
		},
	}, newNopLogger())

	_, err := h.GetMonthlyAmountSuccess(context.Background(), &pb.MonthAmountTransactionRequest{Year: 2025, Month: 1})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetYearlyAmountSuccess(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		yearlyAmountFn: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error) {
			if status != "success" {
				t.Fatalf("expected status=success, got %q", status)
			}
			return []repository.YearlyAmount{
				{Year: "2025", TotalCount: 1200, TotalAmount: 6000000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetYearlyAmountSuccess(context.Background(), &pb.YearAmountTransactionRequest{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalSuccess != 1200 {
		t.Fatalf("expected TotalSuccess=1200, got %d", resp.Data[0].TotalSuccess)
	}
}

func TestGetMonthlyAmountFailed(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyAmountFn: func(_ context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyAmount, error) {
			if status != "failed" {
				t.Fatalf("expected status=failed, got %q", status)
			}
			return []repository.MonthlyAmount{
				{Year: "2025", Month: "Feb", TotalCount: 10, TotalAmount: 50000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetMonthlyAmountFailed(context.Background(), &pb.MonthAmountTransactionRequest{Year: 2025, Month: 2})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalFailed != 10 {
		t.Fatalf("expected TotalFailed=10, got %d", resp.Data[0].TotalFailed)
	}
}

func TestGetYearlyAmountFailed(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		yearlyAmountFn: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error) {
			if status != "failed" {
				t.Fatalf("expected status=failed, got %q", status)
			}
			return []repository.YearlyAmount{
				{Year: "2025", TotalCount: 120, TotalAmount: 600000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetYearlyAmountFailed(context.Background(), &pb.YearAmountTransactionRequest{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalFailed != 120 {
		t.Fatalf("expected TotalFailed=120, got %d", resp.Data[0].TotalFailed)
	}
}

func TestGetMonthlyTransactionMethodSuccess(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyMethodFn: func(_ context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyMethod, error) {
			if status != "success" {
				t.Fatalf("expected status=success, got %q", status)
			}
			return []repository.MonthlyMethod{
				{Month: "Mar", PaymentMethod: "credit_card", TotalCount: 80, TotalAmount: 400000},
				{Month: "Mar", PaymentMethod: "bank_transfer", TotalCount: 20, TotalAmount: 100000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetMonthlyTransactionMethodSuccess(context.Background(), &pb.MonthMethodTransactionRequest{Year: 2025, Month: 3})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Data) != 2 {
		t.Fatalf("expected 2 items, got %d", len(resp.Data))
	}
	if resp.Data[0].PaymentMethod != "credit_card" {
		t.Fatalf("expected payment_method=credit_card, got %q", resp.Data[0].PaymentMethod)
	}
}

func TestGetYearlyTransactionMethodSuccess(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		yearlyMethodFn: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyMethod, error) {
			if status != "success" {
				t.Fatalf("expected status=success, got %q", status)
			}
			return []repository.YearlyMethod{
				{Year: "2025", PaymentMethod: "e_wallet", TotalCount: 500, TotalAmount: 2500000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetYearlyTransactionMethodSuccess(context.Background(), &pb.YearMethodTransactionRequest{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalTransactions != 500 {
		t.Fatalf("expected TotalTransactions=500, got %d", resp.Data[0].TotalTransactions)
	}
}

func TestGetMonthlyTransactionMethodFailed(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyMethodFn: func(_ context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyMethod, error) {
			if status != "failed" {
				t.Fatalf("expected status=failed, got %q", status)
			}
			return []repository.MonthlyMethod{
				{Month: "Apr", PaymentMethod: "credit_card", TotalCount: 5, TotalAmount: 25000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetMonthlyTransactionMethodFailed(context.Background(), &pb.MonthMethodTransactionRequest{Year: 2025, Month: 4})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalTransactions != 5 {
		t.Fatalf("expected TotalTransactions=5, got %d", resp.Data[0].TotalTransactions)
	}
}

func TestGetYearlyTransactionMethodFailed(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		yearlyMethodFn: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyMethod, error) {
			if status != "failed" {
				t.Fatalf("expected status=failed, got %q", status)
			}
			return []repository.YearlyMethod{
				{Year: "2025", PaymentMethod: "bank_transfer", TotalCount: 60, TotalAmount: 300000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetYearlyTransactionMethodFailed(context.Background(), &pb.YearMethodTransactionRequest{Year: 2025})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalAmount != 300000 {
		t.Fatalf("expected TotalAmount=300000, got %d", resp.Data[0].TotalAmount)
	}
}

// --- TransactionStatsByMerchantService ---

func TestGetMonthlyAmountSuccessByMerchant(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyAmountFn: func(_ context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyAmount, error) {
			if status != "success" || merchantID != 99 {
				t.Fatalf("expected status=success merchantID=99, got status=%q merchantID=%d", status, merchantID)
			}
			return []repository.MonthlyAmount{
				{Year: "2025", Month: "May", TotalCount: 30, TotalAmount: 150000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetMonthlyAmountSuccessByMerchant(context.Background(), &pb.MonthAmountTransactionMerchantRequest{
		Year: 2025, Month: 5, MerchantId: 99,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalSuccess != 30 {
		t.Fatalf("expected TotalSuccess=30, got %d", resp.Data[0].TotalSuccess)
	}
}

func TestGetYearlyAmountSuccessByMerchant(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		yearlyAmountFn: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error) {
			if status != "success" || merchantID != 99 {
				t.Fatalf("expected status=success merchantID=99")
			}
			return []repository.YearlyAmount{{Year: "2025", TotalCount: 360, TotalAmount: 1800000}}, nil
		},
	}, newNopLogger())

	_, err := h.GetYearlyAmountSuccessByMerchant(context.Background(), &pb.YearAmountTransactionMerchantRequest{
		Year: 2025, MerchantId: 99,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetMonthlyAmountFailedByMerchant(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyAmountFn: func(_ context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyAmount, error) {
			if status != "failed" || merchantID != 99 {
				t.Fatalf("expected status=failed merchantID=99")
			}
			return []repository.MonthlyAmount{{Year: "2025", Month: "Jun", TotalCount: 3, TotalAmount: 15000}}, nil
		},
	}, newNopLogger())

	resp, err := h.GetMonthlyAmountFailedByMerchant(context.Background(), &pb.MonthAmountTransactionMerchantRequest{
		Year: 2025, Month: 6, MerchantId: 99,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalFailed != 3 {
		t.Fatalf("expected TotalFailed=3, got %d", resp.Data[0].TotalFailed)
	}
}

func TestGetYearlyAmountFailedByMerchant(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		yearlyAmountFn: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyAmount, error) {
			if status != "failed" || merchantID != 99 {
				t.Fatalf("expected status=failed merchantID=99")
			}
			return []repository.YearlyAmount{{Year: "2025", TotalCount: 36, TotalAmount: 180000}}, nil
		},
	}, newNopLogger())

	_, err := h.GetYearlyAmountFailedByMerchant(context.Background(), &pb.YearAmountTransactionMerchantRequest{
		Year: 2025, MerchantId: 99,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetMonthlyTransactionMethodByMerchantSuccess(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyMethodFn: func(_ context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyMethod, error) {
			if status != "success" || merchantID != 99 {
				t.Fatalf("expected status=success merchantID=99")
			}
			return []repository.MonthlyMethod{
				{Month: "Jul", PaymentMethod: "cod", TotalCount: 25, TotalAmount: 125000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetMonthlyTransactionMethodByMerchantSuccess(context.Background(), &pb.MonthMethodTransactionMerchantRequest{
		Year: 2025, Month: 7, MerchantId: 99,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].PaymentMethod != "cod" {
		t.Fatalf("expected PaymentMethod=cod, got %q", resp.Data[0].PaymentMethod)
	}
}

func TestGetYearlyTransactionMethodByMerchantSuccess(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		yearlyMethodFn: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyMethod, error) {
			if status != "success" || merchantID != 99 {
				t.Fatalf("expected status=success merchantID=99")
			}
			return []repository.YearlyMethod{
				{Year: "2025", PaymentMethod: "cod", TotalCount: 300, TotalAmount: 1500000},
			}, nil
		},
	}, newNopLogger())

	_, err := h.GetYearlyTransactionMethodByMerchantSuccess(context.Background(), &pb.YearMethodTransactionMerchantRequest{
		Year: 2025, MerchantId: 99,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetMonthlyTransactionMethodByMerchantFailed(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		monthlyMethodFn: func(_ context.Context, year, month int, status string, merchantID int32) ([]repository.MonthlyMethod, error) {
			if status != "failed" || merchantID != 99 {
				t.Fatalf("expected status=failed merchantID=99")
			}
			return []repository.MonthlyMethod{
				{Month: "Aug", PaymentMethod: "credit_card", TotalCount: 2, TotalAmount: 10000},
			}, nil
		},
	}, newNopLogger())

	resp, err := h.GetMonthlyTransactionMethodByMerchantFailed(context.Background(), &pb.MonthMethodTransactionMerchantRequest{
		Year: 2025, Month: 8, MerchantId: 99,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Data[0].TotalTransactions != 2 {
		t.Fatalf("expected TotalTransactions=2, got %d", resp.Data[0].TotalTransactions)
	}
}

func TestGetYearlyTransactionMethodByMerchantFailed(t *testing.T) {
	h := NewTransactionStatsHandler(&mockRepo{
		yearlyMethodFn: func(_ context.Context, year int, status string, merchantID int32) ([]repository.YearlyMethod, error) {
			if status != "failed" || merchantID != 99 {
				t.Fatalf("expected status=failed merchantID=99")
			}
			return []repository.YearlyMethod{
				{Year: "2025", PaymentMethod: "credit_card", TotalCount: 24, TotalAmount: 120000},
			}, nil
		},
	}, newNopLogger())

	_, err := h.GetYearlyTransactionMethodByMerchantFailed(context.Background(), &pb.YearMethodTransactionMerchantRequest{
		Year: 2025, MerchantId: 99,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
