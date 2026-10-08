package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func okHandler(called *bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*called = true
		w.WriteHeader(http.StatusOK)
	})
}

func TestRateLimiter_AllowsBurstThenRejects(t *testing.T) {
	called := 0
	mw := NewRateLimiter(1, 2).Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called++
		w.WriteHeader(http.StatusOK)
	}))

	for i := 0; i < 2; i++ {
		rec := doRequest(http.MethodPost, "/query", `{"query":"{ getMe { id } }"}`, "", mw)
		if rec.Code != http.StatusOK {
			t.Fatalf("request %d: expected status 200, got %d", i, rec.Code)
		}
	}

	rec := doRequest(http.MethodPost, "/query", `{"query":"{ getMe { id } }"}`, "", mw)
	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("expected status 429 once the burst is exhausted, got %d", rec.Code)
	}
	if called != 2 {
		t.Fatalf("expected next handler to run twice, ran %d times", called)
	}
}

func TestRateLimiter_FallsBackToDefaults(t *testing.T) {
	called := false
	mw := NewRateLimiter(0, 0).Middleware(okHandler(&called))

	rec := doRequest(http.MethodPost, "/query", `{"query":"{ getMe { id } }"}`, "", mw)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 with default settings, got %d", rec.Code)
	}
	if !called {
		t.Fatal("expected next handler to be called")
	}
}

func TestRateLimiter_RejectsWithJSONError(t *testing.T) {
	rec := httptest.NewRecorder()

	limiter := NewRateLimiter(0, 1)
	handler := limiter.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/query", nil))
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/query", nil))

	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("expected status 429, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected JSON content type, got %q", ct)
	}
}
