package repository

import (
	"context"

	dto "github.com/MamangRust/microservice-ecommerce-grpc-transaction/dto"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type orderItemRepository struct {
	client pb.OrderItemQueryServiceClient
}

func NewOrderItemRepository(client pb.OrderItemQueryServiceClient) *orderItemRepository {
	return &orderItemRepository{
		client: client,
	}
}

func (r *orderItemRepository) FindOrderItemByOrder(ctx context.Context, order_id int) ([]*dto.GetOrderItemsByOrderRow, error) {
	res, err := r.client.FindOrderItemByOrder(ctx, &pb.FindByIdOrderItemRequest{Id: int32(order_id)})
	if err != nil {
		// pertahankan status gRPC dari dependency service (NotFound -> 404, dst)
		return nil, err
	}

	var items []*dto.GetOrderItemsByOrderRow
	for _, item := range res.Data {
		items = append(items, &dto.GetOrderItemsByOrderRow{
			OrderItemID: item.Id,
			OrderID:     item.OrderId,
			ProductID:   item.ProductId,
			Quantity:    int32(item.Quantity),
			Price:       int32(item.Price),
		})
	}

	return items, nil
}
