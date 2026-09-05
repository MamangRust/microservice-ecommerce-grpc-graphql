package orderitem_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	orderItemAllCacheKey     = "order_item:all:page:%d:pageSize:%d:search:%s"
	orderItemActiveCacheKey  = "order_item:active:page:%d:pageSize:%d:search:%s"
	orderItemTrashedCacheKey = "order_item:trashed:page:%d:pageSize:%d:search:%s"
	orderItemByOrderCacheKey = "order_item:order:%d"

	ttlDefault = 5 * time.Minute
)

type orderItemQueryCache struct {
	store *cache.CacheStore
}

func NewOrderItemQueryCache(store *cache.CacheStore) *orderItemQueryCache {
	return &orderItemQueryCache{store: store}
}

func (o *orderItemQueryCache) GetCachedOrderItemsAll(ctx context.Context, req *model.FindAllOrderItemInput) (*model.APIResponsePaginationOrderItem, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderItemAllCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationOrderItem](ctx, o.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (o *orderItemQueryCache) SetCachedOrderItemsAll(ctx context.Context, req *model.FindAllOrderItemInput, data *model.APIResponsePaginationOrderItem) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderItemAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, o.store, key, data, ttlDefault)
}

func (o *orderItemQueryCache) GetCachedOrderItemActive(ctx context.Context, req *model.FindAllOrderItemInput) (*model.APIResponsePaginationOrderItemDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderItemActiveCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationOrderItemDeleteAt](ctx, o.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (o *orderItemQueryCache) SetCachedOrderItemActive(ctx context.Context, req *model.FindAllOrderItemInput, data *model.APIResponsePaginationOrderItemDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderItemActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, o.store, key, data, ttlDefault)
}

func (o *orderItemQueryCache) GetCachedOrderItemTrashed(ctx context.Context, req *model.FindAllOrderItemInput) (*model.APIResponsePaginationOrderItemDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderItemTrashedCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationOrderItemDeleteAt](ctx, o.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (o *orderItemQueryCache) SetCachedOrderItemTrashed(ctx context.Context, req *model.FindAllOrderItemInput, data *model.APIResponsePaginationOrderItemDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(orderItemTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, o.store, key, data, ttlDefault)
}

func (o *orderItemQueryCache) GetCachedOrderItems(ctx context.Context, orderID int) (*model.APIResponsesOrderItem, bool) {
	key := fmt.Sprintf(orderItemByOrderCacheKey, orderID)
	result, found := cache.GetFromCache[model.APIResponsesOrderItem](ctx, o.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (o *orderItemQueryCache) SetCachedOrderItems(ctx context.Context, data *model.APIResponsesOrderItem) {
	if data == nil || len(data.Data) == 0 {
		return
	}
	key := fmt.Sprintf(orderItemByOrderCacheKey, data.Data[0].OrderID)
	cache.SetToCache(ctx, o.store, key, data, ttlDefault)
}
