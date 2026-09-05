package slider_cache

import "github.com/MamangRust/microservice-ecommerce-shared/cache"

type sliderMencache struct {
	SliderQueryCache
	SliderCommandCache
}

type SliderMencache interface {
	SliderQueryCache
	SliderCommandCache
}

func NewSliderMencache(cacheStore *cache.CacheStore) SliderMencache {
	return sliderMencache{
		SliderQueryCache:   NewSliderQueryCache(cacheStore),
		SliderCommandCache: NewSliderCommandCache(cacheStore),
	}
}
