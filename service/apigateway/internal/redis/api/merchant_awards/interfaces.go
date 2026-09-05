package merchantawards_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantAwardQueryCache interface {
	GetCachedMerchantAwardAll(ctx context.Context, req *model.FindAllMerchantAwardInput) (*model.APIResponsePaginationMerchantAward, bool)
	SetCachedMerchantAwardAll(ctx context.Context, req *model.FindAllMerchantAwardInput, data *model.APIResponsePaginationMerchantAward)

	GetCachedMerchantAwardActive(ctx context.Context, req *model.FindAllMerchantAwardInput) (*model.APIResponsePaginationMerchantAwardDeleteAt, bool)
	SetCachedMerchantAwardActive(ctx context.Context, req *model.FindAllMerchantAwardInput, data *model.APIResponsePaginationMerchantAwardDeleteAt)

	GetCachedMerchantAwardTrashed(ctx context.Context, req *model.FindAllMerchantAwardInput) (*model.APIResponsePaginationMerchantAwardDeleteAt, bool)
	SetCachedMerchantAwardTrashed(ctx context.Context, req *model.FindAllMerchantAwardInput, data *model.APIResponsePaginationMerchantAwardDeleteAt)

	GetCachedMerchantAward(ctx context.Context, id int) (*model.APIResponseMerchantAward, bool)
	SetCachedMerchantAward(ctx context.Context, data *model.APIResponseMerchantAward)
}

type MerchantAwardCommandCache interface {
	DeleteMerchantAwardCache(ctx context.Context, id int)
}
