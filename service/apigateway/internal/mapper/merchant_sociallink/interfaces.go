package merchant_sociallinkgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantSocialLinkGraphqlMapper interface {
	ToGraphqlResponseMerchantSocialLink(res *pb.ApiResponseMerchantSocial) *model.APIResponseMerchantSocialMediaLink
}
