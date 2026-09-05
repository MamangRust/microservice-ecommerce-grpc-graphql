package transaction_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	transactionAllCacheKey        = "transaction:all:page:%d:pageSize:%d:search:%s"
	transactionByIdCacheKey       = "transaction:id:%d"
	transactionByMerchantCacheKey = "transaction:merchant:%d:page:%d:pageSize:%d:search:%s"
	transactionActiveCacheKey     = "transaction:active:page:%d:pageSize:%d:search:%s"
	transactionTrashedCacheKey    = "transaction:trashed:page:%d:pageSize:%d:search:%s"
	transactionByOrderCacheKey    = "transaction:order:%d"
	ttlDefault                    = 5 * time.Minute
)

type transactionQueryCache struct {
	store *cache.CacheStore
}

func NewTransactionQueryCache(store *cache.CacheStore) *transactionQueryCache {
	return &transactionQueryCache{store: store}
}

func (t *transactionQueryCache) GetCachedTransactionsCache(ctx context.Context, req *model.FindAllTransactionRequest) (*model.APIResponsePaginationTransaction, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(transactionAllCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationTransaction](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (t *transactionQueryCache) SetCachedTransactionsCache(ctx context.Context, req *model.FindAllTransactionRequest, data *model.APIResponsePaginationTransaction) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(transactionAllCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionByMerchant(ctx context.Context, req *model.FindAllTransactionMerchantRequest) (*model.APIResponsePaginationTransaction, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(transactionByMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationTransaction](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (t *transactionQueryCache) SetCachedTransactionByMerchant(ctx context.Context, req *model.FindAllTransactionMerchantRequest, data *model.APIResponsePaginationTransaction) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(transactionByMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionActiveCache(ctx context.Context, req *model.FindAllTransactionRequest) (*model.APIResponsePaginationTransaction, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(transactionActiveCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationTransaction](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (t *transactionQueryCache) SetCachedTransactionActiveCache(ctx context.Context, req *model.FindAllTransactionRequest, data *model.APIResponsePaginationTransaction) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(transactionActiveCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionTrashedCache(ctx context.Context, req *model.FindAllTransactionRequest) (*model.APIResponsePaginationTransactionDeleteAt, bool) {
	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(transactionTrashedCacheKey, req.Page, req.PageSize, search)
	result, found := cache.GetFromCache[model.APIResponsePaginationTransactionDeleteAt](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (t *transactionQueryCache) SetCachedTransactionTrashedCache(ctx context.Context, req *model.FindAllTransactionRequest, data *model.APIResponsePaginationTransactionDeleteAt) {
	if data == nil {
		return
	}

	var search string
	if req.Search != nil {
		search = *req.Search
	}

	key := fmt.Sprintf(transactionTrashedCacheKey, req.Page, req.PageSize, search)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionCache(ctx context.Context, id int) (*model.APIResponseTransaction, bool) {
	key := fmt.Sprintf(transactionByIdCacheKey, id)
	result, found := cache.GetFromCache[model.APIResponseTransaction](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (t *transactionQueryCache) SetCachedTransactionCache(ctx context.Context, data *model.APIResponseTransaction) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactionByIdCacheKey, data.Data.ID)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionQueryCache) GetCachedTransactionByOrderId(ctx context.Context, orderID int) (*model.APIResponseTransaction, bool) {
	key := fmt.Sprintf(transactionByOrderCacheKey, orderID)
	result, found := cache.GetFromCache[model.APIResponseTransaction](ctx, t.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (t *transactionQueryCache) SetCachedTransactionByOrderId(ctx context.Context, orderID int, data *model.APIResponseTransaction) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactionByOrderCacheKey, orderID)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}
