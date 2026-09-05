package merchant_detailgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type merchantDetailResponseMapper struct{}

func NewMerchantDetailResponseMapper() *merchantDetailResponseMapper {
	return &merchantDetailResponseMapper{}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetailDelete(res *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantDetailDelete {
	return &model.APIResponseMerchantDetailDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetailAll(res *pb.ApiResponseMerchantAll) *model.APIResponseMerchantDetailAll {
	return &model.APIResponseMerchantDetailAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetailRelation(res *pb.ApiResponseMerchantDetail) *model.APIResponseMerchantDetailRelation {
	return &model.APIResponseMerchantDetailRelation{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantDetailRelation(res.Data),
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetail(res *pb.ApiResponseMerchantDetail) *model.APIResponseMerchantDetail {
	return &model.APIResponseMerchantDetail{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantDetail(res.Data),
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetailDeleteAt(res *pb.ApiResponseMerchantDetailDeleteAt) *model.APIResponseMerchantDetailDeleteAt {
	return &model.APIResponseMerchantDetailDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantDetailDeleteAt(res.Data),
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponsePaginationMerchantDetail(res *pb.ApiResponsePaginationMerchantDetail) *model.APIResponsePaginationMerchantDetail {
	return &model.APIResponsePaginationMerchantDetail{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantDetailRelation(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponsePaginationMerchantDetailDeleteAt(res *pb.ApiResponsePaginationMerchantDetailDeleteAt) *model.APIResponsePaginationMerchantDetailDeleteAt {
	return &model.APIResponsePaginationMerchantDetailDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantDetailRelationDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantDetailResponseMapper) mapResponseMerchantDetailRelation(merchant *pb.MerchantDetailResponse) *model.MerchantDetailRelationResponse {
	var socialMediaLinks []*model.MerchantSocialMediaLinkResponse
	for _, sm := range merchant.SocialMediaLinks {
		socialMediaLinks = append(socialMediaLinks, &model.MerchantSocialMediaLinkResponse{
			ID:       int32(sm.Id),
			Platform: sm.Platform,
			URL:      sm.Url,
		})
	}

	return &model.MerchantDetailRelationResponse{
		ID:               int32(merchant.Id),
		MerchantID:       int32(merchant.MerchantId),
		DisplayName:      merchant.DisplayName,
		CoverImageURL:    merchant.CoverImageUrl,
		LogoURL:          merchant.LogoUrl,
		ShortDescription: merchant.ShortDescription,
		WebsiteURL:       merchant.WebsiteUrl,
		SocialMediaLinks: socialMediaLinks,
		CreatedAt:        merchant.CreatedAt,
		UpdatedAt:        merchant.UpdatedAt,
	}
}

func (m *merchantDetailResponseMapper) mapResponsesMerchantDetailRelation(merchants []*pb.MerchantDetailResponse) []*model.MerchantDetailRelationResponse {
	var responses []*model.MerchantDetailRelationResponse

	for _, merchant := range merchants {
		responses = append(responses, m.mapResponseMerchantDetailRelation(merchant))
	}

	return responses
}

func (m *merchantDetailResponseMapper) mapResponseMerchantDetailRelationDeleteAt(merchant *pb.MerchantDetailResponseDeleteAt) *model.MerchantDetailRelationResponseDeleteAt {
	var socialMediaLinks []*model.MerchantSocialMediaLinkResponse
	for _, sm := range merchant.SocialMediaLinks {
		socialMediaLinks = append(socialMediaLinks, &model.MerchantSocialMediaLinkResponse{
			ID:       int32(sm.Id),
			Platform: sm.Platform,
			URL:      sm.Url,
		})
	}

	return &model.MerchantDetailRelationResponseDeleteAt{
		ID:               int32(merchant.Id),
		MerchantID:       int32(merchant.MerchantId),
		DisplayName:      merchant.DisplayName,
		CoverImageURL:    merchant.CoverImageUrl,
		LogoURL:          merchant.LogoUrl,
		ShortDescription: merchant.ShortDescription,
		WebsiteURL:       merchant.WebsiteUrl,
		SocialMediaLinks: socialMediaLinks,
		CreatedAt:        merchant.CreatedAt,
		UpdatedAt:        merchant.UpdatedAt,
	}
}

func (m *merchantDetailResponseMapper) mapResponsesMerchantDetailRelationDeleteAt(merchants []*pb.MerchantDetailResponseDeleteAt) []*model.MerchantDetailRelationResponseDeleteAt {
	var responses []*model.MerchantDetailRelationResponseDeleteAt

	for _, merchant := range merchants {
		responses = append(responses, m.mapResponseMerchantDetailRelationDeleteAt(merchant))
	}

	return responses
}

func (m *merchantDetailResponseMapper) mapResponseMerchantDetail(merchant *pb.MerchantDetailResponse) *model.MerchantDetailResponse {

	return &model.MerchantDetailResponse{
		ID:               int32(merchant.Id),
		MerchantID:       int32(merchant.MerchantId),
		DisplayName:      merchant.DisplayName,
		CoverImageURL:    merchant.CoverImageUrl,
		LogoURL:          merchant.LogoUrl,
		ShortDescription: merchant.ShortDescription,
		WebsiteURL:       merchant.WebsiteUrl,
		CreatedAt:        merchant.CreatedAt,
		UpdatedAt:        merchant.UpdatedAt,
	}
}

func (m *merchantDetailResponseMapper) mapResponsesMerchantDetail(merchants []*pb.MerchantDetailResponse) []*model.MerchantDetailResponse {
	var mappedMerchants []*model.MerchantDetailResponse

	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.mapResponseMerchantDetail(merchant))
	}

	return mappedMerchants
}

func (m *merchantDetailResponseMapper) mapResponseMerchantDetailDeleteAt(merchant *pb.MerchantDetailResponseDeleteAt) *model.MerchantDetailResponseDeleteAt {
	var deletedAt string

	if merchant.DeletedAt != nil {
		deletedAt = merchant.DeletedAt.Value
	}

	return &model.MerchantDetailResponseDeleteAt{
		ID:               int32(merchant.Id),
		MerchantID:       int32(merchant.MerchantId),
		DisplayName:      merchant.DisplayName,
		CoverImageURL:    merchant.CoverImageUrl,
		LogoURL:          merchant.LogoUrl,
		ShortDescription: merchant.ShortDescription,
		WebsiteURL:       merchant.WebsiteUrl,
		CreatedAt:        merchant.CreatedAt,
		UpdatedAt:        merchant.UpdatedAt,
		DeletedAt:        &deletedAt,
	}
}

func (m *merchantDetailResponseMapper) mapResponsesMerchantDetailDeleteAt(merchants []*pb.MerchantDetailResponseDeleteAt) []*model.MerchantDetailResponseDeleteAt {
	var mappedMerchants []*model.MerchantDetailResponseDeleteAt

	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.mapResponseMerchantDetailDeleteAt(merchant))
	}

	return mappedMerchants
}
