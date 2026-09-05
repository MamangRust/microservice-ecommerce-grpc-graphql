package handler

import "github.com/MamangRust/microservice-ecommerce-shared/pb"

type ReviewHandleGrpc interface {
	pb.ReviewQueryServiceServer
	pb.ReviewCommandServiceServer
}

type ReviewQueryHandler interface {
	pb.ReviewQueryServiceServer
}

type ReviewCommandHandler interface {
	pb.ReviewCommandServiceServer
}
