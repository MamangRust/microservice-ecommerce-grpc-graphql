package merchantgraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant"
)

type merchantResponseMapper struct{}

func NewMerchantResponseMapper() *merchantResponseMapper {
	return &merchantResponseMapper{}
}

func (m *merchantResponseMapper) ToGraphqlResponseMerchant(res *pb_merchant.ApiResponseMerchant) *model.APIResponseMerchant {
	return &model.APIResponseMerchant{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchant(res.Data),
	}
}

func (m *merchantResponseMapper) ToGraphqlResponsesMerchant(res *pb_merchant.ApiResponsesMerchant) *model.APIResponsesMerchant {
	return &model.APIResponsesMerchant{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponsesMerchant(res.Data),
	}
}

func (m *merchantResponseMapper) ToGraphqlResponseMerchantDeleteAt(res *pb_merchant.ApiResponseMerchantDeleteAt) *model.APIResponseMerchantDeleteAt {
	return &model.APIResponseMerchantDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseMerchantDeleteAt(res.Data),
	}
}

func (m *merchantResponseMapper) ToGraphqlResponseMerchantDelete(res *pb_merchant.ApiResponseMerchantDelete) *model.APIResponseMerchantDelete {
	return &model.APIResponseMerchantDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantResponseMapper) ToGraphqlResponseMerchantAll(res *pb_merchant.ApiResponseMerchantAll) *model.APIResponseMerchantAll {
	return &model.APIResponseMerchantAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *merchantResponseMapper) ToGraphqlResponsePaginationMerchantDeleteAt(
	res *pb_merchant.ApiResponsePaginationMerchantDeleteAt,
) *model.APIResponsePaginationMerchantDeleteAt {
	return &model.APIResponsePaginationMerchantDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchantDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantResponseMapper) ToGraphqlResponsePaginationMerchant(
	res *pb_merchant.ApiResponsePaginationMerchant,
) *model.APIResponsePaginationMerchant {
	return &model.APIResponsePaginationMerchant{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesMerchant(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *merchantResponseMapper) mapResponseMerchant(merchant *pb_merchant.MerchantResponse) *model.MerchantResponse {
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

func (m *merchantResponseMapper) mapResponsesMerchant(merchants []*pb_merchant.MerchantResponse) []*model.MerchantResponse {
	var mapped []*model.MerchantResponse
	for _, merchant := range merchants {
		mapped = append(mapped, m.mapResponseMerchant(merchant))
	}
	return mapped
}

func (m *merchantResponseMapper) mapResponseMerchantDeleteAt(merchant *pb_merchant.MerchantResponseDeleteAt) *model.MerchantResponseDeleteAt {
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

func (m *merchantResponseMapper) mapResponsesMerchantDeleteAt(merchants []*pb_merchant.MerchantResponseDeleteAt) []*model.MerchantResponseDeleteAt {
	var mapped []*model.MerchantResponseDeleteAt
	for _, merchant := range merchants {
		mapped = append(mapped, m.mapResponseMerchantDeleteAt(merchant))
	}
	return mapped
}
