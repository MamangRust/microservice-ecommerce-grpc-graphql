package transaction_cache

import (
	"context"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	transactonMonthAmountSuccessByMerchantKey = "transaction:month:amount:success:merchant:%d:month:%d:year:%d"
	transactonMonthAmountFailedByMerchantKey  = "transaction:month:amount:failed:merchant:%d:month:%d:year:%d"
	transactonYearAmountSuccessByMerchantKey  = "transaction:year:amount:success:merchant:%d:year:%d"
	transactonYearAmountFailedByMerchantKey   = "transaction:year:amount:failed:merchant:%d:year:%d"
	transactonMonthMethodSuccessByMerchantKey = "transaction:month:method:success:merchant:%d:month:%d:year:%d"
	transactonMonthMethodFailedByMerchantKey  = "transaction:month:method:failed:merchant:%d:month:%d:year:%d"
	transactonYearMethodSuccessByMerchantKey  = "transaction:year:method:success:merchant:%d:year:%d"
	transactonYearMethodFailedByMerchantKey   = "transaction:year:method:failed:merchant:%d:year:%d"
)

type transactionStatsByMerchantCache struct {
	store *cache.CacheStore
}

func NewTransactionStatsByMerchantCache(store *cache.CacheStore) *transactionStatsByMerchantCache {
	return &transactionStatsByMerchantCache{store: store}
}

func (t *transactionStatsByMerchantCache) GetCachedMonthAmountSuccessByMerchant(ctx context.Context, req *model.FindMonthlyTransactionStatusByMerchant) (*model.APIResponseTransactionMonthAmountSuccess, bool) {
	key := fmt.Sprintf(transactonMonthAmountSuccessByMerchantKey, req.MerchantID, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseTransactionMonthAmountSuccess](ctx, t.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsByMerchantCache) SetCachedMonthAmountSuccessByMerchant(ctx context.Context, req *model.FindMonthlyTransactionStatusByMerchant, data *model.APIResponseTransactionMonthAmountSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactonMonthAmountSuccessByMerchantKey, req.MerchantID, req.Month, req.Year)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsByMerchantCache) GetCachedYearAmountSuccessByMerchant(ctx context.Context, req *model.FindYearlyTransactionStatusByMerchant) (*model.APIResponseTransactionYearAmountSuccess, bool) {
	key := fmt.Sprintf(transactonYearAmountSuccessByMerchantKey, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseTransactionYearAmountSuccess](ctx, t.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsByMerchantCache) SetCachedYearAmountSuccessByMerchant(ctx context.Context, req *model.FindYearlyTransactionStatusByMerchant, data *model.APIResponseTransactionYearAmountSuccess) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactonYearAmountSuccessByMerchantKey, req.MerchantID, req.Year)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsByMerchantCache) GetCachedMonthAmountFailedByMerchant(ctx context.Context, req *model.FindMonthlyTransactionStatusByMerchant) (*model.APIResponseTransactionMonthAmountFailed, bool) {
	key := fmt.Sprintf(transactonMonthAmountFailedByMerchantKey, req.MerchantID, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseTransactionMonthAmountFailed](ctx, t.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsByMerchantCache) SetCachedMonthAmountFailedByMerchant(ctx context.Context, req *model.FindMonthlyTransactionStatusByMerchant, data *model.APIResponseTransactionMonthAmountFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactonMonthAmountFailedByMerchantKey, req.MerchantID, req.Month, req.Year)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsByMerchantCache) GetCachedYearAmountFailedByMerchant(ctx context.Context, req *model.FindYearlyTransactionStatusByMerchant) (*model.APIResponseTransactionYearAmountFailed, bool) {
	key := fmt.Sprintf(transactonYearAmountFailedByMerchantKey, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseTransactionYearAmountFailed](ctx, t.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsByMerchantCache) SetCachedYearAmountFailedByMerchant(ctx context.Context, req *model.FindYearlyTransactionStatusByMerchant, data *model.APIResponseTransactionYearAmountFailed) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactonYearAmountFailedByMerchantKey, req.MerchantID, req.Year)

	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsByMerchantCache) GetCachedMonthMethodSuccessByMerchant(ctx context.Context, req *model.MonthTransactionMethodByMerchant) (*model.APIResponseTransactionMonthPaymentMethod, bool) {
	key := fmt.Sprintf(transactonMonthMethodSuccessByMerchantKey, req.MerchantID, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseTransactionMonthPaymentMethod](ctx, t.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsByMerchantCache) SetCachedMonthMethodSuccessByMerchant(ctx context.Context, req *model.MonthTransactionMethodByMerchant, data *model.APIResponseTransactionMonthPaymentMethod) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactonMonthMethodSuccessByMerchantKey, req.MerchantID, req.Month, req.Year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsByMerchantCache) GetCachedYearMethodSuccessByMerchant(ctx context.Context, req *model.YearTransactionMethodByMerchant) (*model.APIResponseTransactionYearPaymentMethod, bool) {
	key := fmt.Sprintf(transactonYearMethodSuccessByMerchantKey, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseTransactionYearPaymentMethod](ctx, t.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsByMerchantCache) SetCachedYearMethodSuccessByMerchant(ctx context.Context, req *model.YearTransactionMethodByMerchant, data *model.APIResponseTransactionYearPaymentMethod) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactonYearMethodSuccessByMerchantKey, req.MerchantID, req.Year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsByMerchantCache) GetCachedMonthMethodFailedByMerchant(ctx context.Context, req *model.MonthTransactionMethodByMerchant) (*model.APIResponseTransactionMonthPaymentMethod, bool) {
	key := fmt.Sprintf(transactonMonthMethodFailedByMerchantKey, req.MerchantID, req.Month, req.Year)
	result, found := cache.GetFromCache[model.APIResponseTransactionMonthPaymentMethod](ctx, t.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsByMerchantCache) SetCachedMonthMethodFailedByMerchant(ctx context.Context, req *model.MonthTransactionMethodByMerchant, data *model.APIResponseTransactionMonthPaymentMethod) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactonMonthMethodFailedByMerchantKey, req.MerchantID, req.Month, req.Year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}

func (t *transactionStatsByMerchantCache) GetCachedYearMethodFailedByMerchant(ctx context.Context, req *model.YearTransactionMethodByMerchant) (*model.APIResponseTransactionYearPaymentMethod, bool) {
	key := fmt.Sprintf(transactonYearMethodFailedByMerchantKey, req.MerchantID, req.Year)
	result, found := cache.GetFromCache[model.APIResponseTransactionYearPaymentMethod](ctx, t.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (t *transactionStatsByMerchantCache) SetCachedYearMethodFailedByMerchant(ctx context.Context, req *model.YearTransactionMethodByMerchant, data *model.APIResponseTransactionYearPaymentMethod) {
	if data == nil {
		return
	}
	key := fmt.Sprintf(transactonYearMethodFailedByMerchantKey, req.MerchantID, req.Year)
	cache.SetToCache(ctx, t.store, key, data, ttlDefault)
}
