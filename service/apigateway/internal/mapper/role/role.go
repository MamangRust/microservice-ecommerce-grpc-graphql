package rolegraphqlmapper

import (
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/role"
)

type roleGraphqlMapper struct {
}

func NewRoleGraphqlMapper() *roleGraphqlMapper {
	return &roleGraphqlMapper{}
}

func (s *roleGraphqlMapper) ToGraphqlResponseAll(res *pb_role.ApiResponseRoleAll) *model.APIResponseRoleAll {
	return &model.APIResponseRoleAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponseDelete(res *pb_role.ApiResponseRoleDelete) *model.APIResponseRoleDelete {
	return &model.APIResponseRoleDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponseRole(res *pb_role.ApiResponseRole) *model.APIResponseRole {
	return &model.APIResponseRole{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseRole(res.Data),
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponseRoleDeleteAt(res *pb_role.ApiResponseRole) *model.APIResponseRoleDeleteAt {
	if res == nil {
		return nil
	}
	var deletedAt *string
	var roleData *model.RoleResponseDeleteAt
	if res.Data != nil {
		roleData = &model.RoleResponseDeleteAt{
			ID:        int32(res.Data.Id),
			Name:      res.Data.Name,
			CreatedAt: res.Data.CreatedAt,
			UpdatedAt: res.Data.UpdatedAt,
			DeletedAt: deletedAt,
		}
	}
	return &model.APIResponseRoleDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    roleData,
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponsesRole(res *pb_role.ApiResponsesRole) *model.APIResponsesRole {
	return &model.APIResponsesRole{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponsesRole(res.Data),
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponsePaginationRole(res *pb_role.ApiResponsePaginationRole) *model.APIResponsePaginationRole {
	return &model.APIResponsePaginationRole{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesRole(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponsePaginationRoleDeleteAt(res *pb_role.ApiResponsePaginationRoleDeleteAt) *model.APIResponsePaginationRoleDeleteAt {
	return &model.APIResponsePaginationRoleDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesRoleDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *roleGraphqlMapper) mapResponseRole(role *pb_role.RoleResponse) *model.RoleResponse {
	return &model.RoleResponse{
		ID:        int32(role.Id),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (s *roleGraphqlMapper) mapResponsesRole(roles []*pb_role.RoleResponse) []*model.RoleResponse {
	var responseRoles []*model.RoleResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRole(role))
	}

	return responseRoles
}

func (s *roleGraphqlMapper) mapResponseRoleDeleteAt(role *pb_role.RoleResponseDeleteAt) *model.RoleResponseDeleteAt {
	var deletedAt string

	if role.DeletedAt != nil {
		deletedAt = role.DeletedAt.Value
	}

	return &model.RoleResponseDeleteAt{
		ID:        int32(role.Id),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}

func (s *roleGraphqlMapper) mapResponsesRoleDeleteAt(roles []*pb_role.RoleResponseDeleteAt) []*model.RoleResponseDeleteAt {
	var responseRoles []*model.RoleResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRoleDeleteAt(role))
	}

	return responseRoles
}
