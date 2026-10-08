package order_itemgraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/order_item"
)

type orderItemGraphqlMapper struct{}

func NewOrderItemGraphqlMapper() *orderItemGraphqlMapper {
	return &orderItemGraphqlMapper{}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponseOrderItem(res *pb_order_item.ApiResponseOrderItem) *model.APIResponseOrderItem {
	return &model.APIResponseOrderItem{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponseOrderItem(res.Data),
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponsesOrderItem(res *pb_order_item.ApiResponsesOrderItem) *model.APIResponsesOrderItem {
	return &model.APIResponsesOrderItem{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponsesOrderItem(res.Data),
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponseOrderItemDelete(res *pb_order_item.ApiResponseOrderItemDelete) *model.APIResponseOrderItemDelete {
	return &model.APIResponseOrderItemDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponseOrderItemAll(res *pb_order_item.ApiResponseOrderItemAll) *model.APIResponseOrderItemAll {
	return &model.APIResponseOrderItemAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponsePaginationOrderItem(
	res *pb_order_item.ApiResponsePaginationOrderItem,
) *model.APIResponsePaginationOrderItem {
	return &model.APIResponsePaginationOrderItem{
		Status:     res.Status,
		Message:    res.Message,
		Data:       o.mapResponsesOrderItem(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponsePaginationOrderItemDeleteAt(
	res *pb_order_item.ApiResponsePaginationOrderItemDeleteAt,
) *model.APIResponsePaginationOrderItemDeleteAt {
	return &model.APIResponsePaginationOrderItemDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       o.mapResponsesOrderItemDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (o *orderItemGraphqlMapper) mapResponseOrderItem(orderItem *pb_order_item.OrderItemResponse) *model.OrderItemResponse {
	return &model.OrderItemResponse{
		ID:        int32(orderItem.Id),
		OrderID:   int32(orderItem.OrderId),
		ProductID: int32(orderItem.ProductId),
		Quantity:  int32(orderItem.Quantity),
		Price:     int32(orderItem.Price),
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
	}
}

func (o *orderItemGraphqlMapper) mapResponsesOrderItem(orderItems []*pb_order_item.OrderItemResponse) []*model.OrderItemResponse {
	mapped := make([]*model.OrderItemResponse, 0, len(orderItems))
	for _, oi := range orderItems {
		mapped = append(mapped, o.mapResponseOrderItem(oi))
	}
	return mapped
}

func (o *orderItemGraphqlMapper) mapResponseOrderItemDelete(orderItem *pb_order_item.OrderItemResponseDeleteAt) *model.OrderItemResponseDeleteAt {
	var deletedAt *string
	if orderItem.DeletedAt != nil {
		deletedAt = &orderItem.DeletedAt.Value
	}

	return &model.OrderItemResponseDeleteAt{
		ID:        int32(orderItem.Id),
		OrderID:   int32(orderItem.OrderId),
		ProductID: int32(orderItem.ProductId),
		Quantity:  int32(orderItem.Quantity),
		Price:     int32(orderItem.Price),
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

func (o *orderItemGraphqlMapper) mapResponsesOrderItemDeleteAt(orderItems []*pb_order_item.OrderItemResponseDeleteAt) []*model.OrderItemResponseDeleteAt {
	mapped := make([]*model.OrderItemResponseDeleteAt, 0, len(orderItems))
	for _, oi := range orderItems {
		mapped = append(mapped, o.mapResponseOrderItemDelete(oi))
	}
	return mapped
}
