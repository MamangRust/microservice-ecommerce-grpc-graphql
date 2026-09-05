package review_detailgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type ReviewDetailGraphqlMapper interface {
	ToGraphqlResponseDelete(res *pb.ApiResponseReviewDelete) *model.APIResponseReviewDetailDelete
	ToGraphqlResponseAll(res *pb.ApiResponseReviewAll) *model.APIResponseReviewDetailAll
	ToGraphqlResponseReviewDetail(res *pb.ApiResponseReviewDetail) *model.APIResponseReviewDetail
	ToGraphqlResponsesReviewDetail(res *pb.ApiResponsesReviewDetails) *model.APIResponsesReviewDetails
	ToGraphqlResponseReviewDetailDeleteAt(res *pb.ApiResponseReviewDetailDeleteAt) *model.APIResponseReviewDetailDeleteAt
	ToGraphqlResponsePaginationReviewDetail(res *pb.ApiResponsePaginationReviewDetails) *model.APIResponsePaginationReviewDetails
	ToGraphqlResponsePaginationReviewDetailDeleteAt(res *pb.ApiResponsePaginationReviewDetailsDeleteAt) *model.APIResponsePaginationReviewDetailsDeleteAt
}
