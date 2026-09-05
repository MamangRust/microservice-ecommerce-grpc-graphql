package merchant_policygraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type merchantPolicyResponseMapper struct{}

func NewMerchantPolicyResponseMapper() *merchantPolicyResponseMapper {
	return &merchantPolicyResponseMapper{}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponseMerchantPolicyDelete(res *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantPolicyDelete {
	return &model.APIResponseMerchantPolicyDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponseMerchantPolicyAll(res *pb.ApiResponseMerchantAll) *model.APIResponseMerchantPolicyAll {
	return &model.APIResponseMerchantPolicyAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponseMerchantPolicy(res *pb.ApiResponseMerchantPolicies) *model.APIResponseMerchantPolicy {
	return &model.APIResponseMerchantPolicy{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantPolicy(res.Data),
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponseMerchantPolicyDeleteAt(res *pb.ApiResponseMerchantPoliciesDeleteAt) *model.APIResponseMerchantPolicyDeleteAt {
	return &model.APIResponseMerchantPolicyDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantPolicyDeleteAt(res.Data),
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponsesMerchantPolicy(res *pb.ApiResponsesMerchantPolicies) *model.APIResponsesMerchantPolicy {
	return &model.APIResponsesMerchantPolicy{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponsesMerchantPolicy(res.Data),
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponsePaginationMerchantPolicyDeleteAt(res *pb.ApiResponsePaginationMerchantPoliciesDeleteAt) *model.APIResponsePaginationMerchantPolicyDeleteAt {
	return &model.APIResponsePaginationMerchantPolicyDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantPolicyDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantPolicyResponseMapper) ToGraphqlResponsePaginationMerchantPolicy(res *pb.ApiResponsePaginationMerchantPolicies) *model.APIResponsePaginationMerchantPolicy {
	return &model.APIResponsePaginationMerchantPolicy{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantPolicy(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantPolicyResponseMapper) mapResponseMerchantPolicy(merchant *pb.MerchantPoliciesResponse) *model.MerchantPolicyResponse {
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

func (m *merchantPolicyResponseMapper) mapResponsesMerchantPolicy(merchants []*pb.MerchantPoliciesResponse) []*model.MerchantPolicyResponse {
	var mappedMerchants []*model.MerchantPolicyResponse
	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.mapResponseMerchantPolicy(merchant))
	}
	return mappedMerchants
}

func (m *merchantPolicyResponseMapper) mapResponseMerchantPolicyDeleteAt(merchant *pb.MerchantPoliciesResponseDeleteAt) *model.MerchantPolicyResponseDeleteAt {

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

func (m *merchantPolicyResponseMapper) mapResponsesMerchantPolicyDeleteAt(merchants []*pb.MerchantPoliciesResponseDeleteAt) []*model.MerchantPolicyResponseDeleteAt {
	var mappedMerchants []*model.MerchantPolicyResponseDeleteAt
	for _, merchant := range merchants {
		mappedMerchants = append(mappedMerchants, m.mapResponseMerchantPolicyDeleteAt(merchant))
	}
	return mappedMerchants
}
