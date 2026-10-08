package merchantgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant"
)

type MerchantGraphqlMapper interface {
	ToGraphqlResponseMerchant(res *pb_merchant.ApiResponseMerchant) *model.APIResponseMerchant
	ToGraphqlResponsesMerchant(res *pb_merchant.ApiResponsesMerchant) *model.APIResponsesMerchant
	ToGraphqlResponseMerchantDeleteAt(res *pb_merchant.ApiResponseMerchantDeleteAt) *model.APIResponseMerchantDeleteAt
	ToGraphqlResponseMerchantDelete(res *pb_merchant.ApiResponseMerchantDelete) *model.APIResponseMerchantDelete
	ToGraphqlResponseMerchantAll(res *pb_merchant.ApiResponseMerchantAll) *model.APIResponseMerchantAll
	ToGraphqlResponsePaginationMerchantDeleteAt(res *pb_merchant.ApiResponsePaginationMerchantDeleteAt) *model.APIResponsePaginationMerchantDeleteAt
	ToGraphqlResponsePaginationMerchant(res *pb_merchant.ApiResponsePaginationMerchant) *model.APIResponsePaginationMerchant
}
