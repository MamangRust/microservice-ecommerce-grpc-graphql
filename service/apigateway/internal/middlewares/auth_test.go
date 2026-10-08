package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mycontext "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/context"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"go.uber.org/zap"
)

type fakeTokenManager struct {
	validateFn func(token string) (string, error)
}

func (f *fakeTokenManager) GenerateToken(userId int, audience string) (string, error) {
	return "", nil
}

func (f *fakeTokenManager) ValidateToken(token string) (string, error) {
	if f.validateFn == nil {
		return "", errors.New("no token manager configured")
	}
	return f.validateFn(token)
}

func nopLog() logger.LoggerInterface {
	return &logger.Logger{Log: zap.NewNop()}
}

func doRequest(method, target, body, authHeader string, next http.Handler) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}
	rec := httptest.NewRecorder()
	next.ServeHTTP(rec, req)
	return rec
}

func TestWriteJSONError(t *testing.T) {
	rec := httptest.NewRecorder()
	writeJSONError(rec, "boom", http.StatusBadRequest)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Fatalf("expected JSON content type, got %q", ct)
	}

	var body struct {
		Errors []map[string]interface{} `json:"errors"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}
	if len(body.Errors) != 1 {
		t.Fatalf("expected 1 error entry, got %d", len(body.Errors))
	}
	if body.Errors[0]["message"] != "boom" {
		t.Fatalf("expected message 'boom', got %v", body.Errors[0]["message"])
	}
}

func TestAuthMiddleware_SkipsPublicOperations(t *testing.T) {
	publicQueries := []string{
		`{ loginUser(username: "a", password: "b") { status } }`,
		`{ registerUser(...) { status } }`,
		`{ refreshToken { status } }`,
		`query { loginUser { status } }`,
		`mutation { registerUser { status } }`,
	}

	for _, q := range publicQueries {
		t.Run(q, func(t *testing.T) {
			body := `{"query": ` + jsonString(q) + `}`
			called := false
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = true
				w.WriteHeader(http.StatusOK)
			})
			mw := AuthMiddleware(&fakeTokenManager{}, nopLog())(next)

			rec := doRequest(http.MethodPost, "/query", body, "", mw)
			if !called {
				t.Fatal("expected next handler to be called for public operation")
			}
			if rec.Code != http.StatusOK {
				t.Fatalf("expected status 200, got %d", rec.Code)
			}
		})
	}
}

func TestAuthMiddleware_MissingHeader(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next should not be called")
	})
	mw := AuthMiddleware(&fakeTokenManager{}, nopLog())(next)

	rec := doRequest(http.MethodPost, "/query", `{"query":"{ getMe { id } }"}`, "", mw)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rec.Code)
	}
}

func TestAuthMiddleware_InvalidFormat(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next should not be called")
	})
	mw := AuthMiddleware(&fakeTokenManager{}, nopLog())(next)

	for _, header := range []string{"Basic abc", "BearerOnlyToken"} {
		rec := doRequest(http.MethodPost, "/query", `{"query":"{ getMe { id } }"}`, header, mw)
		if rec.Code != http.StatusUnauthorized {
			t.Fatalf("header %q: expected status 401, got %d", header, rec.Code)
		}
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	tm := &fakeTokenManager{
		validateFn: func(token string) (string, error) {
			return "", errors.New("expired")
		},
	}
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next should not be called")
	})
	mw := AuthMiddleware(tm, nopLog())(next)

	rec := doRequest(http.MethodPost, "/query", `{"query":"{ getMe { id } }"}`, "Bearer bad.token.value", mw)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rec.Code)
	}
}

func TestAuthMiddleware_InvalidUserIDInToken(t *testing.T) {
	tm := &fakeTokenManager{
		validateFn: func(token string) (string, error) {
			return "not-a-number", nil
		},
	}
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next should not be called")
	})
	mw := AuthMiddleware(tm, nopLog())(next)

	rec := doRequest(http.MethodPost, "/query", `{"query":"{ getMe { id } }"}`, "Bearer good.token.value", mw)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rec.Code)
	}
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	tm := &fakeTokenManager{
		validateFn: func(token string) (string, error) {
			return "42", nil
		},
	}

	var gotUserID int
	var gotOK bool
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUserID, gotOK = mycontext.UserForContext(r.Context())
		w.WriteHeader(http.StatusOK)
	})
	mw := AuthMiddleware(tm, nopLog())(next)

	rec := doRequest(http.MethodPost, "/query", `{"query":"{ getMe { id } }"}`, "Bearer valid.token.value", mw)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !gotOK {
		t.Fatal("expected user id in context")
	}
	if gotUserID != 42 {
		t.Fatalf("expected user id 42, got %d", gotUserID)
	}
}

func jsonString(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

