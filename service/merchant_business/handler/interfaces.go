package handler

import "github.com/MamangRust/microservice-ecommerce-shared/pb"

type MerchantBusinessQueryHandler interface {
	pb.MerchantBusinessQueryServiceServer
}

type MerchantBusinessCommandHandler interface {
	pb.MerchantBusinessCommandServiceServer
}
