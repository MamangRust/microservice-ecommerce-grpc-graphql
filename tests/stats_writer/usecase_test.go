package stats_writer_test

import (
	"context"
	"testing"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-writer/repository"
	"github.com/MamangRust/microservice-ecommerce-grpc-stats-writer/usecase"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- mock repo ---

type mockRepo struct {
	insertOrder       func(ctx context.Context, eventID string, eventVersion uint64, event events.OrderEvent) error
	insertOrderItem   func(ctx context.Context, eventID string, eventVersion uint64, event events.OrderItemEvent) error
	insertTransaction func(ctx context.Context, eventID string, eventVersion uint64, event events.TransactionEvent) error
	flush             func(ctx context.Context) error
	close             func() error

	insertOrderCalled       bool
	insertOrderItemCalled   bool
	insertTransactionCalled bool
	flushCalled             bool
	closeCalled             bool
}

func (m *mockRepo) InsertOrderEvent(ctx context.Context, eventID string, eventVersion uint64, event events.OrderEvent) error {
	m.insertOrderCalled = true
	if m.insertOrder != nil {
		return m.insertOrder(ctx, eventID, eventVersion, event)
	}
	return nil
}

func (m *mockRepo) InsertOrderItemEvent(ctx context.Context, eventID string, eventVersion uint64, event events.OrderItemEvent) error {
	m.insertOrderItemCalled = true
	if m.insertOrderItem != nil {
		return m.insertOrderItem(ctx, eventID, eventVersion, event)
	}
	return nil
}

func (m *mockRepo) InsertTransactionEvent(ctx context.Context, eventID string, eventVersion uint64, event events.TransactionEvent) error {
	m.insertTransactionCalled = true
	if m.insertTransaction != nil {
		return m.insertTransaction(ctx, eventID, eventVersion, event)
	}
	return nil
}

func (m *mockRepo) Flush(ctx context.Context) error {
	m.flushCalled = true
	if m.flush != nil {
		return m.flush(ctx)
	}
	return nil
}

func (m *mockRepo) Close() error {
	m.closeCalled = true
	if m.close != nil {
		return m.close()
	}
	return nil
}

func (m *mockRepo) VerifyNoError(t *testing.T) {
	t.Helper()
	assert.NoError(t, m.Flush(context.Background()))
}

var _ repository.Repository = (*mockRepo)(nil)

// --- tests ---

func TestUseCase_SaveOrderEvent(t *testing.T) {
	repo := &mockRepo{}
	uc := usecase.NewStatsUseCase(repo)

	event := events.OrderEvent{
		OrderID:    1,
		UserID:     10,
		MerchantID: 20,
		TotalPrice: 50000,
		Status:     "completed",
		EventTime:  "2026-01-15T10:30:00Z",
	}

	err := uc.SaveOrderEvent(context.Background(), "evt-order-1", event)
	require.NoError(t, err)
	assert.True(t, repo.insertOrderCalled)
}

func TestUseCase_SaveOrderEvent_RepoError(t *testing.T) {
	repo := &mockRepo{
		insertOrder: func(_ context.Context, _ string, _ uint64, _ events.OrderEvent) error {
			return assert.AnError
		},
	}
	uc := usecase.NewStatsUseCase(repo)

	err := uc.SaveOrderEvent(context.Background(), "evt-1", events.OrderEvent{OrderID: 1})
	require.Error(t, err)
	assert.True(t, repo.insertOrderCalled)
}

func TestUseCase_SaveOrderItemEvent(t *testing.T) {
	repo := &mockRepo{}
	uc := usecase.NewStatsUseCase(repo)

	event := events.OrderItemEvent{
		OrderItemID:  100,
		OrderID:      1,
		MerchantID:   20,
		ProductID:    5,
		CategoryID:   3,
		CategoryName: "Electronics",
		Quantity:     2,
		Price:        10000,
		EventTime:    "2026-01-15T10:30:00Z",
	}

	err := uc.SaveOrderItemEvent(context.Background(), "evt-item-100", event)
	require.NoError(t, err)
	assert.True(t, repo.insertOrderItemCalled)
}

func TestUseCase_SaveTransactionEvent(t *testing.T) {
	repo := &mockRepo{}
	uc := usecase.NewStatsUseCase(repo)

	event := events.TransactionEvent{
		TransactionID: 200,
		OrderID:       1,
		MerchantID:    20,
		PaymentMethod: "cash",
		Amount:        50000,
		Status:        "success",
		EventTime:     "2026-01-15T10:31:00Z",
	}

	err := uc.SaveTransactionEvent(context.Background(), "evt-tx-200", event)
	require.NoError(t, err)
	assert.True(t, repo.insertTransactionCalled)
}

func TestUseCase_Close(t *testing.T) {
	repo := &mockRepo{}
	uc := usecase.NewStatsUseCase(repo)

	err := uc.Close()
	require.NoError(t, err)
	assert.True(t, repo.closeCalled)
}

func TestUseCase_Close_Error(t *testing.T) {
	repo := &mockRepo{
		close: func() error { return assert.AnError },
	}
	uc := usecase.NewStatsUseCase(repo)

	err := uc.Close()
	require.Error(t, err)
}

func TestUseCase_EmptyEventTime_YieldsZeroVersion(t *testing.T) {
	var capturedVersion uint64
	repo := &mockRepo{
		insertOrder: func(_ context.Context, _ string, eventVersion uint64, _ events.OrderEvent) error {
			capturedVersion = eventVersion
			return nil
		},
	}
	uc := usecase.NewStatsUseCase(repo)

	_ = uc.SaveOrderEvent(context.Background(), "evt-1", events.OrderEvent{
		OrderID:   1,
		EventTime: "",
	})
	assert.Equal(t, uint64(0), capturedVersion)
}

func TestUseCase_InvalidEventTime_YieldsZeroVersion(t *testing.T) {
	var capturedVersion uint64
	repo := &mockRepo{
		insertOrder: func(_ context.Context, _ string, eventVersion uint64, _ events.OrderEvent) error {
			capturedVersion = eventVersion
			return nil
		},
	}
	uc := usecase.NewStatsUseCase(repo)

	_ = uc.SaveOrderEvent(context.Background(), "evt-1", events.OrderEvent{
		OrderID:   1,
		EventTime: "not-a-date",
	})
	assert.Equal(t, uint64(0), capturedVersion)
}

func TestUseCase_ValidEventTime_YieldsUnixVersion(t *testing.T) {
	var capturedVersion uint64
	repo := &mockRepo{
		insertOrder: func(_ context.Context, _ string, eventVersion uint64, _ events.OrderEvent) error {
			capturedVersion = eventVersion
			return nil
		},
	}
	uc := usecase.NewStatsUseCase(repo)

	_ = uc.SaveOrderEvent(context.Background(), "evt-1", events.OrderEvent{
		OrderID:   1,
		EventTime: "2026-01-15T10:30:00Z",
	})
	assert.Greater(t, capturedVersion, uint64(0))
}
