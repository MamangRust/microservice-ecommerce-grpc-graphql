package pagination

import (
	"testing"

	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/common"
)

func TestMapPaginationMeta_Nil(t *testing.T) {
	result := MapPaginationMeta(nil)
	if result != nil {
		t.Fatalf("expected nil, got %+v", result)
	}
}

func TestMapPaginationMeta_Full(t *testing.T) {
	meta := &pb_common.PaginationMeta{
		CurrentPage:  1,
		PageSize:     20,
		TotalPages:   5,
		TotalRecords: 100,
	}

	result := MapPaginationMeta(meta)
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	expected := &model.PaginationMeta{
		CurrentPage:  1,
		PageSize:     20,
		TotalPages:   5,
		TotalRecords: 100,
	}

	if *result != *expected {
		t.Fatalf("expected %+v, got %+v", expected, result)
	}
}

func TestMapPaginationMeta_ZeroValues(t *testing.T) {
	meta := &pb_common.PaginationMeta{}

	result := MapPaginationMeta(meta)
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	if result.CurrentPage != 0 || result.PageSize != 0 || result.TotalPages != 0 || result.TotalRecords != 0 {
		t.Fatalf("expected all zeros, got %+v", result)
	}
}
