package merchant_awardgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantAwardGraphqlMapper interface {
	ToGraphqlResponseMerchantAwardDelete(res *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantAwardDelete
	ToGraphqlResponseMerchantAwardAll(res *pb.ApiResponseMerchantAll) *model.APIResponseMerchantAwardAll
	ToGraphqlResponseMerchantAward(res *pb.ApiResponseMerchantAward) *model.APIResponseMerchantAward
	ToGraphqlResponseMerchantAwardDeleteAt(res *pb.ApiResponseMerchantAwardDeleteAt) *model.APIResponseMerchantAwardDeleteAt
	ToGraphqlResponseMerchantAwards(res *pb.ApiResponsesMerchantAward) *model.APIResponsesMerchantAward
	ToGraphqlResponsePaginationMerchantAwardDeleteAt(res *pb.ApiResponsePaginationMerchantAwardDeleteAt) *model.APIResponsePaginationMerchantAwardDeleteAt
	ToGraphqlPaginationMerchantAward(res *pb.ApiResponsePaginationMerchantAward) *model.APIResponsePaginationMerchantAward
}
