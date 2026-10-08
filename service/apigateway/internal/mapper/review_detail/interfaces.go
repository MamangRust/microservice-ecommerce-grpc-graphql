package review_detailgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/review"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/review_detail"
)

type ReviewDetailGraphqlMapper interface {
	ToGraphqlResponseDelete(res *pb_review.ApiResponseReviewDelete) *model.APIResponseReviewDetailDelete
	ToGraphqlResponseAll(res *pb_review.ApiResponseReviewAll) *model.APIResponseReviewDetailAll
	ToGraphqlResponseReviewDetail(res *pb_review_detail.ApiResponseReviewDetail) *model.APIResponseReviewDetail
	ToGraphqlResponsesReviewDetail(res *pb_review_detail.ApiResponsesReviewDetails) *model.APIResponsesReviewDetails
	ToGraphqlResponseReviewDetailDeleteAt(res *pb_review_detail.ApiResponseReviewDetailDeleteAt) *model.APIResponseReviewDetailDeleteAt
	ToGraphqlResponsePaginationReviewDetail(res *pb_review_detail.ApiResponsePaginationReviewDetails) *model.APIResponsePaginationReviewDetails
	ToGraphqlResponsePaginationReviewDetailDeleteAt(res *pb_review_detail.ApiResponsePaginationReviewDetailsDeleteAt) *model.APIResponsePaginationReviewDetailsDeleteAt
}
