package stats_writer_test

import (
	"context"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	pkgch	"github.com/MamangRust/microservice-ecommerce-pkg/clickhouse"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-grpc-stats-writer/repository"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	chcontainer "github.com/testcontainers/testcontainers-go/modules/clickhouse"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

func setupClickHouse(t *testing.T) (clickhouse.Conn, func()) {
	t.Helper()
	ctx := context.Background()

	chContainer, err := chcontainer.Run(ctx,
		"clickhouse/clickhouse-server:24.3-alpine",
		chcontainer.WithUsername("testuser"),
		chcontainer.WithPassword("testpass"),
		chcontainer.WithDatabase("testdb"),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = testcontainers.TerminateContainer(chContainer)
	})

	host, err := chContainer.ConnectionHost(ctx)
	require.NoError(t, err)

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{host},
		Auth: clickhouse.Auth{
			Database: "testdb",
			Username: "testuser",
			Password: "testpass",
		},
		DialTimeout: 10 * time.Second,
	})
	require.NoError(t, err)
	require.NoError(t, conn.Ping(ctx))

	log, _ := logger.NewLogger("test", sdklog.NewLoggerProvider())
	require.NoError(t, pkgch.ApplySchema(ctx, conn, log))

	return conn, func() { conn.Close() }
}

func TestRepo_InsertOrderEvent_And_Flush(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	log, _ := logger.NewLogger("test", sdklog.NewLoggerProvider())
	repo := repository.NewClickhouseRepository(conn, log)
	defer repo.Close()

	event := events.OrderEvent{
		OrderID:    1,
		UserID:     10,
		MerchantID: 20,
		TotalPrice: 50000,
		Status:     "completed",
		EventTime:  time.Now().UTC().Format(time.RFC3339),
	}

	err := repo.InsertOrderEvent(context.Background(), "00000000-0000-0000-0000-000000000001", 1, event)
	require.NoError(t, err)

	err = repo.Flush(context.Background())
	require.NoError(t, err)

	var count uint64
	err = conn.QueryRow(context.Background(), "SELECT count() FROM order_events").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), count)
}

func TestRepo_InsertOrderItemEvent_And_Flush(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	log, _ := logger.NewLogger("test", sdklog.NewLoggerProvider())
	repo := repository.NewClickhouseRepository(conn, log)
	defer repo.Close()

	event := events.OrderItemEvent{
		OrderItemID:  100,
		OrderID:      1,
		MerchantID:   20,
		ProductID:    5,
		CategoryID:   3,
		CategoryName: "Electronics",
		Quantity:     2,
		Price:        10000,
		EventTime:    time.Now().UTC().Format(time.RFC3339),
	}

	err := repo.InsertOrderItemEvent(context.Background(), "00000000-0000-0000-0000-000000000002", 1, event)
	require.NoError(t, err)

	err = repo.Flush(context.Background())
	require.NoError(t, err)

	var count uint64
	err = conn.QueryRow(context.Background(), "SELECT count() FROM order_item_events").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), count)
}

func TestRepo_InsertTransactionEvent_And_Flush(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	log, _ := logger.NewLogger("test", sdklog.NewLoggerProvider())
	repo := repository.NewClickhouseRepository(conn, log)
	defer repo.Close()

	event := events.TransactionEvent{
		TransactionID: 200,
		OrderID:       1,
		MerchantID:    20,
		PaymentMethod: "cash",
		Amount:        50000,
		Status:        "success",
		EventTime:     time.Now().UTC().Format(time.RFC3339),
	}

	err := repo.InsertTransactionEvent(context.Background(), "00000000-0000-0000-0000-000000000003", 1, event)
	require.NoError(t, err)

	err = repo.Flush(context.Background())
	require.NoError(t, err)

	var count uint64
	err = conn.QueryRow(context.Background(), "SELECT count() FROM transaction_events").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), count)
}

func TestRepo_MultipleInserts_Flush(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	log, _ := logger.NewLogger("test", sdklog.NewLoggerProvider())
	repo := repository.NewClickhouseRepository(conn, log)
	defer repo.Close()

	now := time.Now().UTC().Format(time.RFC3339)

	for i := int32(1); i <= 5; i++ {
		err := repo.InsertOrderEvent(context.Background(), "00000000-0000-0000-0000-000000000001", 1, events.OrderEvent{
			OrderID:    i,
			UserID:     10,
			MerchantID: 20,
			TotalPrice: 10000,
			Status:     "completed",
			EventTime:  now,
		})
		require.NoError(t, err)
	}

	err := repo.Flush(context.Background())
	require.NoError(t, err)

	var count uint64
	err = conn.QueryRow(context.Background(), "SELECT count() FROM order_events").Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, uint64(5), count)
}

func TestRepo_Flush_EmptyBatches(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	log, _ := logger.NewLogger("test", sdklog.NewLoggerProvider())
	repo := repository.NewClickhouseRepository(conn, log)
	defer repo.Close()

	err := repo.Flush(context.Background())
	assert.NoError(t, err)
}

func TestRepo_Close(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	log, _ := logger.NewLogger("test", sdklog.NewLoggerProvider())
	repo := repository.NewClickhouseRepository(conn, log)

	err := repo.Close()
	assert.NoError(t, err)
}
