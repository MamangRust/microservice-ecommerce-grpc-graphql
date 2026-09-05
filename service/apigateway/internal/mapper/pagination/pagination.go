package pagination

import (
	pbcommon "github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

func MapPaginationMeta(meta *pbcommon.PaginationMeta) *model.PaginationMeta {
	if meta == nil {
		return nil
	}
	return &model.PaginationMeta{
		CurrentPage:  meta.CurrentPage,
		TotalPages:   meta.TotalPages,
		PageSize:     meta.PageSize,
		TotalRecords: meta.TotalRecords,
	}
}
