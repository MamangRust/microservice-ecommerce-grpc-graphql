package merchantawards_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	merchantAwardAllCacheKey     = "merchant_award:all:page:%d:pageSize:%d:search:%s"
	merchantAwardByIdCacheKey    = "merchant_award:id:%d"
	merchantAwardActiveCacheKey  = "merchant_award:active:page:%d:pageSize:%d:search:%s"
	merchantAwardTrashedCacheKey = "merchant_award:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

type merchantAwardQueryCache struct {
	store *cache.CacheStore
}

func NewMerchantAwardQueryCache(store *cache.CacheStore) *merchantAwardQueryCache {
	return &merchantAwardQueryCache{
		store: store,
	}
}

func (m *merchantAwardQueryCache) GetCachedMerchantAwardAll(ctx context.Context, req *model.FindAllMerchantAwardInput) (*model.APIResponsePaginationMerchantAward, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantAwardAllCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantAward](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantAwardQueryCache) SetCachedMerchantAwardAll(ctx context.Context, req *model.FindAllMerchantAwardInput, data *model.APIResponsePaginationMerchantAward) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantAwardAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantAwardQueryCache) GetCachedMerchantAwardActive(ctx context.Context, req *model.FindAllMerchantAwardInput) (*model.APIResponsePaginationMerchantAwardDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantAwardActiveCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantAwardDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantAwardQueryCache) SetCachedMerchantAwardActive(ctx context.Context, req *model.FindAllMerchantAwardInput, data *model.APIResponsePaginationMerchantAwardDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantAwardActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantAwardQueryCache) GetCachedMerchantAwardTrashed(ctx context.Context, req *model.FindAllMerchantAwardInput) (*model.APIResponsePaginationMerchantAwardDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantAwardTrashedCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantAwardDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantAwardQueryCache) SetCachedMerchantAwardTrashed(ctx context.Context, req *model.FindAllMerchantAwardInput, data *model.APIResponsePaginationMerchantAwardDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantAwardTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantAwardQueryCache) GetCachedMerchantAward(ctx context.Context, id int) (*model.APIResponseMerchantAward, bool) {
	key := fmt.Sprintf(merchantAwardByIdCacheKey, id)

	result, found := cache.GetFromCache[model.APIResponseMerchantAward](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantAwardQueryCache) SetCachedMerchantAward(ctx context.Context, data *model.APIResponseMerchantAward) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantAwardByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
