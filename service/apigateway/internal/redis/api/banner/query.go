package banner_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type bannerQueryCache struct {
	store *cache.CacheStore
}

func NewBannerQueryCache(store *cache.CacheStore) *bannerQueryCache {
	return &bannerQueryCache{store: store}
}

func (b *bannerQueryCache) GetCachedBanners(
	ctx context.Context,
	req *model.FindAllBannerInput,
) (*model.APIResponsePaginationBanner, bool) {

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(
		bannerAllCacheKey,
		req.Page,
		req.PageSize,
		search,
	)

	result, found := cache.GetFromCache[model.APIResponsePaginationBanner](
		ctx,
		b.store,
		key,
	)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (b *bannerQueryCache) SetCachedBanners(
	ctx context.Context,
	req *model.FindAllBannerInput,
	data *model.APIResponsePaginationBanner,
) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(
		bannerAllCacheKey,
		req.Page,
		req.PageSize,
		search,
	)

	cache.SetToCache(ctx, b.store, key, data, ttlDefault)
}

func (b *bannerQueryCache) GetCachedActiveBanners(
	ctx context.Context,
	req *model.FindAllBannerInput,
) (*model.APIResponsePaginationBannerDeleteAt, bool) {

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(
		bannerActiveCacheKey,
		req.Page,
		req.PageSize,
		search,
	)

	result, found := cache.GetFromCache[model.APIResponsePaginationBannerDeleteAt](
		ctx,
		b.store,
		key,
	)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (b *bannerQueryCache) SetCachedActiveBanners(
	ctx context.Context,
	req *model.FindAllBannerInput,
	data *model.APIResponsePaginationBannerDeleteAt,
) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(
		bannerActiveCacheKey,
		req.Page,
		req.PageSize,
		search,
	)

	cache.SetToCache(ctx, b.store, key, data, ttlDefault)
}

func (b *bannerQueryCache) GetCachedTrashedBanners(
	ctx context.Context,
	req *model.FindAllBannerInput,
) (*model.APIResponsePaginationBannerDeleteAt, bool) {

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(
		bannerTrashedCacheKey,
		req.Page,
		req.PageSize,
		search,
	)

	result, found := cache.GetFromCache[model.APIResponsePaginationBannerDeleteAt](
		ctx,
		b.store,
		key,
	)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (b *bannerQueryCache) SetCachedTrashedBanners(
	ctx context.Context,
	req *model.FindAllBannerInput,
	data *model.APIResponsePaginationBannerDeleteAt,
) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(
		bannerTrashedCacheKey,
		req.Page,
		req.PageSize,
		search,
	)

	cache.SetToCache(ctx, b.store, key, data, ttlDefault)
}

func (b *bannerQueryCache) GetCachedBanner(
	ctx context.Context,
	id int,
) (*model.APIResponseBanner, bool) {

	key := fmt.Sprintf(bannerByIdCacheKey, id)

	result, found := cache.GetFromCache[model.APIResponseBanner](
		ctx,
		b.store,
		key,
	)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (b *bannerQueryCache) SetCachedBanner(
	ctx context.Context,
	data *model.APIResponseBanner,
) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(bannerByIdCacheKey, data.Data.BannerID)

	cache.SetToCache(ctx, b.store, key, data, ttlDefault)
}
