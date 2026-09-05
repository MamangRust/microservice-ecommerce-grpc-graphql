package stats_reader_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	chdriver "github.com/ClickHouse/clickhouse-go/v2"
	pkgch	"github.com/MamangRust/microservice-ecommerce-pkg/clickhouse"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	chcontainer "github.com/testcontainers/testcontainers-go/modules/clickhouse"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

func setupClickHouse(t *testing.T) (chdriver.Conn, func()) {
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

	conn, err := chdriver.Open(&chdriver.Options{
		Addr: []string{host},
		Auth: chdriver.Auth{
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

func seedOrderEvents(t *testing.T, conn chdriver.Conn, createdAt time.Time, orderID, userID, merchantID uint64, totalPrice int64) {
	t.Helper()
	ctx := context.Background()
	err := conn.Exec(ctx, `INSERT INTO order_events (event_id, order_id, user_id, merchant_id, total_price, created_at, event_version) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		fmt.Sprintf("00000000-0000-0000-0000-%012d", orderID), orderID, userID, merchantID, totalPrice, createdAt, uint64(time.Now().Unix()))
	require.NoError(t, err)
}

func seedOrderItemEvents(t *testing.T, conn chdriver.Conn, createdAt time.Time, orderItemID, orderID, merchantID, productID, categoryID uint64, categoryName string, quantity uint32, price int64) {
	t.Helper()
	ctx := context.Background()
	err := conn.Exec(ctx, `INSERT INTO order_item_events (event_id, order_item_id, order_id, merchant_id, product_id, category_id, category_name, quantity, price, created_at, event_version) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		fmt.Sprintf("00000000-0000-0000-0000-%012d", orderItemID), orderItemID, orderID, merchantID, productID, categoryID, categoryName, quantity, price, createdAt, uint64(time.Now().Unix()))
	require.NoError(t, err)
}

func seedTransactionEvents(t *testing.T, conn chdriver.Conn, createdAt time.Time, txID, orderID, merchantID uint64, paymentMethod, status string, amount int64) {
	t.Helper()
	ctx := context.Background()
	err := conn.Exec(ctx, `INSERT INTO transaction_events (event_id, transaction_id, order_id, merchant_id, payment_method, amount, status, created_at, event_version) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		fmt.Sprintf("00000000-0000-0000-0000-%012d", txID), txID, orderID, merchantID, paymentMethod, amount, status, createdAt, uint64(time.Now().Unix()))
	require.NoError(t, err)
}

// --- Order Stats ---

func TestRepo_GetMonthlyTotalRevenue(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 3, 15, 10, 0, 0, 0, time.UTC)
	seedOrderEvents(t, conn, createdAt, 1, 100, 200, 50000)
	seedOrderEvents(t, conn, createdAt, 2, 100, 200, 30000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)

	results, err := repo.GetMonthlyTotalRevenue(context.Background(), 2026, 3, 0)
	require.NoError(t, err)
	require.Len(t, results, 1)

	assert.Equal(t, "2026", results[0].Year)
	assert.Equal(t, "Mar", results[0].Month)
	assert.Equal(t, int64(80000), results[0].TotalRevenue)
}

func TestRepo_GetMonthlyTotalRevenue_NoData(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyTotalRevenue(context.Background(), 2099, 1, 0)
	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestRepo_GetMonthlyTotalRevenue_ByMerchant(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 4, 10, 10, 0, 0, 0, time.UTC)
	seedOrderEvents(t, conn, createdAt, 1, 100, 200, 60000)
	seedOrderEvents(t, conn, createdAt, 2, 100, 300, 40000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyTotalRevenue(context.Background(), 2026, 4, 200)
	require.NoError(t, err)
	require.Len(t, results, 1)
	assert.Equal(t, int64(60000), results[0].TotalRevenue)
}

func TestRepo_GetYearlyTotalRevenue(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt2026 := time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC)
	createdAt2025 := time.Date(2025, 8, 1, 10, 0, 0, 0, time.UTC)

	seedOrderEvents(t, conn, createdAt2026, 1, 100, 200, 100000)
	seedOrderEvents(t, conn, createdAt2026, 2, 100, 200, 50000)
	seedOrderEvents(t, conn, createdAt2025, 3, 100, 200, 75000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetYearlyTotalRevenue(context.Background(), 2026, 0)
	require.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestRepo_GetMonthlyOrderStats(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 5, 20, 10, 0, 0, 0, time.UTC)
	seedOrderEvents(t, conn, createdAt, 1, 100, 200, 50000)
	seedOrderEvents(t, conn, createdAt, 2, 100, 200, 30000)
	seedOrderItemEvents(t, conn, createdAt, 1, 1, 200, 10, 1, "Electronics", 3, 10000)
	seedOrderItemEvents(t, conn, createdAt, 2, 2, 200, 20, 2, "Fashion", 2, 5000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_events FINAL")
	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_item_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyOrderStats(context.Background(), 2026, 0)
	require.NoError(t, err)
	require.NotEmpty(t, results)
}

func TestRepo_GetYearlyOrderStats(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 3, 1, 10, 0, 0, 0, time.UTC)
	seedOrderEvents(t, conn, createdAt, 1, 100, 200, 80000)
	seedOrderItemEvents(t, conn, createdAt, 1, 1, 200, 10, 1, "Electronics", 5, 16000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_events FINAL")
	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_item_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetYearlyOrderStats(context.Background(), 2026, 0)
	require.NoError(t, err)
	require.NotEmpty(t, results)
	assert.Equal(t, "2026", results[0].Year)
}

// --- Category Stats ---

func TestRepo_GetMonthlyTotalPricing(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 2, 20, 10, 0, 0, 0, time.UTC)
	seedOrderItemEvents(t, conn, createdAt, 1, 1, 200, 10, 1, "Electronics", 2, 10000)
	seedOrderItemEvents(t, conn, createdAt, 2, 1, 200, 20, 2, "Fashion", 1, 5000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_item_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyTotalPricing(context.Background(), 2026, 2, "", 0)
	require.NoError(t, err)
	require.Len(t, results, 1)
	assert.Equal(t, int64(25000), results[0].TotalRevenue)
}

func TestRepo_GetMonthlyTotalPricing_NoData(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyTotalPricing(context.Background(), 2099, 1, "", 0)
	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestRepo_GetYearlyTotalPricing(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt2026 := time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC)
	createdAt2025 := time.Date(2025, 6, 1, 10, 0, 0, 0, time.UTC)

	seedOrderItemEvents(t, conn, createdAt2026, 1, 1, 200, 10, 1, "Electronics", 3, 10000)
	seedOrderItemEvents(t, conn, createdAt2025, 2, 2, 200, 20, 2, "Fashion", 2, 5000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_item_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetYearlyTotalPricing(context.Background(), 2026, "", 0)
	require.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestRepo_GetMonthlyCategoryStats(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 5, 12, 10, 0, 0, 0, time.UTC)
	seedOrderItemEvents(t, conn, createdAt, 1, 1, 200, 10, 1, "Electronics", 3, 10000)
	seedOrderItemEvents(t, conn, createdAt, 2, 2, 200, 20, 2, "Fashion", 2, 15000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_item_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyCategoryStats(context.Background(), 2026, "", 0)
	require.NoError(t, err)
	require.Len(t, results, 2)
}

func TestRepo_GetYearlyCategoryStats(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 7, 1, 10, 0, 0, 0, time.UTC)
	seedOrderItemEvents(t, conn, createdAt, 1, 1, 200, 10, 1, "Electronics", 5, 10000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE order_item_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetYearlyCategoryStats(context.Background(), 2026, "", 0)
	require.NoError(t, err)
	require.NotEmpty(t, results)
}

// --- Transaction Stats ---

func TestRepo_GetMonthlyAmount(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 7, 8, 10, 0, 0, 0, time.UTC)
	seedTransactionEvents(t, conn, createdAt, 1, 1, 200, "credit_card", "success", 50000)
	seedTransactionEvents(t, conn, createdAt, 2, 2, 200, "bank_transfer", "success", 30000)
	seedTransactionEvents(t, conn, createdAt, 3, 3, 200, "cod", "failed", 0)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE transaction_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyAmount(context.Background(), 2026, 7, "success", 0)
	require.NoError(t, err)
	require.Len(t, results, 1)
	assert.Equal(t, uint64(2), results[0].TotalCount)
	assert.Equal(t, int64(80000), results[0].TotalAmount)
}

func TestRepo_GetMonthlyAmount_NoData(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyAmount(context.Background(), 2099, 1, "success", 0)
	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestRepo_GetYearlyAmount(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt2026 := time.Date(2026, 6, 1, 10, 0, 0, 0, time.UTC)
	createdAt2025 := time.Date(2025, 6, 1, 10, 0, 0, 0, time.UTC)

	seedTransactionEvents(t, conn, createdAt2026, 1, 1, 200, "credit_card", "success", 50000)
	seedTransactionEvents(t, conn, createdAt2025, 2, 2, 200, "bank_transfer", "success", 30000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE transaction_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetYearlyAmount(context.Background(), 2026, "success", 0)
	require.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestRepo_GetMonthlyMethod(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 8, 10, 10, 0, 0, 0, time.UTC)
	seedTransactionEvents(t, conn, createdAt, 1, 1, 200, "credit_card", "success", 50000)
	seedTransactionEvents(t, conn, createdAt, 2, 2, 200, "credit_card", "success", 30000)
	seedTransactionEvents(t, conn, createdAt, 3, 3, 200, "cod", "success", 10000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE transaction_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyMethod(context.Background(), 2026, 8, "success", 0)
	require.NoError(t, err)
	require.Len(t, results, 2)
}

func TestRepo_GetYearlyMethod(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 9, 1, 10, 0, 0, 0, time.UTC)
	seedTransactionEvents(t, conn, createdAt, 1, 1, 200, "credit_card", "success", 50000)
	seedTransactionEvents(t, conn, createdAt, 2, 2, 200, "bank_transfer", "success", 30000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE transaction_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetYearlyMethod(context.Background(), 2026, "success", 0)
	require.NoError(t, err)
	require.NotEmpty(t, results)
}

func TestRepo_GetMonthlyAmount_ByMerchant(t *testing.T) {
	conn, cleanup := setupClickHouse(t)
	defer cleanup()

	createdAt := time.Date(2026, 10, 5, 10, 0, 0, 0, time.UTC)
	seedTransactionEvents(t, conn, createdAt, 1, 1, 200, "credit_card", "success", 50000)
	seedTransactionEvents(t, conn, createdAt, 2, 2, 300, "cod", "success", 30000)

	_ = conn.Exec(context.Background(), "OPTIMIZE TABLE transaction_events FINAL")

	repo := repository.NewClickHouseReaderRepository(conn)
	results, err := repo.GetMonthlyAmount(context.Background(), 2026, 10, "success", 200)
	require.NoError(t, err)
	require.Len(t, results, 1)
	assert.Equal(t, int64(50000), results[0].TotalAmount)
}
