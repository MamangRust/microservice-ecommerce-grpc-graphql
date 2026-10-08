package merchant_sociallinkgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_detail"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_social_link"
)

type merchantSocialLinkResponseMapper struct {
}

func NewMerchantSocialLinkResponseMapper() *merchantSocialLinkResponseMapper {
	return &merchantSocialLinkResponseMapper{}
}

func (m *merchantSocialLinkResponseMapper) ToGraphqlResponseMerchantSocialLink(res *pb_merchant_social_link.ApiResponseMerchantSocial) *model.APIResponseMerchantSocialMediaLink {
	return &model.APIResponseMerchantSocialMediaLink{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantSocialLink(res.Data),
	}
}

func (m *merchantSocialLinkResponseMapper) mapResponseMerchantSocialLink(response *pb_merchant_detail.MerchantSocialMediaLinkResponse) *model.MerchantSocialMediaLinkResponse {
	if response == nil {
		return nil
	}
	return &model.MerchantSocialMediaLinkResponse{
		ID:       int32(response.Id),
		Platform: response.Platform,
		URL:      response.Url,
	}
}
