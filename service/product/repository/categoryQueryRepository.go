package repository

import (
	"context"

	dto "github.com/MamangRust/microservice-ecommerce-grpc-product/dto"
	category_errors "github.com/MamangRust/microservice-ecommerce-shared/errors/category_errors"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type categoryQueryRepository struct {
	client pb.CategoryQueryServiceClient
}

func NewCategoryQueryRepository(client pb.CategoryQueryServiceClient) *categoryQueryRepository {
	return &categoryQueryRepository{
		client: client,
	}
}

func (r *categoryQueryRepository) FindByID(ctx context.Context, category_id int) (*dto.GetCategoryByIDRow, error) {
	res, err := r.client.FindById(ctx, &pb.FindByIdCategoryRequest{Id: int32(category_id)})
	if err != nil {
		return nil, category_errors.ErrFindCategoryById.WithInternal(err)
	}

	return &dto.GetCategoryByIDRow{
		CategoryID:  res.Data.Id,
		Name:        res.Data.Name,
		Description: &res.Data.Description,
	}, nil
}
