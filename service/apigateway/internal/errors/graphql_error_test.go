package graphqlerror

import (
	"fmt"
	"net/http"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewGraphqlError(t *testing.T) {
	err := NewGraphqlError("Not Found", "resource not found", http.StatusNotFound)
	expected := "graphql error: [404] Not Found - resource not found"
	if err.Error() != expected {
		t.Fatalf("expected %q, got %q", expected, err.Error())
	}
}

func TestNewGraphqlError_EmptyMessage(t *testing.T) {
	err := NewGraphqlError("OK", "", http.StatusOK)
	expected := "graphql error: [200] OK - "
	if err.Error() != expected {
		t.Fatalf("expected %q, got %q", expected, err.Error())
	}
}

func TestToGraphqlErrorFromErrorResponse_Nil(t *testing.T) {
	err := ToGraphqlErrorFromErrorResponse(nil)
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestToGraphqlErrorFromErrorResponse_NonGRPCError(t *testing.T) {
	original := fmt.Errorf("some random error")
	err := ToGraphqlErrorFromErrorResponse(original)
	expected := fmt.Sprintf("graphql error: %v", original)
	if err.Error() != expected {
		t.Fatalf("expected %q, got %q", expected, err.Error())
	}
}

func TestToGraphqlErrorFromErrorResponse_GRPCErrors(t *testing.T) {
	tests := []struct {
		name       string
		grpcCode   codes.Code
		grpcMsg    string
		wantStatus string
		wantCode   int
	}{
		{
			name:       "InvalidArgument→400",
			grpcCode:   codes.InvalidArgument,
			grpcMsg:    "bad request",
			wantStatus: "Bad Request",
			wantCode:   http.StatusBadRequest,
		},
		{
			name:       "Unauthenticated→401",
			grpcCode:   codes.Unauthenticated,
			grpcMsg:    "unauthorized",
			wantStatus: "Unauthorized",
			wantCode:   http.StatusUnauthorized,
		},
		{
			name:       "PermissionDenied→403",
			grpcCode:   codes.PermissionDenied,
			grpcMsg:    "forbidden",
			wantStatus: "Forbidden",
			wantCode:   http.StatusForbidden,
		},
		{
			name:       "NotFound→404",
			grpcCode:   codes.NotFound,
			grpcMsg:    "not found",
			wantStatus: "Not Found",
			wantCode:   http.StatusNotFound,
		},
		{
			name:       "AlreadyExists→409",
			grpcCode:   codes.AlreadyExists,
			grpcMsg:    "conflict",
			wantStatus: "Conflict",
			wantCode:   http.StatusConflict,
		},
		{
			name:       "ResourceExhausted→429",
			grpcCode:   codes.ResourceExhausted,
			grpcMsg:    "too many requests",
			wantStatus: "Too Many Requests",
			wantCode:   http.StatusTooManyRequests,
		},
		{
			name:       "DeadlineExceeded→504",
			grpcCode:   codes.DeadlineExceeded,
			grpcMsg:    "timeout",
			wantStatus: "Gateway Timeout",
			wantCode:   http.StatusGatewayTimeout,
		},
		{
			name:       "Internal→500",
			grpcCode:   codes.Internal,
			grpcMsg:    "internal error",
			wantStatus: "Internal Server Error",
			wantCode:   http.StatusInternalServerError,
		},
		{
			name:       "Unknown→500",
			grpcCode:   codes.Unknown,
			grpcMsg:    "unknown",
			wantStatus: "Internal Server Error",
			wantCode:   http.StatusInternalServerError,
		},
		{
			name:       "Unavailable→500",
			grpcCode:   codes.Unavailable,
			grpcMsg:    "unavailable",
			wantStatus: "Internal Server Error",
			wantCode:   http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpcErr := status.Error(tt.grpcCode, tt.grpcMsg)
			err := ToGraphqlErrorFromErrorResponse(grpcErr)
			expected := fmt.Sprintf("graphql error: [%d] %s - %s", tt.wantCode, tt.wantStatus, tt.grpcMsg)
			if err.Error() != expected {
				t.Fatalf("expected %q, got %q", expected, err.Error())
			}
		})
	}
}

func TestGrpcToHttpCode(t *testing.T) {
	tests := []struct {
		code     codes.Code
		expected int
	}{
		{codes.InvalidArgument, http.StatusBadRequest},
		{codes.Unauthenticated, http.StatusUnauthorized},
		{codes.PermissionDenied, http.StatusForbidden},
		{codes.NotFound, http.StatusNotFound},
		{codes.AlreadyExists, http.StatusConflict},
		{codes.ResourceExhausted, http.StatusTooManyRequests},
		{codes.DeadlineExceeded, http.StatusGatewayTimeout},
		{codes.OK, http.StatusInternalServerError},
		{codes.Canceled, http.StatusInternalServerError},
		{codes.DataLoss, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.code.String(), func(t *testing.T) {
			got := grpcToHttpCode(tt.code)
			if got != tt.expected {
				t.Fatalf("grpcToHttpCode(%s) = %d, want %d", tt.code, got, tt.expected)
			}
		})
	}
}
