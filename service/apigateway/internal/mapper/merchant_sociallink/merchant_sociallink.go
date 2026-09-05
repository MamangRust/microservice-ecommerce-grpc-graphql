package merchant_sociallinkgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type merchantSocialLinkResponseMapper struct {
}

func NewMerchantSocialLinkResponseMapper() *merchantSocialLinkResponseMapper {
	return &merchantSocialLinkResponseMapper{}
}

func (m *merchantSocialLinkResponseMapper) ToGraphqlResponseMerchantSocialLink(res *pb.ApiResponseMerchantSocial) *model.APIResponseMerchantSocialMediaLink {
	return &model.APIResponseMerchantSocialMediaLink{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantSocialLink(res.Data),
	}
}

func (m *merchantSocialLinkResponseMapper) mapResponseMerchantSocialLink(response *pb.MerchantSocialMediaLinkResponse) *model.MerchantSocialMediaLinkResponse {
	if response == nil {
		return nil
	}
	return &model.MerchantSocialMediaLinkResponse{
		ID:       int32(response.Id),
		Platform: response.Platform,
		URL:      response.Url,
	}
}
