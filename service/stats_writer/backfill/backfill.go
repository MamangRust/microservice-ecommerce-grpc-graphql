// Package backfill implements the stats-writer `backfill` command: it reads
// historical OLTP rows from the per-service PostgreSQL databases (orders,
// order_items joined with products/categories, transactions) and materializes
// them into ClickHouse through the same batch repository used for live events.
//
// This is the bootstrap path for the stats pipeline — it lets the ClickHouse
// tables reflect pre-existing data without replaying every domain event.
package backfill

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-writer/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/database"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/events"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// backfillEventID derives a deterministic UUID per entity so re-running the
// backfill replaces the same ReplacingMergeTree key (with a newer version)
// instead of appending duplicates.
func backfillEventID(kind string, id int32) string {
	return uuid.NewSHA1(uuid.NameSpaceDNS, []byte(fmt.Sprintf("backfill:%s:%d", kind, id))).String()
}

// Backfiller reads OLTP rows and pushes them into ClickHouse.
type Backfiller struct {
	log    logger.LoggerInterface
	repo   repository.Repository
	order  *pgxpool.Pool
	item   *pgxpool.Pool
	product *pgxpool.Pool
	category *pgxpool.Pool
	tx     *pgxpool.Pool
}

// New opens one connection per service database that owns a stats source and
// returns a ready Backfiller. Call Close to release them.
func New(log logger.LoggerInterface, repo repository.Repository) (*Backfiller, error) {
	open := func(prefix string) (*pgxpool.Pool, error) {
		conn, err := database.NewClientWithPrefix(log, prefix)
		if err != nil {
			return nil, fmt.Errorf("connect %s: %w", prefix, err)
		}
		return conn, nil
	}

	order, err := open("DB_ORDER")
	if err != nil {
		return nil, err
	}
	item, err := open("DB_ORDER_ITEM")
	if err != nil {
		order.Close()
		return nil, err
	}
	product, err := open("DB_PRODUCT")
	if err != nil {
		order.Close()
		item.Close()
		return nil, err
	}
	category, err := open("DB_CATEGORY")
	if err != nil {
		order.Close()
		item.Close()
		product.Close()
		return nil, err
	}
	tx, err := open("DB_TRANSACTION")
	if err != nil {
		order.Close()
		item.Close()
		product.Close()
		category.Close()
		return nil, err
	}

	return &Backfiller{
		log:      log,
		repo:     repo,
		order:    order,
		item:     item,
		product:  product,
		category: category,
		tx:       tx,
	}, nil
}

func (b *Backfiller) Close() {
	for _, conn := range []*pgxpool.Pool{b.order, b.item, b.product, b.category, b.tx} {
		if conn != nil {
			conn.Close()
		}
	}
}

// Run streams all stats sources into ClickHouse. The event version is the
// backfill run timestamp so re-running supersedes previous rows.
func (b *Backfiller) Run(ctx context.Context) error {
	version := uint64(time.Now().Unix())
	counts := map[string]int{}

	if err := b.backfillOrders(ctx, version, counts); err != nil {
		return err
	}
	if err := b.backfillOrderItems(ctx, version, counts); err != nil {
		return err
	}
	if err := b.backfillTransactions(ctx, version, counts); err != nil {
		return err
	}

	if err := b.repo.Flush(ctx); err != nil {
		return fmt.Errorf("flush backfill batches: %w", err)
	}

	b.log.Info("backfill complete",
		zap.Int("orders", counts["order"]),
		zap.Int("order_items", counts["order_item"]),
		zap.Int("transactions", counts["transaction"]),
	)
	return nil
}

type orderRow struct {
	OrderID    int32
	UserID     int32
	MerchantID int32
	TotalPrice int32
	CreatedAt  time.Time
}

type itemRow struct {
	OrderItemID int32
	OrderID     int32
	ProductID   int32
	Quantity    int32
	Price       int32
	CreatedAt   time.Time
}

type txRow struct {
	TransactionID int32
	OrderID       int32
	MerchantID    int32
	PaymentMethod string
	Amount        int32
	Status        string
	CreatedAt     time.Time
}

func (b *Backfiller) backfillOrders(ctx context.Context, version uint64, counts map[string]int) error {
	rows, err := b.order.Query(ctx, `SELECT order_id, user_id, merchant_id, total_price, created_at
		FROM orders WHERE deleted_at IS NULL`)
	if err != nil {
		return fmt.Errorf("query orders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var r orderRow
		if err := rows.Scan(&r.OrderID, &r.UserID, &r.MerchantID, &r.TotalPrice, &r.CreatedAt); err != nil {
			return fmt.Errorf("scan order: %w", err)
		}
		event := events.OrderEvent{
			OrderID:    r.OrderID,
			UserID:     r.UserID,
			MerchantID: r.MerchantID,
			TotalPrice: r.TotalPrice,
			Status:     "created",
			EventTime:  r.CreatedAt.UTC().Format(time.RFC3339),
		}
		if err := b.repo.InsertOrderEvent(ctx, backfillEventID("order", r.OrderID), version, event); err != nil {
			return fmt.Errorf("insert order %d: %w", r.OrderID, err)
		}
		counts["order"]++
	}
	return rows.Err()
}

func (b *Backfiller) backfillOrderItems(ctx context.Context, version uint64, counts map[string]int) error {
	// Load order -> merchant_id from the order service DB.
	orderMerchant := map[int32]int32{}
	orderRows, err := b.order.Query(ctx, `SELECT order_id, merchant_id FROM orders WHERE deleted_at IS NULL`)
	if err != nil {
		return fmt.Errorf("query order merchant map: %w", err)
	}
	for orderRows.Next() {
		var orderID, merchantID int32
		if err := orderRows.Scan(&orderID, &merchantID); err != nil {
			orderRows.Close()
			return fmt.Errorf("scan order merchant map: %w", err)
		}
		orderMerchant[orderID] = merchantID
	}
	orderRows.Close()
	if err := orderRows.Err(); err != nil {
		return err
	}

	// Load product -> category_id from the product service DB.
	productCategory := map[int32]int32{}
	productRows, err := b.product.Query(ctx, `SELECT product_id, category_id FROM products WHERE deleted_at IS NULL`)
	if err != nil {
		return fmt.Errorf("query product category map: %w", err)
	}
	for productRows.Next() {
		var productID, categoryID int32
		if err := productRows.Scan(&productID, &categoryID); err != nil {
			productRows.Close()
			return fmt.Errorf("scan product category map: %w", err)
		}
		productCategory[productID] = categoryID
	}
	productRows.Close()
	if err := productRows.Err(); err != nil {
		return err
	}

	// Load category_id -> name from the category service DB.
	categoryName := map[int32]string{}
	categoryRows, err := b.category.Query(ctx, `SELECT category_id, name FROM categories WHERE deleted_at IS NULL`)
	if err != nil {
		return fmt.Errorf("query category name map: %w", err)
	}
	for categoryRows.Next() {
		var categoryID int32
		var name string
		if err := categoryRows.Scan(&categoryID, &name); err != nil {
			categoryRows.Close()
			return fmt.Errorf("scan category name map: %w", err)
		}
		categoryName[categoryID] = name
	}
	categoryRows.Close()
	if err := categoryRows.Err(); err != nil {
		return err
	}

	itemRows, err := b.item.Query(ctx, `SELECT order_item_id, order_id, product_id, quantity, price, created_at
		FROM order_items WHERE deleted_at IS NULL`)
	if err != nil {
		return fmt.Errorf("query order items: %w", err)
	}
	defer itemRows.Close()

	for itemRows.Next() {
		var r itemRow
		if err := itemRows.Scan(&r.OrderItemID, &r.OrderID, &r.ProductID, &r.Quantity, &r.Price, &r.CreatedAt); err != nil {
			return fmt.Errorf("scan order item: %w", err)
		}
		catID := productCategory[r.ProductID]
		event := events.OrderItemEvent{
			OrderItemID:  r.OrderItemID,
			OrderID:      r.OrderID,
			MerchantID:   orderMerchant[r.OrderID],
			ProductID:    r.ProductID,
			CategoryID:   catID,
			CategoryName: categoryName[catID],
			Quantity:     r.Quantity,
			Price:        r.Price,
			EventTime:    r.CreatedAt.UTC().Format(time.RFC3339),
		}
		if err := b.repo.InsertOrderItemEvent(ctx, backfillEventID("order_item", r.OrderItemID), version, event); err != nil {
			return fmt.Errorf("insert order item %d: %w", r.OrderItemID, err)
		}
		counts["order_item"]++
	}
	return itemRows.Err()
}

func (b *Backfiller) backfillTransactions(ctx context.Context, version uint64, counts map[string]int) error {
	rows, err := b.tx.Query(ctx, `SELECT transaction_id, order_id, merchant_id, payment_method, amount, payment_status, created_at
		FROM transactions WHERE deleted_at IS NULL`)
	if err != nil {
		return fmt.Errorf("query transactions: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var r txRow
		if err := rows.Scan(&r.TransactionID, &r.OrderID, &r.MerchantID, &r.PaymentMethod, &r.Amount, &r.Status, &r.CreatedAt); err != nil {
			return fmt.Errorf("scan transaction: %w", err)
		}
		event := events.TransactionEvent{
			TransactionID: r.TransactionID,
			OrderID:       r.OrderID,
			MerchantID:    r.MerchantID,
			PaymentMethod: r.PaymentMethod,
			Amount:        r.Amount,
			Status:        r.Status,
			EventTime:     r.CreatedAt.UTC().Format(time.RFC3339),
		}
		if err := b.repo.InsertTransactionEvent(ctx, backfillEventID("transaction", r.TransactionID), version, event); err != nil {
			return fmt.Errorf("insert transaction %d: %w", r.TransactionID, err)
		}
		counts["transaction"]++
	}
	return rows.Err()
}
