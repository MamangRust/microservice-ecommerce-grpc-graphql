package review_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type ReviewQueryCache interface {
	GetReviewAllCache(ctx context.Context, req *model.FindAllReviewRequest) (*model.APIResponsePaginationReview, bool)
	SetReviewAllCache(ctx context.Context, req *model.FindAllReviewRequest, data *model.APIResponsePaginationReview)

	GetReviewByProductCache(ctx context.Context, req *model.FindAllReviewProductRequest) (*model.APIResponsePaginationReviewRelationDetail, bool)
	SetReviewByProductCache(ctx context.Context, req *model.FindAllReviewProductRequest, data *model.APIResponsePaginationReviewRelationDetail)

	GetReviewByMerchantCache(ctx context.Context, req *model.FindAllReviewMerchantRequest) (*model.APIResponsePaginationReviewRelationDetail, bool)
	SetReviewByMerchantCache(ctx context.Context, req *model.FindAllReviewMerchantRequest, data *model.APIResponsePaginationReviewRelationDetail)

	GetReviewActiveCache(ctx context.Context, req *model.FindAllReviewRequest) (*model.APIResponsePaginationReviewDeleteAt, bool)
	SetReviewActiveCache(ctx context.Context, req *model.FindAllReviewRequest, data *model.APIResponsePaginationReviewDeleteAt)

	GetReviewTrashedCache(ctx context.Context, req *model.FindAllReviewRequest) (*model.APIResponsePaginationReviewDeleteAt, bool)
	SetReviewTrashedCache(ctx context.Context, req *model.FindAllReviewRequest, data *model.APIResponsePaginationReviewDeleteAt)

	GetReviewByIdCache(ctx context.Context, id int) (*model.APIResponseReview, bool)
	SetReviewByIdCache(ctx context.Context, data *model.APIResponseReview)
}

type ReviewCommandCache interface {
	DeleteReviewCache(ctx context.Context, review_id int)
}
