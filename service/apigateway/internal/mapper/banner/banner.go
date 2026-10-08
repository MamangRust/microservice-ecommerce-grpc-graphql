package bannergraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/banner"
)

type bannerResponseMapper struct {
}

func NewBannerResponseMapper() *bannerResponseMapper {
	return &bannerResponseMapper{}
}

func (s *bannerResponseMapper) ToGraphqlResponseAll(res *pb_banner.ApiResponseBannerAll) *model.APIResponseBannerAll {
	return &model.APIResponseBannerAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *bannerResponseMapper) ToGraphqlResponseDelete(res *pb_banner.ApiResponseBannerDelete) *model.APIResponseBannerDelete {
	return &model.APIResponseBannerDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *bannerResponseMapper) ToGraphqlResponseBanner(res *pb_banner.ApiResponseBanner) *model.APIResponseBanner {
	return &model.APIResponseBanner{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseBanner(res.Data),
	}
}

func (s *bannerResponseMapper) ToGraphqlResponseBannerDeleteAt(res *pb_banner.ApiResponseBannerDeleteAt) *model.APIResponseBannerDeleteAt {
	return &model.APIResponseBannerDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseBannerDeleteAt(res.Data),
	}
}

func (s *bannerResponseMapper) ToGraphqlResponsesBanner(res *pb_banner.ApiResponsesBanner) *model.APIResponsesBanner {
	return &model.APIResponsesBanner{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponsesBanner(res.Data),
	}
}

func (s *bannerResponseMapper) ToGraphqlResponsePaginationBanner(res *pb_banner.ApiResponsePaginationBanner) *model.APIResponsePaginationBanner {
	return &model.APIResponsePaginationBanner{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesBanner(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *bannerResponseMapper) ToGraphqlResponsePaginationBannerDeleteAt(res *pb_banner.ApiResponsePaginationBannerDeleteAt) *model.APIResponsePaginationBannerDeleteAt {
	return &model.APIResponsePaginationBannerDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesBannerDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *bannerResponseMapper) mapResponseBanner(banner *pb_banner.BannerResponse) *model.BannerResponse {
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

func (s *bannerResponseMapper) mapResponsesBanner(banners []*pb_banner.BannerResponse) []*model.BannerResponse {
	var responses []*model.BannerResponse

	for _, banner := range banners {
		responses = append(responses, s.mapResponseBanner(banner))
	}

	return responses
}

func (s *bannerResponseMapper) mapResponseBannerDeleteAt(banner *pb_banner.BannerResponseDeleteAt) *model.BannerResponseDeleteAt {
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

func (s *bannerResponseMapper) mapResponsesBannerDeleteAt(banners []*pb_banner.BannerResponseDeleteAt) []*model.BannerResponseDeleteAt {
	var responses []*model.BannerResponseDeleteAt

	for _, banner := range banners {
		responses = append(responses, s.mapResponseBannerDeleteAt(banner))
	}

	return responses
}
