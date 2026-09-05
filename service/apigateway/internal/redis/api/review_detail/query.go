package reviewdetail_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	reviewDetailAllCacheKey         = "review_detail:all:page:%d:pageSize:%d:search:%s"
	reviewDetailByIdCacheKey        = "review_detail:id:%d"
	reviewDetailActiveCacheKey      = "review_detail:active:page:%d:pageSize:%d:search:%s"
	reviewDetailTrashedCacheKey     = "review_detail:trashed:page:%d:pageSize:%d:search:%s"
	reviewDetailByIdTrashedCacheKey = "review_detail:id_trashed:%d"

	ttlDefault = 5 * time.Minute
)

type reviewDetailQueryCache struct {
	store *cache.CacheStore
}

func NewReviewDetailQueryCache(store *cache.CacheStore) *reviewDetailQueryCache {
	return &reviewDetailQueryCache{store: store}
}

func (r *reviewDetailQueryCache) GetReviewDetailAllCache(ctx context.Context, req *model.FindAllReviewDetailInput) (*model.APIResponsePaginationReviewDetails, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewDetailAllCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationReviewDetails](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewDetailQueryCache) SetReviewDetailAllCache(ctx context.Context, req *model.FindAllReviewDetailInput, data *model.APIResponsePaginationReviewDetails) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewDetailAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewDetailQueryCache) GetReviewDetailActiveCache(ctx context.Context, req *model.FindAllReviewDetailInput) (*model.APIResponsePaginationReviewDetailsDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewDetailActiveCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationReviewDetailsDeleteAt](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewDetailQueryCache) SetReviewDetailActiveCache(ctx context.Context, req *model.FindAllReviewDetailInput, data *model.APIResponsePaginationReviewDetailsDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewDetailActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewDetailQueryCache) GetReviewDetailTrashedCache(ctx context.Context, req *model.FindAllReviewDetailInput) (*model.APIResponsePaginationReviewDetailsDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewDetailTrashedCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationReviewDetailsDeleteAt](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewDetailQueryCache) SetReviewDetailTrashedCache(ctx context.Context, req *model.FindAllReviewDetailInput, data *model.APIResponsePaginationReviewDetailsDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewDetailTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewDetailQueryCache) GetCachedReviewDetailCache(ctx context.Context, reviewID int) (*model.APIResponseReviewDetail, bool) {
	key := fmt.Sprintf(reviewDetailByIdCacheKey, reviewID)
	result, found := cache.GetFromCache[model.APIResponseReviewDetail](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewDetailQueryCache) SetCachedReviewDetailCache(ctx context.Context, data *model.APIResponseReviewDetail) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(reviewDetailByIdCacheKey, data.Data.ID)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewDetailQueryCache) GetCachedReviewDetailTrashedCache(ctx context.Context, reviewID int) (*model.APIResponseReviewDetailDeleteAt, bool) {
	key := fmt.Sprintf(reviewDetailByIdTrashedCacheKey, reviewID)
	result, found := cache.GetFromCache[model.APIResponseReviewDetailDeleteAt](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewDetailQueryCache) SetCachedReviewDetailTrashedCache(ctx context.Context, data *model.APIResponseReviewDetailDeleteAt) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(reviewDetailByIdTrashedCacheKey, data.Data.ID)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}
