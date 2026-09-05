package review_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	reviewAllCacheKey      = "review:all:page:%d:pageSize:%d:search:%s"
	reviewProductCacheKey  = "review:product:%d:page:%d:pageSize:%d:search:%s"
	reviewMerchantCacheKey = "review:merchant:%d:page:%d:pageSize:%d:search:%s"

	reviewByIdCacheKey    = "review:id:%d"
	reviewActiveCacheKey  = "review:active:page:%d:pageSize:%d:search:%s"
	reviewTrashedCacheKey = "review:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

type reviewQueryCache struct {
	store *cache.CacheStore
}

func NewReviewQueryCache(store *cache.CacheStore) *reviewQueryCache {
	return &reviewQueryCache{store: store}
}

func (r *reviewQueryCache) GetReviewAllCache(ctx context.Context, req *model.FindAllReviewRequest) (*model.APIResponsePaginationReview, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewAllCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationReview](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewQueryCache) SetReviewAllCache(ctx context.Context, req *model.FindAllReviewRequest, data *model.APIResponsePaginationReview) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewAllCacheKey, req.Page, req.PageSize, search)
	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewQueryCache) GetReviewByProductCache(ctx context.Context, req *model.FindAllReviewProductRequest) (*model.APIResponsePaginationReviewRelationDetail, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewProductCacheKey, req.ProductID, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationReviewRelationDetail](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewQueryCache) SetReviewByProductCache(ctx context.Context, req *model.FindAllReviewProductRequest, data *model.APIResponsePaginationReviewRelationDetail) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewProductCacheKey, req.ProductID, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewQueryCache) GetReviewByMerchantCache(ctx context.Context, req *model.FindAllReviewMerchantRequest) (*model.APIResponsePaginationReviewRelationDetail, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationReviewRelationDetail](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewQueryCache) SetReviewByMerchantCache(ctx context.Context, req *model.FindAllReviewMerchantRequest, data *model.APIResponsePaginationReviewRelationDetail) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewQueryCache) GetReviewActiveCache(ctx context.Context, req *model.FindAllReviewRequest) (*model.APIResponsePaginationReviewDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewActiveCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationReviewDeleteAt](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewQueryCache) SetReviewActiveCache(ctx context.Context, req *model.FindAllReviewRequest, data *model.APIResponsePaginationReviewDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewQueryCache) GetReviewTrashedCache(ctx context.Context, req *model.FindAllReviewRequest) (*model.APIResponsePaginationReviewDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewTrashedCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationReviewDeleteAt](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewQueryCache) SetReviewTrashedCache(ctx context.Context, req *model.FindAllReviewRequest, data *model.APIResponsePaginationReviewDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(reviewTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}

func (r *reviewQueryCache) GetReviewByIdCache(ctx context.Context, id int) (*model.APIResponseReview, bool) {
	key := fmt.Sprintf(reviewByIdCacheKey, id)
	result, found := cache.GetFromCache[model.APIResponseReview](ctx, r.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (r *reviewQueryCache) SetReviewByIdCache(ctx context.Context, data *model.APIResponseReview) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(reviewByIdCacheKey, data.Data.ID)
	cache.SetToCache(ctx, r.store, key, data, ttlDefault)
}
