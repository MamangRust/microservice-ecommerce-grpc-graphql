package repository

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// Row shapes returned by the ClickHouse queries. Amounts are int64 because
// ClickHouse SUM returns a 64-bit integer for Int64 columns.
type MonthlyRevenue struct {
	Year         string
	Month        string
	TotalRevenue int64
}

type YearlyRevenue struct {
	Year         string
	TotalRevenue int64
}

type MonthlyOrder struct {
	Month          string
	OrderCount     uint64
	TotalRevenue   int64
	TotalItemsSold uint64
}

type YearlyOrder struct {
	Year               string
	OrderCount         uint64
	TotalRevenue       int64
	TotalItemsSold     uint64
	UniqueProductsSold uint64
}

type MonthlyCategory struct {
	Month        string
	CategoryID   uint64
	CategoryName string
	OrderCount   uint64
	ItemsSold    uint64
	TotalRevenue int64
}

type YearlyCategory struct {
	Year               string
	CategoryID         uint64
	CategoryName       string
	OrderCount         uint64
	ItemsSold          uint64
	TotalRevenue       int64
	UniqueProductsSold uint64
}

type MonthlyAmount struct {
	Year        string
	Month       string
	TotalCount  uint64
	TotalAmount int64
}

type YearlyAmount struct {
	Year        string
	TotalCount  uint64
	TotalAmount int64
}

type MonthlyMethod struct {
	Month          string
	PaymentMethod  string
	TotalCount     uint64
	TotalAmount    int64
}

type YearlyMethod struct {
	Year          string
	PaymentMethod string
	TotalCount    uint64
	TotalAmount   int64
}

// Repository reads aggregated statistics from ClickHouse. The category pricing
// queries accept a filterField ("", "merchant_id" or "category_id") so the
// all/by-merchant/by-category handlers share one query builder, mirroring the
// payment-gateway stats-reader repository.
type Repository interface {
	// Order stats (order_events, optionally joined with order_item_events).
	GetMonthlyTotalRevenue(ctx context.Context, year, month int, merchantID int32) ([]MonthlyRevenue, error)
	GetYearlyTotalRevenue(ctx context.Context, year int, merchantID int32) ([]YearlyRevenue, error)
	GetMonthlyOrderStats(ctx context.Context, year int, merchantID int32) ([]MonthlyOrder, error)
	GetYearlyOrderStats(ctx context.Context, year int, merchantID int32) ([]YearlyOrder, error)

	// Category pricing stats (order_item_events).
	GetMonthlyTotalPricing(ctx context.Context, year, month int, filterField string, filterValue int32) ([]MonthlyRevenue, error)
	GetYearlyTotalPricing(ctx context.Context, year int, filterField string, filterValue int32) ([]YearlyRevenue, error)
	GetMonthlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]MonthlyCategory, error)
	GetYearlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]YearlyCategory, error)

	// Transaction stats (transaction_events).
	GetMonthlyAmount(ctx context.Context, year, month int, status string, merchantID int32) ([]MonthlyAmount, error)
	GetYearlyAmount(ctx context.Context, year int, status string, merchantID int32) ([]YearlyAmount, error)
	GetMonthlyMethod(ctx context.Context, year, month int, status string, merchantID int32) ([]MonthlyMethod, error)
	GetYearlyMethod(ctx context.Context, year int, status string, merchantID int32) ([]YearlyMethod, error)
}

func NewRepository(conn clickhouse.Conn) Repository {
	return NewClickHouseReaderRepository(conn)
}
