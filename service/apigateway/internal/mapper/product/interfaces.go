package productgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/product"
)

type ProductGraphqlMapper interface {
	ToGraphqlResponseProduct(res *pb_product.ApiResponseProduct) *model.APIResponseProduct
	ToGraphqlResponsesProduct(res *pb_product.ApiResponsesProduct) *model.APIResponsesProduct
	ToGraphqlResponseProductDeleteAt(res *pb_product.ApiResponseProductDeleteAt) *model.APIResponseProductDeleteAt
	ToGraphqlResponseProductDelete(res *pb_product.ApiResponseProductDelete) *model.APIResponseProductDelete
	ToGraphqlResponseProductAll(res *pb_product.ApiResponseProductAll) *model.APIResponseProductAll
	ToGraphqlResponsePaginationProduct(res *pb_product.ApiResponsePaginationProduct) *model.APIResponsePaginationProduct
	ToGraphqlResponsePaginationProductDeleteAt(res *pb_product.ApiResponsePaginationProductDeleteAt) *model.APIResponsePaginationProductDeleteAt
}
