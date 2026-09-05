package merchantpolicies_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantPolicyQueryCache interface {
	GetCachedMerchantPolicyAll(ctx context.Context, req *model.FindAllMerchantPoliciesInput) (*model.APIResponsePaginationMerchantPolicy, bool)
	SetCachedMerchantPolicyAll(ctx context.Context, req *model.FindAllMerchantPoliciesInput, data *model.APIResponsePaginationMerchantPolicy)

	GetCachedMerchantPolicyActive(ctx context.Context, req *model.FindAllMerchantPoliciesInput) (*model.APIResponsePaginationMerchantPolicyDeleteAt, bool)
	SetCachedMerchantPolicyActive(ctx context.Context, req *model.FindAllMerchantPoliciesInput, data *model.APIResponsePaginationMerchantPolicyDeleteAt)

	GetCachedMerchantPolicyTrashed(ctx context.Context, req *model.FindAllMerchantPoliciesInput) (*model.APIResponsePaginationMerchantPolicyDeleteAt, bool)
	SetCachedMerchantPolicyTrashed(ctx context.Context, req *model.FindAllMerchantPoliciesInput, data *model.APIResponsePaginationMerchantPolicyDeleteAt)

	GetCachedMerchantPolicy(ctx context.Context, id int) (*model.APIResponseMerchantPolicy, bool)
	SetCachedMerchantPolicy(ctx context.Context, data *model.APIResponseMerchantPolicy)
}

type MerchantPolicyCommandCache interface {
	DeleteMerchantPolicyCache(ctx context.Context, merchantID int)
}
