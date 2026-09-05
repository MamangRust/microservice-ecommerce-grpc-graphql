package handler

import "github.com/MamangRust/microservice-ecommerce-shared/pb"

type BannerQueryHandler interface {
	pb.BannerQueryServiceServer
}

type BannerCommandHandler interface {
	pb.BannerCommandServiceServer
}

