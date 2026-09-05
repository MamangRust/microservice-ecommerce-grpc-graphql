package handler

import "github.com/MamangRust/microservice-ecommerce-shared/pb"

type CartQueryHandler interface {
	pb.CartQueryServiceServer
}

type CartCommandHandler interface {
	pb.CartCommandServiceServer
}
