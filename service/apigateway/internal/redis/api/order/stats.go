package order_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	monthlyTotalRevenueCacheKey = "order:monthly:totalRevenue:month:%d:year:%d"
	yearlyTotalRevenueCacheKey  = "order:yearly:totalRevenue:year:%d"
	monthlyOrderCacheKey        = "order:monthly:order:month:%d"
	yearlyOrderCacheKey         = "order:yearly:order:year:%d"
)

type orderStatsCache struct {
	store *cache.CacheStore
}

func NewOrderStatsCache(store *cache.CacheStore) *orderStatsCache {
	return &orderStatsCache{store: store}
}

func (s *orderStatsCache) GetMonthlyTotalRevenueCache(ctx context.Context, req *model.FindYearMonthTotalRevenue) (*model.APIResponseOrderMonthlyTotalRevenue, bool) {
	key := fmt.Sprintf(monthlyTotalRevenueCacheKey, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseOrderMonthlyTotalRevenue](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsCache) SetMonthlyTotalRevenueCache(ctx context.Context, req *model.FindYearMonthTotalRevenue, data *model.APIResponseOrderMonthlyTotalRevenue) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthlyTotalRevenueCacheKey, req.Month, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderStatsCache) GetYearlyTotalRevenueCache(ctx context.Context, year int) (*model.APIResponseOrderYearlyTotalRevenue, bool) {
	key := fmt.Sprintf(yearlyTotalRevenueCacheKey, year)
	result, found := cache.GetFromCache[model.APIResponseOrderYearlyTotalRevenue](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsCache) SetYearlyTotalRevenueCache(ctx context.Context, year int, data *model.APIResponseOrderYearlyTotalRevenue) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearlyTotalRevenueCacheKey, year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderStatsCache) GetMonthlyOrderCache(ctx context.Context, year int) (*model.APIResponseOrderMonthly, bool) {
	key := fmt.Sprintf(monthlyOrderCacheKey, year)
	result, found := cache.GetFromCache[model.APIResponseOrderMonthly](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsCache) SetMonthlyOrderCache(ctx context.Context, year int, data *model.APIResponseOrderMonthly) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthlyOrderCacheKey, year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderStatsCache) GetYearlyOrderCache(ctx context.Context, year int) (*model.APIResponseOrderYearly, bool) {
	key := fmt.Sprintf(yearlyOrderCacheKey, year)
	result, found := cache.GetFromCache[model.APIResponseOrderYearly](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsCache) SetYearlyOrderCache(ctx context.Context, year int, data *model.APIResponseOrderYearly) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearlyOrderCacheKey, year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
