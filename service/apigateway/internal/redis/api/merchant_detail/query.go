package merchantdetail_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	merchantDetailAllCacheKey          = "merchant_detail:all:page:%d:pageSize:%d:search:%s"
	merchantDetailByIdCacheKey         = "merchant_detail:id:%d"
	merchantDetailActiveCacheKey       = "merchant_detail:active:page:%d:pageSize:%d:search:%s"
	merchantDetailTrashedCacheKey      = "merchant_detail:trashed:page:%d:pageSize:%d:search:%s"
	merchantDetailRelationByIdCacheKey = "merchant:detail:relation:id:%d"

	ttlDefault = 5 * time.Minute
)

type merchantDetailQueryCache struct {
	store *cache.CacheStore
}

func NewMerchantDetailQueryCache(store *cache.CacheStore) *merchantDetailQueryCache {
	return &merchantDetailQueryCache{
		store: store,
	}
}

func (m *merchantDetailQueryCache) GetCachedMerchantDetailAll(ctx context.Context, req *model.FindAllMerchantDetailInput) (*model.APIResponsePaginationMerchantDetail, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantDetailAllCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantDetail](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantDetailQueryCache) SetCachedMerchantDetailAll(ctx context.Context, req *model.FindAllMerchantDetailInput, data *model.APIResponsePaginationMerchantDetail) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantDetailAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantDetailQueryCache) GetCachedMerchantDetailActive(ctx context.Context, req *model.FindAllMerchantDetailInput) (*model.APIResponsePaginationMerchantDetailDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantDetailActiveCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantDetailDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantDetailQueryCache) SetCachedMerchantDetailActive(ctx context.Context, req *model.FindAllMerchantDetailInput, data *model.APIResponsePaginationMerchantDetailDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantDetailActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantDetailQueryCache) GetCachedMerchantDetailTrashed(ctx context.Context, req *model.FindAllMerchantDetailInput) (*model.APIResponsePaginationMerchantDetailDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantDetailTrashedCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationMerchantDetailDeleteAt](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantDetailQueryCache) SetCachedMerchantDetailTrashed(ctx context.Context, req *model.FindAllMerchantDetailInput, data *model.APIResponsePaginationMerchantDetailDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(merchantDetailTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantDetailQueryCache) GetCachedMerchantDetail(ctx context.Context, id int) (*model.APIResponseMerchantDetail, bool) {
	key := fmt.Sprintf(merchantDetailByIdCacheKey, id)

	result, found := cache.GetFromCache[model.APIResponseMerchantDetail](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantDetailQueryCache) SetCachedMerchantDetail(ctx context.Context, data *model.APIResponseMerchantDetail) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(merchantDetailByIdCacheKey, data.Data.ID)

	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}

func (m *merchantDetailQueryCache) GetCachedMerchantDetailRelation(
	ctx context.Context,
	merchantID int,
) (*model.APIResponseMerchantDetailRelation, bool) {

	key := fmt.Sprintf(merchantDetailRelationByIdCacheKey, merchantID)

	result, found := cache.GetFromCache[model.APIResponseMerchantDetailRelation](ctx, m.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (m *merchantDetailQueryCache) SetCachedMerchantDetailRelation(
	ctx context.Context,
	merchantID int,
	data *model.APIResponseMerchantDetailRelation,
) {
	if merchantID <= 0 || data == nil || data.Data.ID != int32(merchantID) {
		return
	}

	key := fmt.Sprintf(merchantDetailRelationByIdCacheKey, merchantID)
	cache.SetToCache(ctx, m.store, key, data, ttlDefault)
}
