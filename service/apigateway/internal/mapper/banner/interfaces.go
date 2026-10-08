package bannergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/banner"
)

type BannerGraphqlMapper interface {
	ToGraphqlResponseBanner(res *pb_banner.ApiResponseBanner) *model.APIResponseBanner
	ToGraphqlResponseBannerDeleteAt(res *pb_banner.ApiResponseBannerDeleteAt) *model.APIResponseBannerDeleteAt
	ToGraphqlResponsesBanner(res *pb_banner.ApiResponsesBanner) *model.APIResponsesBanner
	ToGraphqlResponseDelete(res *pb_banner.ApiResponseBannerDelete) *model.APIResponseBannerDelete
	ToGraphqlResponseAll(res *pb_banner.ApiResponseBannerAll) *model.APIResponseBannerAll
	ToGraphqlResponsePaginationBanner(res *pb_banner.ApiResponsePaginationBanner) *model.APIResponsePaginationBanner
	ToGraphqlResponsePaginationBannerDeleteAt(res *pb_banner.ApiResponsePaginationBannerDeleteAt) *model.APIResponsePaginationBannerDeleteAt
}
