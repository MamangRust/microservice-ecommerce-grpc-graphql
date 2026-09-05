package banner_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type BannerQueryCache interface {
	GetCachedBanners(
		ctx context.Context,
		req *model.FindAllBannerInput,
	) (*model.APIResponsePaginationBanner, bool)
	SetCachedBanners(
		ctx context.Context,
		req *model.FindAllBannerInput,
		data *model.APIResponsePaginationBanner,
	)
	GetCachedActiveBanners(
		ctx context.Context,
		req *model.FindAllBannerInput,
	) (*model.APIResponsePaginationBannerDeleteAt, bool)
	SetCachedActiveBanners(
		ctx context.Context,
		req *model.FindAllBannerInput,
		data *model.APIResponsePaginationBannerDeleteAt,
	)
	GetCachedTrashedBanners(
		ctx context.Context,
		req *model.FindAllBannerInput,
	) (*model.APIResponsePaginationBannerDeleteAt, bool)
	SetCachedTrashedBanners(
		ctx context.Context,
		req *model.FindAllBannerInput,
		data *model.APIResponsePaginationBannerDeleteAt,
	)
	GetCachedBanner(
		ctx context.Context,
		id int,
	) (*model.APIResponseBanner, bool)
	SetCachedBanner(
		ctx context.Context,
		data *model.APIResponseBanner,
	)
}

type BannerCommandCache interface {
	DeleteBannerCache(ctx context.Context, id int)
}
