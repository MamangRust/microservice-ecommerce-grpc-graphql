package order_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	orderAllCacheKey     = "order:all:page:%d:pageSize:%d:search:%s"
	orderByIdCacheKey    = "order:id:%d"
	orderActiveCacheKey  = "order:active:page:%d:pageSize:%d:search:%s"
	orderTrashedCacheKey = "order:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

type orderQueryCache struct {
	store *cache.CacheStore
}

func NewOrderQueryCache(store *cache.CacheStore) *orderQueryCache {
	return &orderQueryCache{store: store}
}

func (s *orderQueryCache) GetOrderAllCache(ctx context.Context, req *model.FindAllOrderInput) (*model.APIResponsePaginationOrder, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderAllCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationOrder](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetOrderAllCache(ctx context.Context, req *model.FindAllOrderInput, data *model.APIResponsePaginationOrder) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderQueryCache) GetOrderActiveCache(ctx context.Context, req *model.FindAllOrderInput) (*model.APIResponsePaginationOrderDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderActiveCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationOrderDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetOrderActiveCache(ctx context.Context, req *model.FindAllOrderInput, data *model.APIResponsePaginationOrderDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderQueryCache) GetOrderTrashedCache(ctx context.Context, req *model.FindAllOrderInput) (*model.APIResponsePaginationOrderDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderTrashedCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationOrderDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetOrderTrashedCache(ctx context.Context, req *model.FindAllOrderInput, data *model.APIResponsePaginationOrderDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderQueryCache) GetCachedOrderCache(ctx context.Context, order_id int) (*model.APIResponseOrder, bool) {
	key := fmt.Sprintf(orderByIdCacheKey, order_id)
	result, found := cache.GetFromCache[model.APIResponseOrder](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetCachedOrderCache(ctx context.Context, data *model.APIResponseOrder) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(orderByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
