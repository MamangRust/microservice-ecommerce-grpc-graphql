package merchant_policygraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type MerchantPolicyGraphqlMapper interface {
	ToGraphqlResponseMerchantPolicyDelete(res *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantPolicyDelete
	ToGraphqlResponseMerchantPolicyAll(res *pb.ApiResponseMerchantAll) *model.APIResponseMerchantPolicyAll
	ToGraphqlResponseMerchantPolicy(res *pb.ApiResponseMerchantPolicies) *model.APIResponseMerchantPolicy
	ToGraphqlResponseMerchantPolicyDeleteAt(res *pb.ApiResponseMerchantPoliciesDeleteAt) *model.APIResponseMerchantPolicyDeleteAt
	ToGraphqlResponsesMerchantPolicy(res *pb.ApiResponsesMerchantPolicies) *model.APIResponsesMerchantPolicy
	ToGraphqlResponsePaginationMerchantPolicyDeleteAt(res *pb.ApiResponsePaginationMerchantPoliciesDeleteAt) *model.APIResponsePaginationMerchantPolicyDeleteAt
	ToGraphqlResponsePaginationMerchantPolicy(res *pb.ApiResponsePaginationMerchantPolicies) *model.APIResponsePaginationMerchantPolicy
}
