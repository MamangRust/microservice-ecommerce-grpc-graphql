package shipping_addressgraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/shipping_address"
)

type shippingAddresResponseMapper struct{}

func NewshippingAddresResponseMapper() *shippingAddresResponseMapper {
	return &shippingAddresResponseMapper{}
}

func (s *shippingAddresResponseMapper) ToGraphResponseShippingAddress(res *pb_shipping_address.ApiResponseShipping) *model.APIResponseShipping {
	return &model.APIResponseShipping{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseShippingAddress(res.Data),
	}
}

func (s *shippingAddresResponseMapper) ToGraphResponseShippingAddressDeleteAt(res *pb_shipping_address.ApiResponseShippingDeleteAt) *model.APIResponseShippingDeleteAt {
	return &model.APIResponseShippingDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseShippingAddressDeleteAt(res.Data),
	}
}

func (s *shippingAddresResponseMapper) ToGraphResponsesShippingAddress(res *pb_shipping_address.ApiResponsesShipping) *model.APIResponsesShipping {
	return &model.APIResponsesShipping{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponsesShippingAddress(res.Data),
	}
}

func (s *shippingAddresResponseMapper) ToGraphResponseShippingAddressDelete(res *pb_shipping_address.ApiResponseShippingDelete) *model.APIResponseShippingDelete {
	return &model.APIResponseShippingDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *shippingAddresResponseMapper) ToGraphResponseShippingAddressAll(res *pb_shipping_address.ApiResponseShippingAll) *model.APIResponseShippingAll {
	return &model.APIResponseShippingAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *shippingAddresResponseMapper) ToGraphResponsePaginationShippingAddressDeleteAt(
	res *pb_shipping_address.ApiResponsePaginationShippingDeleteAt,
) *model.APIResponsePaginationShippingDeleteAt {
	return &model.APIResponsePaginationShippingDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesShippingAddressDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *shippingAddresResponseMapper) ToGraphResponsePaginationShippingAddress(
	res *pb_shipping_address.ApiResponsePaginationShipping,
) *model.APIResponsePaginationShipping {
	return &model.APIResponsePaginationShipping{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesShippingAddress(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *shippingAddresResponseMapper) mapResponseShippingAddress(address *pb_shipping_address.ShippingResponse) *model.ShippingResponse {
	return &model.ShippingResponse{
		ID:             int32(address.Id),
		OrderID:        int32(address.OrderId),
		Alamat:         address.Alamat,
		Provinsi:       address.Provinsi,
		Negara:         address.Negara,
		Kota:           address.Kota,
		ShippingMethod: address.ShippingMethod,
		ShippingCost:   int32(address.ShippingCost),
		CreatedAt:      address.CreatedAt,
		UpdatedAt:      address.UpdatedAt,
	}
}

func (s *shippingAddresResponseMapper) mapResponsesShippingAddress(addresses []*pb_shipping_address.ShippingResponse) []*model.ShippingResponse {
	var mappedAddresses []*model.ShippingResponse
	for _, address := range addresses {
		mappedAddresses = append(mappedAddresses, s.mapResponseShippingAddress(address))
	}
	return mappedAddresses
}

func (s *shippingAddresResponseMapper) mapResponseShippingAddressDeleteAt(address *pb_shipping_address.ShippingResponseDeleteAt) *model.ShippingResponseDeleteAt {
	var deletedAt *string
	if address.DeletedAt != nil {
		deletedAt = &address.DeletedAt.Value
	}

	return &model.ShippingResponseDeleteAt{
		ID:             int32(address.Id),
		OrderID:        int32(address.OrderId),
		Alamat:         address.Alamat,
		Provinsi:       address.Provinsi,
		Negara:         address.Negara,
		Kota:           address.Kota,
		ShippingMethod: address.ShippingMethod,
		ShippingCost:   int32(address.ShippingCost),
		CreatedAt:      address.CreatedAt,
		UpdatedAt:      address.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}

func (s *shippingAddresResponseMapper) mapResponsesShippingAddressDeleteAt(addresses []*pb_shipping_address.ShippingResponseDeleteAt) []*model.ShippingResponseDeleteAt {
	var mappedAddresses []*model.ShippingResponseDeleteAt
	for _, address := range addresses {
		mappedAddresses = append(mappedAddresses, s.mapResponseShippingAddressDeleteAt(address))
	}
	return mappedAddresses
}
