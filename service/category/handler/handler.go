package handler

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-category/service"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type Deps struct {
	Service *service.Service
	Logger  logger.LoggerInterface
}

// F5: legacy OLTP category stats handlers were removed; stats are served by
// service/stats_reader from ClickHouse.
type Handler struct {
	CategoryQuery   pb.CategoryQueryServiceServer
	CategoryCommand pb.CategoryCommandServiceServer
}

func NewHandler(deps *Deps) *Handler {
	return &Handler{
		CategoryQuery:   NewCategoryQueryHandler(deps.Service.CategoryQuery, deps.Logger),
		CategoryCommand: NewCategoryCommandHandler(deps.Service.CategoryCommand, deps.Logger),
	}
}
