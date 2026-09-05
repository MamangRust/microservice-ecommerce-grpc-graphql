package merchantgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantGraphqlMapper interface {
	ToGraphqlResponseMerchant(res *pb.ApiResponseMerchant) *model.APIResponseMerchant
	ToGraphqlResponsesMerchant(res *pb.ApiResponsesMerchant) *model.APIResponsesMerchant
	ToGraphqlResponseMerchantDeleteAt(res *pb.ApiResponseMerchantDeleteAt) *model.APIResponseMerchantDeleteAt
	ToGraphqlResponseMerchantDelete(res *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantDelete
	ToGraphqlResponseMerchantAll(res *pb.ApiResponseMerchantAll) *model.APIResponseMerchantAll
	ToGraphqlResponsePaginationMerchantDeleteAt(res *pb.ApiResponsePaginationMerchantDeleteAt) *model.APIResponsePaginationMerchantDeleteAt
	ToGraphqlResponsePaginationMerchant(res *pb.ApiResponsePaginationMerchant) *model.APIResponsePaginationMerchant
}
