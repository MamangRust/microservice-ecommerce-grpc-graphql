package mycontext

import (
	"context"
	"testing"
)

func TestWithUserID_UserForContext_RoundTrip(t *testing.T) {
	ctx := context.Background()
	ctx = WithUserID(ctx, 42)

	id, ok := UserForContext(ctx)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if id != 42 {
		t.Fatalf("expected id=42, got %d", id)
	}
}

func TestUserForContext_Missing(t *testing.T) {
	_, ok := UserForContext(context.Background())
	if ok {
		t.Fatal("expected ok=false for background context")
	}
}

func TestUserForContext_WrongType(t *testing.T) {
	ctx := context.WithValue(context.Background(), UserIDContextKey, "not_an_int")
	_, ok := UserForContext(ctx)
	if ok {
		t.Fatal("expected ok=false when value is not an int")
	}
}

func TestUserForContext_ZeroValue(t *testing.T) {
	ctx := WithUserID(context.Background(), 0)
	id, ok := UserForContext(ctx)
	if !ok {
		t.Fatal("expected ok=true for zero value")
	}
	if id != 0 {
		t.Fatalf("expected id=0, got %d", id)
	}
}

func TestApiKeyFromContext_Present(t *testing.T) {
	ctx := context.WithValue(context.Background(), ApiKeyContextKey, "my-api-key")
	key, ok := ApiKeyFromContext(ctx)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if key != "my-api-key" {
		t.Fatalf("expected key='my-api-key', got %q", key)
	}
}

func TestApiKeyFromContext_Missing(t *testing.T) {
	_, ok := ApiKeyFromContext(context.Background())
	if ok {
		t.Fatal("expected ok=false")
	}
}
