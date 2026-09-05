package merchant_awardgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type merchantAwardResponseMapper struct{}

func NewMerchantAwardResponseMapper() *merchantAwardResponseMapper {
	return &merchantAwardResponseMapper{}
}

func (m *merchantAwardResponseMapper) ToGraphqlResponseMerchantAwardDelete(res *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantAwardDelete {
	return &model.APIResponseMerchantAwardDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantAwardResponseMapper) ToGraphqlResponseMerchantAwardAll(res *pb.ApiResponseMerchantAll) *model.APIResponseMerchantAwardAll {
	return &model.APIResponseMerchantAwardAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantAwardResponseMapper) ToGraphqlResponseMerchantAward(res *pb.ApiResponseMerchantAward) *model.APIResponseMerchantAward {
	return &model.APIResponseMerchantAward{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantAward(res.Data),
	}
}

func (m *merchantAwardResponseMapper) ToGraphqlResponseMerchantAwardDeleteAt(res *pb.ApiResponseMerchantAwardDeleteAt) *model.APIResponseMerchantAwardDeleteAt {
	return &model.APIResponseMerchantAwardDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantAwardDeleteAt(res.Data),
	}
}

func (m *merchantAwardResponseMapper) ToGraphqlResponseMerchantAwards(res *pb.ApiResponsesMerchantAward) *model.APIResponsesMerchantAward {
	return &model.APIResponsesMerchantAward{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponsesMerchantAward(res.Data),
	}
}

func (m *merchantAwardResponseMapper) ToGraphqlResponsePaginationMerchantAwardDeleteAt(
	res *pb.ApiResponsePaginationMerchantAwardDeleteAt,
) *model.APIResponsePaginationMerchantAwardDeleteAt {
	return &model.APIResponsePaginationMerchantAwardDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantAwardDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantAwardResponseMapper) ToGraphqlPaginationMerchantAward(
	res *pb.ApiResponsePaginationMerchantAward,
) *model.APIResponsePaginationMerchantAward {
	return &model.APIResponsePaginationMerchantAward{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantAward(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantAwardResponseMapper) mapResponseMerchantAward(merchantAward *pb.MerchantAwardResponse) *model.MerchantAwardResponse {
	return &model.MerchantAwardResponse{
		ID:             int32(merchantAward.Id),
		MerchantID:     int32(merchantAward.MerchantId),
		Title:          merchantAward.Title,
		Description:    merchantAward.Description,
		IssuedBy:       merchantAward.IssuedBy,
		IssueDate:      merchantAward.IssueDate,
		ExpiryDate:     merchantAward.ExpiryDate,
		CertificateURL: merchantAward.CertificateUrl,
		CreatedAt:      merchantAward.CreatedAt,
		UpdatedAt:      merchantAward.UpdatedAt,
	}
}

func (m *merchantAwardResponseMapper) mapResponsesMerchantAward(merchantsAward []*pb.MerchantAwardResponse) []*model.MerchantAwardResponse {
	var responses []*model.MerchantAwardResponse

	for _, merchant := range merchantsAward {
		responses = append(responses, m.mapResponseMerchantAward(merchant))
	}

	return responses
}

func (m *merchantAwardResponseMapper) mapResponseMerchantAwardDeleteAt(merchantAward *pb.MerchantAwardResponseDeleteAt) *model.MerchantAwardResponseDeleteAt {
	return &model.MerchantAwardResponseDeleteAt{
		ID:             int32(merchantAward.Id),
		MerchantID:     int32(merchantAward.MerchantId),
		Title:          merchantAward.Title,
		Description:    merchantAward.Description,
		IssuedBy:       merchantAward.IssuedBy,
		IssueDate:      merchantAward.IssueDate,
		ExpiryDate:     merchantAward.ExpiryDate,
		CertificateURL: merchantAward.CertificateUrl,
		CreatedAt:      merchantAward.CreatedAt,
		UpdatedAt:      merchantAward.UpdatedAt,
	}
}

func (m *merchantAwardResponseMapper) mapResponsesMerchantAwardDeleteAt(merchantsAward []*pb.MerchantAwardResponseDeleteAt) []*model.MerchantAwardResponseDeleteAt {
	var responses []*model.MerchantAwardResponseDeleteAt

	for _, merchant := range merchantsAward {
		responses = append(responses, m.mapResponseMerchantAwardDeleteAt(merchant))
	}

	return responses
}
