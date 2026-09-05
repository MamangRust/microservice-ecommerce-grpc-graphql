package reviewdetail_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type ReviewDetailQueryCache interface {
	GetReviewDetailAllCache(ctx context.Context, req *model.FindAllReviewDetailInput) (*model.APIResponsePaginationReviewDetails, bool)
	SetReviewDetailAllCache(ctx context.Context, req *model.FindAllReviewDetailInput, data *model.APIResponsePaginationReviewDetails)

	GetReviewDetailActiveCache(ctx context.Context, req *model.FindAllReviewDetailInput) (*model.APIResponsePaginationReviewDetailsDeleteAt, bool)
	SetReviewDetailActiveCache(ctx context.Context, req *model.FindAllReviewDetailInput, data *model.APIResponsePaginationReviewDetailsDeleteAt)

	GetReviewDetailTrashedCache(ctx context.Context, req *model.FindAllReviewDetailInput) (*model.APIResponsePaginationReviewDetailsDeleteAt, bool)
	SetReviewDetailTrashedCache(ctx context.Context, req *model.FindAllReviewDetailInput, data *model.APIResponsePaginationReviewDetailsDeleteAt)

	GetCachedReviewDetailCache(ctx context.Context, reviewID int) (*model.APIResponseReviewDetail, bool)
	SetCachedReviewDetailCache(ctx context.Context, data *model.APIResponseReviewDetail)

	GetCachedReviewDetailTrashedCache(ctx context.Context, reviewID int) (*model.APIResponseReviewDetailDeleteAt, bool)
	SetCachedReviewDetailTrashedCache(ctx context.Context, data *model.APIResponseReviewDetailDeleteAt)
}

type ReviewDetailCommandCache interface {
	DeleteReviewDetailCache(ctx context.Context, reviewID int)
}
