package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mycontext "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/context"
)

func TestPyroscopeMiddleware_PassesThroughWithContext(t *testing.T) {
	var gotUserID int
	var gotOK bool

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUserID, gotOK = mycontext.UserForContext(r.Context())
		w.WriteHeader(http.StatusOK)
	})

	mw := PyroscopeMiddleware()(next)
	req := httptest.NewRequest(http.MethodPost, "/query", nil)
	req = req.WithContext(mycontext.WithUserID(req.Context(), 11))

	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !gotOK || gotUserID != 11 {
		t.Fatalf("expected tagged context to keep user id 11, got %d (ok=%v)", gotUserID, gotOK)
	}
}
