package category_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	categoryStatsByIdMonthTotalPriceCacheKey = "category:stats:byid:%d:month:%d:year:%d"
	categoryStatsByIdYearTotalPriceCacheKey  = "category:stats:byid:%d:year:%d"
	categoryStatsByIdMonthPriceCacheKey      = "category:stats:byid:%d:month:%d"
	categoryStatsByIdYearPriceCacheKey       = "category:stats:byid:%d:year:%d"
)

type categoryStatsByIdCache struct {
	store *cache.CacheStore
}

func NewCategoryStatsByIdCache(store *cache.CacheStore) *categoryStatsByIdCache {
	return &categoryStatsByIdCache{store: store}
}

func (s *categoryStatsByIdCache) GetCachedMonthTotalPriceByIdCache(ctx context.Context, req *model.FindYearMonthTotalPriceByIDInput) (*model.APIResponseCategoryMonthlyTotalPrice, bool) {
	key := fmt.Sprintf(categoryStatsByIdMonthTotalPriceCacheKey, req.CategoryID, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryMonthlyTotalPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsByIdCache) SetCachedMonthTotalPriceByIdCache(ctx context.Context, req *model.FindYearMonthTotalPriceByIDInput, data *model.APIResponseCategoryMonthlyTotalPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsByIdMonthTotalPriceCacheKey, req.CategoryID, req.Month, req.Year)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsByIdCache) GetCachedYearTotalPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput) (*model.APIResponseCategoryYearlyTotalPrice, bool) {
	key := fmt.Sprintf(categoryStatsByIdYearTotalPriceCacheKey, req.CategoryID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryYearlyTotalPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsByIdCache) SetCachedYearTotalPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput, data *model.APIResponseCategoryYearlyTotalPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsByIdYearTotalPriceCacheKey, req.CategoryID, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsByIdCache) GetCachedMonthPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput) (*model.APIResponseCategoryMonthPrice, bool) {
	key := fmt.Sprintf(categoryStatsByIdMonthPriceCacheKey, req.CategoryID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryMonthPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsByIdCache) SetCachedMonthPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput, data *model.APIResponseCategoryMonthPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsByIdMonthPriceCacheKey, req.CategoryID, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsByIdCache) GetCachedYearPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput) (*model.APIResponseCategoryYearPrice, bool) {
	key := fmt.Sprintf(categoryStatsByIdYearPriceCacheKey, req.CategoryID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryYearPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsByIdCache) SetCachedYearPriceByIdCache(ctx context.Context, req *model.FindYearCategoryByIDInput, data *model.APIResponseCategoryYearPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsByIdYearPriceCacheKey, req.CategoryID, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
