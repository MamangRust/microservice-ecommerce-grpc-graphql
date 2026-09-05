package slider_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	sliderAllCacheKey     = "slider:all:page:%d:pageSize:%d:search:%s"
	sliderActiveCacheKey  = "slider:active:page:%d:pageSize:%d:search:%s"
	sliderTrashedCacheKey = "slider:trashed:page:%d:pageSize:%d:search:%s"
	sliderIdKey           = "slider:id:%d"

	ttlDefault = 5 * time.Minute
)

type sliderQueryCache struct {
	store *cache.CacheStore
}

func NewSliderQueryCache(store *cache.CacheStore) *sliderQueryCache {
	return &sliderQueryCache{store: store}
}

func (s *sliderQueryCache) GetSliderAllCache(ctx context.Context, req *model.FindAllSliderRequest) (*model.APIResponsePaginationSlider, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(sliderAllCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationSlider](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *sliderQueryCache) SetSliderAllCache(ctx context.Context, req *model.FindAllSliderRequest, data *model.APIResponsePaginationSlider) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(sliderAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *sliderQueryCache) GetSliderActiveCache(ctx context.Context, req *model.FindAllSliderRequest) (*model.APIResponsePaginationSliderDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(sliderActiveCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationSliderDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *sliderQueryCache) SetSliderActiveCache(ctx context.Context, req *model.FindAllSliderRequest, data *model.APIResponsePaginationSliderDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(sliderActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *sliderQueryCache) GetSliderTrashedCache(ctx context.Context, req *model.FindAllSliderRequest) (*model.APIResponsePaginationSliderDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(sliderTrashedCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationSliderDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *sliderQueryCache) SetSliderTrashedCache(ctx context.Context, req *model.FindAllSliderRequest, data *model.APIResponsePaginationSliderDeleteAt) {
	if data == nil {
		return
	}
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(sliderTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *sliderQueryCache) GetCachedSliderById(ctx context.Context, id int) (*model.APIResponseSlider, bool) {
	key := fmt.Sprintf(sliderIdKey, id)

	result, found := cache.GetFromCache[model.APIResponseSlider](ctx, s.store, key)
	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *sliderQueryCache) SetCachedSliderById(ctx context.Context, data *model.APIResponseSlider) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(sliderIdKey, data.Data.ID)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
