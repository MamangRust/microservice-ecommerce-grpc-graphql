package reviewgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type reviewResponseMapper struct{}

func NewReviewResponseMapper() *reviewResponseMapper {
	return &reviewResponseMapper{}
}

func (r *reviewResponseMapper) ToGraphqlResponseReview(res *pb.ApiResponseReview) *model.APIResponseReview {
	return &model.APIResponseReview{
		Status:  res.Status,
		Message: res.Message,
		Data:    r.mapResponseReview(res.Data),
	}
}

func (r *reviewResponseMapper) ToGraphqlResponseReviewDeleteAt(res *pb.ApiResponseReviewDeleteAt) *model.APIResponseReviewDeleteAt {
	return &model.APIResponseReviewDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    r.mapResponseReviewDeleteAt(res.Data),
	}
}

func (r *reviewResponseMapper) ToGraphqlResponsesReview(res *pb.ApiResponsesReview) *model.APIResponsesReview {
	return &model.APIResponsesReview{
		Status:  res.Status,
		Message: res.Message,
		Data:    r.mapResponsesReview(res.Data),
	}
}

func (r *reviewResponseMapper) ToGraphqlResponseReviewDelete(res *pb.ApiResponseReviewDelete) *model.APIResponseReviewDelete {
	return &model.APIResponseReviewDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (r *reviewResponseMapper) ToGraphqlResponsePaginationReviewDeleteAt(
	res *pb.ApiResponsePaginationReviewDeleteAt,
) *model.APIResponsePaginationReviewDeleteAt {
	return &model.APIResponsePaginationReviewDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       r.mapResponsesReviewDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (r *reviewResponseMapper) ToGraphqlResponseReviewAll(res *pb.ApiResponseReviewAll) *model.APIResponseReviewAll {
	return &model.APIResponseReviewAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (r *reviewResponseMapper) ToGraphqlResponsePaginationReview(
	res *pb.ApiResponsePaginationReview,
) *model.APIResponsePaginationReview {
	return &model.APIResponsePaginationReview{
		Status:     res.Status,
		Message:    res.Message,
		Data:       r.mapResponsesReview(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (r *reviewResponseMapper) ToGraphqlResponsePaginationReviewRelationDetail(
	res *pb.ApiResponsePaginationReviewDetail,
) *model.APIResponsePaginationReviewRelationDetail {
	return &model.APIResponsePaginationReviewRelationDetail{
		Status:     res.Status,
		Message:    res.Message,
		Data:       r.mapResponsesReviewsRelationDetail(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (r *reviewResponseMapper) mapResponseReview(review *pb.ReviewResponse) *model.ReviewResponse {
	return &model.ReviewResponse{
		ID:        int32(review.Id),
		UserID:    int32(review.UserId),
		ProductID: int32(review.ProductId),
		Name:      review.Name,
		Comment:   review.Comment,
		Rating:    int32(review.Rating),
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}
}

func (r *reviewResponseMapper) mapResponsesReview(reviews []*pb.ReviewResponse) []*model.ReviewResponse {
	var mapped []*model.ReviewResponse
	for _, review := range reviews {
		mapped = append(mapped, r.mapResponseReview(review))
	}
	return mapped
}

func (r *reviewResponseMapper) mapResponseReviewRelationDetail(review *pb.ReviewsDetailResponse) *model.ReviewRelationDetailResponse {
	if review == nil {
		return nil
	}

	var reviewDetail *model.ReviewDetailResponse
	if review.ReviewDetail != nil {
		reviewDetail = &model.ReviewDetailResponse{
			ID:        int32(review.ReviewDetail.Id),
			Type:      &review.ReviewDetail.Type,
			URL:       &review.ReviewDetail.Url,
			Caption:   &review.ReviewDetail.Caption,
			CreatedAt: &review.ReviewDetail.CreatedAt,
		}
	}

	var deletedAt *string
	if review.DeletedAt != "" {
		deletedAt = &review.DeletedAt
	}

	return &model.ReviewRelationDetailResponse{
		ID:           int32(review.Id),
		UserID:       int32(review.UserId),
		ProductID:    int32(review.ProductId),
		Name:         review.Name,
		Comment:      review.Comment,
		Rating:       int32(review.Rating),
		ReviewDetail: reviewDetail,
		CreatedAt:    review.CreatedAt,
		UpdatedAt:    review.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}

func (r *reviewResponseMapper) mapResponsesReviewsRelationDetail(reviews []*pb.ReviewsDetailResponse) []*model.ReviewRelationDetailResponse {
	var mapped []*model.ReviewRelationDetailResponse
	for _, review := range reviews {
		mapped = append(mapped, r.mapResponseReviewRelationDetail(review))
	}
	return mapped
}

func (r *reviewResponseMapper) mapResponseReviewDeleteAt(review *pb.ReviewResponseDeleteAt) *model.ReviewResponseDeleteAt {
	var deletedAt *string
	if review.DeletedAt != nil {
		deletedAt = &review.DeletedAt.Value
	}

	return &model.ReviewResponseDeleteAt{
		ID:        int32(review.Id),
		UserID:    int32(review.UserId),
		ProductID: int32(review.ProductId),
		Name:      review.Name,
		Comment:   review.Comment,
		Rating:    int32(review.Rating),
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

func (r *reviewResponseMapper) mapResponsesReviewDeleteAt(reviews []*pb.ReviewResponseDeleteAt) []*model.ReviewResponseDeleteAt {
	var mapped []*model.ReviewResponseDeleteAt
	for _, review := range reviews {
		mapped = append(mapped, r.mapResponseReviewDeleteAt(review))
	}
	return mapped
}
