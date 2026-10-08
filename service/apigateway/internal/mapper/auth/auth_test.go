package authgraphqlmapper

import (
	"testing"

	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/model"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/auth"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/user"
)

func TestToGraphqlResponseLogin(t *testing.T) {
	mapper := NewAuthGraphqlMapper()

	pbResp := &pb_auth.ApiResponseLogin{
		Status:  "success",
		Message: "login successful",
		Data: &pb_auth.TokenResponse{
			AccessToken:  "access-123",
			RefreshToken: "refresh-456",
		},
	}

	result := mapper.ToGraphqlResponseLogin(pbResp)
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	if result.Status != "success" {
		t.Errorf("Status = %q, want %q", result.Status, "success")
	}
	if result.Message != "login successful" {
		t.Errorf("Message = %q, want %q", result.Message, "login successful")
	}
	if result.Data == nil {
		t.Fatal("expected non-nil Data")
	}
	if result.Data.AccessToken != "access-123" {
		t.Errorf("AccessToken = %q, want %q", result.Data.AccessToken, "access-123")
	}
	if result.Data.RefreshToken != "refresh-456" {
		t.Errorf("RefreshToken = %q, want %q", result.Data.RefreshToken, "refresh-456")
	}
}

func TestToGraphqlResponseLogin_PanicsOnNilData(t *testing.T) {
	mapper := NewAuthGraphqlMapper()

	pbResp := &pb_auth.ApiResponseLogin{
		Status:  "error",
		Message: "invalid credentials",
	}

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic when Data is nil")
		}
	}()
	_ = mapper.ToGraphqlResponseLogin(pbResp)
}

func TestToGraphqlResponseRegister(t *testing.T) {
	mapper := NewAuthGraphqlMapper()

	pbResp := &pb_auth.ApiResponseRegister{
		Status:  "success",
		Message: "registration successful",
		Data: &pb_user.UserResponse{
			Id:        1,
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "john@example.com",
			CreatedAt: "2024-01-01T00:00:00Z",
			UpdatedAt: "2024-01-01T00:00:00Z",
		},
	}

	result := mapper.ToGraphqlResponseRegister(pbResp)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Status != "success" {
		t.Errorf("Status = %q, want %q", result.Status, "success")
	}
	if result.Data == nil {
		t.Fatal("expected non-nil Data")
	}
	if result.Data.ID != 1 {
		t.Errorf("ID = %d, want %d", result.Data.ID, 1)
	}
	if result.Data.Email != "john@example.com" {
		t.Errorf("Email = %q, want %q", result.Data.Email, "john@example.com")
	}
}

func TestToGraphqlResponseGetMe(t *testing.T) {
	mapper := NewAuthGraphqlMapper()

	pbResp := &pb_auth.ApiResponseGetMe{
		Status:  "success",
		Message: "user found",
		Data: &pb_user.UserResponse{
			Id:        42,
			Firstname: "Jane",
			Lastname:  "Smith",
			Email:     "jane@example.com",
			CreatedAt: "2024-06-15T10:30:00Z",
			UpdatedAt: "2024-06-15T10:30:00Z",
		},
	}

	result := mapper.ToGraphqlResponseGetMe(pbResp)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Status != "success" {
		t.Errorf("Status = %q, want %q", result.Status, "success")
	}
	if result.Data == nil {
		t.Fatal("expected non-nil Data")
	}
	if result.Data.ID != 42 {
		t.Errorf("ID = %d, want %d", result.Data.ID, 42)
	}
	if result.Data.Firstname != "Jane" {
		t.Errorf("Firstname = %q, want %q", result.Data.Firstname, "Jane")
	}
	if result.Data.UpdatedAt != "2024-06-15T10:30:00Z" {
		t.Errorf("UpdatedAt = %q, want %q", result.Data.UpdatedAt, "2024-06-15T10:30:00Z")
	}
}

func TestToGraphqlRefreshToken(t *testing.T) {
	mapper := NewAuthGraphqlMapper()

	pbResp := &pb_auth.ApiResponseRefreshToken{
		Status:  "success",
		Message: "token refreshed",
		Data: &pb_auth.TokenResponse{
			AccessToken:  "new-access",
			RefreshToken: "new-refresh",
		},
	}

	result := mapper.ToGraphqlResponseRefreshToken(pbResp)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Data == nil {
		t.Fatal("expected non-nil Data")
	}
	if result.Data.AccessToken != "new-access" {
		t.Errorf("AccessToken = %q, want %q", result.Data.AccessToken, "new-access")
	}
	if result.Data.RefreshToken != "new-refresh" {
		t.Errorf("RefreshToken = %q, want %q", result.Data.RefreshToken, "new-refresh")
	}
}

func TestToGraphqlForgotPassword(t *testing.T) {
	mapper := NewAuthGraphqlMapper()

	pbResp := &pb_auth.ApiResponseForgotPassword{
		Status:  "success",
		Message: "check your email",
	}

	result := mapper.ToGraphqlForgotPassword(pbResp)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Status != "success" {
		t.Errorf("Status = %q, want %q", result.Status, "success")
	}
	if result.Message != "check your email" {
		t.Errorf("Message = %q, want %q", result.Message, "check your email")
	}
}

func TestToGraphqlResetPassword(t *testing.T) {
	mapper := NewAuthGraphqlMapper()

	pbResp := &pb_auth.ApiResponseResetPassword{
		Status:  "success",
		Message: "password reset",
	}

	result := mapper.ToGraphqlResetPassword(pbResp)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Status != "success" {
		t.Errorf("Status = %q, want %q", result.Status, "success")
	}
	if result.Message != "password reset" {
		t.Errorf("Message = %q, want %q", result.Message, "password reset")
	}
}

func TestToGraphqlVerifyCode(t *testing.T) {
	mapper := NewAuthGraphqlMapper()

	pbResp := &pb_auth.ApiResponseVerifyCode{
		Status:  "success",
		Message: "code verified",
	}

	result := mapper.ToGraphqlVerifyCode(pbResp)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Status != "success" {
		t.Errorf("Status = %q, want %q", result.Status, "success")
	}
	if result.Message != "code verified" {
		t.Errorf("Message = %q, want %q", result.Message, "code verified")
	}
}

func TestMapResponseUser(t *testing.T) {
	mapper := &authGraphqlMapper{}

	pbUser := &pb_user.UserResponse{
		Id:        99,
		Firstname: "Alice",
		Lastname:  "Wonderland",
		Email:     "alice@example.com",
		CreatedAt: "2024-03-10T08:00:00Z",
		UpdatedAt: "2024-03-10T09:00:00Z",
	}

	result := mapper.mapResponseUser(pbUser)
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	expected := &model.UserResponse{
		ID:        99,
		Firstname: "Alice",
		Lastname:  "Wonderland",
		Email:     "alice@example.com",
		CreatedAt: "2024-03-10T08:00:00Z",
		UpdatedAt: "2024-03-10T09:00:00Z",
	}

	if *result != *expected {
		t.Fatalf("expected %+v, got %+v", expected, result)
	}
}

func TestMapResponseToken(t *testing.T) {
	mapper := &authGraphqlMapper{}

	pbToken := &pb_auth.TokenResponse{
		AccessToken:  "at-111",
		RefreshToken: "rt-222",
	}

	result := mapper.mapResponseToken(pbToken)
	if result == nil {
		t.Fatal("expected non-nil result")
	}

	expected := &model.TokenResponse{
		AccessToken:  "at-111",
		RefreshToken: "rt-222",
	}

	if *result != *expected {
		t.Fatalf("expected %+v, got %+v", expected, result)
	}
}
