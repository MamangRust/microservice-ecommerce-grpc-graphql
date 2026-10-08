package usergraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/user"
)

type UserGraphqlMapper interface {
	ToGraphqlResponseUser(res *pb_user.ApiResponseUser) *model.APIResponseUserResponse
	ToGraphqlResponseUserDeleteAt(res *pb_user.ApiResponseUserDeleteAt) *model.APIResponseUserResponseDeleteAt
	ToGraphqlResponseUsers(res *pb_user.ApiResponsesUser) *model.APIResponsesUser
	ToGraphqlResponseUserDelete(res *pb_user.ApiResponseUserDelete) *model.APIResponseUserDelete
	ToGraphqlResponseUserAll(res *pb_user.ApiResponseUserAll) *model.APIResponseUserAll
	ToGraphqlResponsePaginationUser(res *pb_user.ApiResponsePaginationUser) *model.APIResponsePaginationUser
	ToGraphqlResponsePaginationUserDeleteAt(res *pb_user.ApiResponsePaginationUserDeleteAt) *model.APIResponsePaginationUserDeleteAt
}
