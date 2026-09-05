package bannergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type BannerGraphqlMapper interface {
	ToGraphqlResponseBanner(res *pb.ApiResponseBanner) *model.APIResponseBanner
	ToGraphqlResponseBannerDeleteAt(res *pb.ApiResponseBannerDeleteAt) *model.APIResponseBannerDeleteAt
	ToGraphqlResponsesBanner(res *pb.ApiResponsesBanner) *model.APIResponsesBanner
	ToGraphqlResponseDelete(res *pb.ApiResponseBannerDelete) *model.APIResponseBannerDelete
	ToGraphqlResponseAll(res *pb.ApiResponseBannerAll) *model.APIResponseBannerAll
	ToGraphqlResponsePaginationBanner(res *pb.ApiResponsePaginationBanner) *model.APIResponsePaginationBanner
	ToGraphqlResponsePaginationBannerDeleteAt(res *pb.ApiResponsePaginationBannerDeleteAt) *model.APIResponsePaginationBannerDeleteAt
}
