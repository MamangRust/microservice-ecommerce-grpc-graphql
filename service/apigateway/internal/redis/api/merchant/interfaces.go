package merchant_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantQueryCache interface {
	GetCachedMerchants(ctx context.Context, req *model.FindAllMerchantInput) (*model.APIResponsePaginationMerchant, bool)
	SetCachedMerchants(ctx context.Context, req *model.FindAllMerchantInput, data *model.APIResponsePaginationMerchant)

	GetCachedMerchantActive(ctx context.Context, req *model.FindAllMerchantInput) (*model.APIResponsePaginationMerchantDeleteAt, bool)
	SetCachedMerchantActive(ctx context.Context, req *model.FindAllMerchantInput, data *model.APIResponsePaginationMerchantDeleteAt)

	GetCachedMerchantTrashed(ctx context.Context, req *model.FindAllMerchantInput) (*model.APIResponsePaginationMerchantDeleteAt, bool)
	SetCachedMerchantTrashed(ctx context.Context, req *model.FindAllMerchantInput, data *model.APIResponsePaginationMerchantDeleteAt)

	GetCachedMerchant(ctx context.Context, id int) (*model.APIResponseMerchant, bool)
	SetCachedMerchant(ctx context.Context, data *model.APIResponseMerchant)

	GetCachedMerchantsByUserId(ctx context.Context, id int) (*model.APIResponsePaginationMerchant, bool)
	SetCachedMerchantsByUserId(ctx context.Context, userId int, data *model.APIResponsePaginationMerchant)
}

type MerchantCommandCache interface {
	DeleteCachedMerchant(ctx context.Context, id int)
}
