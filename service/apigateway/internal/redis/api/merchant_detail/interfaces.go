package merchantdetail_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantDetailQueryCache interface {
	GetCachedMerchantDetailAll(ctx context.Context, req *model.FindAllMerchantDetailInput) (*model.APIResponsePaginationMerchantDetail, bool)
	SetCachedMerchantDetailAll(ctx context.Context, req *model.FindAllMerchantDetailInput, data *model.APIResponsePaginationMerchantDetail)

	GetCachedMerchantDetailActive(ctx context.Context, req *model.FindAllMerchantDetailInput) (*model.APIResponsePaginationMerchantDetailDeleteAt, bool)
	SetCachedMerchantDetailActive(ctx context.Context, req *model.FindAllMerchantDetailInput, data *model.APIResponsePaginationMerchantDetailDeleteAt)

	GetCachedMerchantDetailTrashed(ctx context.Context, req *model.FindAllMerchantDetailInput) (*model.APIResponsePaginationMerchantDetailDeleteAt, bool)
	SetCachedMerchantDetailTrashed(ctx context.Context, req *model.FindAllMerchantDetailInput, data *model.APIResponsePaginationMerchantDetailDeleteAt)

	GetCachedMerchantDetail(ctx context.Context, id int) (*model.APIResponseMerchantDetail, bool)
	SetCachedMerchantDetail(ctx context.Context, data *model.APIResponseMerchantDetail)

	GetCachedMerchantDetailRelation(
		ctx context.Context,
		merchantID int,
	) (*model.APIResponseMerchantDetailRelation, bool)

	SetCachedMerchantDetailRelation(
		ctx context.Context,
		merchantID int,
		data *model.APIResponseMerchantDetailRelation,
	)
}

type MerchantDetailCommandCache interface {
	DeleteMerchantDetailCache(ctx context.Context, id int)
}
