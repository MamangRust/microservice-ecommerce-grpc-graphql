package order_itemgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/order_item"
)

type OrderItemGraphqlMapper interface {
	ToGraphqlResponseOrderItem(res *pb_order_item.ApiResponseOrderItem) *model.APIResponseOrderItem
	ToGraphqlResponsesOrderItem(res *pb_order_item.ApiResponsesOrderItem) *model.APIResponsesOrderItem
	ToGraphqlResponseOrderItemDelete(res *pb_order_item.ApiResponseOrderItemDelete) *model.APIResponseOrderItemDelete
	ToGraphqlResponseOrderItemAll(res *pb_order_item.ApiResponseOrderItemAll) *model.APIResponseOrderItemAll
	ToGraphqlResponsePaginationOrderItem(res *pb_order_item.ApiResponsePaginationOrderItem) *model.APIResponsePaginationOrderItem
	ToGraphqlResponsePaginationOrderItemDeleteAt(res *pb_order_item.ApiResponsePaginationOrderItemDeleteAt) *model.APIResponsePaginationOrderItemDeleteAt
}
