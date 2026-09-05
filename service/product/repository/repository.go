package repository

import (
	db "github.com/MamangRust/microservice-ecommerce-grpc-product/database/schema"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type Repositories struct {
	ProductQuery   ProductQueryRepository
	ProductCommand ProductCommandRepository
	CategoryQuery  CategoryQueryRepository
	MerchantQuery  MerchantQueryRepository
}

func NewRepositories(DB *db.Queries, categoryQueryClient pb.CategoryQueryServiceClient, merchantQueryClient pb.MerchantQueryServiceClient) *Repositories {
	return &Repositories{
		ProductQuery:   NewProductQueryRepository(DB, categoryQueryClient),
		ProductCommand: NewProductCommandRepository(DB),
		CategoryQuery:  NewCategoryQueryRepository(categoryQueryClient),
		MerchantQuery:  NewMerchantQueryRepository(merchantQueryClient),
	}
}
