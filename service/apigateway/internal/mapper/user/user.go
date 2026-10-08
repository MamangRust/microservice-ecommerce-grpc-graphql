package usergraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/user"
)

type userGraphqlMapper struct {
}

func NewUserGraphqlMapper() *userGraphqlMapper {
	return &userGraphqlMapper{}
}

func (u *userGraphqlMapper) ToGraphqlResponseUserDelete(res *pb_user.ApiResponseUserDelete) *model.APIResponseUserDelete {
	return &model.APIResponseUserDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (u *userGraphqlMapper) ToGraphqlResponseUserAll(res *pb_user.ApiResponseUserAll) *model.APIResponseUserAll {
	return &model.APIResponseUserAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (u *userGraphqlMapper) ToGraphqlResponseUser(res *pb_user.ApiResponseUser) *model.APIResponseUserResponse {
	return &model.APIResponseUserResponse{
		Status:  res.Status,
		Message: res.Message,
		Data:    u.mapUserResponse(res.Data),
	}
}

func (u *userGraphqlMapper) ToGraphqlResponseUserDeleteAt(res *pb_user.ApiResponseUserDeleteAt) *model.APIResponseUserResponseDeleteAt {
	return &model.APIResponseUserResponseDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    u.mapUserResponseDeleteAt(res.Data),
	}
}

func (u *userGraphqlMapper) ToGraphqlResponseUsers(res *pb_user.ApiResponsesUser) *model.APIResponsesUser {
	return &model.APIResponsesUser{
		Status:  res.Status,
		Message: res.Message,
		Data:    u.mapUserResponses(res.Data),
	}
}

func (u *userGraphqlMapper) ToGraphqlResponsePaginationUser(res *pb_user.ApiResponsePaginationUser) *model.APIResponsePaginationUser {
	return &model.APIResponsePaginationUser{
		Status:     res.Status,
		Message:    res.Message,
		Data:       u.mapUserResponses(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (u *userGraphqlMapper) ToGraphqlResponsePaginationUserDeleteAt(res *pb_user.ApiResponsePaginationUserDeleteAt) *model.APIResponsePaginationUserDeleteAt {
	return &model.APIResponsePaginationUserDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       u.mapUserResponsesDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (u *userGraphqlMapper) mapUserResponse(user *pb_user.UserResponse) *model.UserResponse {
	return &model.UserResponse{
		ID:        int32(user.Id),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u *userGraphqlMapper) mapUserResponses(users []*pb_user.UserResponse) []*model.UserResponse {
	var responses []*model.UserResponse

	for _, user := range users {
		responses = append(responses, u.mapUserResponse(user))
	}

	return responses
}

func (u *userGraphqlMapper) mapUserResponseDeleteAt(user *pb_user.UserResponseDeleteAt) *model.UserResponseDeleteAt {
	var deletedAt string

	if user.DeletedAt != nil {
		deletedAt = user.DeletedAt.Value
	}

	return &model.UserResponseDeleteAt{
		ID:        int32(user.Id),
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}

func (u *userGraphqlMapper) mapUserResponsesDeleteAt(users []*pb_user.UserResponseDeleteAt) []*model.UserResponseDeleteAt {
	var responses []*model.UserResponseDeleteAt

	for _, user := range users {
		responses = append(responses, u.mapUserResponseDeleteAt(user))
	}

	return responses
}
