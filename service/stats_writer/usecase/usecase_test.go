package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
)

// mockRepo implements repository.Repository for usecase tests.
type mockRepo struct {
	insertOrderFn        func(ctx context.Context, eventID string, eventVersion uint64, event events.OrderEvent) error
	insertOrderItemFn    func(ctx context.Context, eventID string, eventVersion uint64, event events.OrderItemEvent) error
	insertTransactionFn  func(ctx context.Context, eventID string, eventVersion uint64, event events.TransactionEvent) error
	flushFn              func(ctx context.Context) error
	closeFn              func() error
}

func (m *mockRepo) InsertOrderEvent(ctx context.Context, eventID string, eventVersion uint64, event events.OrderEvent) error {
	if m.insertOrderFn != nil {
		return m.insertOrderFn(ctx, eventID, eventVersion, event)
	}
	return nil
}

func (m *mockRepo) InsertOrderItemEvent(ctx context.Context, eventID string, eventVersion uint64, event events.OrderItemEvent) error {
	if m.insertOrderItemFn != nil {
		return m.insertOrderItemFn(ctx, eventID, eventVersion, event)
	}
	return nil
}

func (m *mockRepo) InsertTransactionEvent(ctx context.Context, eventID string, eventVersion uint64, event events.TransactionEvent) error {
	if m.insertTransactionFn != nil {
		return m.insertTransactionFn(ctx, eventID, eventVersion, event)
	}
	return nil
}

func (m *mockRepo) Flush(ctx context.Context) error {
	if m.flushFn != nil {
		return m.flushFn(ctx)
	}
	return nil
}

func (m *mockRepo) Close() error {
	if m.closeFn != nil {
		return m.closeFn()
	}
	return nil
}

// --- SaveOrderEvent ---

func TestSaveOrderEvent(t *testing.T) {
	var gotID string
	var gotVersion uint64
	var gotEvent events.OrderEvent

	uc := NewStatsUseCase(&mockRepo{
		insertOrderFn: func(_ context.Context, eventID string, eventVersion uint64, event events.OrderEvent) error {
			gotID = eventID
			gotVersion = eventVersion
			gotEvent = event
			return nil
		},
	})

	event := events.OrderEvent{
		OrderID:    1,
		UserID:     2,
		MerchantID: 3,
		TotalPrice: 5000,
		Status:     "created",
		EventTime:  "2025-06-15T10:30:00Z",
	}

	err := uc.SaveOrderEvent(context.Background(), "evt-001", event)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotID != "evt-001" {
		t.Fatalf("expected eventID=evt-001, got %q", gotID)
	}
	if gotVersion == 0 {
		t.Fatal("expected non-zero eventVersion")
	}
	if gotEvent.OrderID != 1 || gotEvent.TotalPrice != 5000 {
		t.Fatalf("unexpected event: %+v", gotEvent)
	}
}

func TestSaveOrderEvent_Error(t *testing.T) {
	uc := NewStatsUseCase(&mockRepo{
		insertOrderFn: func(_ context.Context, _ string, _ uint64, _ events.OrderEvent) error {
			return errors.New("insert failed")
		},
	})

	err := uc.SaveOrderEvent(context.Background(), "evt-002", events.OrderEvent{
		EventTime: "2025-06-15T10:30:00Z",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// --- SaveOrderItemEvent ---

func TestSaveOrderItemEvent(t *testing.T) {
	var gotID string
	var gotVersion uint64
	var gotEvent events.OrderItemEvent

	uc := NewStatsUseCase(&mockRepo{
		insertOrderItemFn: func(_ context.Context, eventID string, eventVersion uint64, event events.OrderItemEvent) error {
			gotID = eventID
			gotVersion = eventVersion
			gotEvent = event
			return nil
		},
	})

	event := events.OrderItemEvent{
		OrderItemID:  10,
		OrderID:      1,
		MerchantID:   3,
		ProductID:    5,
		CategoryID:   2,
		CategoryName: "Electronics",
		Quantity:     3,
		Price:        1000,
		EventTime:    "2025-06-15T11:00:00Z",
	}

	err := uc.SaveOrderItemEvent(context.Background(), "evt-item-001", event)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotID != "evt-item-001" {
		t.Fatalf("expected eventID=evt-item-001, got %q", gotID)
	}
	if gotVersion == 0 {
		t.Fatal("expected non-zero eventVersion")
	}
	if gotEvent.CategoryName != "Electronics" {
		t.Fatalf("expected CategoryName=Electronics, got %q", gotEvent.CategoryName)
	}
}

func TestSaveOrderItemEvent_Error(t *testing.T) {
	uc := NewStatsUseCase(&mockRepo{
		insertOrderItemFn: func(_ context.Context, _ string, _ uint64, _ events.OrderItemEvent) error {
			return errors.New("insert item failed")
		},
	})

	err := uc.SaveOrderItemEvent(context.Background(), "evt-item-002", events.OrderItemEvent{
		EventTime: "2025-06-15T11:00:00Z",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// --- SaveTransactionEvent ---

func TestSaveTransactionEvent(t *testing.T) {
	var gotID string
	var gotVersion uint64
	var gotEvent events.TransactionEvent

	uc := NewStatsUseCase(&mockRepo{
		insertTransactionFn: func(_ context.Context, eventID string, eventVersion uint64, event events.TransactionEvent) error {
			gotID = eventID
			gotVersion = eventVersion
			gotEvent = event
			return nil
		},
	})

	event := events.TransactionEvent{
		TransactionID: 20,
		OrderID:       1,
		MerchantID:    3,
		PaymentMethod: "credit_card",
		Amount:        5000,
		Status:        "success",
		EventTime:     "2025-06-15T12:00:00Z",
	}

	err := uc.SaveTransactionEvent(context.Background(), "evt-tx-001", event)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotID != "evt-tx-001" {
		t.Fatalf("expected eventID=evt-tx-001, got %q", gotID)
	}
	if gotVersion == 0 {
		t.Fatal("expected non-zero eventVersion")
	}
	if gotEvent.PaymentMethod != "credit_card" || gotEvent.Status != "success" {
		t.Fatalf("unexpected event: %+v", gotEvent)
	}
}

func TestSaveTransactionEvent_Error(t *testing.T) {
	uc := NewStatsUseCase(&mockRepo{
		insertTransactionFn: func(_ context.Context, _ string, _ uint64, _ events.TransactionEvent) error {
			return errors.New("insert tx failed")
		},
	})

	err := uc.SaveTransactionEvent(context.Background(), "evt-tx-002", events.TransactionEvent{
		EventTime: "2025-06-15T12:00:00Z",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// --- Close ---

func TestClose(t *testing.T) {
	called := false
	uc := NewStatsUseCase(&mockRepo{
		closeFn: func() error {
			called = true
			return nil
		},
	})

	err := uc.Close()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !called {
		t.Fatal("expected Close to be called")
	}
}

func TestClose_Error(t *testing.T) {
	uc := NewStatsUseCase(&mockRepo{
		closeFn: func() error {
			return errors.New("close failed")
		},
	})

	err := uc.Close()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// --- eventVersion edge cases ---

func TestSaveOrderEvent_EmptyEventTime(t *testing.T) {
	var gotVersion uint64

	uc := NewStatsUseCase(&mockRepo{
		insertOrderFn: func(_ context.Context, _ string, eventVersion uint64, _ events.OrderEvent) error {
			gotVersion = eventVersion
			return nil
		},
	})

	err := uc.SaveOrderEvent(context.Background(), "evt-empty", events.OrderEvent{
		EventTime: "",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotVersion != 0 {
		t.Fatalf("expected eventVersion=0 for empty EventTime, got %d", gotVersion)
	}
}

func TestSaveOrderEvent_InvalidEventTime(t *testing.T) {
	var gotVersion uint64

	uc := NewStatsUseCase(&mockRepo{
		insertOrderFn: func(_ context.Context, _ string, eventVersion uint64, _ events.OrderEvent) error {
			gotVersion = eventVersion
			return nil
		},
	})

	err := uc.SaveOrderEvent(context.Background(), "evt-invalid", events.OrderEvent{
		EventTime: "not-a-date",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotVersion != 0 {
		t.Fatalf("expected eventVersion=0 for invalid EventTime, got %d", gotVersion)
	}
}
