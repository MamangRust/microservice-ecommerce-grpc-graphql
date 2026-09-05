package handler

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type OrderQueryHandler interface {
	pb.OrderQueryServiceServer
}

type OrderCommandHandler interface {
	pb.OrderCommandServiceServer
}
