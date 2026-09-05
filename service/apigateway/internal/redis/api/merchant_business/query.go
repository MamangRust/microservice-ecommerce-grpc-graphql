package merchantbusiness_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	merchantBusinessAllCacheKey     = "merchant_business:all:page:%d:pageSize:%d:search:%s"
	merchantBusinessByIdCacheKey    = "merchant_business:id:%d"
	merchantBusinessActiveCacheKey  = "merchant_business:active:page:%d:pageSize:%d:search:%s"
	merchantBusinessTrashedCacheKey = "merchant_business:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

type merchantBusinessQueryCache struct {
	store *cache.CacheStore
}

func NewMerchantBusinessQueryCache(store *cache.CacheStore) *merchantBusinessQueryCache {
	return &merchantBusinessQueryCache{
		store: store,
	}
}

func (m *merchantBusinessQueryCache) GetCachedMerchantBusinessAll(ctx context.Context, req *model.FindAllMerchantBusinessInput) (*model.APIResponsePaginationMerchantBusiness, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantBusinessAllCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantBusiness](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantBusinessQueryCache) SetCachedMerchantBusinessAll(ctx context.Context, req *model.FindAllMerchantBusinessInput, data *model.APIResponsePaginationMerchantBusiness) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantBusinessAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantBusinessQueryCache) GetCachedMerchantBusinessActive(ctx context.Context, req *model.FindAllMerchantBusinessInput) (*model.APIResponsePaginationMerchantBusinessDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantBusinessActiveCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantBusinessDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantBusinessQueryCache) SetCachedMerchantBusinessActive(ctx context.Context, req *model.FindAllMerchantBusinessInput, data *model.APIResponsePaginationMerchantBusinessDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantBusinessActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantBusinessQueryCache) GetCachedMerchantBusinessTrashed(ctx context.Context, req *model.FindAllMerchantBusinessInput) (*model.APIResponsePaginationMerchantBusinessDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantBusinessTrashedCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantBusinessDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantBusinessQueryCache) SetCachedMerchantBusinessTrashed(ctx context.Context, req *model.FindAllMerchantBusinessInput, data *model.APIResponsePaginationMerchantBusinessDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantBusinessTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantBusinessQueryCache) GetCachedMerchantBusiness(ctx context.Context, id int) (*model.APIResponseMerchantBusiness, bool) {
	key := fmt.Sprintf(merchantBusinessByIdCacheKey, id)

	result, found := cache.GetFromCache[model.APIResponseMerchantBusiness](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantBusinessQueryCache) SetCachedMerchantBusiness(ctx context.Context, data *model.APIResponseMerchantBusiness) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantBusinessByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
