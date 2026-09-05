package categorygraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type categoryGraphqlMapper struct{}

func NewCategoryGraphqlMapper() *categoryGraphqlMapper {
	return &categoryGraphqlMapper{}
}

func (m *categoryGraphqlMapper) ToGraphqlResponseCategory(res *pb.ApiResponseCategory) *model.APIResponseCategory {
	return &model.APIResponseCategory{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseCategory(res.Data),
	}
}

func (m *categoryGraphqlMapper) ToGraphqlResponseCategoryDeleteAt(res *pb.ApiResponseCategoryDeleteAt) *model.APIResponseCategoryDeleteAt {
	return &model.APIResponseCategoryDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.mapResponseCategoryDeleteAt(res.Data),
	}
}

func (m *categoryGraphqlMapper) ToGraphqlResponseCategoryDelete(res *pb.ApiResponseCategoryDelete) *model.APIResponseCategoryDelete {
	return &model.APIResponseCategoryDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *categoryGraphqlMapper) ToGraphqlResponseCategoryAll(res *pb.ApiResponseCategoryAll) *model.APIResponseCategoryAll {
	return &model.APIResponseCategoryAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (m *categoryGraphqlMapper) ToGraphqlResponsePaginationCategoryDeleteAt(
	res *pb.ApiResponsePaginationCategoryDeleteAt,
) *model.APIResponsePaginationCategoryDeleteAt {
	return &model.APIResponsePaginationCategoryDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesCategoryDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *categoryGraphqlMapper) ToGraphqlResponsePaginationCategory(
	res *pb.ApiResponsePaginationCategory,
) *model.APIResponsePaginationCategory {
	return &model.APIResponsePaginationCategory{
		Status:     res.Status,
		Message:    res.Message,
		Data:       m.mapResponsesCategory(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (m *categoryGraphqlMapper) ToGraphqlCategoryMonthlyPrice(res *pb.ApiResponseCategoryMonthPrice) *model.APIResponseCategoryMonthPrice {
	return &model.APIResponseCategoryMonthPrice{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.MapResponsesCategoryMonthlyPrices(res.Data),
	}
}

func (m *categoryGraphqlMapper) ToGraphqlCategoryYearlyPrice(res *pb.ApiResponseCategoryYearPrice) *model.APIResponseCategoryYearPrice {
	return &model.APIResponseCategoryYearPrice{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.MapResponsesCategoryYearlyPrices(res.Data),
	}
}

func (m *categoryGraphqlMapper) ToGraphqlMonthlyTotalPrice(res *pb.ApiResponseCategoryMonthlyTotalPrice) *model.APIResponseCategoryMonthlyTotalPrice {
	return &model.APIResponseCategoryMonthlyTotalPrice{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.MapResponseCategoryMonthlyTotalPrices(res.Data),
	}
}

func (m *categoryGraphqlMapper) ToGraphqlYearlyTotalPrice(res *pb.ApiResponseCategoryYearlyTotalPrice) *model.APIResponseCategoryYearlyTotalPrice {
	return &model.APIResponseCategoryYearlyTotalPrice{
		Status:  res.Status,
		Message: res.Message,
		Data:    m.MapResponseCategoryYearlyTotalPrices(res.Data),
	}
}

func (c *categoryGraphqlMapper) mapResponseCategory(category *pb.CategoryResponse) *model.CategoryResponse {
	return &model.CategoryResponse{
		ID:            int32(category.Id),
		Name:          category.Name,
		Description:   category.Description,
		SlugCategory:  category.SlugCategory,
		ImageCategory: category.SlugCategory,
		CreatedAt:     category.CreatedAt,
		UpdatedAt:     category.UpdatedAt,
	}
}

func (c *categoryGraphqlMapper) mapResponsesCategory(categories []*pb.CategoryResponse) []*model.CategoryResponse {
	var responses []*model.CategoryResponse

	for _, category := range categories {
		responses = append(responses, c.mapResponseCategory(category))
	}

	return responses
}

func (c *categoryGraphqlMapper) mapResponseCategoryDeleteAt(category *pb.CategoryResponseDeleteAt) *model.CategoryResponseDeleteAt {
	var deletedAt string

	if category.DeletedAt != nil {
		deletedAt = category.DeletedAt.Value
	}

	return &model.CategoryResponseDeleteAt{
		ID:            int32(category.Id),
		Name:          category.Name,
		Description:   category.Description,
		SlugCategory:  category.SlugCategory,
		ImageCategory: category.SlugCategory,
		CreatedAt:     category.CreatedAt,
		UpdatedAt:     category.UpdatedAt,
		DeletedAt:     &deletedAt,
	}
}

func (c *categoryGraphqlMapper) mapResponsesCategoryDeleteAt(categories []*pb.CategoryResponseDeleteAt) []*model.CategoryResponseDeleteAt {
	var responses []*model.CategoryResponseDeleteAt

	for _, category := range categories {
		responses = append(responses, c.mapResponseCategoryDeleteAt(category))
	}

	return responses
}

func (m *categoryGraphqlMapper) MapResponseCategoryMonthlyPrice(category *pb.CategoryMonthPriceResponse) *model.CategoryMonthPriceResponse {
	if category == nil {
		return nil
	}
	return &model.CategoryMonthPriceResponse{
		Month:        category.Month,
		CategoryID:   int32(category.CategoryId),
		CategoryName: category.CategoryName,
		OrderCount:   int32(category.OrderCount),
		ItemsSold:    int32(category.ItemsSold),
		TotalRevenue: int32(category.TotalRevenue),
	}
}

func (m *categoryGraphqlMapper) MapResponsesCategoryMonthlyPrices(categories []*pb.CategoryMonthPriceResponse) []*model.CategoryMonthPriceResponse {
	var records []*model.CategoryMonthPriceResponse
	for _, c := range categories {
		records = append(records, m.MapResponseCategoryMonthlyPrice(c))
	}
	return records
}

func (m *categoryGraphqlMapper) MapResponseCategoryYearlyPrice(category *pb.CategoryYearPriceResponse) *model.CategoryYearPriceResponse {
	if category == nil {
		return nil
	}
	return &model.CategoryYearPriceResponse{
		Year:               category.Year,
		CategoryID:         int32(category.CategoryId),
		CategoryName:       category.CategoryName,
		OrderCount:         int32(category.OrderCount),
		ItemsSold:          int32(category.ItemsSold),
		TotalRevenue:       int32(category.TotalRevenue),
		UniqueProductsSold: int32(category.UniqueProductsSold),
	}
}

func (m *categoryGraphqlMapper) MapResponsesCategoryYearlyPrices(categories []*pb.CategoryYearPriceResponse) []*model.CategoryYearPriceResponse {
	var records []*model.CategoryYearPriceResponse
	for _, c := range categories {
		records = append(records, m.MapResponseCategoryYearlyPrice(c))
	}
	return records
}

func (m *categoryGraphqlMapper) MapResponseCategoryMonthlyTotalPrice(c *pb.CategoriesMonthlyTotalPriceResponse) *model.CategoryMonthlyTotalPriceResponse {
	if c == nil {
		return nil
	}
	return &model.CategoryMonthlyTotalPriceResponse{
		Year:         c.Year,
		Month:        c.Month,
		TotalRevenue: int32(c.TotalRevenue),
	}
}

func (m *categoryGraphqlMapper) MapResponseCategoryMonthlyTotalPrices(categories []*pb.CategoriesMonthlyTotalPriceResponse) []*model.CategoryMonthlyTotalPriceResponse {
	var records []*model.CategoryMonthlyTotalPriceResponse
	for _, c := range categories {
		records = append(records, m.MapResponseCategoryMonthlyTotalPrice(c))
	}
	return records
}

func (m *categoryGraphqlMapper) MapResponseCategoryYearlyTotalPrice(c *pb.CategoriesYearlyTotalPriceResponse) *model.CategoryYearlyTotalPriceResponse {
	if c == nil {
		return nil
	}
	return &model.CategoryYearlyTotalPriceResponse{
		Year:         c.Year,
		TotalRevenue: int32(c.TotalRevenue),
	}
}

func (m *categoryGraphqlMapper) MapResponseCategoryYearlyTotalPrices(categories []*pb.CategoriesYearlyTotalPriceResponse) []*model.CategoryYearlyTotalPriceResponse {
	var records []*model.CategoryYearlyTotalPriceResponse
	for _, c := range categories {
		records = append(records, m.MapResponseCategoryYearlyTotalPrice(c))
	}
	return records
}
