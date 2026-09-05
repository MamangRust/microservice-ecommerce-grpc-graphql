package order_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	monthlyTotalRevenueCacheKeyByMerchant = "order:monthly:totalRevenue:merchant:%d:month:%d:year:%d"
	yearlyTotalRevenueCacheKeyByMerchant  = "order:yearly:totalRevenue:merchant:%d:year:%d"
	monthlyOrderCacheKeyByMerchant        = "order:monthly:order:merchant:%d:year:%d"
	yearlyOrderCacheKeyByMerchant         = "order:yearly:order:merchant:%d:year:%d"
)

type orderStatsByMerchantCache struct {
	store *cache.CacheStore
}

func NewOrderStatsByMerchantCache(store *cache.CacheStore) *orderStatsByMerchantCache {
	return &orderStatsByMerchantCache{store: store}
}

func (s *orderStatsByMerchantCache) GetMonthlyTotalRevenueByMerchantCache(ctx context.Context, req *model.FindYearMonthTotalRevenueByMerchant) (*model.APIResponseOrderMonthlyTotalRevenue, bool) {
	key := fmt.Sprintf(monthlyTotalRevenueCacheKeyByMerchant, req.MerchantID, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseOrderMonthlyTotalRevenue](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsByMerchantCache) SetMonthlyTotalRevenueByMerchantCache(ctx context.Context, req *model.FindYearMonthTotalRevenueByMerchant, data *model.APIResponseOrderMonthlyTotalRevenue) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthlyTotalRevenueCacheKeyByMerchant, req.MerchantID, req.Month, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderStatsByMerchantCache) GetYearlyTotalRevenueByMerchantCache(ctx context.Context, req *model.FindYearTotalRevenueByMerchant) (*model.APIResponseOrderYearlyTotalRevenue, bool) {
	key := fmt.Sprintf(yearlyTotalRevenueCacheKeyByMerchant, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseOrderYearlyTotalRevenue](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsByMerchantCache) SetYearlyTotalRevenueByMerchantCache(ctx context.Context, req *model.FindYearTotalRevenueByMerchant, data *model.APIResponseOrderYearlyTotalRevenue) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearlyTotalRevenueCacheKeyByMerchant, req.MerchantID, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderStatsByMerchantCache) GetMonthlyOrderByMerchantCache(ctx context.Context, req *model.FindYearOrderByMerchantInput) (*model.APIResponseOrderMonthly, bool) {
	key := fmt.Sprintf(monthlyOrderCacheKeyByMerchant, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseOrderMonthly](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsByMerchantCache) SetMonthlyOrderByMerchantCache(ctx context.Context, req *model.FindYearOrderByMerchantInput, data *model.APIResponseOrderMonthly) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(monthlyOrderCacheKeyByMerchant, req.MerchantID, req.Year)

	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderStatsByMerchantCache) GetYearlyOrderByMerchantCache(ctx context.Context, req *model.FindYearOrderByMerchantInput) (*model.APIResponseOrderYearly, bool) {
	key := fmt.Sprintf(yearlyOrderCacheKeyByMerchant, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseOrderYearly](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsByMerchantCache) SetYearlyOrderByMerchantCache(ctx context.Context, req *model.FindYearOrderByMerchantInput, data *model.APIResponseOrderYearly) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(yearlyOrderCacheKeyByMerchant, req.MerchantID, req.Year)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}
