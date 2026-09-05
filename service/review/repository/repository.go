package repository

import (
	db "github.com/MamangRust/microservice-ecommerce-grpc-review/database/schema"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type Repositories struct {
	ProductQuery  ProductQueryRepository
	ReviewQuery   ReviewQueryRepository
	UserQuery     UserQueryRepository
	ReviewCommand ReviewCommandRepository
}

func NewRepositories(DB *db.Queries, userQueryClient pb.UserQueryServiceClient, productQueryClient pb.ProductQueryServiceClient) *Repositories {
	return &Repositories{
		ProductQuery:  NewProductQueryRepository(productQueryClient),
		ReviewQuery:   NewReviewQueryRepository(DB, productQueryClient),
		UserQuery:     NewUserQueryRepository(userQueryClient),
		ReviewCommand: NewReviewCommandRepository(DB),
	}
}
