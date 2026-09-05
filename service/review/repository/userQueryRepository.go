package repository

import (
	"context"

	dto "github.com/MamangRust/microservice-ecommerce-grpc-review/dto"
	user_errors "github.com/MamangRust/microservice-ecommerce-shared/errors/user_errors"
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
		return nil, user_errors.ErrUserInternal.WithInternal(err)
	}

	return &dto.GetUserByIDRow{
		UserID:    res.Data.Id,
		Firstname: "", // UserResponse does not provide Firstname
		Lastname:  "", // UserResponse does not provide Lastname
		Email:     res.Data.Email,
	}, nil
}
