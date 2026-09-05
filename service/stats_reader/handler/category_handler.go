package handler

import (
	"context"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"go.uber.org/zap"
)

// CategoryStatsHandler serves CategoryStatsService, CategoryStatsByIdService
// and CategoryStatsByMerchantService from ClickHouse.
type CategoryStatsHandler struct {
	pb.UnimplementedCategoryStatsServiceServer
	pb.UnimplementedCategoryStatsByIdServiceServer
	pb.UnimplementedCategoryStatsByMerchantServiceServer
	repo repository.Repository
	log  logger.LoggerInterface
}

func NewCategoryStatsHandler(repo repository.Repository, log logger.LoggerInterface) *CategoryStatsHandler {
	return &CategoryStatsHandler{repo: repo, log: log}
}

// --- CategoryStatsService ---

func (h *CategoryStatsHandler) FindMonthlyTotalPrices(ctx context.Context, req *pb.FindYearMonthTotalPrices) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	data, err := h.repo.GetMonthlyTotalPricing(ctx, int(req.GetYear()), int(req.GetMonth()), "", 0)
	if err != nil {
		h.log.Error("FindMonthlyTotalPrices failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapMonthlyPricing(data),
	}, nil
}

func (h *CategoryStatsHandler) FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	data, err := h.repo.GetYearlyTotalPricing(ctx, int(req.GetYear()), "", 0)
	if err != nil {
		h.log.Error("FindYearlyTotalPrices failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapYearlyPricing(data),
	}, nil
}

func (h *CategoryStatsHandler) FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryMonthPrice, error) {
	data, err := h.repo.GetMonthlyCategoryStats(ctx, int(req.GetYear()), "", 0)
	if err != nil {
		h.log.Error("FindMonthPrice failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Monthly payment methods retrieved successfully",
		Data:    mapMonthlyCategory(data),
	}, nil
}

func (h *CategoryStatsHandler) FindYearPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryYearPrice, error) {
	data, err := h.repo.GetYearlyCategoryStats(ctx, int(req.GetYear()), "", 0)
	if err != nil {
		h.log.Error("FindYearPrice failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapYearlyCategory(data),
	}, nil
}

// --- CategoryStatsByIdService ---

func (h *CategoryStatsHandler) FindMonthlyTotalPricesById(ctx context.Context, req *pb.FindYearMonthTotalPriceById) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	data, err := h.repo.GetMonthlyTotalPricing(ctx, int(req.GetYear()), int(req.GetMonth()), "category_id", req.GetCategoryId())
	if err != nil {
		h.log.Error("FindMonthlyTotalPricesById failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapMonthlyPricing(data),
	}, nil
}

func (h *CategoryStatsHandler) FindYearlyTotalPricesById(ctx context.Context, req *pb.FindYearTotalPriceById) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	data, err := h.repo.GetYearlyTotalPricing(ctx, int(req.GetYear()), "category_id", req.GetCategoryId())
	if err != nil {
		h.log.Error("FindYearlyTotalPricesById failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapYearlyPricing(data),
	}, nil
}

func (h *CategoryStatsHandler) FindMonthPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryMonthPrice, error) {
	data, err := h.repo.GetMonthlyCategoryStats(ctx, int(req.GetYear()), "category_id", req.GetCategoryId())
	if err != nil {
		h.log.Error("FindMonthPriceById failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Monthly payment methods retrieved successfully",
		Data:    mapMonthlyCategory(data),
	}, nil
}

func (h *CategoryStatsHandler) FindYearPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryYearPrice, error) {
	data, err := h.repo.GetYearlyCategoryStats(ctx, int(req.GetYear()), "category_id", req.GetCategoryId())
	if err != nil {
		h.log.Error("FindYearPriceById failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapYearlyCategory(data),
	}, nil
}

// --- CategoryStatsByMerchantService ---

func (h *CategoryStatsHandler) FindMonthlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalPriceByMerchant) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	data, err := h.repo.GetMonthlyTotalPricing(ctx, int(req.GetYear()), int(req.GetMonth()), "merchant_id", req.GetMerchantId())
	if err != nil {
		h.log.Error("FindMonthlyTotalPricesByMerchant failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapMonthlyPricing(data),
	}, nil
}

func (h *CategoryStatsHandler) FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	data, err := h.repo.GetYearlyTotalPricing(ctx, int(req.GetYear()), "merchant_id", req.GetMerchantId())
	if err != nil {
		h.log.Error("FindYearlyTotalPricesByMerchant failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapYearlyPricing(data),
	}, nil
}

func (h *CategoryStatsHandler) FindMonthPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryMonthPrice, error) {
	data, err := h.repo.GetMonthlyCategoryStats(ctx, int(req.GetYear()), "merchant_id", req.GetMerchantId())
	if err != nil {
		h.log.Error("FindMonthPriceByMerchant failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Monthly payment methods retrieved successfully",
		Data:    mapMonthlyCategory(data),
	}, nil
}

func (h *CategoryStatsHandler) FindYearPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryYearPrice, error) {
	data, err := h.repo.GetYearlyCategoryStats(ctx, int(req.GetYear()), "merchant_id", req.GetMerchantId())
	if err != nil {
		h.log.Error("FindYearPriceByMerchant failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapYearlyCategory(data),
	}, nil
}

// --- Mappers ---

func mapMonthlyPricing(data []repository.MonthlyRevenue) []*pb.CategoriesMonthlyTotalPriceResponse {
	var out []*pb.CategoriesMonthlyTotalPriceResponse
	for _, d := range data {
		out = append(out, &pb.CategoriesMonthlyTotalPriceResponse{
			Year:         d.Year,
			Month:        d.Month,
			TotalRevenue: int32(d.TotalRevenue),
		})
	}
	return out
}

func mapYearlyPricing(data []repository.YearlyRevenue) []*pb.CategoriesYearlyTotalPriceResponse {
	var out []*pb.CategoriesYearlyTotalPriceResponse
	for _, d := range data {
		out = append(out, &pb.CategoriesYearlyTotalPriceResponse{
			Year:         d.Year,
			TotalRevenue: int32(d.TotalRevenue),
		})
	}
	return out
}

func mapMonthlyCategory(data []repository.MonthlyCategory) []*pb.CategoryMonthPriceResponse {
	var out []*pb.CategoryMonthPriceResponse
	for _, d := range data {
		out = append(out, &pb.CategoryMonthPriceResponse{
			Month:        d.Month,
			CategoryId:   int32(d.CategoryID),
			CategoryName: d.CategoryName,
			OrderCount:   int32(d.OrderCount),
			ItemsSold:    int32(d.ItemsSold),
			TotalRevenue: int32(d.TotalRevenue),
		})
	}
	return out
}

func mapYearlyCategory(data []repository.YearlyCategory) []*pb.CategoryYearPriceResponse {
	var out []*pb.CategoryYearPriceResponse
	for _, d := range data {
		out = append(out, &pb.CategoryYearPriceResponse{
			Year:               d.Year,
			CategoryId:         int32(d.CategoryID),
			CategoryName:       d.CategoryName,
			OrderCount:         int32(d.OrderCount),
			ItemsSold:          int32(d.ItemsSold),
			TotalRevenue:       int32(d.TotalRevenue),
			UniqueProductsSold: int32(d.UniqueProductsSold),
		})
	}
	return out
}
