package ordergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/order"
)

type OrderGraphqlMapper interface {
	ToGraphqlResponseOrder(res *pb_order.ApiResponseOrder) *model.APIResponseOrder
	ToGraphqlResponsesOrder(res *pb_order.ApiResponsesOrder) *model.APIResponsesOrder
	ToGraphqlResponseOrderDeleteAt(res *pb_order.ApiResponseOrderDeleteAt) *model.APIResponseOrderDeleteAt
	ToGraphqlResponseOrderDelete(res *pb_order.ApiResponseOrderDelete) *model.APIResponseOrderDelete
	ToGraphqlResponseOrderAll(res *pb_order.ApiResponseOrderAll) *model.APIResponseOrderAll
	ToGraphqlResponsePaginationOrderDeleteAt(res *pb_order.ApiResponsePaginationOrderDeleteAt) *model.APIResponsePaginationOrderDeleteAt
	ToGraphqlResponsePaginationOrder(res *pb_order.ApiResponsePaginationOrder) *model.APIResponsePaginationOrder
	ToGraphqlResponseOrderMonthlyRevenue(res *pb_order.ApiResponseOrderMonthly) *model.APIResponseOrderMonthly
	ToGraphqlResponseOrderYearlyRevenue(res *pb_order.ApiResponseOrderYearly) *model.APIResponseOrderYearly
	ToGraphqlResponseOrderMonthlyTotalRevenue(res *pb_order.ApiResponseOrderMonthlyTotalRevenue) *model.APIResponseOrderMonthlyTotalRevenue
	ToGraphqlResponseOrderYearlyTotalRevenue(res *pb_order.ApiResponseOrderYearlyTotalRevenue) *model.APIResponseOrderYearlyTotalRevenue
}
