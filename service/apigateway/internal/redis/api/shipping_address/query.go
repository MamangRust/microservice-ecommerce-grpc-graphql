package shippingaddress_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	shippingAddressAllCacheKey     = "shipping_address:all:page:%d:pageSize:%d:search:%s"
	shippingAddressActiveCacheKey  = "shipping_address:active:page:%d:pageSize:%d:search:%s"
	shippingAddressTrashedCacheKey = "shipping_address:trashed:page:%d:pageSize:%d:search:%s"

	shippingAddressByOrderIdCacheKey = "shipping_address:order_id:%d"
	shippingAddressByIdCacheKey      = "shipping_address:id:%d"

	ttlDefault = 5 * time.Minute
)

type shippingAddressQueryCache struct {
	store *cache.CacheStore
}

func NewShippingAddressQueryCache(store *cache.CacheStore) *shippingAddressQueryCache {
	return &shippingAddressQueryCache{store: store}
}

func (r *shippingAddressQueryCache) GetShippingAddressAllCache(ctx context.Context, req *model.FindAllShippingRequest) (*model.APIResponsePaginationShipping, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(shippingAddressAllCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationShipping](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *shippingAddressQueryCache) SetShippingAddressAllCache(ctx context.Context, req *model.FindAllShippingRequest, data *model.APIResponsePaginationShipping) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(shippingAddressAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *shippingAddressQueryCache) GetShippingAddressActiveCache(ctx context.Context, req *model.FindAllShippingRequest) (*model.APIResponsePaginationShippingDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(shippingAddressActiveCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationShippingDeleteAt](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *shippingAddressQueryCache) SetShippingAddressActiveCache(ctx context.Context, req *model.FindAllShippingRequest, data *model.APIResponsePaginationShippingDeleteAt) {
	if data == nil {
		return
	}
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(shippingAddressActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *shippingAddressQueryCache) GetShippingAddressTrashedCache(ctx context.Context, req *model.FindAllShippingRequest) (*model.APIResponsePaginationShippingDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(shippingAddressTrashedCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationShippingDeleteAt](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *shippingAddressQueryCache) SetShippingAddressTrashedCache(ctx context.Context, req *model.FindAllShippingRequest, data *model.APIResponsePaginationShippingDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(shippingAddressTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *shippingAddressQueryCache) GetCachedShippingAddressCache(ctx context.Context, shipping_id int) (*model.APIResponseShipping, bool) {
	key := fmt.Sprintf(shippingAddressByIdCacheKey, shipping_id)
	result, found := cache.GetFromCache[model.APIResponseShipping](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *shippingAddressQueryCache) SetCachedShippingAddressCache(ctx context.Context, data *model.APIResponseShipping) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(shippingAddressByIdCacheKey, data.Data.ID)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *shippingAddressQueryCache) GetCachedShippingAddressByOrderCache(ctx context.Context, order_id int) (*model.APIResponseShipping, bool) {
	key := fmt.Sprintf(shippingAddressByOrderIdCacheKey, order_id)
	result, found := cache.GetFromCache[model.APIResponseShipping](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *shippingAddressQueryCache) SetCachedShippingAddressByOrderCache(ctx context.Context, data *model.APIResponseShipping) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(shippingAddressByOrderIdCacheKey, data.Data.OrderID)
	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}
