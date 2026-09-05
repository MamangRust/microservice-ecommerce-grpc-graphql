package rolegraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type roleGraphqlMapper struct {
}

func NewRoleGraphqlMapper() *roleGraphqlMapper {
	return &roleGraphqlMapper{}
}

func (s *roleGraphqlMapper) ToGraphqlResponseAll(res *pb.ApiResponseRoleAll) *model.APIResponseRoleAll {
	return &model.APIResponseRoleAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponseDelete(res *pb.ApiResponseRoleDelete) *model.APIResponseRoleDelete {
	return &model.APIResponseRoleDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponseRole(res *pb.ApiResponseRole) *model.APIResponseRole {
	return &model.APIResponseRole{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponseRole(res.Data),
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponseRoleDeleteAt(res *pb.ApiResponseRole) *model.APIResponseRoleDeleteAt {
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

func (s *roleGraphqlMapper) ToGraphqlResponsesRole(res *pb.ApiResponsesRole) *model.APIResponsesRole {
	return &model.APIResponsesRole{
		Status:  res.Status,
		Message: res.Message,
		Data:    s.mapResponsesRole(res.Data),
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponsePaginationRole(res *pb.ApiResponsePaginationRole) *model.APIResponsePaginationRole {
	return &model.APIResponsePaginationRole{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesRole(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *roleGraphqlMapper) ToGraphqlResponsePaginationRoleDeleteAt(res *pb.ApiResponsePaginationRoleDeleteAt) *model.APIResponsePaginationRoleDeleteAt {
	return &model.APIResponsePaginationRoleDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       s.mapResponsesRoleDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (s *roleGraphqlMapper) mapResponseRole(role *pb.RoleResponse) *model.RoleResponse {
	return &model.RoleResponse{
		ID:        int32(role.Id),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (s *roleGraphqlMapper) mapResponsesRole(roles []*pb.RoleResponse) []*model.RoleResponse {
	var responseRoles []*model.RoleResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRole(role))
	}

	return responseRoles
}

func (s *roleGraphqlMapper) mapResponseRoleDeleteAt(role *pb.RoleResponseDeleteAt) *model.RoleResponseDeleteAt {
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

func (s *roleGraphqlMapper) mapResponsesRoleDeleteAt(roles []*pb.RoleResponseDeleteAt) []*model.RoleResponseDeleteAt {
	var responseRoles []*model.RoleResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRoleDeleteAt(role))
	}

	return responseRoles
}
