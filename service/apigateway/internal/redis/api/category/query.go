package category_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	categoryAllCacheKey     = "category:all:page:%d:pageSize:%d:search:%s"
	categoryByIdCacheKey    = "category:id:%d"
	categoryActiveCacheKey  = "category:active:page:%d:pageSize:%d:search:%s"
	categoryTrashedCacheKey = "category:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

type categoryQueryCache struct {
	store *cache.CacheStore
}

func NewCategoryQueryCache(store *cache.CacheStore) *categoryQueryCache {
	return &categoryQueryCache{store: store}
}

func (s *categoryQueryCache) GetCachedCategoriesCache(
	ctx context.Context,
	req *model.FindAllCategoryInput,
) (*model.APIResponsePaginationCategory, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(categoryAllCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationCategory](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryQueryCache) SetCachedCategoriesCache(
	ctx context.Context,
	req *model.FindAllCategoryInput,
	data *model.APIResponsePaginationCategory,
) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(categoryAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryQueryCache) GetCachedCategoryActiveCache(
	ctx context.Context,
	req *model.FindAllCategoryInput,
) (*model.APIResponsePaginationCategoryDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(categoryActiveCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationCategoryDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryQueryCache) SetCachedCategoryActiveCache(
	ctx context.Context,
	req *model.FindAllCategoryInput,
	data *model.APIResponsePaginationCategoryDeleteAt,
) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(categoryActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryQueryCache) GetCachedCategoryTrashedCache(
	ctx context.Context,
	req *model.FindAllCategoryInput,
) (*model.APIResponsePaginationCategoryDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(categoryTrashedCacheKey, req.Page, req.PageSize, search)

	result, found := cache.GetFromCache[model.APIResponsePaginationCategoryDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryQueryCache) SetCachedCategoryTrashedCache(
	ctx context.Context,
	req *model.FindAllCategoryInput,
	data *model.APIResponsePaginationCategoryDeleteAt,
) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(categoryTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryQueryCache) GetCachedCategoryCache(
	ctx context.Context,
	id int,
) (*model.APIResponseCategory, bool) {
	key := fmt.Sprintf(categoryByIdCacheKey, id)
	result, found := cache.GetFromCache[model.APIResponseCategory](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryQueryCache) SetCachedCategoryCache(
	ctx context.Context,
	data *model.APIResponseCategory,
) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(categoryByIdCacheKey, data.Data.ID)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
