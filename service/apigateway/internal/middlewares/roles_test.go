package middlewares

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mycontext "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/context"
)

type fakeRoleChecker struct {
	calls      int
	lastUserID int
	lastRoles  []string
	err        error
}

func (f *fakeRoleChecker) CheckRole(ctx context.Context, userID int, requiredRoles ...string) error {
	f.calls++
	f.lastUserID = userID
	f.lastRoles = requiredRoles
	return f.err
}

func requestWithUser(t *testing.T, userID int) *http.Request {
	t.Helper()

	req := httptest.NewRequest(http.MethodPost, "/query", nil)
	return req.WithContext(mycontext.WithUserID(req.Context(), userID))
}

func TestRequireRoles_SkipsWhenNoUserInContext(t *testing.T) {
	checker := &fakeRoleChecker{}
	called := false

	mw := RequireRoles(checker, "admin")(okHandler(&called))
	req := httptest.NewRequest(http.MethodPost, "/query", nil)
	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, req)

	if !called {
		t.Fatal("expected public (unauthenticated) request to pass through")
	}
	if checker.calls != 0 {
		t.Fatalf("expected checker not to run, ran %d times", checker.calls)
	}
}

func TestRequireRoles_AllowsMatchingRole(t *testing.T) {
	checker := &fakeRoleChecker{}
	called := false

	mw := RequireRoles(checker, "admin", "superadmin")(okHandler(&called))
	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, requestWithUser(t, 42))

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !called {
		t.Fatal("expected next handler to be called")
	}
	if checker.lastUserID != 42 {
		t.Fatalf("expected user id 42 to reach the checker, got %d", checker.lastUserID)
	}
	if len(checker.lastRoles) != 2 || checker.lastRoles[0] != "admin" {
		t.Fatalf("expected required roles to be forwarded, got %v", checker.lastRoles)
	}
}

func TestRequireRoles_RejectsWhenCheckerDenies(t *testing.T) {
	checker := &fakeRoleChecker{err: errors.New("role not permitted")}
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next should not be called")
	})

	mw := RequireRoles(checker, "admin")(next)
	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, requestWithUser(t, 7))

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d", rec.Code)
	}
}

func TestRequireRoles_NoAllowedRolesStillAuthenticates(t *testing.T) {
	checker := &fakeRoleChecker{}
	called := false

	mw := RequireRoles(checker)(okHandler(&called))
	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, requestWithUser(t, 3))

	if rec.Code != http.StatusOK || !called {
		t.Fatalf("expected request with an authenticated user to pass, got status %d", rec.Code)
	}
	if checker.calls != 0 {
		t.Fatalf("expected checker not to run without allowed roles, ran %d times", checker.calls)
	}
}
