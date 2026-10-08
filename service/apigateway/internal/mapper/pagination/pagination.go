package pagination

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/common"
)

func MapPaginationMeta(meta *pb_common.PaginationMeta) *model.PaginationMeta {
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
