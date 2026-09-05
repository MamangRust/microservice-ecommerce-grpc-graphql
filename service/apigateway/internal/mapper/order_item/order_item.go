package order_itemgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type orderItemGraphqlMapper struct{}

func NewOrderItemGraphqlMapper() *orderItemGraphqlMapper {
	return &orderItemGraphqlMapper{}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponseOrderItem(res *pb.ApiResponseOrderItem) *model.APIResponseOrderItem {
	return &model.APIResponseOrderItem{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponseOrderItem(res.Data),
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponsesOrderItem(res *pb.ApiResponsesOrderItem) *model.APIResponsesOrderItem {
	return &model.APIResponsesOrderItem{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponsesOrderItem(res.Data),
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponseOrderItemDelete(res *pb.ApiResponseOrderItemDelete) *model.APIResponseOrderItemDelete {
	return &model.APIResponseOrderItemDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponseOrderItemAll(res *pb.ApiResponseOrderItemAll) *model.APIResponseOrderItemAll {
	return &model.APIResponseOrderItemAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponsePaginationOrderItem(
	res *pb.ApiResponsePaginationOrderItem,
) *model.APIResponsePaginationOrderItem {
	return &model.APIResponsePaginationOrderItem{
		Status:     res.Status,
		Message:    res.Message,
		Data:       o.mapResponsesOrderItem(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (o *orderItemGraphqlMapper) ToGraphqlResponsePaginationOrderItemDeleteAt(
	res *pb.ApiResponsePaginationOrderItemDeleteAt,
) *model.APIResponsePaginationOrderItemDeleteAt {
	return &model.APIResponsePaginationOrderItemDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       o.mapResponsesOrderItemDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (o *orderItemGraphqlMapper) mapResponseOrderItem(orderItem *pb.OrderItemResponse) *model.OrderItemResponse {
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

func (o *orderItemGraphqlMapper) mapResponsesOrderItem(orderItems []*pb.OrderItemResponse) []*model.OrderItemResponse {
	mapped := make([]*model.OrderItemResponse, 0, len(orderItems))
	for _, oi := range orderItems {
		mapped = append(mapped, o.mapResponseOrderItem(oi))
	}
	return mapped
}

func (o *orderItemGraphqlMapper) mapResponseOrderItemDelete(orderItem *pb.OrderItemResponseDeleteAt) *model.OrderItemResponseDeleteAt {
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

func (o *orderItemGraphqlMapper) mapResponsesOrderItemDeleteAt(orderItems []*pb.OrderItemResponseDeleteAt) []*model.OrderItemResponseDeleteAt {
	mapped := make([]*model.OrderItemResponseDeleteAt, 0, len(orderItems))
	for _, oi := range orderItems {
		mapped = append(mapped, o.mapResponseOrderItemDelete(oi))
	}
	return mapped
}
