package shippingaddress_cache

import "github.com/MamangRust/microservice-ecommerce-shared/cache"

type ShippingAddressMencache interface {
	ShippingAddressQueryCache
	ShippingAddressCommandCache
}

type shippingAddressMencache struct {
	ShippingAddressQueryCache
	ShippingAddressCommandCache
}

func NewShippingAddressMencache(cacheStore *cache.CacheStore) ShippingAddressMencache {
	return &shippingAddressMencache{
		ShippingAddressQueryCache:   NewShippingAddressQueryCache(cacheStore),
		ShippingAddressCommandCache: NewShippingAddressCommandCache(cacheStore),
	}
}
