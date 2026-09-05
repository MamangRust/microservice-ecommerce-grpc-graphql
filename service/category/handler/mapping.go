package handler

import (
	"encoding/json"
	"log"

	db "github.com/MamangRust/microservice-ecommerce-grpc-category/database/schema"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (h *Handler) mapToCategoryResponse(data interface{}) interface{} {
	switch v := data.(type) {
	case *db.GetCategoryByIDRow:
		return &pb.CategoryResponse{
			Id:            int32(v.CategoryID),
			Name:          v.Name,
			Description:   *v.Description,
			SlugCategory:  *v.SlugCategory,
			ImageCategory: *v.ImageCategory,
			CreatedAt:     v.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt:     v.UpdatedAt.Time.Format("2006-01-02"),
		}
	case *db.GetCategoriesRow:
		return &pb.CategoryResponse{
			Id:            int32(v.CategoryID),
			Name:          v.Name,
			Description:   *v.Description,
			SlugCategory:  *v.SlugCategory,
			ImageCategory: *v.ImageCategory,
			CreatedAt:     v.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt:     v.UpdatedAt.Time.Format("2006-01-02"),
		}
	case *db.GetCategoriesActiveRow:
		var deletedAt string
		if v.DeletedAt.Valid {
			deletedAt = v.DeletedAt.Time.Format("2006-01-02")
		}
		return &pb.CategoryResponseDeleteAt{
			Id:            int32(v.CategoryID),
			Name:          v.Name,
			Description:   *v.Description,
			SlugCategory:  *v.SlugCategory,
			ImageCategory: *v.ImageCategory,
			CreatedAt:     v.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt:     v.UpdatedAt.Time.Format("2006-01-02"),
			DeletedAt:     &wrapperspb.StringValue{Value: deletedAt},
		}
	case *db.GetCategoriesTrashedRow:
		var deletedAt string
		if v.DeletedAt.Valid {
			deletedAt = v.DeletedAt.Time.Format("2006-01-02")
		}
		return &pb.CategoryResponseDeleteAt{
			Id:            int32(v.CategoryID),
			Name:          v.Name,
			Description:   *v.Description,
			SlugCategory:  *v.SlugCategory,
			ImageCategory: *v.ImageCategory,
			CreatedAt:     v.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt:     v.UpdatedAt.Time.Format("2006-01-02"),
			DeletedAt:     &wrapperspb.StringValue{Value: deletedAt},
		}
	case *db.CreateCategoryRow:
		return &pb.CategoryResponse{
			Id:            int32(v.CategoryID),
			Name:          v.Name,
			Description:   *v.Description,
			SlugCategory:  *v.SlugCategory,
			ImageCategory: *v.ImageCategory,
			CreatedAt:     v.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt:     v.UpdatedAt.Time.Format("2006-01-02"),
		}
	case *db.UpdateCategoryRow:
		return &pb.CategoryResponse{
			Id:            int32(v.CategoryID),
			Name:          v.Name,
			Description:   *v.Description,
			SlugCategory:  *v.SlugCategory,
			ImageCategory: *v.ImageCategory,
			CreatedAt:     v.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt:     v.UpdatedAt.Time.Format("2006-01-02"),
		}
	case *db.Category:
		var deletedAt string
		if v.DeletedAt.Valid {
			deletedAt = v.DeletedAt.Time.Format("2006-01-02")
		}
		return &pb.CategoryResponseDeleteAt{
			Id:            int32(v.CategoryID),
			Name:          v.Name,
			Description:   *v.Description,
			SlugCategory:  *v.SlugCategory,
			ImageCategory: *v.ImageCategory,
			CreatedAt:     v.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt:     v.UpdatedAt.Time.Format("2006-01-02"),
			DeletedAt:     &wrapperspb.StringValue{Value: deletedAt},
		}
	default:
		log.Printf("Unknown type for mapping: %T", v)
		return nil
	}
}

func (h *Handler) mapToPayload(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
