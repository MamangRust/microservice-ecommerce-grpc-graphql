package repository

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type ClickHouseReaderRepository struct {
	conn clickhouse.Conn
}

func NewClickHouseReaderRepository(conn clickhouse.Conn) *ClickHouseReaderRepository {
	return &ClickHouseReaderRepository{conn: conn}
}

// --- Order stats ---

func (r *ClickHouseReaderRepository) GetMonthlyTotalRevenue(ctx context.Context, year, month int, merchantID int32) ([]MonthlyRevenue, error) {
	where := "toYear(created_at) = ? AND toMonth(created_at) = ?"
	args := []interface{}{year, month}
	if merchantID > 0 {
		where += " AND merchant_id = ?"
		args = append(args, merchantID)
	}
	query := fmt.Sprintf(`
		SELECT toString(toYear(created_at)) AS year, formatDateTime(created_at, '%%b') AS month, sum(total_price) AS total_revenue
		FROM order_events
		WHERE %s
		GROUP BY year, month, toMonth(created_at)
		ORDER BY year, toMonth(created_at)
	`, where)
	return r.queryMonthlyRevenue(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetYearlyTotalRevenue(ctx context.Context, year int, merchantID int32) ([]YearlyRevenue, error) {
	where := "(toYear(created_at) = ? OR toYear(created_at) = ?)"
	args := []interface{}{year, year - 1}
	if merchantID > 0 {
		where += " AND merchant_id = ?"
		args = append(args, merchantID)
	}
	query := fmt.Sprintf(`
		SELECT toString(toYear(created_at)) AS year, sum(total_price) AS total_revenue
		FROM order_events
		WHERE %s
		GROUP BY year
		ORDER BY year DESC
	`, where)
	return r.queryYearlyRevenue(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetMonthlyOrderStats(ctx context.Context, year int, merchantID int32) ([]MonthlyOrder, error) {
	where := "toYear(o.created_at) = ?"
	args := []interface{}{year}
	if merchantID > 0 {
		where += " AND o.merchant_id = ?"
		args = append(args, merchantID)
	}
	query := fmt.Sprintf(`
		SELECT
			formatDateTime(o.created_at, '%%b') AS month,
			countDistinct(o.order_id) AS order_count,
			sum(o.total_price) AS total_revenue,
			sum(i.quantity) AS total_items_sold
		FROM order_events o
		LEFT JOIN order_item_events i ON o.order_id = i.order_id
		WHERE %s
		GROUP BY month, toMonth(o.created_at)
		ORDER BY toMonth(o.created_at)
	`, where)
	return r.queryMonthlyOrder(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetYearlyOrderStats(ctx context.Context, year int, merchantID int32) ([]YearlyOrder, error) {
	where := "toYear(o.created_at) >= ? AND toYear(o.created_at) <= ?"
	args := []interface{}{year - 4, year}
	if merchantID > 0 {
		where += " AND o.merchant_id = ?"
		args = append(args, merchantID)
	}
	query := fmt.Sprintf(`
		SELECT
			toString(toYear(o.created_at)) AS year,
			countDistinct(o.order_id) AS order_count,
			sum(o.total_price) AS total_revenue,
			sum(i.quantity) AS total_items_sold,
			uniqExact(i.product_id) AS unique_products_sold
		FROM order_events o
		LEFT JOIN order_item_events i ON o.order_id = i.order_id
		WHERE %s
		GROUP BY year
		ORDER BY year
	`, where)
	return r.queryYearlyOrder(ctx, query, args...)
}

// --- Category pricing stats ---

func (r *ClickHouseReaderRepository) GetMonthlyTotalPricing(ctx context.Context, year, month int, filterField string, filterValue int32) ([]MonthlyRevenue, error) {
	where := "toYear(created_at) = ? AND toMonth(created_at) = ?"
	args := []interface{}{year, month}
	if filterField != "" {
		where += " AND " + filterField + " = ?"
		args = append(args, filterValue)
	}
	query := fmt.Sprintf(`
		SELECT toString(toYear(created_at)) AS year, formatDateTime(created_at, '%%b') AS month, sum(price * quantity) AS total_revenue
		FROM order_item_events
		WHERE %s
		GROUP BY year, month, toMonth(created_at)
		ORDER BY year, toMonth(created_at)
	`, where)
	return r.queryMonthlyRevenue(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetYearlyTotalPricing(ctx context.Context, year int, filterField string, filterValue int32) ([]YearlyRevenue, error) {
	where := "(toYear(created_at) = ? OR toYear(created_at) = ?)"
	args := []interface{}{year, year - 1}
	if filterField != "" {
		where += " AND " + filterField + " = ?"
		args = append(args, filterValue)
	}
	query := fmt.Sprintf(`
		SELECT toString(toYear(created_at)) AS year, sum(price * quantity) AS total_revenue
		FROM order_item_events
		WHERE %s
		GROUP BY year
		ORDER BY year DESC
	`, where)
	return r.queryYearlyRevenue(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetMonthlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]MonthlyCategory, error) {
	where := "toYear(created_at) = ?"
	args := []interface{}{year}
	if filterField != "" {
		where += " AND " + filterField + " = ?"
		args = append(args, filterValue)
	}
	query := fmt.Sprintf(`
		SELECT
			formatDateTime(created_at, '%%b') AS month,
			category_id,
			category_name,
			countDistinct(order_id) AS order_count,
			sum(quantity) AS items_sold,
			sum(price * quantity) AS total_revenue
		FROM order_item_events
		WHERE %s
		GROUP BY month, toMonth(created_at), category_id, category_name
		ORDER BY toMonth(created_at), total_revenue DESC
	`, where)
	return r.queryMonthlyCategory(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetYearlyCategoryStats(ctx context.Context, year int, filterField string, filterValue int32) ([]YearlyCategory, error) {
	where := "toYear(created_at) >= ? AND toYear(created_at) <= ?"
	args := []interface{}{year - 4, year}
	if filterField != "" {
		where += " AND " + filterField + " = ?"
		args = append(args, filterValue)
	}
	query := fmt.Sprintf(`
		SELECT
			toString(toYear(created_at)) AS year,
			category_id,
			category_name,
			countDistinct(order_id) AS order_count,
			sum(quantity) AS items_sold,
			sum(price * quantity) AS total_revenue,
			uniqExact(product_id) AS unique_products_sold
		FROM order_item_events
		WHERE %s
		GROUP BY year, category_id, category_name
		ORDER BY year, total_revenue DESC
	`, where)
	return r.queryYearlyCategory(ctx, query, args...)
}

// --- Transaction stats ---

func (r *ClickHouseReaderRepository) GetMonthlyAmount(ctx context.Context, year, month int, status string, merchantID int32) ([]MonthlyAmount, error) {
	where := "toYear(created_at) = ? AND toMonth(created_at) = ? AND status = ?"
	args := []interface{}{year, month, status}
	if merchantID > 0 {
		where += " AND merchant_id = ?"
		args = append(args, merchantID)
	}
	query := fmt.Sprintf(`
		SELECT toString(toYear(created_at)) AS year, formatDateTime(created_at, '%%b') AS month, count() AS total_count, sum(amount) AS total_amount
		FROM transaction_events
		WHERE %s
		GROUP BY year, month, toMonth(created_at)
		ORDER BY year, toMonth(created_at)
	`, where)
	return r.queryMonthlyAmount(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetYearlyAmount(ctx context.Context, year int, status string, merchantID int32) ([]YearlyAmount, error) {
	where := "(toYear(created_at) = ? OR toYear(created_at) = ?) AND status = ?"
	args := []interface{}{year, year - 1, status}
	if merchantID > 0 {
		where += " AND merchant_id = ?"
		args = append(args, merchantID)
	}
	query := fmt.Sprintf(`
		SELECT toString(toYear(created_at)) AS year, count() AS total_count, sum(amount) AS total_amount
		FROM transaction_events
		WHERE %s
		GROUP BY year
		ORDER BY year DESC
	`, where)
	return r.queryYearlyAmount(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetMonthlyMethod(ctx context.Context, year, month int, status string, merchantID int32) ([]MonthlyMethod, error) {
	where := "toYear(created_at) = ? AND toMonth(created_at) = ? AND status = ?"
	args := []interface{}{year, month, status}
	if merchantID > 0 {
		where += " AND merchant_id = ?"
		args = append(args, merchantID)
	}
	query := fmt.Sprintf(`
		SELECT formatDateTime(created_at, '%%b') AS month, payment_method, count() AS total_count, sum(amount) AS total_amount
		FROM transaction_events
		WHERE %s
		GROUP BY month, toMonth(created_at), payment_method
		ORDER BY toMonth(created_at), payment_method
	`, where)
	return r.queryMonthlyMethod(ctx, query, args...)
}

func (r *ClickHouseReaderRepository) GetYearlyMethod(ctx context.Context, year int, status string, merchantID int32) ([]YearlyMethod, error) {
	where := "(toYear(created_at) = ? OR toYear(created_at) = ?) AND status = ?"
	args := []interface{}{year, year - 1, status}
	if merchantID > 0 {
		where += " AND merchant_id = ?"
		args = append(args, merchantID)
	}
	query := fmt.Sprintf(`
		SELECT toString(toYear(created_at)) AS year, payment_method, count() AS total_count, sum(amount) AS total_amount
		FROM transaction_events
		WHERE %s
		GROUP BY year, payment_method
		ORDER BY year DESC, payment_method
	`, where)
	return r.queryYearlyMethod(ctx, query, args...)
}

// --- Scan helpers ---

func (r *ClickHouseReaderRepository) queryMonthlyRevenue(ctx context.Context, query string, args ...interface{}) ([]MonthlyRevenue, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []MonthlyRevenue
	for rows.Next() {
		var m MonthlyRevenue
		if err := rows.Scan(&m.Year, &m.Month, &m.TotalRevenue); err != nil {
			return nil, err
		}
		results = append(results, m)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryYearlyRevenue(ctx context.Context, query string, args ...interface{}) ([]YearlyRevenue, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []YearlyRevenue
	for rows.Next() {
		var y YearlyRevenue
		if err := rows.Scan(&y.Year, &y.TotalRevenue); err != nil {
			return nil, err
		}
		results = append(results, y)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryMonthlyOrder(ctx context.Context, query string, args ...interface{}) ([]MonthlyOrder, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []MonthlyOrder
	for rows.Next() {
		var m MonthlyOrder
		if err := rows.Scan(&m.Month, &m.OrderCount, &m.TotalRevenue, &m.TotalItemsSold); err != nil {
			return nil, err
		}
		results = append(results, m)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryYearlyOrder(ctx context.Context, query string, args ...interface{}) ([]YearlyOrder, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []YearlyOrder
	for rows.Next() {
		var y YearlyOrder
		if err := rows.Scan(&y.Year, &y.OrderCount, &y.TotalRevenue, &y.TotalItemsSold, &y.UniqueProductsSold); err != nil {
			return nil, err
		}
		results = append(results, y)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryMonthlyCategory(ctx context.Context, query string, args ...interface{}) ([]MonthlyCategory, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []MonthlyCategory
	for rows.Next() {
		var m MonthlyCategory
		if err := rows.Scan(&m.Month, &m.CategoryID, &m.CategoryName, &m.OrderCount, &m.ItemsSold, &m.TotalRevenue); err != nil {
			return nil, err
		}
		results = append(results, m)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryYearlyCategory(ctx context.Context, query string, args ...interface{}) ([]YearlyCategory, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []YearlyCategory
	for rows.Next() {
		var y YearlyCategory
		if err := rows.Scan(&y.Year, &y.CategoryID, &y.CategoryName, &y.OrderCount, &y.ItemsSold, &y.TotalRevenue, &y.UniqueProductsSold); err != nil {
			return nil, err
		}
		results = append(results, y)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryMonthlyAmount(ctx context.Context, query string, args ...interface{}) ([]MonthlyAmount, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []MonthlyAmount
	for rows.Next() {
		var m MonthlyAmount
		if err := rows.Scan(&m.Year, &m.Month, &m.TotalCount, &m.TotalAmount); err != nil {
			return nil, err
		}
		results = append(results, m)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryYearlyAmount(ctx context.Context, query string, args ...interface{}) ([]YearlyAmount, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []YearlyAmount
	for rows.Next() {
		var y YearlyAmount
		if err := rows.Scan(&y.Year, &y.TotalCount, &y.TotalAmount); err != nil {
			return nil, err
		}
		results = append(results, y)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryMonthlyMethod(ctx context.Context, query string, args ...interface{}) ([]MonthlyMethod, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []MonthlyMethod
	for rows.Next() {
		var m MonthlyMethod
		if err := rows.Scan(&m.Month, &m.PaymentMethod, &m.TotalCount, &m.TotalAmount); err != nil {
			return nil, err
		}
		results = append(results, m)
	}
	return results, rows.Err()
}

func (r *ClickHouseReaderRepository) queryYearlyMethod(ctx context.Context, query string, args ...interface{}) ([]YearlyMethod, error) {
	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []YearlyMethod
	for rows.Next() {
		var m YearlyMethod
		if err := rows.Scan(&m.Year, &m.PaymentMethod, &m.TotalCount, &m.TotalAmount); err != nil {
			return nil, err
		}
		results = append(results, m)
	}
	return results, rows.Err()
}

