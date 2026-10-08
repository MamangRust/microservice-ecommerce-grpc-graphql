package review_detailgraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/review"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/review_detail"
)

type reviewDetailResponseMapper struct{}

func NewReviewDetailResponseMapper() *reviewDetailResponseMapper {
	return &reviewDetailResponseMapper{}
}

func (s *reviewDetailResponseMapper) ToGraphqlResponseAll(res *pb_review.ApiResponseReviewAll) *model.APIResponseReviewDetailAll {
	return &model.APIResponseReviewDetailAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *reviewDetailResponseMapper) ToGraphqlResponseDelete(res *pb_review.ApiResponseReviewDelete) *model.APIResponseReviewDetailDelete {
	return &model.APIResponseReviewDetailDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *reviewDetailResponseMapper) ToGraphqlResponseReviewDetail(res *pb_review_detail.ApiResponseReviewDetail) *model.APIResponseReviewDetail {
	return &model.APIResponseReviewDetail{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseReviewDetail(res.Data),
	}
}

func (m *reviewDetailResponseMapper) ToGraphqlResponsesReviewDetail(res *pb_review_detail.ApiResponsesReviewDetails) *model.APIResponsesReviewDetails {
	return &model.APIResponsesReviewDetails{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponsesReviewDetail(res.Data),
	}
}

func (m *reviewDetailResponseMapper) ToGraphqlResponseReviewDetailDeleteAt(res *pb_review_detail.ApiResponseReviewDetailDeleteAt) *model.APIResponseReviewDetailDeleteAt {
	return &model.APIResponseReviewDetailDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseReviewDetailDeleteAt(res.Data),
	}
}

func (m *reviewDetailResponseMapper) ToGraphqlResponsePaginationReviewDetail(
	res *pb_review_detail.ApiResponsePaginationReviewDetails,
) *model.APIResponsePaginationReviewDetails {
	return &model.APIResponsePaginationReviewDetails{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesReviewDetail(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *reviewDetailResponseMapper) ToGraphqlResponsePaginationReviewDetailDeleteAt(
	res *pb_review_detail.ApiResponsePaginationReviewDetailsDeleteAt,
) *model.APIResponsePaginationReviewDetailsDeleteAt {
	return &model.APIResponsePaginationReviewDetailsDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesReviewDetailDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *reviewDetailResponseMapper) mapResponseReviewDetail(reviewDetail *pb_review_detail.ReviewDetailsResponse) *model.ReviewDetailResponse {
	if reviewDetail == nil {
		return nil
	}

	return &model.ReviewDetailResponse{
		ID:        int32(reviewDetail.Id),
		ReviewID:  int32(reviewDetail.ReviewId),
		Type:      &reviewDetail.Type,
		URL:       &reviewDetail.Url,
		Caption:   &reviewDetail.Caption,
		CreatedAt: &reviewDetail.CreatedAt,
		UpdatedAt: &reviewDetail.UpdatedAt,
	}
}

func (m *reviewDetailResponseMapper) mapResponsesReviewDetail(reviewDetails []*pb_review_detail.ReviewDetailsResponse) []*model.ReviewDetailResponse {
	mapped := make([]*model.ReviewDetailResponse, 0, len(reviewDetails))
	for _, r := range reviewDetails {
		mapped = append(mapped, m.mapResponseReviewDetail(r))
	}
	return mapped
}

func (m *reviewDetailResponseMapper) mapResponseReviewDetailDeleteAt(reviewDetail *pb_review_detail.ReviewDetailsResponseDeleteAt) *model.ReviewDetailResponseDeletedAt {
	var deletedAt string

	if reviewDetail.DeletedAt != nil {
		deletedAt = reviewDetail.DeletedAt.Value
	}

	return &model.ReviewDetailResponseDeletedAt{
		ID:        int32(reviewDetail.Id),
		ReviewID:  int32(reviewDetail.ReviewId),
		Type:      &reviewDetail.Type,
		URL:       &reviewDetail.Url,
		Caption:   &reviewDetail.Caption,
		CreatedAt: &reviewDetail.CreatedAt,
		UpdatedAt: &reviewDetail.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}

func (m *reviewDetailResponseMapper) mapResponsesReviewDetailDeleteAt(reviewDetails []*pb_review_detail.ReviewDetailsResponseDeleteAt) []*model.ReviewDetailResponseDeletedAt {
	mapped := make([]*model.ReviewDetailResponseDeletedAt, 0, len(reviewDetails))
	for _, r := range reviewDetails {
		mapped = append(mapped, m.mapResponseReviewDetailDeleteAt(r))
	}
	return mapped
}
