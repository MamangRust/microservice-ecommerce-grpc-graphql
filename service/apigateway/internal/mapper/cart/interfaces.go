package cartgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/cart"
)

type CartGraphqlMapper interface {
	ToGraphqlResponseCartDelete(res *pb_cart.ApiResponseCartDelete) *model.APIResponseCartDelete
	ToGraphqlResponseCartAll(res *pb_cart.ApiResponseCartAll) *model.APIResponseCartAll
	ToGraphqlResponseCart(res *pb_cart.ApiResponseCart) *model.APIResponseCart
	ToGraphqlResponsePaginationCart(res *pb_cart.ApiResponsePaginationCart) *model.APIResponsePaginationCart
}
