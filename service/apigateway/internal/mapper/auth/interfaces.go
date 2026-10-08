package authgraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/auth"
)

type AuthGraphqlMapper interface {
	ToGraphqlVerifyCode(res *pb_auth.ApiResponseVerifyCode) *model.APIResponseVerifyCode
	ToGraphqlForgotPassword(res *pb_auth.ApiResponseForgotPassword) *model.APIResponseForgotPassword
	ToGraphqlResetPassword(res *pb_auth.ApiResponseResetPassword) *model.APIResponseResetPassword
	ToGraphqlResponseLogin(res *pb_auth.ApiResponseLogin) *model.APIResponseLogin
	ToGraphqlResponseRegister(res *pb_auth.ApiResponseRegister) *model.APIResponseRegister
	ToGraphqlResponseRefreshToken(res *pb_auth.ApiResponseRefreshToken) *model.APIResponseRefreshToken
	ToGraphqlResponseGetMe(res *pb_auth.ApiResponseGetMe) *model.APIResponseGetMe
}
