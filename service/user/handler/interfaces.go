package handler

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type UserQueryHandler interface {
	pb.UserQueryServiceServer
}

type UserCommandHandler interface {
	pb.UserCommandServiceServer
}
