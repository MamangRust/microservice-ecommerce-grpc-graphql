package merchant_policygraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_policy"
)

type merchantPolicyResponseMapper struct{}

func NewMerchantPolicyResponseMapper() *merchantPolicyResponseMapper {
	return &merchantPolicyResponseMapper{}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponseMerchantPolicyDelete(res *pb_merchant.ApiResponseMerchantDelete) *model.APIResponseMerchantPolicyDelete {
	return &model.APIResponseMerchantPolicyDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponseMerchantPolicyAll(res *pb_merchant.ApiResponseMerchantAll) *model.APIResponseMerchantPolicyAll {
	return &model.APIResponseMerchantPolicyAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponseMerchantPolicy(res *pb_merchant_policy.ApiResponseMerchantPolicies) *model.APIResponseMerchantPolicy {
	return &model.APIResponseMerchantPolicy{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantPolicy(res.Data),
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponseMerchantPolicyDeleteAt(res *pb_merchant_policy.ApiResponseMerchantPoliciesDeleteAt) *model.APIResponseMerchantPolicyDeleteAt {
	return &model.APIResponseMerchantPolicyDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantPolicyDeleteAt(res.Data),
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponsesMerchantPolicy(res *pb_merchant_policy.ApiResponsesMerchantPolicies) *model.APIResponsesMerchantPolicy {
	return &model.APIResponsesMerchantPolicy{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponsesMerchantPolicy(res.Data),
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponsePaginationMerchantPolicyDeleteAt(res *pb_merchant_policy.ApiResponsePaginationMerchantPoliciesDeleteAt) *model.APIResponsePaginationMerchantPolicyDeleteAt {
	return &model.APIResponsePaginationMerchantPolicyDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantPolicyDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponsePaginationMerchantPolicy(res *pb_merchant_policy.ApiResponsePaginationMerchantPolicies) *model.APIResponsePaginationMerchantPolicy {
	return &model.APIResponsePaginationMerchantPolicy{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantPolicy(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantPolicyResponseMapper) mapResponseMerchantPolicy(merchant *pb_merchant_policy.MerchantPoliciesResponse) *model.MerchantPolicyResponse {
	return &model.MerchantPolicyResponse{
		ID:          int32(merchant.Id),
		MerchantID:  int32(merchant.MerchantId),
		PolicyType:  merchant.PolicyType,
		Title:       merchant.Title,
		Description: merchant.Description,
		CreatedAt:   merchant.CreatedAt,
		UpdatedAt:   merchant.UpdatedAt,
	}
}

func (m *merchantPolicyResponseMapper) mapResponsesMerchantPolicy(merchants []*pb_merchant_policy.MerchantPoliciesResponse) []*model.MerchantPolicyResponse {
	var mappedMerchants []*model.MerchantPolicyResponse
	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.mapResponseMerchantPolicy(merchant))
	}
	return mappedMerchants
}

func (m *merchantPolicyResponseMapper) mapResponseMerchantPolicyDeleteAt(merchant *pb_merchant_policy.MerchantPoliciesResponseDeleteAt) *model.MerchantPolicyResponseDeleteAt {

	return &model.MerchantPolicyResponseDeleteAt{
		ID:          int32(merchant.Id),
		MerchantID:  int32(merchant.MerchantId),
		PolicyType:  merchant.PolicyType,
		Title:       merchant.Title,
		Description: merchant.Description,
		CreatedAt:   merchant.CreatedAt,
		UpdatedAt:   merchant.UpdatedAt,
	}
}

func (m *merchantPolicyResponseMapper) mapResponsesMerchantPolicyDeleteAt(merchants []*pb_merchant_policy.MerchantPoliciesResponseDeleteAt) []*model.MerchantPolicyResponseDeleteAt {
	var mappedMerchants []*model.MerchantPolicyResponseDeleteAt
	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.mapResponseMerchantPolicyDeleteAt(merchant))
	}
	return mappedMerchants
}
