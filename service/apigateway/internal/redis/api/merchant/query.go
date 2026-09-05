package merchant_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	merchantAllCacheKey      = "merchant:all:page:%d:pageSize:%d:search:%s"
	merchantByIdCacheKey     = "merchant:id:%d"
	merchantActiveCacheKey   = "merchant:active:page:%d:pageSize:%d:search:%s"
	merchantTrashedCacheKey  = "merchant:trashed:page:%d:pageSize:%d:search:%s"
	merchantByUserIdCacheKey = "merchant:user_id:%d"

	ttlDefault = 5 * time.Minute
)

type merchantQueryCache struct {
	store *cache.CacheStore
}

func NewMerchantQueryCache(store *cache.CacheStore) *merchantQueryCache {
	return &merchantQueryCache{store: store}
}

func (m *merchantQueryCache) GetCachedMerchants(ctx context.Context, req *model.FindAllMerchantInput) (*model.APIResponsePaginationMerchant, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantAllCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchant](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantQueryCache) SetCachedMerchants(ctx context.Context, req *model.FindAllMerchantInput, data *model.APIResponsePaginationMerchant) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantAllCacheKey, req.Page, req.PageSize, search)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchantActive(ctx context.Context, req *model.FindAllMerchantInput) (*model.APIResponsePaginationMerchantDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantActiveCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantQueryCache) SetCachedMerchantActive(ctx context.Context, req *model.FindAllMerchantInput, data *model.APIResponsePaginationMerchantDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchantTrashed(ctx context.Context, req *model.FindAllMerchantInput) (*model.APIResponsePaginationMerchantDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantTrashedCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantQueryCache) SetCachedMerchantTrashed(ctx context.Context, req *model.FindAllMerchantInput, data *model.APIResponsePaginationMerchantDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchant(ctx context.Context, id int) (*model.APIResponseMerchant, bool) {
	key := fmt.Sprintf(merchantByIdCacheKey, id)

	result, found := cache.GetFromCache[model.APIResponseMerchant](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantQueryCache) SetCachedMerchant(ctx context.Context, data *model.APIResponseMerchant) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantQueryCache) GetCachedMerchantsByUserId(ctx context.Context, id int) (*model.APIResponsePaginationMerchant, bool) {
	key := fmt.Sprintf(merchantByUserIdCacheKey, id)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchant](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantQueryCache) SetCachedMerchantsByUserId(ctx context.Context, userId int, data *model.APIResponsePaginationMerchant) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantByUserIdCacheKey, userId)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
