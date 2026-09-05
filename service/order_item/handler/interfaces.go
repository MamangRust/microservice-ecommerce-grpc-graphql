package handler

import "github.com/MamangRust/microservice-ecommerce-shared/pb"

type OrderItemQueryHandler interface {
	pb.OrderItemQueryServiceServer
}

type OrderItemCommandHandler interface {
	pb.OrderItemCommandServiceServer
}
