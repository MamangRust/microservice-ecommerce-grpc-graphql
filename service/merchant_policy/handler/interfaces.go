package handler

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type MerchantPolicyQueryHandler interface {
	pb.MerchantPolicyQueryServiceServer
}

type MerchantPolicyCommandHandler interface {
	pb.MerchantPolicyCommandServiceServer
}
