package handler

import (
	"math"

	db "github.com/MamangRust/microservice-ecommerce-grpc-order/database/schema"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func normalizePage(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}

func createPaginationMeta(page, pageSize, totalRecords int) *pb.PaginationMeta {
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))
	return &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}
}

func formatTimestamp(v interface{}) string {
	switch t := v.(type) {
	case pgtype.Timestamptz:
		if t.Valid {
			return t.Time.Format("2006-01-02 15:04:05.000")
		}
	case pgtype.Timestamp:
		if t.Valid {
			return t.Time.Format("2006-01-02 15:04:05.000")
		}
	}
	return ""
}

func mapToProtoOrderResponse(m interface{}) *pb.OrderResponse {
	switch v := m.(type) {
	case *db.Order:
		return &pb.OrderResponse{
			Id:         v.OrderID,
			MerchantId: v.MerchantID,
			UserId:     v.UserID,
			TotalPrice: int32(v.TotalPrice),
			CreatedAt:  formatTimestamp(v.CreatedAt),
			UpdatedAt:  formatTimestamp(v.UpdatedAt),
		}
	case *db.GetOrdersRow:
		return &pb.OrderResponse{
			Id:         v.OrderID,
			MerchantId: v.MerchantID,
			UserId:     v.UserID,
			TotalPrice: int32(v.TotalPrice),
			CreatedAt:  formatTimestamp(v.CreatedAt),
			UpdatedAt:  formatTimestamp(v.UpdatedAt),
		}
	case *db.GetOrderByIDRow:
		return &pb.OrderResponse{
			Id:         v.OrderID,
			MerchantId: v.MerchantID,
			UserId:     v.UserID,
			TotalPrice: int32(v.TotalPrice),
			CreatedAt:  formatTimestamp(v.CreatedAt),
			UpdatedAt:  formatTimestamp(v.UpdatedAt),
		}
	case *db.CreateOrderRow:
		return &pb.OrderResponse{
			Id:         v.OrderID,
			MerchantId: v.MerchantID,
			UserId:     v.UserID,
			TotalPrice: int32(v.TotalPrice),
			CreatedAt:  formatTimestamp(v.CreatedAt),
			UpdatedAt:  formatTimestamp(v.UpdatedAt),
		}
	case *db.UpdateOrderRow:
		return &pb.OrderResponse{
			Id:         v.OrderID,
			MerchantId: v.MerchantID,
			UserId:     v.UserID,
			TotalPrice: int32(v.TotalPrice),
			CreatedAt:  formatTimestamp(v.CreatedAt),
			UpdatedAt:  formatTimestamp(v.UpdatedAt),
		}
	default:
		return nil
	}
}

func mapToProtoOrderResponseDeleteAt(m interface{}) *pb.OrderResponseDeleteAt {
	var res *pb.OrderResponseDeleteAt
	var deletedAt interface{}

	switch v := m.(type) {
	case *db.Order:
		res = &pb.OrderResponseDeleteAt{
			Id:         v.OrderID,
			MerchantId: v.MerchantID,
			UserId:     v.UserID,
			TotalPrice: int32(v.TotalPrice),
			CreatedAt:  formatTimestamp(v.CreatedAt),
			UpdatedAt:  formatTimestamp(v.UpdatedAt),
		}
		deletedAt = v.DeletedAt
	case *db.GetOrdersActiveRow:
		res = &pb.OrderResponseDeleteAt{
			Id:         v.OrderID,
			MerchantId: v.MerchantID,
			UserId:     v.UserID,
			TotalPrice: int32(v.TotalPrice),
			CreatedAt:  formatTimestamp(v.CreatedAt),
			UpdatedAt:  formatTimestamp(v.UpdatedAt),
		}
		deletedAt = v.DeletedAt
	case *db.GetOrdersTrashedRow:
		res = &pb.OrderResponseDeleteAt{
			Id:         v.OrderID,
			MerchantId: v.MerchantID,
			UserId:     v.UserID,
			TotalPrice: int32(v.TotalPrice),
			CreatedAt:  formatTimestamp(v.CreatedAt),
			UpdatedAt:  formatTimestamp(v.UpdatedAt),
		}
		deletedAt = v.DeletedAt
	default:
		return nil
	}

	if val := formatTimestamp(deletedAt); val != "" {
		res.DeletedAt = &wrapperspb.StringValue{Value: val}
	}

	return res
}

