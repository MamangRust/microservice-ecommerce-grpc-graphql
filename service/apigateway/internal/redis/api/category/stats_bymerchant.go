package category_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	categoryStatsByMerchantMonthTotalPriceCacheKey = "category:stats:bymerchant:%d:month:%d:year:%d"
	categoryStatsByMerchantYearTotalPriceCacheKey  = "category:stats:bymerchant:%d:year:%d"
	categoryStatsByMerchantMonthPriceCacheKey      = "category:stats:bymerchant:%d:month:%d"
	categoryStatsByMerchantYearPriceCacheKey       = "category:stats:bymerchant:%d:year:%d"
)

type categoryStatsByMerchantCache struct {
	store *cache.CacheStore
}

func NewCategoryStatsByMerchantCache(store *cache.CacheStore) *categoryStatsByMerchantCache {
	return &categoryStatsByMerchantCache{store: store}
}

func (s *categoryStatsByMerchantCache) GetCachedMonthTotalPriceByMerchantCache(ctx context.Context, req *model.FindYearMonthTotalPriceByMerchantInput) (*model.APIResponseCategoryMonthlyTotalPrice, bool) {
	key := fmt.Sprintf(categoryStatsByMerchantMonthTotalPriceCacheKey, req.MerchantID, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryMonthlyTotalPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsByMerchantCache) SetCachedMonthTotalPriceByMerchantCache(ctx context.Context, req *model.FindYearMonthTotalPriceByMerchantInput, data *model.APIResponseCategoryMonthlyTotalPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsByMerchantMonthTotalPriceCacheKey, req.MerchantID, req.Month, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsByMerchantCache) GetCachedYearTotalPriceByMerchantCache(ctx context.Context, req *model.FindYearTotalPriceByMerchantInput) (*model.APIResponseCategoryYearlyTotalPrice, bool) {
	key := fmt.Sprintf(categoryStatsByMerchantYearTotalPriceCacheKey, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryYearlyTotalPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsByMerchantCache) SetCachedYearTotalPriceByMerchantCache(ctx context.Context, req *model.FindYearTotalPriceByMerchantInput, data *model.APIResponseCategoryYearlyTotalPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsByMerchantYearTotalPriceCacheKey, req.MerchantID, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsByMerchantCache) GetCachedMonthPriceByMerchantCache(ctx context.Context, req *model.FindYearCategoryByMerchantInput) (*model.APIResponseCategoryMonthPrice, bool) {
	key := fmt.Sprintf(categoryStatsByMerchantMonthPriceCacheKey, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryMonthPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsByMerchantCache) SetCachedMonthPriceByMerchantCache(ctx context.Context, req *model.FindYearCategoryByMerchantInput, data *model.APIResponseCategoryMonthPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsByMerchantMonthPriceCacheKey, req.MerchantID, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *categoryStatsByMerchantCache) GetCachedYearPriceByMerchantCache(ctx context.Context, req *model.FindYearCategoryByMerchantInput) (*model.APIResponseCategoryYearPrice, bool) {
	key := fmt.Sprintf(categoryStatsByMerchantYearPriceCacheKey, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseCategoryYearPrice](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *categoryStatsByMerchantCache) SetCachedYearPriceByMerchantCache(ctx context.Context, req *model.FindYearCategoryByMerchantInput, data *model.APIResponseCategoryYearPrice) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(categoryStatsByMerchantYearPriceCacheKey, req.MerchantID, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
