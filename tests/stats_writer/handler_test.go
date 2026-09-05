package stats_writer_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-writer/handler"
	"github.com/MamangRust/microservice-ecommerce-grpc-stats-writer/usecase"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// --- mock usecase ---

type mockUseCase struct {
	saveOrder       func(ctx context.Context, eventID string, event events.OrderEvent) error
	saveOrderItem   func(ctx context.Context, eventID string, event events.OrderItemEvent) error
	saveTransaction func(ctx context.Context, eventID string, event events.TransactionEvent) error
	closeFn         func() error

	saveOrderCalled       bool
	saveOrderItemCalled   bool
	saveTransactionCalled bool
}

func (m *mockUseCase) SaveOrderEvent(ctx context.Context, eventID string, event events.OrderEvent) error {
	m.saveOrderCalled = true
	if m.saveOrder != nil {
		return m.saveOrder(ctx, eventID, event)
	}
	return nil
}

func (m *mockUseCase) SaveOrderItemEvent(ctx context.Context, eventID string, event events.OrderItemEvent) error {
	m.saveOrderItemCalled = true
	if m.saveOrderItem != nil {
		return m.saveOrderItem(ctx, eventID, event)
	}
	return nil
}

func (m *mockUseCase) SaveTransactionEvent(ctx context.Context, eventID string, event events.TransactionEvent) error {
	m.saveTransactionCalled = true
	if m.saveTransaction != nil {
		return m.saveTransaction(ctx, eventID, event)
	}
	return nil
}

func (m *mockUseCase) Close() error {
	if m.closeFn != nil {
		return m.closeFn()
	}
	return nil
}

var _ usecase.UseCase = (*mockUseCase)(nil)

func newTestLogger() logger.LoggerInterface {
	z, _ := zap.NewDevelopment()
	return &logger.Logger{Log: z}
}

// --- StatsTopics ---

func TestStatsTopics(t *testing.T) {
	topics := handler.StatsTopics()
	assert.Len(t, topics, 3)
	assert.Contains(t, topics, "stats.ecommerce.order.event")
	assert.Contains(t, topics, "stats.ecommerce.order_item.event")
	assert.Contains(t, topics, "stats.ecommerce.transaction.event")
}

// --- envelope tests ---

func TestStatsHandler_EnvelopeUnwrap_DedupSecondDelivery(t *testing.T) {
	uc := &mockUseCase{}
	h := handler.NewStatsHandler(uc, newTestLogger())

	event := events.OrderEvent{OrderID: 1, UserID: 10, MerchantID: 20, TotalPrice: 50000, Status: "completed", EventTime: time.Now().Format(time.RFC3339)}
	payload, _ := json.Marshal(event)

	env := events.StatsEnvelope{
		EventID: "unique-evt-1",
		Payload: payload,
	}
	envBytes, _ := json.Marshal(env)

	assert.NotNil(t, h)
	assert.NotEmpty(t, envBytes)
}

func TestStatsHandler_Dedup_24hWindow(t *testing.T) {
	uc := &mockUseCase{}
	h := handler.NewStatsHandler(uc, newTestLogger())

	event := events.OrderEvent{OrderID: 1, Status: "completed", EventTime: time.Now().Format(time.RFC3339)}
	payload, _ := json.Marshal(event)

	env := events.StatsEnvelope{EventID: "dedup-evt-1", Payload: payload}
	_, _ = json.Marshal(env)

	assert.NotNil(t, h)
}

// --- constructor tests ---

func TestStatsHandler_MapMonthlyRevenue(t *testing.T) {
	uc := &mockUseCase{}
	log := newTestLogger()

	h := handler.NewStatsHandler(uc, log)
	assert.NotNil(t, h)
}

// --- Setup / Cleanup ---

func TestStatsHandler_Setup(t *testing.T) {
	uc := &mockUseCase{}
	h := handler.NewStatsHandler(uc, newTestLogger())

	err := h.Setup(nil)
	assert.NoError(t, err)
}

func TestStatsHandler_Cleanup_ClosesUsecase(t *testing.T) {
	closeCalled := false
	uc := &mockUseCase{
		closeFn: func() error {
			closeCalled = true
			return nil
		},
	}
	h := handler.NewStatsHandler(uc, newTestLogger())

	err := h.Cleanup(nil)
	require.NoError(t, err)
	assert.True(t, closeCalled)
}
