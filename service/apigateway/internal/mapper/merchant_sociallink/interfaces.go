package merchant_sociallinkgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_social_link"
)

type MerchantSocialLinkGraphqlMapper interface {
	ToGraphqlResponseMerchantSocialLink(res *pb_merchant_social_link.ApiResponseMerchantSocial) *model.APIResponseMerchantSocialMediaLink
}
