package bannergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type bannerResponseMapper struct {
}

func NewBannerResponseMapper() *bannerResponseMapper {
	return &bannerResponseMapper{}
}

func (s *bannerResponseMapper) ToGraphqlResponseAll(res *pb.ApiResponseBannerAll) *model.APIResponseBannerAll {
	return &model.APIResponseBannerAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *bannerResponseMapper) ToGraphqlResponseDelete(res *pb.ApiResponseBannerDelete) *model.APIResponseBannerDelete {
	return &model.APIResponseBannerDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *bannerResponseMapper) ToGraphqlResponseBanner(res *pb.ApiResponseBanner) *model.APIResponseBanner {
	return &model.APIResponseBanner{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseBanner(res.Data),
	}
}

func (s *bannerResponseMapper) ToGraphqlResponseBannerDeleteAt(res *pb.ApiResponseBannerDeleteAt) *model.APIResponseBannerDeleteAt {
	return &model.APIResponseBannerDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseBannerDeleteAt(res.Data),
	}
}

func (s *bannerResponseMapper) ToGraphqlResponsesBanner(res *pb.ApiResponsesBanner) *model.APIResponsesBanner {
	return &model.APIResponsesBanner{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponsesBanner(res.Data),
	}
}

func (s *bannerResponseMapper) ToGraphqlResponsePaginationBanner(res *pb.ApiResponsePaginationBanner) *model.APIResponsePaginationBanner {
	return &model.APIResponsePaginationBanner{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesBanner(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *bannerResponseMapper) ToGraphqlResponsePaginationBannerDeleteAt(res *pb.ApiResponsePaginationBannerDeleteAt) *model.APIResponsePaginationBannerDeleteAt {
	return &model.APIResponsePaginationBannerDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesBannerDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *bannerResponseMapper) mapResponseBanner(banner *pb.BannerResponse) *model.BannerResponse {
	return &model.BannerResponse{
		BannerID:  banner.BannerId,
		Name:      banner.Name,
		StartDate: banner.StartDate,
		EndDate:   banner.EndDate,
		StartTime: banner.StartTime,
		EndTime:   banner.EndTime,
		IsActive:  banner.IsActive,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}
}

func (s *bannerResponseMapper) mapResponsesBanner(banners []*pb.BannerResponse) []*model.BannerResponse {
	var responses []*model.BannerResponse

	for _, banner := range banners {
		responses = append(responses, s.mapResponseBanner(banner))
	}

	return responses
}

func (s *bannerResponseMapper) mapResponseBannerDeleteAt(banner *pb.BannerResponseDeleteAt) *model.BannerResponseDeleteAt {
	var deletedAt string

	if banner.DeletedAt != nil {
		deletedAt = banner.DeletedAt.Value
	}

	return &model.BannerResponseDeleteAt{
		BannerID:  banner.BannerId,
		Name:      banner.Name,
		StartDate: banner.StartDate,
		EndDate:   banner.EndDate,
		StartTime: banner.StartTime,
		EndTime:   banner.EndTime,
		IsActive:  banner.IsActive,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

func (s *bannerResponseMapper) mapResponsesBannerDeleteAt(banners []*pb.BannerResponseDeleteAt) []*model.BannerResponseDeleteAt {
	var responses []*model.BannerResponseDeleteAt

	for _, banner := range banners {
		responses = append(responses, s.mapResponseBannerDeleteAt(banner))
	}

	return responses
}
