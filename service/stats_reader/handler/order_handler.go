package handler

import (
	"context"

	"github.com/MamangRust/microservice-ecommerce-grpc-stats-reader/repository"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"go.uber.org/zap"
)

// OrderStatsHandler serves OrderStatsService (revenue + order aggregates) from
// ClickHouse.
type OrderStatsHandler struct {
	pb.UnimplementedOrderStatsServiceServer
	repo repository.Repository
	log  logger.LoggerInterface
}

func NewOrderStatsHandler(repo repository.Repository, log logger.LoggerInterface) *OrderStatsHandler {
	return &OrderStatsHandler{repo: repo, log: log}
}

func (h *OrderStatsHandler) FindMonthlyTotalRevenue(ctx context.Context, req *pb.FindYearMonthTotalRevenue) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	data, err := h.repo.GetMonthlyTotalRevenue(ctx, int(req.GetYear()), int(req.GetMonth()), 0)
	if err != nil {
		h.log.Error("FindMonthlyTotalRevenue failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseOrderMonthlyTotalRevenue{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapMonthlyRevenue(data),
	}, nil
}

func (h *OrderStatsHandler) FindYearlyTotalRevenue(ctx context.Context, req *pb.FindYearTotalRevenue) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	data, err := h.repo.GetYearlyTotalRevenue(ctx, int(req.GetYear()), 0)
	if err != nil {
		h.log.Error("FindYearlyTotalRevenue failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseOrderYearlyTotalRevenue{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapYearlyRevenue(data),
	}, nil
}

func (h *OrderStatsHandler) FindMonthlyRevenue(ctx context.Context, req *pb.FindYearOrder) (*pb.ApiResponseOrderMonthly, error) {
	data, err := h.repo.GetMonthlyOrderStats(ctx, int(req.GetYear()), 0)
	if err != nil {
		h.log.Error("FindMonthlyRevenue failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseOrderMonthly{
		Status:  "success",
		Message: "Monthly revenue data retrieved",
		Data:    mapMonthlyOrder(data),
	}, nil
}

func (h *OrderStatsHandler) FindYearlyRevenue(ctx context.Context, req *pb.FindYearOrder) (*pb.ApiResponseOrderYearly, error) {
	data, err := h.repo.GetYearlyOrderStats(ctx, int(req.GetYear()), 0)
	if err != nil {
		h.log.Error("FindYearlyRevenue failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseOrderYearly{
		Status:  "success",
		Message: "Yearly revenue data retrieved",
		Data:    mapYearlyOrder(data),
	}, nil
}

func (h *OrderStatsHandler) FindMonthlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearMonthTotalRevenueByMerchant) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	data, err := h.repo.GetMonthlyTotalRevenue(ctx, int(req.GetYear()), int(req.GetMonth()), req.GetMerchantId())
	if err != nil {
		h.log.Error("FindMonthlyTotalRevenueByMerchant failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseOrderMonthlyTotalRevenue{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    mapMonthlyRevenue(data),
	}, nil
}

func (h *OrderStatsHandler) FindYearlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearTotalRevenueByMerchant) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	data, err := h.repo.GetYearlyTotalRevenue(ctx, int(req.GetYear()), req.GetMerchantId())
	if err != nil {
		h.log.Error("FindYearlyTotalRevenueByMerchant failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseOrderYearlyTotalRevenue{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    mapYearlyRevenue(data),
	}, nil
}

func (h *OrderStatsHandler) FindMonthlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) (*pb.ApiResponseOrderMonthly, error) {
	data, err := h.repo.GetMonthlyOrderStats(ctx, int(req.GetYear()), req.GetMerchantId())
	if err != nil {
		h.log.Error("FindMonthlyRevenueByMerchant failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseOrderMonthly{
		Status:  "success",
		Message: "Monthly revenue by merchant data retrieved",
		Data:    mapMonthlyOrder(data),
	}, nil
}

func (h *OrderStatsHandler) FindYearlyRevenueByMerchant(ctx context.Context, req *pb.FindYearOrderByMerchant) (*pb.ApiResponseOrderYearly, error) {
	data, err := h.repo.GetYearlyOrderStats(ctx, int(req.GetYear()), req.GetMerchantId())
	if err != nil {
		h.log.Error("FindYearlyRevenueByMerchant failed", zap.Error(err))
		return nil, err
	}
	return &pb.ApiResponseOrderYearly{
		Status:  "success",
		Message: "Yearly revenue by merchant data retrieved",
		Data:    mapYearlyOrder(data),
	}, nil
}

// --- Mappers ---

func mapMonthlyRevenue(data []repository.MonthlyRevenue) []*pb.OrderMonthlyTotalRevenueResponse {
	var out []*pb.OrderMonthlyTotalRevenueResponse
	for _, d := range data {
		out = append(out, &pb.OrderMonthlyTotalRevenueResponse{
			Year:         d.Year,
			Month:        d.Month,
			TotalRevenue: int32(d.TotalRevenue),
		})
	}
	return out
}

func mapYearlyRevenue(data []repository.YearlyRevenue) []*pb.OrderYearlyTotalRevenueResponse {
	var out []*pb.OrderYearlyTotalRevenueResponse
	for _, d := range data {
		out = append(out, &pb.OrderYearlyTotalRevenueResponse{
			Year:         d.Year,
			TotalRevenue: int32(d.TotalRevenue),
		})
	}
	return out
}

func mapMonthlyOrder(data []repository.MonthlyOrder) []*pb.OrderMonthlyResponse {
	var out []*pb.OrderMonthlyResponse
	for _, d := range data {
		out = append(out, &pb.OrderMonthlyResponse{
			Month:          d.Month,
			OrderCount:     int32(d.OrderCount),
			TotalRevenue:   int32(d.TotalRevenue),
			TotalItemsSold: int32(d.TotalItemsSold),
		})
	}
	return out
}

func mapYearlyOrder(data []repository.YearlyOrder) []*pb.OrderYearlyResponse {
	var out []*pb.OrderYearlyResponse
	for _, d := range data {
		out = append(out, &pb.OrderYearlyResponse{
			Year:               d.Year,
			OrderCount:         int32(d.OrderCount),
			TotalRevenue:       int32(d.TotalRevenue),
			TotalItemsSold:     int32(d.TotalItemsSold),
			UniqueProductsSold: int32(d.UniqueProductsSold),
		})
	}
	return out
}
