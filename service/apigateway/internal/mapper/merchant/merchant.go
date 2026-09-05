package merchantgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type merchantResponseMapper struct{}

func NewMerchantResponseMapper() *merchantResponseMapper {
	return &merchantResponseMapper{}
}

func (m *merchantResponseMapper) ToGraphqlResponseMerchant(res *pb.ApiResponseMerchant) *model.APIResponseMerchant {
	return &model.APIResponseMerchant{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchant(res.Data),
	}
}

func (m *merchantResponseMapper) ToGraphqlResponsesMerchant(res *pb.ApiResponsesMerchant) *model.APIResponsesMerchant {
	return &model.APIResponsesMerchant{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponsesMerchant(res.Data),
	}
}

func (m *merchantResponseMapper) ToGraphqlResponseMerchantDeleteAt(res *pb.ApiResponseMerchantDeleteAt) *model.APIResponseMerchantDeleteAt {
	return &model.APIResponseMerchantDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantDeleteAt(res.Data),
	}
}

func (m *merchantResponseMapper) ToGraphqlResponseMerchantDelete(res *pb.ApiResponseMerchantDelete) *model.APIResponseMerchantDelete {
	return &model.APIResponseMerchantDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantResponseMapper) ToGraphqlResponseMerchantAll(res *pb.ApiResponseMerchantAll) *model.APIResponseMerchantAll {
	return &model.APIResponseMerchantAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantResponseMapper) ToGraphqlResponsePaginationMerchantDeleteAt(
	res *pb.ApiResponsePaginationMerchantDeleteAt,
) *model.APIResponsePaginationMerchantDeleteAt {
	return &model.APIResponsePaginationMerchantDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantResponseMapper) ToGraphqlResponsePaginationMerchant(
	res *pb.ApiResponsePaginationMerchant,
) *model.APIResponsePaginationMerchant {
	return &model.APIResponsePaginationMerchant{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchant(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantResponseMapper) mapResponseMerchant(merchant *pb.MerchantResponse) *model.MerchantResponse {
	return &model.MerchantResponse{
		ID:           int32(merchant.Id),
		UserID:       int32(merchant.UserId),
		Name:         merchant.Name,
		Description:  merchant.Description,
		Address:      merchant.Address,
		ContactEmail: merchant.ContactEmail,
		ContactPhone: merchant.ContactPhone,
		Status:       merchant.Status,
		CreatedAt:    merchant.CreatedAt,
		UpdatedAt:    merchant.UpdatedAt,
	}
}

func (m *merchantResponseMapper) mapResponsesMerchant(merchants []*pb.MerchantResponse) []*model.MerchantResponse {
	var mapped []*model.MerchantResponse
	for _, merchant := range merchants {
		mapped = append(mapped, m.mapResponseMerchant(merchant))
	}
	return mapped
}

func (m *merchantResponseMapper) mapResponseMerchantDeleteAt(merchant *pb.MerchantResponseDeleteAt) *model.MerchantResponseDeleteAt {
	var deletedAt string

	if merchant.DeletedAt != nil {
		deletedAt = merchant.DeletedAt.Value
	}

	return &model.MerchantResponseDeleteAt{
		ID:           int32(merchant.Id),
		UserID:       int32(merchant.UserId),
		Name:         merchant.Name,
		Description:  merchant.Description,
		Address:      merchant.Address,
		ContactEmail: merchant.ContactEmail,
		ContactPhone: merchant.ContactPhone,
		Status:       merchant.Status,
		CreatedAt:    merchant.CreatedAt,
		UpdatedAt:    merchant.UpdatedAt,
		DeletedAt:    &deletedAt,
	}
}

func (m *merchantResponseMapper) mapResponsesMerchantDeleteAt(merchants []*pb.MerchantResponseDeleteAt) []*model.MerchantResponseDeleteAt {
	var mapped []*model.MerchantResponseDeleteAt
	for _, merchant := range merchants {
		mapped = append(mapped, m.mapResponseMerchantDeleteAt(merchant))
	}
	return mapped
}
