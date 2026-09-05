package ordergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type OrderGraphqlMapper interface {
	ToGraphqlResponseOrder(res *pb.ApiResponseOrder) *model.APIResponseOrder
	ToGraphqlResponsesOrder(res *pb.ApiResponsesOrder) *model.APIResponsesOrder
	ToGraphqlResponseOrderDeleteAt(res *pb.ApiResponseOrderDeleteAt) *model.APIResponseOrderDeleteAt
	ToGraphqlResponseOrderDelete(res *pb.ApiResponseOrderDelete) *model.APIResponseOrderDelete
	ToGraphqlResponseOrderAll(res *pb.ApiResponseOrderAll) *model.APIResponseOrderAll
	ToGraphqlResponsePaginationOrderDeleteAt(res *pb.ApiResponsePaginationOrderDeleteAt) *model.APIResponsePaginationOrderDeleteAt
	ToGraphqlResponsePaginationOrder(res *pb.ApiResponsePaginationOrder) *model.APIResponsePaginationOrder
	ToGraphqlResponseOrderMonthlyRevenue(res *pb.ApiResponseOrderMonthly) *model.APIResponseOrderMonthly
	ToGraphqlResponseOrderYearlyRevenue(res *pb.ApiResponseOrderYearly) *model.APIResponseOrderYearly
	ToGraphqlResponseOrderMonthlyTotalRevenue(res *pb.ApiResponseOrderMonthlyTotalRevenue) *model.APIResponseOrderMonthlyTotalRevenue
	ToGraphqlResponseOrderYearlyTotalRevenue(res *pb.ApiResponseOrderYearlyTotalRevenue) *model.APIResponseOrderYearlyTotalRevenue
}
