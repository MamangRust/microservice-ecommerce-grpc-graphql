package category_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type CategoryQueryCache interface {
	GetCachedCategoriesCache(ctx context.Context, req *model.FindAllCategoryInput) (*model.APIResponsePaginationCategory, bool)
	SetCachedCategoriesCache(ctx context.Context, req *model.FindAllCategoryInput, data *model.APIResponsePaginationCategory)
	GetCachedCategoryActiveCache(ctx context.Context, req *model.FindAllCategoryInput) (*model.APIResponsePaginationCategoryDeleteAt, bool)
	SetCachedCategoryActiveCache(ctx context.Context, req *model.FindAllCategoryInput, data *model.APIResponsePaginationCategoryDeleteAt)
	GetCachedCategoryTrashedCache(ctx context.Context, req *model.FindAllCategoryInput) (*model.APIResponsePaginationCategoryDeleteAt, bool)
	SetCachedCategoryTrashedCache(ctx context.Context, req *model.FindAllCategoryInput, data *model.APIResponsePaginationCategoryDeleteAt)
	GetCachedCategoryCache(ctx context.Context, id int) (*model.APIResponseCategory, bool)
	SetCachedCategoryCache(ctx context.Context, data *model.APIResponseCategory)
}

type CategoryCommandCache interface {
	DeleteCachedCategoryCache(ctx context.Context, id int)
}

type CategoryStatsCache interface {
	GetCachedMonthTotalPriceCache(ctx context.Context, req *model.FindYearMonthTotalPricesInput) (*model.APIResponseCategoryMonthlyTotalPrice, bool)
	SetCachedMonthTotalPriceCache(ctx context.Context, req *model.FindYearMonthTotalPricesInput, data *model.APIResponseCategoryMonthlyTotalPrice)

	GetCachedYearTotalPriceCache(ctx context.Context, year int) (*model.APIResponseCategoryYearlyTotalPrice, bool)
	SetCachedYearTotalPriceCache(ctx context.Context, year int, data *model.APIResponseCategoryYearlyTotalPrice)

	GetCachedMonthPriceCache(ctx context.Context, year int) (*model.APIResponseCategoryMonthPrice, bool)
	SetCachedMonthPriceCache(ctx context.Context, year int, data *model.APIResponseCategoryMonthPrice)

	GetCachedYearPriceCache(ctx context.Context, year int) (*model.APIResponseCategoryYearPrice, bool)
	SetCachedYearPriceCache(ctx context.Context, year int, data *model.APIResponseCategoryYearPrice)
}

type CategoryStatsByIdCache interface {
	GetCachedMonthTotalPriceByIdCache(ctx context.Context, req *model.FindYearMonthTotalPriceByIDInput) (*model.APIResponseCategoryMonthlyTotalPrice, bool)
	SetCachedMonthTotalPriceByIdCache(ctx context.Context, req *model.FindYearMonthTotalPriceByIDInput, data *model.APIResponseCategoryMonthlyTotalPrice)

	GetCachedYearTotalPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput) (*model.APIResponseCategoryYearlyTotalPrice, bool)
	SetCachedYearTotalPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput, data *model.APIResponseCategoryYearlyTotalPrice)

	GetCachedMonthPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput) (*model.APIResponseCategoryMonthPrice, bool)
	SetCachedMonthPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput, data *model.APIResponseCategoryMonthPrice)

	GetCachedYearPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput) (*model.APIResponseCategoryYearPrice, bool)
	SetCachedYearPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput, data *model.APIResponseCategoryYearPrice)
}

type CategoryStatsByMerchantCache interface {
	GetCachedMonthTotalPriceByMerchantCache(ctx context.Context, req *model.FindYearMonthTotalPriceByMerchantInput) (*model.APIResponseCategoryMonthlyTotalPrice, bool)
	SetCachedMonthTotalPriceByMerchantCache(ctx context.Context, req *model.FindYearMonthTotalPriceByMerchantInput, data *model.APIResponseCategoryMonthlyTotalPrice)

	GetCachedYearTotalPriceByMerchantCache(ctx context.Context, req *model.FindYearTotalPriceByMerchantInput) (*model.APIResponseCategoryYearlyTotalPrice, bool)
	SetCachedYearTotalPriceByMerchantCache(ctx context.Context, req *model.FindYearTotalPriceByMerchantInput, data *model.APIResponseCategoryYearlyTotalPrice)

	GetCachedMonthPriceByMerchantCache(ctx context.Context, req *model.FindYearCategoryByMerchantInput) (*model.APIResponseCategoryMonthPrice, bool)
	SetCachedMonthPriceByMerchantCache(ctx context.Context, req *model.FindYearCategoryByMerchantInput, data *model.APIResponseCategoryMonthPrice)

	GetCachedYearPriceByMerchantCache(ctx context.Context, req *model.FindYearCategoryByMerchantInput) (*model.APIResponseCategoryYearPrice, bool)
	SetCachedYearPriceByMerchantCache(ctx context.Context, req *model.FindYearCategoryByMerchantInput, data *model.APIResponseCategoryYearPrice)
}
