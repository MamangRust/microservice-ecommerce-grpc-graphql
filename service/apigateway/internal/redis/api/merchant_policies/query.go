package merchantpolicies_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	merchantPolicyAllCacheKey     = "merchant_policy:all:page:%d:pageSize:%d:search:%s"
	merchantPolicyByIdCacheKey    = "merchant_policy:id:%d"
	merchantPolicyActiveCacheKey  = "merchant_policy:active:page:%d:pageSize:%d:search:%s"
	merchantPolicyTrashedCacheKey = "merchant_policy:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

type merchantPolicyQueryCache struct {
	store *cache.CacheStore
}

func NewMerchantPolicyQueryCache(store *cache.CacheStore) *merchantPolicyQueryCache {
	return &merchantPolicyQueryCache{
		store: store,
	}
}

func (m *merchantPolicyQueryCache) GetCachedMerchantPolicyAll(ctx context.Context, req *model.FindAllMerchantPoliciesInput) (*model.APIResponsePaginationMerchantPolicy, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantPolicyAllCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantPolicy](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantPolicyQueryCache) SetCachedMerchantPolicyAll(ctx context.Context, req *model.FindAllMerchantPoliciesInput, data *model.APIResponsePaginationMerchantPolicy) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantPolicyAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantPolicyQueryCache) GetCachedMerchantPolicyActive(ctx context.Context, req *model.FindAllMerchantPoliciesInput) (*model.APIResponsePaginationMerchantPolicyDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantPolicyActiveCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantPolicyDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantPolicyQueryCache) SetCachedMerchantPolicyActive(ctx context.Context, req *model.FindAllMerchantPoliciesInput, data *model.APIResponsePaginationMerchantPolicyDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantPolicyActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantPolicyQueryCache) GetCachedMerchantPolicyTrashed(ctx context.Context, req *model.FindAllMerchantPoliciesInput) (*model.APIResponsePaginationMerchantPolicyDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}
	key := fmt.Sprintf(merchantPolicyTrashedCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantPolicyDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantPolicyQueryCache) SetCachedMerchantPolicyTrashed(ctx context.Context, req *model.FindAllMerchantPoliciesInput, data *model.APIResponsePaginationMerchantPolicyDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantPolicyTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantPolicyQueryCache) GetCachedMerchantPolicy(ctx context.Context, id int) (*model.APIResponseMerchantPolicy, bool) {
	key := fmt.Sprintf(merchantPolicyByIdCacheKey, id)

	result, found := cache.GetFromCache[model.APIResponseMerchantPolicy](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantPolicyQueryCache) SetCachedMerchantPolicy(ctx context.Context, data *model.APIResponseMerchantPolicy) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantPolicyByIdCacheKey, data.Data.ID)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
