package repository

import (
	"context"

	dto "github.com/MamangRust/microservice-ecommerce-grpc-cart/dto"
	"github.com/MamangRust/microservice-ecommerce-shared/errors/product_errors"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type productQueryRepository struct {
	client pb.ProductQueryServiceClient
}

func NewProductQueryRepository(client pb.ProductQueryServiceClient) ProductQueryRepository {
	return &productQueryRepository{
		client: client,
	}
}

func (r *productQueryRepository) FindById(ctx context.Context, id int) (*dto.GetProductByIDRow, error) {
	res, err := r.client.FindById(ctx, &pb.FindByIdProductRequest{Id: int32(id)})
	if err != nil {
		return nil, product_errors.ErrProductNotFound.WithInternal(err)
	}

	rating := float64(res.Data.Rating)

	return &dto.GetProductByIDRow{
		ProductID:    res.Data.Id,
		MerchantID:   res.Data.MerchantId,
		CategoryID:   res.Data.CategoryId,
		Name:         res.Data.Name,
		Description:  &res.Data.Description,
		Price:        res.Data.Price,
		CountInStock: res.Data.CountInStock,
		Brand:        &res.Data.Brand,
		Weight:       &res.Data.Weight,
		Rating:       &rating,
		SlugProduct:  &res.Data.SlugProduct,
		ImageProduct: &res.Data.ImageProduct,
	}, nil
}
