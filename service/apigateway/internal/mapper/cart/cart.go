package cartgraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/cart"
)

type cartResponseMapper struct{}

func NewCartResponseMapper() *cartResponseMapper {
	return &cartResponseMapper{}
}

func (c *cartResponseMapper) ToGraphqlResponseCartDelete(res *pb_cart.ApiResponseCartDelete) *model.APIResponseCartDelete {
	return &model.APIResponseCartDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (c *cartResponseMapper) ToGraphqlResponseCartAll(res *pb_cart.ApiResponseCartAll) *model.APIResponseCartAll {
	return &model.APIResponseCartAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (c *cartResponseMapper) ToGraphqlResponseCart(res *pb_cart.ApiResponseCart) *model.APIResponseCart {
	return &model.APIResponseCart{
		Status:  res.Status,
		Message: res.Message,
		Data:    c.mapResponseCart(res.Data),
	}
}

func (c *cartResponseMapper) ToGraphqlResponsePaginationCart(res *pb_cart.ApiResponsePaginationCart) *model.APIResponsePaginationCart {
	return &model.APIResponsePaginationCart{
		Status:     res.Status,
		Message:    res.Message,
		Data:       c.mapResponsesCart(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (c *cartResponseMapper) mapResponseCart(cart *pb_cart.CartResponse) *model.CartResponse {
	return &model.CartResponse{
		ID:        int32(cart.Id),
		UserID:    int32(cart.UserId),
		ProductID: int32(cart.ProductId),
		Name:      cart.Name,
		Price:     int32(cart.Price),
		Image:     cart.Image,
		Quantity:  int32(cart.Quantity),
		Weight:    int32(cart.Weight),
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}

func (c *cartResponseMapper) mapResponsesCart(carts []*pb_cart.CartResponse) []*model.CartResponse {
	var responses []*model.CartResponse

	for _, cart := range carts {
		responses = append(responses, c.mapResponseCart(cart))
	}

	return responses
}
