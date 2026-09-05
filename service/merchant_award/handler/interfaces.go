package handler

import "github.com/MamangRust/microservice-ecommerce-shared/pb"

type MerchantAwardQueryHandler interface {
	pb.MerchantAwardQueryServiceServer
}

type MerchantAwardCommandHandler interface {
	pb.MerchantAwardCommandServiceServer
}
