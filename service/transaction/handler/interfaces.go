package handler

import "github.com/MamangRust/microservice-ecommerce-shared/pb"

type TransactionQueryHandler interface {
	pb.TransactionQueryServiceServer
}

type TransactionCommandHandler interface {
	pb.TransactionCommandServiceServer
}
