package rolegraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/role"
)

type RoleGraphqlMapper interface {
	ToGraphqlResponseRole(res *pb_role.ApiResponseRole) *model.APIResponseRole
	ToGraphqlResponseRoleDeleteAt(res *pb_role.ApiResponseRole) *model.APIResponseRoleDeleteAt
	ToGraphqlResponsesRole(res *pb_role.ApiResponsesRole) *model.APIResponsesRole
	ToGraphqlResponseDelete(res *pb_role.ApiResponseRoleDelete) *model.APIResponseRoleDelete
	ToGraphqlResponseAll(res *pb_role.ApiResponseRoleAll) *model.APIResponseRoleAll
	ToGraphqlResponsePaginationRole(res *pb_role.ApiResponsePaginationRole) *model.APIResponsePaginationRole
	ToGraphqlResponsePaginationRoleDeleteAt(res *pb_role.ApiResponsePaginationRoleDeleteAt) *model.APIResponsePaginationRoleDeleteAt
}
