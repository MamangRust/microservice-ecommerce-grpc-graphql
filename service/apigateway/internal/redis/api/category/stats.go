package category_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	categoryStatsMonthTotalPriceCacheKey = "category:stats:month:%d:year:%d"
	categoryStatsYearTotalPriceCacheKey  = "category:stats:year:%d"
	categoryStatsMonthPriceCacheKey      = "category:stats:month:%d"
	categoryStatsYearPriceCacheKey       = "category:stats:year:%d"
)

type categoryStatsCache struct {
	store *cache.CacheStore
}

func NewCategoryStatsCache(store *cache.CacheStore) *categoryStatsCache {
	return &categoryStatsCache{store: store}
}

func (s *categoryStatsCache) GetCachedMonthTotalPriceCache(ctx context.Context, req *model.FindYearMonthTotalPricesInput) (*model.APIResponseCategoryMonthlyTotalPrice, bool) {
	key := fmt.Sprintf(categoryStatsMonthTotalPriceCacheKey, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryMonthlyTotalPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsCache) SetCachedMonthTotalPriceCache(ctx context.Context, req *model.FindYearMonthTotalPricesInput, data *model.APIResponseCategoryMonthlyTotalPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsMonthTotalPriceCacheKey, req.Month, req.Year)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsCache) GetCachedYearTotalPriceCache(ctx context.Context, year int) (*model.APIResponseCategoryYearlyTotalPrice, bool) {
	key := fmt.Sprintf(categoryStatsYearTotalPriceCacheKey, year)
	result, found := cache.GetFromCache[model.APIResponseCategoryYearlyTotalPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsCache) SetCachedYearTotalPriceCache(ctx context.Context, year int, data *model.APIResponseCategoryYearlyTotalPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsYearTotalPriceCacheKey, year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsCache) GetCachedMonthPriceCache(ctx context.Context, year int) (*model.APIResponseCategoryMonthPrice, bool) {
	key := fmt.Sprintf(categoryStatsMonthPriceCacheKey, year)
	result, found := cache.GetFromCache[model.APIResponseCategoryMonthPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsCache) SetCachedMonthPriceCache(ctx context.Context, year int, data *model.APIResponseCategoryMonthPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsMonthPriceCacheKey, year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsCache) GetCachedYearPriceCache(ctx context.Context, year int) (*model.APIResponseCategoryYearPrice, bool) {
	key := fmt.Sprintf(categoryStatsYearPriceCacheKey, year)
	result, found := cache.GetFromCache[model.APIResponseCategoryYearPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsCache) SetCachedYearPriceCache(ctx context.Context, year int, data *model.APIResponseCategoryYearPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsYearPriceCacheKey, year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
