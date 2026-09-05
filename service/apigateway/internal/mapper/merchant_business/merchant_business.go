package merchant_businessgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type merchantBusinessResponseMapper struct {
}

func NewMerchantBusinessResponseMapper() *merchantBusinessResponseMapper {
	return &merchantBusinessResponseMapper{}
}

func (m *merchantBusinessResponseMapper) ToGraphqlResponseMerchantBusinessDelete(res *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantBusinessDelete {
	return &model.APIResponseMerchantBusinessDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantBusinessResponseMapper) ToGraphqlResponseMerchantBusinessAll(res *pb.ApiResponseMerchantAll) *model.APIResponseMerchantBusinessAll {
	return &model.APIResponseMerchantBusinessAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantBusinessResponseMapper) ToGraphqlResponseMerchantBusiness(res *pb.ApiResponseMerchantBusiness) *model.APIResponseMerchantBusiness {
	return &model.APIResponseMerchantBusiness{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapMerchantBusiness(res.Data),
	}
}

func (m *merchantBusinessResponseMapper) ToGraphqlResponseMerchantBusinessDeleteAt(
	res *pb.ApiResponseMerchantBusinessDeleteAt,
) *model.APIResponseMerchantBusinessDeleteAt {
	return &model.APIResponseMerchantBusinessDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapMerchantBusinessDeleteAt(res.Data),
	}
}

func (m *merchantBusinessResponseMapper) ToGraphqlResponsesMerchantBusiness(
	res *pb.ApiResponsesMerchantBusiness,
) *model.APIResponsesMerchantBusiness {
	return &model.APIResponsesMerchantBusiness{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapMerchantBusinesses(res.Data),
	}
}

func (m *merchantBusinessResponseMapper) ToGraphqlResponsePaginationMerchantBusinessDeleteAt(
	res *pb.ApiResponsePaginationMerchantBusinessDeleteAt,
) *model.APIResponsePaginationMerchantBusinessDeleteAt {
	return &model.APIResponsePaginationMerchantBusinessDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapMerchantBusinessesDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantBusinessResponseMapper) ToGraphqlResponsePaginationMerchantBusiness(
	res *pb.ApiResponsePaginationMerchantBusiness,
) *model.APIResponsePaginationMerchantBusiness {
	return &model.APIResponsePaginationMerchantBusiness{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapMerchantBusinesses(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantBusinessResponseMapper) mapMerchantBusiness(res *pb.MerchantBusinessResponse) *model.MerchantBusinessResponse {
	if res == nil {
		return nil
	}
	return &model.MerchantBusinessResponse{
		ID:                int32(res.Id),
		MerchantID:        int32(res.MerchantId),
		BusinessType:      res.BusinessType,
		TaxID:             res.TaxId,
		EstablishedYear:   int32(res.EstablishedYear),
		NumberOfEmployees: int32(res.NumberOfEmployees),
		WebsiteURL:        res.WebsiteUrl,
		CreatedAt:         res.CreatedAt,
		UpdatedAt:         res.UpdatedAt,
	}
}

func (m *merchantBusinessResponseMapper) mapMerchantBusinesses(responses []*pb.MerchantBusinessResponse) []*model.MerchantBusinessResponse {
	if responses == nil {
		return nil
	}
	var result []*model.MerchantBusinessResponse
	for _, r := range responses {
		result = append(result, m.mapMerchantBusiness(r))
	}
	return result
}

func (m *merchantBusinessResponseMapper) mapMerchantBusinessDeleteAt(res *pb.MerchantBusinessResponseDeleteAt) *model.MerchantBusinessResponseDeleteAt {
	var deletedAt string

	if res.DeletedAt != nil {
		deletedAt = res.DeletedAt.Value
	}

	return &model.MerchantBusinessResponseDeleteAt{
		ID:                int32(res.Id),
		MerchantID:        int32(res.MerchantId),
		BusinessType:      res.BusinessType,
		TaxID:             res.TaxId,
		EstablishedYear:   int32(res.EstablishedYear),
		NumberOfEmployees: int32(res.NumberOfEmployees),
		WebsiteURL:        res.WebsiteUrl,
		CreatedAt:         res.CreatedAt,
		UpdatedAt:         res.UpdatedAt,
		DeletedAt:         &deletedAt,
	}
}

func (m *merchantBusinessResponseMapper) mapMerchantBusinessesDeleteAt(responses []*pb.MerchantBusinessResponseDeleteAt) []*model.MerchantBusinessResponseDeleteAt {
	if responses == nil {
		return nil
	}
	var result []*model.MerchantBusinessResponseDeleteAt
	for _, r := range responses {
		result = append(result, m.mapMerchantBusinessDeleteAt(r))
	}
	return result
}
