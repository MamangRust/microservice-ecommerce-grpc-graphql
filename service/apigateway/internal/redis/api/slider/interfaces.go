package slider_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type SliderQueryCache interface {
	GetSliderAllCache(ctx context.Context, req *model.FindAllSliderRequest) (*model.APIResponsePaginationSlider, bool)
	SetSliderAllCache(ctx context.Context, req *model.FindAllSliderRequest, data *model.APIResponsePaginationSlider)

	GetSliderActiveCache(ctx context.Context, req *model.FindAllSliderRequest) (*model.APIResponsePaginationSliderDeleteAt, bool)
	SetSliderActiveCache(ctx context.Context, req *model.FindAllSliderRequest, data *model.APIResponsePaginationSliderDeleteAt)

	GetSliderTrashedCache(ctx context.Context, req *model.FindAllSliderRequest) (*model.APIResponsePaginationSliderDeleteAt, bool)
	SetSliderTrashedCache(ctx context.Context, req *model.FindAllSliderRequest, data *model.APIResponsePaginationSliderDeleteAt)
}

type SliderCommandCache interface {
	DeleteSliderCache(ctx context.Context, slider_id int)
}
