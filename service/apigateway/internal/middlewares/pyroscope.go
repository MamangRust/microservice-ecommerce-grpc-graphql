package middlewares

import (
	"context"
	"net/http"

	"github.com/grafana/pyroscope-go"
)

// PyroscopeMiddleware tags the profiling context of every request with its path
// and method so CPU profiles can be broken down per endpoint. Previously an
// Echo-only middleware, ported to net/http for the gqlgen server.
func PyroscopeMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pyroscope.TagWrapper(r.Context(), pyroscope.Labels(
				"path", r.URL.Path,
				"method", r.Method,
			), func(ctx context.Context) {
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		})
	}
}
