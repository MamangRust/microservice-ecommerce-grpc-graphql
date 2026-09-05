package order_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type OrderStatsCache interface {
	GetMonthlyTotalRevenueCache(ctx context.Context, req *model.FindYearMonthTotalRevenue) (*model.APIResponseOrderMonthlyTotalRevenue, bool)
	SetMonthlyTotalRevenueCache(ctx context.Context, req *model.FindYearMonthTotalRevenue, data *model.APIResponseOrderMonthlyTotalRevenue)

	GetYearlyTotalRevenueCache(ctx context.Context, year int) (*model.APIResponseOrderYearlyTotalRevenue, bool)
	SetYearlyTotalRevenueCache(ctx context.Context, year int, data *model.APIResponseOrderYearlyTotalRevenue)

	GetMonthlyOrderCache(ctx context.Context, year int) (*model.APIResponseOrderMonthly, bool)
	SetMonthlyOrderCache(ctx context.Context, year int, data *model.APIResponseOrderMonthly)

	GetYearlyOrderCache(ctx context.Context, year int) (*model.APIResponseOrderYearly, bool)
	SetYearlyOrderCache(ctx context.Context, year int, data *model.APIResponseOrderYearly)
}

type OrderStatsByMerchantCache interface {
	GetMonthlyTotalRevenueByMerchantCache(ctx context.Context, req *model.FindYearMonthTotalRevenueByMerchant) (*model.APIResponseOrderMonthlyTotalRevenue, bool)
	SetMonthlyTotalRevenueByMerchantCache(ctx context.Context, req *model.FindYearMonthTotalRevenueByMerchant, data *model.APIResponseOrderMonthlyTotalRevenue)

	GetYearlyTotalRevenueByMerchantCache(ctx context.Context, req *model.FindYearTotalRevenueByMerchant) (*model.APIResponseOrderYearlyTotalRevenue, bool)
	SetYearlyTotalRevenueByMerchantCache(ctx context.Context, req *model.FindYearTotalRevenueByMerchant, data *model.APIResponseOrderYearlyTotalRevenue)

	GetMonthlyOrderByMerchantCache(ctx context.Context, req *model.FindYearOrderByMerchantInput) (*model.APIResponseOrderMonthly, bool)
	SetMonthlyOrderByMerchantCache(ctx context.Context, req *model.FindYearOrderByMerchantInput, data *model.APIResponseOrderMonthly)

	GetYearlyOrderByMerchantCache(ctx context.Context, req *model.FindYearOrderByMerchantInput) (*model.APIResponseOrderYearly, bool)
	SetYearlyOrderByMerchantCache(ctx context.Context, req *model.FindYearOrderByMerchantInput, data *model.APIResponseOrderYearly)
}

type OrderQueryCache interface {
	GetOrderAllCache(ctx context.Context, req *model.FindAllOrderInput) (*model.APIResponsePaginationOrder, bool)
	SetOrderAllCache(ctx context.Context, req *model.FindAllOrderInput, data *model.APIResponsePaginationOrder)

	GetOrderActiveCache(ctx context.Context, req *model.FindAllOrderInput) (*model.APIResponsePaginationOrderDeleteAt, bool)
	SetOrderActiveCache(ctx context.Context, req *model.FindAllOrderInput, data *model.APIResponsePaginationOrderDeleteAt)

	GetOrderTrashedCache(ctx context.Context, req *model.FindAllOrderInput) (*model.APIResponsePaginationOrderDeleteAt, bool)
	SetOrderTrashedCache(ctx context.Context, req *model.FindAllOrderInput, data *model.APIResponsePaginationOrderDeleteAt)

	GetCachedOrderCache(ctx context.Context, order_id int) (*model.APIResponseOrder, bool)
	SetCachedOrderCache(ctx context.Context, data *model.APIResponseOrder)
}

type OrderCommandCache interface {
	DeleteOrderCache(ctx context.Context, orderID int)
}
