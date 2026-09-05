package handler

import "github.com/MamangRust/microservice-ecommerce-shared/pb"

type ProductQueryHandler interface {
	pb.ProductQueryServiceServer
}

type ProductCommandHandler interface {
	pb.ProductCommandServiceServer
}
