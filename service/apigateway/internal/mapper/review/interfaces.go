package reviewgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/review"
)

type ReviewGraphqlMapper interface {
	ToGraphqlResponseReview(res *pb_review.ApiResponseReview) *model.APIResponseReview
	ToGraphqlResponseReviewDeleteAt(res *pb_review.ApiResponseReviewDeleteAt) *model.APIResponseReviewDeleteAt
	ToGraphqlResponsesReview(res *pb_review.ApiResponsesReview) *model.APIResponsesReview
	ToGraphqlResponseReviewDelete(res *pb_review.ApiResponseReviewDelete) *model.APIResponseReviewDelete
	ToGraphqlResponseReviewAll(res *pb_review.ApiResponseReviewAll) *model.APIResponseReviewAll
	ToGraphqlResponsePaginationReviewDeleteAt(res *pb_review.ApiResponsePaginationReviewDeleteAt) *model.APIResponsePaginationReviewDeleteAt
	ToGraphqlResponsePaginationReview(res *pb_review.ApiResponsePaginationReview) *model.APIResponsePaginationReview
	ToGraphqlResponsePaginationReviewRelationDetail(res *pb_review.ApiResponsePaginationReviewDetail) *model.APIResponsePaginationReviewRelationDetail
}
