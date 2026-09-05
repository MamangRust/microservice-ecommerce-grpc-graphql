package cart_cache

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

const (
	cartAllCacheKey = "cart:all:page:%d:pageSize:%d:search:%s"
	ttlDefault      = 5 * time.Minute
)

type cartQueryCache struct {
	store *cache.CacheStore
}

func NewCartQueryCache(store *cache.CacheStore) *cartQueryCache {
	return &cartQueryCache{store: store}
}

func (c *cartQueryCache) GetCachedCarts(
	ctx context.Context,
	request *model.FindAllCartInput,
) (*model.APIResponsePaginationCart, bool) {

	key := fmt.Sprintf(
		cartAllCacheKey,
		request.Page,
		request.PageSize,
		safeString(request.Search),
	)

	result, found := cache.GetFromCache[model.APIResponsePaginationCart](
		ctx,
		c.store,
		key,
	)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (c *cartQueryCache) SetCachedCarts(
	ctx context.Context,
	request *model.FindAllCartInput,
	resp *model.APIResponsePaginationCart,
) {
	if resp == nil {
		return
	}

	key := fmt.Sprintf(
		cartAllCacheKey,
		request.Page,
		request.PageSize,
		safeString(request.Search),
	)

	cache.SetToCache(ctx, c.store, key, resp, ttlDefault)
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
