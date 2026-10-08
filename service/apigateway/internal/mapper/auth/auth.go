package authgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/auth"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/user"
)

type authGraphqlMapper struct {
}

func NewAuthGraphqlMapper() *authGraphqlMapper {
	return &authGraphqlMapper{}
}

func (s *authGraphqlMapper) ToGraphqlVerifyCode(res *pb_auth.ApiResponseVerifyCode) *model.APIResponseVerifyCode {
	return &model.APIResponseVerifyCode{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *authGraphqlMapper) ToGraphqlForgotPassword(res *pb_auth.ApiResponseForgotPassword) *model.APIResponseForgotPassword {
	return &model.APIResponseForgotPassword{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *authGraphqlMapper) ToGraphqlResetPassword(res *pb_auth.ApiResponseResetPassword) *model.APIResponseResetPassword {
	return &model.APIResponseResetPassword{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *authGraphqlMapper) ToGraphqlResponseLogin(res *pb_auth.ApiResponseLogin) *model.APIResponseLogin {
	return &model.APIResponseLogin{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseToken(res.Data),
	}
}

func (s *authGraphqlMapper) ToGraphqlResponseRegister(res *pb_auth.ApiResponseRegister) *model.APIResponseRegister {
	return &model.APIResponseRegister{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseUser(res.Data),
	}
}

func (s *authGraphqlMapper) ToGraphqlResponseRefreshToken(res *pb_auth.ApiResponseRefreshToken) *model.APIResponseRefreshToken {
	return &model.APIResponseRefreshToken{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseToken(res.Data),
	}
}

func (s *authGraphqlMapper) ToGraphqlResponseGetMe(res *pb_auth.ApiResponseGetMe) *model.APIResponseGetMe {
	return &model.APIResponseGetMe{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseUser(res.Data),
	}
}

func (s *authGraphqlMapper) mapResponseUser(res *pb_user.UserResponse) *model.UserResponse {
	return &model.UserResponse{
		ID:        res.Id,
		Firstname: res.Firstname,
		Lastname:  res.Lastname,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
}

func (s *authGraphqlMapper) mapResponseToken(res *pb_auth.TokenResponse) *model.TokenResponse {
	return &model.TokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
}
