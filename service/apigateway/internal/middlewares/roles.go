package middlewares

import (
	"context"
	"net/http"

	mycontext "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/context"
)

// RoleChecker verifies that a user owns at least one of the required roles.
// *rolepermission.RolePermission (Kafka + cache) satisfies it.
type RoleChecker interface {
	CheckRole(ctx context.Context, userID int, requiredRoles ...string) error
}

// RequireRoles rejects the request with HTTP 403 unless the authenticated user
// holds one of the allowed roles. It is the net/http equivalent of the Echo
// RequireRoles helper that used to live in service/apigateway/middlewares.
//
// Requests without a user in the context pass through untouched: those are the
// public GraphQL operations that AuthMiddleware already whitelisted, so there is
// nothing to authorise.
func RequireRoles(checker RoleChecker, allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := mycontext.UserForContext(r.Context())
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			if len(allowedRoles) > 0 {
				if err := checker.CheckRole(r.Context(), userID, allowedRoles...); err != nil {
					writeJSONError(w, "role not permitted", http.StatusForbidden)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
