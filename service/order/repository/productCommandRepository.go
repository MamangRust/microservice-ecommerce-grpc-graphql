package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"

	dto "github.com/MamangRust/microservice-ecommerce-grpc-order/dto"
	product_errors "github.com/MamangRust/microservice-ecommerce-shared/errors/product_errors"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type productCommandRepository struct {
	client pb.ProductCommandServiceClient
}

func NewProductCommandRepository(client pb.ProductCommandServiceClient) *productCommandRepository {
	return &productCommandRepository{
		client: client,
	}
}

func (r *productCommandRepository) UpdateProductCountStock(ctx context.Context, productID int, stock int) (*dto.UpdateProductCountStockRow, error) {
	res, err := r.client.UpdateProductCountStock(ctx, &pb.UpdateProductCountStockRequest{
		ProductId: int32(productID),
		Stock:     int32(stock),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, product_errors.ErrProductNotFound
		}
		return nil, product_errors.ErrProductInternal.WithInternal(err)
	}

	return &dto.UpdateProductCountStockRow{
		ProductID:    res.Data.Id,
		CountInStock: res.Data.CountInStock,
	}, nil
}

func (r *productCommandRepository) AdjustProductStock(ctx context.Context, productID int, delta int, operationID string) (*dto.AdjustProductStockRow, error) {
	res, err := r.client.AdjustProductStock(ctx, &pb.AdjustProductStockRequest{
		ProductId:   int32(productID),
		Delta:       int32(delta),
		OperationId: operationID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, product_errors.ErrProductNotFound
		}
		return nil, product_errors.ErrProductInternal.WithInternal(err)
	}

	return &dto.AdjustProductStockRow{
		ProductID:    res.Data.Id,
		CountInStock: res.Data.CountInStock,
	}, nil
}
