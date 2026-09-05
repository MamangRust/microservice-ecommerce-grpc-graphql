package repository

import (
	"context"

	db "github.com/MamangRust/microservice-ecommerce-grpc-order/database/schema"
	dto "github.com/MamangRust/microservice-ecommerce-grpc-order/dto"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/requests"
)

type UserQueryRepository interface {
	FindByID(ctx context.Context, user_id int) (*dto.GetUserByIDRow, error)
}

type ProductQueryRepository interface {
	FindByID(ctx context.Context, product_id int) (*dto.GetProductByIDRow, error)
}

type MerchantQueryRepository interface {
	FindByID(ctx context.Context, user_id int) (*dto.GetMerchantByIDRow, error)
}

type ProductCommandRepository interface {
	UpdateProductCountStock(ctx context.Context, product_id int, stock int) (*dto.UpdateProductCountStockRow, error)
	AdjustProductStock(ctx context.Context, product_id int, delta int, operationID string) (*dto.AdjustProductStockRow, error)
}

type ShippingAddressCommandRepository interface {
	Create(
		ctx context.Context,
		request *requests.CreateShippingAddressRequest,
	) (*dto.CreateShippingAddressRow, error)

	Update(
		ctx context.Context,
		request *requests.UpdateShippingAddressRequest,
	) (*dto.UpdateShippingAddressRow, error)

	DeleteByOrderIDPermanent(
		ctx context.Context,
		order_id int,
	) (bool, error)
	DeleteAll(ctx context.Context) (bool, error)
}

type TransactionCommandRepository interface {
	DeleteByOrderIDPermanent(
		ctx context.Context,
		order_id int,
	) (bool, error)
	DeleteAll(ctx context.Context) (bool, error)
}

type OrderItemQueryRepository interface {
	FindOrderItemByOrder(
		ctx context.Context,
		order_id int,
	) ([]*dto.GetOrderItemsByOrderRow, error)
	CalculateTotalPrice(
		ctx context.Context,
		order_id int,
	) (*int32, error)
}

type OrderItemCommandRepository interface {
	Create(
		ctx context.Context,
		req *requests.CreateOrderItemRecordRequest,
	) (*dto.CreateOrderItemRow, error)

	Update(
		ctx context.Context,
		req *requests.UpdateOrderItemRecordRequest,
	) (*dto.UpdateOrderItemRow, error)

	Trash(
		ctx context.Context,
		order_id int,
	) (*dto.OrderItem, error)

	Restore(
		ctx context.Context,
		order_id int,
	) (*dto.OrderItem, error)

	DeletePermanent(
		ctx context.Context,
		order_id int,
	) (bool, error)
	DeleteByOrderIDPermanent(
		ctx context.Context,
		order_id int,
	) (bool, error)
	RestoreAll(ctx context.Context) (bool, error)
	DeleteAll(ctx context.Context) (bool, error)
}

// F5: legacy OLTP order stats repositories were removed; stats are served by
// service/stats_reader from ClickHouse.

type OrderCommandRepository interface {
	Create(
		ctx context.Context,
		request *requests.CreateOrderRecordRequest,
	) (*db.CreateOrderRow, error)

	Update(
		ctx context.Context,
		request *requests.UpdateOrderRecordRequest,
	) (*db.UpdateOrderRow, error)

	Trash(
		ctx context.Context,
		order_id int,
	) (*db.Order, error)

	Restore(
		ctx context.Context,
		order_id int,
	) (*db.Order, error)

	DeletePermanent(
		ctx context.Context,
		order_id int,
	) (bool, error)

	// DeletePermanentWithChildren atomically removes a trashed order and all of
	// its child rows (stock reservations, order items, transactions, shipping
	// addresses) in a single SQL statement. It returns ErrOrderNotFound when the
	// order is not trashed.
	DeletePermanentWithChildren(
		ctx context.Context,
		order_id int,
	) (bool, error)
	FindTrashedByID(ctx context.Context, order_id int) (*db.Order, error)
	FindTrashed(ctx context.Context) ([]*db.Order, error)

	RestoreAll(ctx context.Context) (bool, error)
	DeleteAll(ctx context.Context) (bool, error)
}

type OrderQueryRepository interface {
	FindAll(
		ctx context.Context,
		req *requests.FindAllOrder,
	) ([]*db.GetOrdersRow, error)

	FindActive(
		ctx context.Context,
		req *requests.FindAllOrder,
	) ([]*db.GetOrdersActiveRow, error)

	FindTrashed(
		ctx context.Context,
		req *requests.FindAllOrder,
	) ([]*db.GetOrdersTrashedRow, error)

	FindByMerchant(
		ctx context.Context,
		req *requests.FindAllOrderByMerchant,
	) ([]*db.GetOrdersByMerchantRow, error)

	FindByID(
		ctx context.Context,
		order_id int,
	) (*db.GetOrderByIDRow, error)
}
