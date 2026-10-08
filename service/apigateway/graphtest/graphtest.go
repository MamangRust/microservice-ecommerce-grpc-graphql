package graphtest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	mycontext "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/context"
	graph "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/handler"
	graphqlmapper "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/middlewares"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/auth"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/banner"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/cart"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/category"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_award"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_business"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_detail"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_policy"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/merchant_social_link"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/order"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/order_item"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/product"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/review"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/review_detail"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/role"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/shipping_address"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/slider"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/transaction"
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/user"
	pb_user_role "github.com/MamangRust/microservice-ecommerce-grpc-pb/user_role"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-pkg/upload_image"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
)

// Resolver is an alias for the internal graph.Resolver type.
type Resolver = graph.Resolver

// ServiceConnections is an alias for the internal ServiceConnections type.
type ServiceConnections = graph.ServiceConnections

// ConnMap converts a *ServiceConnections to a map[string]*grpc.ClientConn.
func ConnMap(c *ServiceConnections) map[string]*grpc.ClientConn {
	if c == nil {
		return nil
	}
	m := map[string]*grpc.ClientConn{
		"auth":              c.AuthClient,
		"role":              c.RoleClient,
		"user":              c.UserClient,
		"category":          c.CategoryClient,
		"merchant":          c.MerchantClient,
		"order-item":        c.OrderItemClient,
		"order":             c.OrderClient,
		"product":           c.ProductClient,
		"transaction":       c.TransactionClient,
		"cart":              c.CartClient,
		"review":            c.ReviewClient,
		"slider":            c.SliderClient,
		"shipping-address":  c.ShippingClient,
		"banner":            c.BannerClient,
		"merchant_award":    c.MerchantAwardClient,
		"merchant_business": c.MerchantBusinessClient,
		"merchant_detail":   c.MerchantDetailClient,
		"merchant_policy":   c.MerchantPolicyClient,
		"review-detail":     c.ReviewDetailClient,
		"merchant-social":   c.MerchantSocialLinkClient,
		"stats_reader":      c.StatsReaderClient,
	}
	return m
}

// NewResolver creates a new Resolver from the given connections, logger, and cache.
func NewResolver(conns map[string]*grpc.ClientConn, log logger.LoggerInterface, cacheStore *cache.CacheStore) *Resolver {
	mapper := graphqlmapper.NewGraphqlMapper()
	imageUpload := upload_image.NewImageUpload(log)
	clients := buildGRPCClients(conns)

	return graph.NewResolver(&graph.Deps{
		Clients:     clients,
		Logger:      log,
		Mapping:     mapper,
		Cache:       cacheStore,
		ImageUpload: imageUpload,
	})
}

// NewTestHandler builds a GraphQL HTTP handler directly from a connection map,
// cache store and logger. It is the convenience entry point used by the
// external test suites (see tests/*/handler_graphql_test.go).
func NewTestHandler(conns map[string]*grpc.ClientConn, cacheStore *cache.CacheStore, log logger.LoggerInterface) http.Handler {
	return NewHandler(NewResolver(conns, log, cacheStore))
}

// permissiveRoleChecker stands in for the RBAC checker of the running gateway.
//
// The harness mounts the schema without AuthMiddleware, and that middleware is
// what puts the authenticated user in the request context, so there is no user
// for the @hasRole directive to authorise and it passes every field through.
// The checker is wired anyway so the directive is never left nil, and so that a
// test which injects a user (see WithUser) gets a predictable answer instead of
// an "rbac: role checker is not configured" error.
type permissiveRoleChecker struct{}

func (permissiveRoleChecker) CheckRole(context.Context, int, ...string) error { return nil }

// NewHandler creates a GraphQL HTTP handler from a resolver.
func NewHandler(resolver *Resolver) *handler.Server {
	return NewHandlerWithRoleChecker(resolver, permissiveRoleChecker{})
}

// NewHandlerWithRoleChecker builds a GraphQL HTTP handler with an explicit RBAC
// checker, so a test can exercise the @hasRole directive (pair it with WithUser
// to simulate an authenticated caller).
func NewHandlerWithRoleChecker(resolver *Resolver, checker middlewares.RoleChecker) *handler.Server {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers:  resolver,
		Directives: graph.DirectiveRoot{HasRole: middlewares.HasRole(checker)},
	}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	return srv
}

// WithUser emulates what AuthMiddleware does for authenticated traffic by
// putting the given user id in the request context, so the @hasRole directive
// can be exercised end to end.
func WithUser(next http.Handler, userID int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(mycontext.WithUserID(r.Context(), userID)))
	})
}

// GraphQLQuery represents a GraphQL request payload.
type GraphQLQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQLResponse wraps the standard GraphQL response structure.
type GraphQLResponse struct {
	Data   map[string]interface{} `json:"data"`
	Errors []GraphQLError         `json:"errors,omitempty"`
}

// GraphQLError represents a single error in the GraphQL response.
type GraphQLError struct {
	Message string `json:"message"`
}

// ExecuteGraphQL executes a query against the handler.
func ExecuteGraphQL(srv http.Handler, query string, variables map[string]interface{}, authToken string) (*GraphQLResponse, error) {
	gqlReq := GraphQLQuery{
		Query:     query,
		Variables: variables,
	}
	body, err := json.Marshal(gqlReq)
	if err != nil {
		return nil, err
	}

	req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if authToken != "" {
		req.Header.Set("Authorization", "Bearer "+authToken)
	}

	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	respBody, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		return nil, err
	}

	var resp GraphQLResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// buildGRPCClients creates GRPCClients from a connection map.
func buildGRPCClients(conns map[string]*grpc.ClientConn) *graph.GRPCClients {
	get := func(key string) *grpc.ClientConn {
		if c, ok := conns[key]; ok && c != nil {
			return c
		}
		return nil
	}

	return &graph.GRPCClients{
		AuthClient:                       newAuthServiceClient(get("auth")),
		RoleCommandClient:                newRoleCommandServiceClient(get("role")),
		RoleQueryClient:                  newRoleQueryServiceClient(get("role")),
		UserRoleQueryClient:              newUserRoleQueryServiceClient(get("role")),
		UserCommandClient:                newUserCommandServiceClient(get("user")),
		UserQueryClient:                  newUserQueryServiceClient(get("user")),
		BannerCommandClient:              newBannerCommandServiceClient(get("banner")),
		BannerQueryClient:                newBannerQueryServiceClient(get("banner")),
		CartCommandClient:                newCartCommandServiceClient(get("cart")),
		CartQueryClient:                  newCartQueryServiceClient(get("cart")),
		CategoryCommandClient:            newCategoryCommandServiceClient(get("category")),
		CategoryQueryClient:              newCategoryQueryServiceClient(get("category")),
		CategoryStatsClient:              newCategoryStatsServiceClient(get("stats_reader")),
		CategoryStatsByMerchantClient:    newCategoryStatsByMerchantServiceClient(get("stats_reader")),
		CategoryStatsByIdClient:          newCategoryStatsByIdServiceClient(get("stats_reader")),
		MerchantCommandClient:            newMerchantCommandServiceClient(get("merchant")),
		MerchantQueryClient:              newMerchantQueryServiceClient(get("merchant")),
		MerchantAwardCommandClient:       newMerchantAwardCommandServiceClient(get("merchant_award")),
		MerchantAwardQueryClient:         newMerchantAwardQueryServiceClient(get("merchant_award")),
		MerchantBusinessCommandClient:    newMerchantBusinessCommandServiceClient(get("merchant_business")),
		MerchantBusinessQueryClient:      newMerchantBusinessQueryServiceClient(get("merchant_business")),
		MerchantDetailCommandClient:      newMerchantDetailCommandServiceClient(get("merchant_detail")),
		MerchantDetailQueryClient:        newMerchantDetailQueryServiceClient(get("merchant_detail")),
		MerchantPolicyCommandClient:      newMerchantPolicyCommandServiceClient(get("merchant_policy")),
		MerchantPolicyQueryClient:        newMerchantPolicyQueryServiceClient(get("merchant_policy")),
		MerchantSocialLinkClient:         newMerchantSocialCommandServiceClient(get("merchant-social")),
		OrderCommandClient:               newOrderCommandServiceClient(get("order")),
		OrderQueryClient:                 newOrderQueryServiceClient(get("order")),
		OrderStatsClient:                 newOrderStatsServiceClient(get("stats_reader")),
		OrderStatsByMerchantClient:       newOrderStatsByMerchantServiceClient(get("stats_reader")),
		OrderItemCommandClient:           newOrderItemCommandServiceClient(get("order-item")),
		OrderItemQueryClient:             newOrderItemQueryServiceClient(get("order-item")),
		ProductCommandClient:             newProductCommandServiceClient(get("product")),
		ProductQueryClient:               newProductQueryServiceClient(get("product")),
		ReviewCommandClient:              newReviewCommandServiceClient(get("review")),
		ReviewQueryClient:                newReviewQueryServiceClient(get("review")),
		ReviewDetailCommandClient:        newReviewDetailCommandServiceClient(get("review-detail")),
		ReviewDetailQueryClient:          newReviewDetailQueryServiceClient(get("review-detail")),
		ShippingCommandClient:            newShippingCommandServiceClient(get("shipping-address")),
		ShippingQueryClient:              newShippingQueryServiceClient(get("shipping-address")),
		SliderCommandClient:              newSliderCommandServiceClient(get("slider")),
		SliderQueryClient:                newSliderQueryServiceClient(get("slider")),
		TransactionCommandClient:         newTransactionCommandServiceClient(get("transaction")),
		TransactionQueryClient:           newTransactionQueryServiceClient(get("transaction")),
		TransactionStatsClient:           newTransactionStatsServiceClient(get("stats_reader")),
		TransactionStatsByMerchantClient: newTransactionStatsByMerchantServiceClient(get("stats_reader")),
	}
}

// Nil-safe gRPC client constructors

func newAuthServiceClient(c *grpc.ClientConn) pb_auth.AuthServiceClient {
	if c == nil {
		return nil
	}
	return pb_auth.NewAuthServiceClient(c)
}

func newRoleCommandServiceClient(c *grpc.ClientConn) pb_role.RoleCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_role.NewRoleCommandServiceClient(c)
}

func newRoleQueryServiceClient(c *grpc.ClientConn) pb_role.RoleQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_role.NewRoleQueryServiceClient(c)
}

func newUserRoleQueryServiceClient(c *grpc.ClientConn) pb_user_role.UserRoleQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_user_role.NewUserRoleQueryServiceClient(c)
}

func newUserCommandServiceClient(c *grpc.ClientConn) pb_user.UserCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_user.NewUserCommandServiceClient(c)
}

func newUserQueryServiceClient(c *grpc.ClientConn) pb_user.UserQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_user.NewUserQueryServiceClient(c)
}

func newBannerCommandServiceClient(c *grpc.ClientConn) pb_banner.BannerCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_banner.NewBannerCommandServiceClient(c)
}

func newBannerQueryServiceClient(c *grpc.ClientConn) pb_banner.BannerQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_banner.NewBannerQueryServiceClient(c)
}

func newCartCommandServiceClient(c *grpc.ClientConn) pb_cart.CartCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_cart.NewCartCommandServiceClient(c)
}

func newCartQueryServiceClient(c *grpc.ClientConn) pb_cart.CartQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_cart.NewCartQueryServiceClient(c)
}

func newCategoryCommandServiceClient(c *grpc.ClientConn) pb_category.CategoryCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_category.NewCategoryCommandServiceClient(c)
}

func newCategoryQueryServiceClient(c *grpc.ClientConn) pb_category.CategoryQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_category.NewCategoryQueryServiceClient(c)
}

func newCategoryStatsServiceClient(c *grpc.ClientConn) pb_category.CategoryStatsServiceClient {
	if c == nil {
		return nil
	}
	return pb_category.NewCategoryStatsServiceClient(c)
}

func newCategoryStatsByMerchantServiceClient(c *grpc.ClientConn) pb_category.CategoryStatsByMerchantServiceClient {
	if c == nil {
		return nil
	}
	return pb_category.NewCategoryStatsByMerchantServiceClient(c)
}

func newCategoryStatsByIdServiceClient(c *grpc.ClientConn) pb_category.CategoryStatsByIdServiceClient {
	if c == nil {
		return nil
	}
	return pb_category.NewCategoryStatsByIdServiceClient(c)
}

func newMerchantCommandServiceClient(c *grpc.ClientConn) pb_merchant.MerchantCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant.NewMerchantCommandServiceClient(c)
}

func newMerchantQueryServiceClient(c *grpc.ClientConn) pb_merchant.MerchantQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant.NewMerchantQueryServiceClient(c)
}

func newMerchantAwardCommandServiceClient(c *grpc.ClientConn) pb_merchant_award.MerchantAwardCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_award.NewMerchantAwardCommandServiceClient(c)
}

func newMerchantAwardQueryServiceClient(c *grpc.ClientConn) pb_merchant_award.MerchantAwardQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_award.NewMerchantAwardQueryServiceClient(c)
}

func newMerchantBusinessCommandServiceClient(c *grpc.ClientConn) pb_merchant_business.MerchantBusinessCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_business.NewMerchantBusinessCommandServiceClient(c)
}

func newMerchantBusinessQueryServiceClient(c *grpc.ClientConn) pb_merchant_business.MerchantBusinessQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_business.NewMerchantBusinessQueryServiceClient(c)
}

func newMerchantDetailCommandServiceClient(c *grpc.ClientConn) pb_merchant_detail.MerchantDetailCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_detail.NewMerchantDetailCommandServiceClient(c)
}

func newMerchantDetailQueryServiceClient(c *grpc.ClientConn) pb_merchant_detail.MerchantDetailQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_detail.NewMerchantDetailQueryServiceClient(c)
}

func newMerchantPolicyCommandServiceClient(c *grpc.ClientConn) pb_merchant_policy.MerchantPolicyCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_policy.NewMerchantPolicyCommandServiceClient(c)
}

func newMerchantPolicyQueryServiceClient(c *grpc.ClientConn) pb_merchant_policy.MerchantPolicyQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_policy.NewMerchantPolicyQueryServiceClient(c)
}

func newMerchantSocialCommandServiceClient(c *grpc.ClientConn) pb_merchant_social_link.MerchantSocialCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_merchant_social_link.NewMerchantSocialCommandServiceClient(c)
}

func newOrderCommandServiceClient(c *grpc.ClientConn) pb_order.OrderCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_order.NewOrderCommandServiceClient(c)
}

func newOrderQueryServiceClient(c *grpc.ClientConn) pb_order.OrderQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_order.NewOrderQueryServiceClient(c)
}

func newOrderStatsServiceClient(c *grpc.ClientConn) pb_order.OrderStatsServiceClient {
	if c == nil {
		return nil
	}
	return pb_order.NewOrderStatsServiceClient(c)
}

func newOrderStatsByMerchantServiceClient(c *grpc.ClientConn) pb_order.OrderStatsByMerchantServiceClient {
	if c == nil {
		return nil
	}
	return pb_order.NewOrderStatsByMerchantServiceClient(c)
}

func newOrderItemCommandServiceClient(c *grpc.ClientConn) pb_order_item.OrderItemCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_order_item.NewOrderItemCommandServiceClient(c)
}

func newOrderItemQueryServiceClient(c *grpc.ClientConn) pb_order_item.OrderItemQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_order_item.NewOrderItemQueryServiceClient(c)
}

func newProductCommandServiceClient(c *grpc.ClientConn) pb_product.ProductCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_product.NewProductCommandServiceClient(c)
}

func newProductQueryServiceClient(c *grpc.ClientConn) pb_product.ProductQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_product.NewProductQueryServiceClient(c)
}

func newReviewCommandServiceClient(c *grpc.ClientConn) pb_review.ReviewCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_review.NewReviewCommandServiceClient(c)
}

func newReviewQueryServiceClient(c *grpc.ClientConn) pb_review.ReviewQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_review.NewReviewQueryServiceClient(c)
}

func newReviewDetailCommandServiceClient(c *grpc.ClientConn) pb_review_detail.ReviewDetailCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_review_detail.NewReviewDetailCommandServiceClient(c)
}

func newReviewDetailQueryServiceClient(c *grpc.ClientConn) pb_review_detail.ReviewDetailQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_review_detail.NewReviewDetailQueryServiceClient(c)
}

func newShippingCommandServiceClient(c *grpc.ClientConn) pb_shipping_address.ShippingCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_shipping_address.NewShippingCommandServiceClient(c)
}

func newShippingQueryServiceClient(c *grpc.ClientConn) pb_shipping_address.ShippingQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_shipping_address.NewShippingQueryServiceClient(c)
}

func newSliderCommandServiceClient(c *grpc.ClientConn) pb_slider.SliderCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_slider.NewSliderCommandServiceClient(c)
}

func newSliderQueryServiceClient(c *grpc.ClientConn) pb_slider.SliderQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_slider.NewSliderQueryServiceClient(c)
}

func newTransactionCommandServiceClient(c *grpc.ClientConn) pb_transaction.TransactionCommandServiceClient {
	if c == nil {
		return nil
	}
	return pb_transaction.NewTransactionCommandServiceClient(c)
}

func newTransactionQueryServiceClient(c *grpc.ClientConn) pb_transaction.TransactionQueryServiceClient {
	if c == nil {
		return nil
	}
	return pb_transaction.NewTransactionQueryServiceClient(c)
}

func newTransactionStatsServiceClient(c *grpc.ClientConn) pb_transaction.TransactionStatsServiceClient {
	if c == nil {
		return nil
	}
	return pb_transaction.NewTransactionStatsServiceClient(c)
}

func newTransactionStatsByMerchantServiceClient(c *grpc.ClientConn) pb_transaction.TransactionStatsByMerchantServiceClient {
	if c == nil {
		return nil
	}
	return pb_transaction.NewTransactionStatsByMerchantServiceClient(c)
}
