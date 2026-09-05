package rolegraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type RoleGraphqlMapper interface {
	ToGraphqlResponseRole(res *pb.ApiResponseRole) *model.APIResponseRole
	ToGraphqlResponseRoleDeleteAt(res *pb.ApiResponseRole) *model.APIResponseRoleDeleteAt
	ToGraphqlResponsesRole(res *pb.ApiResponsesRole) *model.APIResponsesRole
	ToGraphqlResponseDelete(res *pb.ApiResponseRoleDelete) *model.APIResponseRoleDelete
	ToGraphqlResponseAll(res *pb.ApiResponseRoleAll) *model.APIResponseRoleAll
	ToGraphqlResponsePaginationRole(res *pb.ApiResponsePaginationRole) *model.APIResponsePaginationRole
	ToGraphqlResponsePaginationRoleDeleteAt(res *pb.ApiResponsePaginationRoleDeleteAt) *model.APIResponsePaginationRoleDeleteAt
}
