package merchant_detailgraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_detail"
)

type merchantDetailResponseMapper struct{}

func NewMerchantDetailResponseMapper() *merchantDetailResponseMapper {
	return &merchantDetailResponseMapper{}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetailDelete(res *pb_merchant.ApiResponseMerchantDelete) *model.APIResponseMerchantDetailDelete {
	return &model.APIResponseMerchantDetailDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetailAll(res *pb_merchant.ApiResponseMerchantAll) *model.APIResponseMerchantDetailAll {
	return &model.APIResponseMerchantDetailAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetailRelation(res *pb_merchant_detail.ApiResponseMerchantDetail) *model.APIResponseMerchantDetailRelation {
	return &model.APIResponseMerchantDetailRelation{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantDetailRelation(res.Data),
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetail(res *pb_merchant_detail.ApiResponseMerchantDetail) *model.APIResponseMerchantDetail {
	return &model.APIResponseMerchantDetail{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantDetail(res.Data),
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponseMerchantDetailDeleteAt(res *pb_merchant_detail.ApiResponseMerchantDetailDeleteAt) *model.APIResponseMerchantDetailDeleteAt {
	return &model.APIResponseMerchantDetailDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantDetailDeleteAt(res.Data),
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponsePaginationMerchantDetail(res *pb_merchant_detail.ApiResponsePaginationMerchantDetail) *model.APIResponsePaginationMerchantDetail {
	return &model.APIResponsePaginationMerchantDetail{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantDetailRelation(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantDetailResponseMapper) ToGraphqlResponsePaginationMerchantDetailDeleteAt(res *pb_merchant_detail.ApiResponsePaginationMerchantDetailDeleteAt) *model.APIResponsePaginationMerchantDetailDeleteAt {
	return &model.APIResponsePaginationMerchantDetailDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantDetailRelationDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantDetailResponseMapper) mapResponseMerchantDetailRelation(merchant *pb_merchant_detail.MerchantDetailResponse) *model.MerchantDetailRelationResponse {
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

func (m *merchantDetailResponseMapper) mapResponsesMerchantDetailRelation(merchants []*pb_merchant_detail.MerchantDetailResponse) []*model.MerchantDetailRelationResponse {
	var responses []*model.MerchantDetailRelationResponse

	for _, merchant := range merchants {
		responses = append(responses, m.mapResponseMerchantDetailRelation(merchant))
	}

	return responses
}

func (m *merchantDetailResponseMapper) mapResponseMerchantDetailRelationDeleteAt(merchant *pb_merchant_detail.MerchantDetailResponseDeleteAt) *model.MerchantDetailRelationResponseDeleteAt {
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

func (m *merchantDetailResponseMapper) mapResponsesMerchantDetailRelationDeleteAt(merchants []*pb_merchant_detail.MerchantDetailResponseDeleteAt) []*model.MerchantDetailRelationResponseDeleteAt {
	var responses []*model.MerchantDetailRelationResponseDeleteAt

	for _, merchant := range merchants {
		responses = append(responses, m.mapResponseMerchantDetailRelationDeleteAt(merchant))
	}

	return responses
}

func (m *merchantDetailResponseMapper) mapResponseMerchantDetail(merchant *pb_merchant_detail.MerchantDetailResponse) *model.MerchantDetailResponse {

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

func (m *merchantDetailResponseMapper) mapResponsesMerchantDetail(merchants []*pb_merchant_detail.MerchantDetailResponse) []*model.MerchantDetailResponse {
	var mappedMerchants []*model.MerchantDetailResponse

	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.mapResponseMerchantDetail(merchant))
	}

	return mappedMerchants
}

func (m *merchantDetailResponseMapper) mapResponseMerchantDetailDeleteAt(merchant *pb_merchant_detail.MerchantDetailResponseDeleteAt) *model.MerchantDetailResponseDeleteAt {
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

func (m *merchantDetailResponseMapper) mapResponsesMerchantDetailDeleteAt(merchants []*pb_merchant_detail.MerchantDetailResponseDeleteAt) []*model.MerchantDetailResponseDeleteAt {
	var mappedMerchants []*model.MerchantDetailResponseDeleteAt

	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.mapResponseMerchantDetailDeleteAt(merchant))
	}

	return mappedMerchants
}
