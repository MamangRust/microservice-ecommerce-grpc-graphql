package repository

import (
	"context"

	dto "github.com/MamangRust/microservice-ecommerce-grpc-transaction/dto"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type userQueryRepository struct {
	client pb.UserQueryServiceClient
}

func NewUserQueryRepository(client pb.UserQueryServiceClient) *userQueryRepository {
	return &userQueryRepository{
		client: client,
	}
}

func (r *userQueryRepository) FindByID(ctx context.Context, user_id int) (*dto.GetUserByIDRow, error) {
	res, err := r.client.FindById(ctx, &pb.FindByIdUserRequest{Id: int32(user_id)})
	if err != nil {
		// pertahankan status gRPC dari dependency service (NotFound -> 404, dst)
		return nil, err
	}

	return &dto.GetUserByIDRow{
		UserID:    res.Data.Id,
		Firstname: res.Data.Firstname,
		Lastname:  res.Data.Lastname,
		Email:     res.Data.Email,
	}, nil
}
