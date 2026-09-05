package product_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	productAllCacheKey      = "product:all:page:%d:pageSize:%d:search:%s"
	productCategoryCacheKey = "product:category:%s:page:%d:pageSize:%d:search:%s"
	productMerchantCacheKey = "product:merchant:%d:page:%d:pageSize:%d:search:%s"

	productActiveCacheKey  = "product:active:page:%d:pageSize:%d:search:%s"
	productTrashedCacheKey = "product:trashed:page:%d:pageSize:%d:search:%s"
	productByIdCacheKey    = "product:id:%d"

	ttlDefault = 5 * time.Minute
)

type productQueryCache struct {
	store *cache.CacheStore
}

func NewProductQueryCache(store *cache.CacheStore) *productQueryCache {
	return &productQueryCache{store: store}
}

func (p *productQueryCache) GetCachedProducts(ctx context.Context, req *model.FindAllProductInput) (*model.APIResponsePaginationProduct, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productAllCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationProduct](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (p *productQueryCache) SetCachedProducts(ctx context.Context, req *model.FindAllProductInput, data *model.APIResponsePaginationProduct) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, p.store, key, data, ttlDefault)
}

func (p *productQueryCache) GetCachedProductsByMerchant(ctx context.Context, req *model.FindAllProductMerchantInput) (*model.APIResponsePaginationProduct, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationProduct](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProductsByMerchant(ctx context.Context, req *model.FindAllProductMerchantInput, data *model.APIResponsePaginationProduct) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, p.store, key, data, ttlDefault)
}

func (p *productQueryCache) GetCachedProductsByCategory(ctx context.Context, req *model.FindAllProductCategoryInput) (*model.APIResponsePaginationProduct, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productCategoryCacheKey, safeString(req.CategoryName), req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationProduct](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProductsByCategory(ctx context.Context, req *model.FindAllProductCategoryInput, data *model.APIResponsePaginationProduct) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	var category string

	if req.CategoryName != nil {
		category = *req.CategoryName
	}

	key := fmt.Sprintf(productCategoryCacheKey, category, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, p.store, key, data, ttlDefault)
}

func (p *productQueryCache) GetCachedProductActive(ctx context.Context, req *model.FindAllProductInput) (*model.APIResponsePaginationProductDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productActiveCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationProductDeleteAt](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProductActive(ctx context.Context, req *model.FindAllProductInput, data *model.APIResponsePaginationProductDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, p.store, key, data, ttlDefault)
}

func (p *productQueryCache) GetCachedProductTrashed(ctx context.Context, req *model.FindAllProductInput) (*model.APIResponsePaginationProductDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productTrashedCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationProductDeleteAt](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProductTrashed(ctx context.Context, req *model.FindAllProductInput, data *model.APIResponsePaginationProductDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(productTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, p.store, key, data, ttlDefault)
}

func (p *productQueryCache) GetCachedProduct(ctx context.Context, productID int) (*model.APIResponseProduct, bool) {
	key := fmt.Sprintf(productByIdCacheKey, productID)
	result, found := cache.GetFromCache[model.APIResponseProduct](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProduct(ctx context.Context, data *model.APIResponseProduct) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(productByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, p.store, key, data, ttlDefault)
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
