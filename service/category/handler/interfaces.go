package handler

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type CategoryQueryHandler interface {
	pb.CategoryQueryServiceServer
}

type CategoryCommandHandler interface {
	pb.CategoryCommandServiceServer
}
