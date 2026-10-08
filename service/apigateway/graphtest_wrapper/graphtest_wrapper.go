// Package graphtest_wrapper re-exports graphtest for external test packages.
package graphtest_wrapper

import (
	"net/http"

	gt "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/graphtest"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/middlewares"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"google.golang.org/grpc"
)

// Resolver is an alias for graphtest.Resolver.
type Resolver = gt.Resolver

// ServiceConnections is an alias for graphtest.ServiceConnections.
type ServiceConnections = gt.ServiceConnections

// ConnMap converts a *ServiceConnections to a map[string]*grpc.ClientConn.
func ConnMap(c *ServiceConnections) map[string]*grpc.ClientConn {
	return gt.ConnMap(c)
}

// NewResolver creates a new Resolver from the given connections.
func NewResolver(conns map[string]*grpc.ClientConn, log logger.LoggerInterface, cacheStore *cache.CacheStore) *Resolver {
	return gt.NewResolver(conns, log, cacheStore)
}

// NewHandler creates a GraphQL HTTP handler from a resolver.
func NewHandler(resolver *Resolver) http.Handler {
	return gt.NewHandler(resolver)
}

// NewTestHandler builds a GraphQL HTTP handler directly from connections, cache
// and logger.
func NewTestHandler(conns map[string]*grpc.ClientConn, cacheStore *cache.CacheStore, log logger.LoggerInterface) http.Handler {
	return gt.NewTestHandler(conns, cacheStore, log)
}

// NewHandlerWithRoleChecker builds a GraphQL HTTP handler with an explicit RBAC
// checker, so a test can exercise the @hasRole directive.
func NewHandlerWithRoleChecker(resolver *Resolver, checker middlewares.RoleChecker) http.Handler {
	return gt.NewHandlerWithRoleChecker(resolver, checker)
}

// WithUser injects the given user id into the request context, emulating what
// AuthMiddleware does for authenticated traffic.
func WithUser(next http.Handler, userID int) http.Handler {
	return gt.WithUser(next, userID)
}

// ExecuteGraphQL executes a query against the handler.
func ExecuteGraphQL(handler http.Handler, query string, variables map[string]interface{}, authToken string) (*gt.GraphQLResponse, error) {
	return gt.ExecuteGraphQL(handler, query, variables, authToken)
}
