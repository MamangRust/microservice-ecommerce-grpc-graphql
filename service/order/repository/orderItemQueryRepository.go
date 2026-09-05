package repository

import (
	"context"

	dto "github.com/MamangRust/microservice-ecommerce-grpc-order/dto"
	order_item_errors "github.com/MamangRust/microservice-ecommerce-shared/errors/order_item_errors"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type orderItemQueryRepository struct {
	queryClient   pb.OrderItemQueryServiceClient
	commandClient pb.OrderItemCommandServiceClient
}

func NewOrderItemQueryRepository(queryClient pb.OrderItemQueryServiceClient, commandClient pb.OrderItemCommandServiceClient) *orderItemQueryRepository {
	return &orderItemQueryRepository{
		queryClient:   queryClient,
		commandClient: commandClient,
	}
}

func (r *orderItemQueryRepository) FindOrderItemByOrder(ctx context.Context, order_id int) ([]*dto.GetOrderItemsByOrderRow, error) {
	res, err := r.queryClient.FindOrderItemByOrder(ctx, &pb.FindByIdOrderItemRequest{Id: int32(order_id)})
	if err != nil {
		return nil, order_item_errors.ErrFindOrderItemByOrder.WithInternal(err)
	}

	var items []*dto.GetOrderItemsByOrderRow
	for _, item := range res.Data {
		items = append(items, &dto.GetOrderItemsByOrderRow{
			OrderItemID: item.Id,
			OrderID:     item.OrderId,
			ProductID:   item.ProductId,
			Quantity:    item.Quantity,
			Price:       item.Price,
		})
	}

	return items, nil
}

func (r *orderItemQueryRepository) CalculateTotalPrice(ctx context.Context, order_id int) (*int32, error) {
	res, err := r.commandClient.CalculateTotalPrice(ctx, &pb.CalculateTotalPriceRequest{OrderId: int32(order_id)})
	if err != nil {
		return nil, order_item_errors.ErrCalculateTotalPrice.WithInternal(err)
	}

	total := int32(res.TotalPrice)
	return &total, nil
}
