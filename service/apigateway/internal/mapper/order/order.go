package ordergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type orderGraphqlMapper struct{}

func NewOrderGraphqlMapper() *orderGraphqlMapper {
	return &orderGraphqlMapper{}
}

func (o *orderGraphqlMapper) ToGraphqlResponseOrder(res *pb.ApiResponseOrder) *model.APIResponseOrder {
	return &model.APIResponseOrder{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponseOrder(res.Data),
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponsesOrder(res *pb.ApiResponsesOrder) *model.APIResponsesOrder {
	return &model.APIResponsesOrder{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponsesOrder(res.Data),
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponseOrderDeleteAt(res *pb.ApiResponseOrderDeleteAt) *model.APIResponseOrderDeleteAt {
	return &model.APIResponseOrderDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponseOrderDeleteAt(res.Data),
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponseOrderDelete(res *pb.ApiResponseOrderDelete) *model.APIResponseOrderDelete {
	return &model.APIResponseOrderDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponseOrderAll(res *pb.ApiResponseOrderAll) *model.APIResponseOrderAll {
	return &model.APIResponseOrderAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponsePaginationOrderDeleteAt(
	res *pb.ApiResponsePaginationOrderDeleteAt,
) *model.APIResponsePaginationOrderDeleteAt {
	return &model.APIResponsePaginationOrderDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       o.mapResponsesOrderDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponsePaginationOrder(
	res *pb.ApiResponsePaginationOrder,
) *model.APIResponsePaginationOrder {
	return &model.APIResponsePaginationOrder{
		Status:     res.Status,
		Message:    res.Message,
		Data:       o.mapResponsesOrder(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponseOrderYearlyTotalRevenue(
	res *pb.ApiResponseOrderYearlyTotalRevenue,
) *model.APIResponseOrderYearlyTotalRevenue {

	return &model.APIResponseOrderYearlyTotalRevenue{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponseOrderYearlyTotalRevenues(res.Data),
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponseOrderMonthlyTotalRevenue(
	res *pb.ApiResponseOrderMonthlyTotalRevenue,
) *model.APIResponseOrderMonthlyTotalRevenue {

	return &model.APIResponseOrderMonthlyTotalRevenue{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponseOrderMonthlyTotalRevenues(res.Data),
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponseOrderMonthlyRevenue(
	res *pb.ApiResponseOrderMonthly,
) *model.APIResponseOrderMonthly {
	return &model.APIResponseOrderMonthly{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponsesOrderMonthlyPrices(res.Data),
	}
}

func (o *orderGraphqlMapper) ToGraphqlResponseOrderYearlyRevenue(
	res *pb.ApiResponseOrderYearly,
) *model.APIResponseOrderYearly {

	return &model.APIResponseOrderYearly{
		Status:  res.Status,
		Message: res.Message,
		Data:    o.mapResponsesOrderYearlyPrices(res.Data),
	}
}

func (o *orderGraphqlMapper) mapResponseOrder(order *pb.OrderResponse) *model.OrderResponse {
	return &model.OrderResponse{
		ID:         int32(order.Id),
		MerchantID: int32(order.MerchantId),
		UserID:     int32(order.UserId),
		TotalPrice: int32(order.TotalPrice),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func (o *orderGraphqlMapper) mapResponsesOrder(orders []*pb.OrderResponse) []*model.OrderResponse {
	mapped := make([]*model.OrderResponse, 0, len(orders))
	for _, order := range orders {
		mapped = append(mapped, o.mapResponseOrder(order))
	}
	return mapped
}

func (o *orderGraphqlMapper) mapResponseOrderDeleteAt(order *pb.OrderResponseDeleteAt) *model.OrderResponseDeleteAt {
	var deletedAt *string
	if order.DeletedAt != nil {
		deletedAt = &order.DeletedAt.Value
	}

	return &model.OrderResponseDeleteAt{
		ID:         int32(order.Id),
		MerchantID: int32(order.MerchantId),
		UserID:     int32(order.UserId),
		TotalPrice: int32(order.TotalPrice),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		DeletedAt:  deletedAt,
	}
}

func (o *orderGraphqlMapper) mapResponsesOrderDeleteAt(orders []*pb.OrderResponseDeleteAt) []*model.OrderResponseDeleteAt {
	mapped := make([]*model.OrderResponseDeleteAt, 0, len(orders))
	for _, order := range orders {
		mapped = append(mapped, o.mapResponseOrderDeleteAt(order))
	}
	return mapped
}

func (o *orderGraphqlMapper) mapResponseOrderMonthlyPrice(res *pb.OrderMonthlyResponse) *model.OrderMonthlyResponse {
	return &model.OrderMonthlyResponse{
		Month:          res.Month,
		OrderCount:     int32(res.OrderCount),
		TotalRevenue:   int32(res.TotalRevenue),
		TotalItemsSold: int32(res.TotalItemsSold),
	}
}

func (o *orderGraphqlMapper) mapResponsesOrderMonthlyPrices(res []*pb.OrderMonthlyResponse) []*model.OrderMonthlyResponse {
	mapped := make([]*model.OrderMonthlyResponse, 0, len(res))
	for _, r := range res {
		mapped = append(mapped, o.mapResponseOrderMonthlyPrice(r))
	}
	return mapped
}

func (o *orderGraphqlMapper) mapResponseOrderYearlyPrice(res *pb.OrderYearlyResponse) *model.OrderYearlyResponse {
	return &model.OrderYearlyResponse{
		Year:               res.Year,
		OrderCount:         int32(res.OrderCount),
		TotalRevenue:       int32(res.TotalRevenue),
		TotalItemsSold:     int32(res.TotalItemsSold),
		ActiveCashiers:     int32(res.ActiveCashiers),
		UniqueProductsSold: int32(res.UniqueProductsSold),
	}
}

func (o *orderGraphqlMapper) mapResponsesOrderYearlyPrices(res []*pb.OrderYearlyResponse) []*model.OrderYearlyResponse {
	mapped := make([]*model.OrderYearlyResponse, 0, len(res))
	for _, r := range res {
		mapped = append(mapped, o.mapResponseOrderYearlyPrice(r))
	}
	return mapped
}

func (o *orderGraphqlMapper) mapResponseOrderMonthlyTotalRevenue(res *pb.OrderMonthlyTotalRevenueResponse) *model.OrderMonthlyTotalRevenueResponse {
	return &model.OrderMonthlyTotalRevenueResponse{
		Year:           res.Year,
		Month:          res.Month,
		TotalRevenue:   int32(res.TotalRevenue),
		TotalItemsSold: int32(res.TotalItemsSold),
	}
}

func (o *orderGraphqlMapper) mapResponseOrderMonthlyTotalRevenues(res []*pb.OrderMonthlyTotalRevenueResponse) []*model.OrderMonthlyTotalRevenueResponse {
	mapped := make([]*model.OrderMonthlyTotalRevenueResponse, 0, len(res))
	for _, r := range res {
		mapped = append(mapped, o.mapResponseOrderMonthlyTotalRevenue(r))
	}
	return mapped
}

func (o *orderGraphqlMapper) mapResponseOrderYearlyTotalRevenue(res *pb.OrderYearlyTotalRevenueResponse) *model.OrderYearlyTotalRevenueResponse {
	return &model.OrderYearlyTotalRevenueResponse{
		Year:         res.Year,
		TotalRevenue: int32(res.TotalRevenue),
	}
}

func (o *orderGraphqlMapper) mapResponseOrderYearlyTotalRevenues(res []*pb.OrderYearlyTotalRevenueResponse) []*model.OrderYearlyTotalRevenueResponse {
	mapped := make([]*model.OrderYearlyTotalRevenueResponse, 0, len(res))
	for _, r := range res {
		mapped = append(mapped, o.mapResponseOrderYearlyTotalRevenue(r))
	}
	return mapped
}
