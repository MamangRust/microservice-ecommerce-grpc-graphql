package product_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type ProductQueryCache interface {
	GetCachedProducts(ctx context.Context, req *model.FindAllProductInput) (*model.APIResponsePaginationProduct, bool)
	SetCachedProducts(ctx context.Context, req *model.FindAllProductInput, data *model.APIResponsePaginationProduct)

	GetCachedProductsByMerchant(ctx context.Context, req *model.FindAllProductMerchantInput) (*model.APIResponsePaginationProduct, bool)
	SetCachedProductsByMerchant(ctx context.Context, req *model.FindAllProductMerchantInput, data *model.APIResponsePaginationProduct)

	GetCachedProductsByCategory(ctx context.Context, req *model.FindAllProductCategoryInput) (*model.APIResponsePaginationProduct, bool)
	SetCachedProductsByCategory(ctx context.Context, req *model.FindAllProductCategoryInput, data *model.APIResponsePaginationProduct)

	GetCachedProductActive(ctx context.Context, req *model.FindAllProductInput) (*model.APIResponsePaginationProductDeleteAt, bool)
	SetCachedProductActive(ctx context.Context, req *model.FindAllProductInput, data *model.APIResponsePaginationProductDeleteAt)

	GetCachedProductTrashed(ctx context.Context, req *model.FindAllProductInput) (*model.APIResponsePaginationProductDeleteAt, bool)
	SetCachedProductTrashed(ctx context.Context, req *model.FindAllProductInput, data *model.APIResponsePaginationProductDeleteAt)

	GetCachedProduct(ctx context.Context, productID int) (*model.APIResponseProduct, bool)
	SetCachedProduct(ctx context.Context, data *model.APIResponseProduct)
}

type ProductCommandCache interface {
	DeleteCachedProduct(ctx context.Context, productID int)
}
