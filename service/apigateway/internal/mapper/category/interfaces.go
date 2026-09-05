package categorygraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type CategoryGraphqlMapper interface {
	ToGraphqlResponseCategory(res *pb.ApiResponseCategory) *model.APIResponseCategory
	ToGraphqlResponseCategoryDeleteAt(res *pb.ApiResponseCategoryDeleteAt) *model.APIResponseCategoryDeleteAt
	ToGraphqlResponseCategoryDelete(res *pb.ApiResponseCategoryDelete) *model.APIResponseCategoryDelete
	ToGraphqlResponseCategoryAll(res *pb.ApiResponseCategoryAll) *model.APIResponseCategoryAll
	ToGraphqlResponsePaginationCategoryDeleteAt(res *pb.ApiResponsePaginationCategoryDeleteAt) *model.APIResponsePaginationCategoryDeleteAt
	ToGraphqlResponsePaginationCategory(res *pb.ApiResponsePaginationCategory) *model.APIResponsePaginationCategory
	ToGraphqlCategoryMonthlyPrice(res *pb.ApiResponseCategoryMonthPrice) *model.APIResponseCategoryMonthPrice
	ToGraphqlCategoryYearlyPrice(res *pb.ApiResponseCategoryYearPrice) *model.APIResponseCategoryYearPrice
	ToGraphqlMonthlyTotalPrice(res *pb.ApiResponseCategoryMonthlyTotalPrice) *model.APIResponseCategoryMonthlyTotalPrice
	ToGraphqlYearlyTotalPrice(res *pb.ApiResponseCategoryYearlyTotalPrice) *model.APIResponseCategoryYearlyTotalPrice
}
