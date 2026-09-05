package order_itemgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type OrderItemGraphqlMapper interface {
	ToGraphqlResponseOrderItem(res *pb.ApiResponseOrderItem) *model.APIResponseOrderItem
	ToGraphqlResponsesOrderItem(res *pb.ApiResponsesOrderItem) *model.APIResponsesOrderItem
	ToGraphqlResponseOrderItemDelete(res *pb.ApiResponseOrderItemDelete) *model.APIResponseOrderItemDelete
	ToGraphqlResponseOrderItemAll(res *pb.ApiResponseOrderItemAll) *model.APIResponseOrderItemAll
	ToGraphqlResponsePaginationOrderItem(res *pb.ApiResponsePaginationOrderItem) *model.APIResponsePaginationOrderItem
	ToGraphqlResponsePaginationOrderItemDeleteAt(res *pb.ApiResponsePaginationOrderItemDeleteAt) *model.APIResponsePaginationOrderItemDeleteAt
}
