package merchantbusiness_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantBusinessQueryCache interface {
	GetCachedMerchantBusinessAll(ctx context.Context, req *model.FindAllMerchantBusinessInput) (*model.APIResponsePaginationMerchantBusiness, bool)
	SetCachedMerchantBusinessAll(ctx context.Context, req *model.FindAllMerchantBusinessInput, data *model.APIResponsePaginationMerchantBusiness)

	GetCachedMerchantBusinessActive(ctx context.Context, req *model.FindAllMerchantBusinessInput) (*model.APIResponsePaginationMerchantBusinessDeleteAt, bool)
	SetCachedMerchantBusinessActive(ctx context.Context, req *model.FindAllMerchantBusinessInput, data *model.APIResponsePaginationMerchantBusinessDeleteAt)

	GetCachedMerchantBusinessTrashed(ctx context.Context, req *model.FindAllMerchantBusinessInput) (*model.APIResponsePaginationMerchantBusinessDeleteAt, bool)
	SetCachedMerchantBusinessTrashed(ctx context.Context, req *model.FindAllMerchantBusinessInput, data *model.APIResponsePaginationMerchantBusinessDeleteAt)

	GetCachedMerchantBusiness(ctx context.Context, id int) (*model.APIResponseMerchantBusiness, bool)
	SetCachedMerchantBusiness(ctx context.Context, data *model.APIResponseMerchantBusiness)
}

type MerchantBusinessCommandCache interface {
	DeleteMerchantBusinessCache(ctx context.Context, id int)
}
