package productgraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/product"
)

type productGraphqlMapper struct {
}

func NewProductGraphqlMapper() *productGraphqlMapper {
	return &productGraphqlMapper{}
}

func (p *productGraphqlMapper) ToGraphqlResponseProduct(res *pb_product.ApiResponseProduct) *model.APIResponseProduct {
	return &model.APIResponseProduct{
		Status:  res.Status,
		Message: res.Message,
		Data:    p.mapResponseProduct(res.Data),
	}
}

func (p *productGraphqlMapper) ToGraphqlResponsesProduct(res *pb_product.ApiResponsesProduct) *model.APIResponsesProduct {
	return &model.APIResponsesProduct{
		Status:  res.Status,
		Message: res.Message,
		Data:    p.mapResponsesProduct(res.Data),
	}
}

func (p *productGraphqlMapper) ToGraphqlResponseProductDeleteAt(res *pb_product.ApiResponseProductDeleteAt) *model.APIResponseProductDeleteAt {
	return &model.APIResponseProductDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    p.mapResponseProductDeleteAt(res.Data),
	}
}

func (p *productGraphqlMapper) ToGraphqlResponseProductDelete(res *pb_product.ApiResponseProductDelete) *model.APIResponseProductDelete {
	return &model.APIResponseProductDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (p *productGraphqlMapper) ToGraphqlResponseProductAll(res *pb_product.ApiResponseProductAll) *model.APIResponseProductAll {
	return &model.APIResponseProductAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (p *productGraphqlMapper) ToGraphqlResponsePaginationProduct(res *pb_product.ApiResponsePaginationProduct) *model.APIResponsePaginationProduct {
	return &model.APIResponsePaginationProduct{
		Status:     res.Status,
		Message:    res.Message,
		Data:       p.mapResponsesProduct(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (p *productGraphqlMapper) ToGraphqlResponsePaginationProductDeleteAt(res *pb_product.ApiResponsePaginationProductDeleteAt) *model.APIResponsePaginationProductDeleteAt {
	return &model.APIResponsePaginationProductDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       p.mapResponsesProductDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (p *productGraphqlMapper) mapResponseProduct(product *pb_product.ProductResponse) *model.ProductResponse {
	id := int32(product.Id)
	merchantID := int32(product.MerchantId)
	categoryID := int32(product.CategoryId)
	price := int32(product.Price)
	countInStock := int32(product.CountInStock)
	weight := int32(product.Weight)

	return &model.ProductResponse{
		ID:           id,
		MerchantID:   merchantID,
		CategoryID:   categoryID,
		Name:         product.Name,
		Description:  &product.Description,
		Price:        price,
		CountInStock: countInStock,
		Brand:        product.Brand,
		Weight:       weight,
		SlugProduct:  product.SlugProduct,
		ImageProduct: product.ImageProduct,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
	}
}

func (p *productGraphqlMapper) mapResponsesProduct(products []*pb_product.ProductResponse) []*model.ProductResponse {
	var mapped []*model.ProductResponse
	for _, product := range products {
		mapped = append(mapped, p.mapResponseProduct(product))
	}
	return mapped
}

func (p *productGraphqlMapper) mapResponseProductDeleteAt(product *pb_product.ProductResponseDeleteAt) *model.ProductResponseDeleteAt {
	id := int32(product.Id)
	merchantID := int32(product.MerchantId)
	categoryID := int32(product.CategoryId)
	price := int32(product.Price)
	countInStock := int32(product.CountInStock)
	weight := int32(product.Weight)

	var deletedAt string
	if product.DeletedAt != nil {
		deletedAt = product.DeletedAt.Value
	}

	return &model.ProductResponseDeleteAt{
		ID:           id,
		MerchantID:   merchantID,
		CategoryID:   categoryID,
		Name:         product.Name,
		Description:  &product.Description,
		Price:        price,
		CountInStock: countInStock,
		Brand:        product.Brand,
		Weight:       weight,
		SlugProduct:  product.SlugProduct,
		ImageProduct: product.ImageProduct,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}

func (p *productGraphqlMapper) mapResponsesProductDeleteAt(products []*pb_product.ProductResponseDeleteAt) []*model.ProductResponseDeleteAt {
	var mapped []*model.ProductResponseDeleteAt
	for _, product := range products {
		mapped = append(mapped, p.mapResponseProductDeleteAt(product))
	}
	return mapped
}
