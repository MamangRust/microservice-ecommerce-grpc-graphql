package apps

import (
	"github.com/MamangRust/microservice-ecommerce-grpc-review-detail/cache"
	db "github.com/MamangRust/microservice-ecommerce-grpc-review-detail/database/schema"
	"github.com/MamangRust/microservice-ecommerce-grpc-review-detail/handler"
	"github.com/MamangRust/microservice-ecommerce-grpc-review-detail/repository"
	"github.com/MamangRust/microservice-ecommerce-grpc-review-detail/service"
	"github.com/MamangRust/microservice-ecommerce-pkg/server"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"google.golang.org/grpc"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	queries := db.New(srv.Pool)

	repos := repository.NewRepositories(queries)

	obs, _ := observability.NewObservability("review_detail-server", srv.Logger)
	cache := cache.NewMencache(srv.CacheStore)

	svc := service.NewService(&service.Deps{
		Observability: obs,
		Cache:         cache,
		Repositories:  repos,
		Logger:        srv.Logger,
	})

	h := handler.NewHandler(&handler.Deps{Service: svc, Logger: srv.Logger})

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterReviewDetailQueryServiceServer(gs, h.ReviewDetailQuery)
		pb.RegisterReviewDetailCommandServiceServer(gs, h.ReviewDetailCommand)
	}

	return srv, nil
}
