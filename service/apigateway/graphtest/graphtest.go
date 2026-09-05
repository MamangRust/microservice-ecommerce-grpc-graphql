// Package graphtest exposes a test-friendly GraphQL HTTP handler for the
// apigateway module.  External test packages import this instead of internal/.
package graphtest

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-pkg/upload_image"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	pb "github.com/MamangRust/microservice-ecommerce-shared/pb"
	graph "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/handler"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
)

// NewTestHandler builds a full gqlgen HTTP handler wired to the given gRPC
// connections and returns it as a plain http.Handler suitable for httptest.
func NewTestHandler(
	conns map[string]*grpc.ClientConn,
	cacheStore *cache.CacheStore,
	log logger.LoggerInterface,
) http.Handler {
	grpcClients := buildGRPCClients(conns)
	mapper := graphqlmapper.NewGraphqlMapper()
	imageUpload := upload_image.NewImageUpload(log)

	resolver := graph.NewResolver(&graph.Deps{
		Clients:     grpcClients,
		Logger:      log,
		Mapping:     mapper,
		Cache:       cacheStore,
		ImageUpload: imageUpload,
	})

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})
}

func buildGRPCClients(conns map[string]*grpc.ClientConn) *graph.GRPCClients {
	get := func(key string) *grpc.ClientConn {
		if c, ok := conns[key]; ok && c != nil {
			return c
		}
		return nil
	}

	return &graph.GRPCClients{
		AuthClient:                       pb.NewAuthServiceClient(get("auth")),
		RoleCommandClient:                pb.NewRoleCommandServiceClient(get("role")),
		RoleQueryClient:                  pb.NewRoleQueryServiceClient(get("role")),
		UserCommandClient:                pb.NewUserCommandServiceClient(get("user")),
		UserQueryClient:                  pb.NewUserQueryServiceClient(get("user")),
		BannerCommandClient:              pb.NewBannerCommandServiceClient(get("banner")),
		BannerQueryClient:                pb.NewBannerQueryServiceClient(get("banner")),
		CartCommandClient:                pb.NewCartCommandServiceClient(get("cart")),
		CartQueryClient:                  pb.NewCartQueryServiceClient(get("cart")),
		CategoryCommandClient:            pb.NewCategoryCommandServiceClient(get("category")),
		CategoryQueryClient:              pb.NewCategoryQueryServiceClient(get("category")),
		CategoryStatsClient:              pb.NewCategoryStatsServiceClient(get("stats_reader")),
		CategoryStatsByMerchantClient:    pb.NewCategoryStatsByMerchantServiceClient(get("stats_reader")),
		CategoryStatsByIdClient:          pb.NewCategoryStatsByIdServiceClient(get("stats_reader")),
		MerchantCommandClient:            pb.NewMerchantCommandServiceClient(get("merchant")),
		MerchantQueryClient:              pb.NewMerchantQueryServiceClient(get("merchant")),
		MerchantAwardCommandClient:       pb.NewMerchantAwardCommandServiceClient(get("merchant_award")),
		MerchantAwardQueryClient:         pb.NewMerchantAwardQueryServiceClient(get("merchant_award")),
		MerchantBusinessCommandClient:    pb.NewMerchantBusinessCommandServiceClient(get("merchant_business")),
		MerchantBusinessQueryClient:      pb.NewMerchantBusinessQueryServiceClient(get("merchant_business")),
		MerchantDetailCommandClient:      pb.NewMerchantDetailCommandServiceClient(get("merchant_detail")),
		MerchantDetailQueryClient:        pb.NewMerchantDetailQueryServiceClient(get("merchant_detail")),
		MerchantPolicyCommandClient:      pb.NewMerchantPolicyCommandServiceClient(get("merchant_policy")),
		MerchantPolicyQueryClient:        pb.NewMerchantPolicyQueryServiceClient(get("merchant_policy")),
		MerchantSocialLinkClient:         pb.NewMerchantSocialCommandServiceClient(get("merchant")),
		OrderCommandClient:               pb.NewOrderCommandServiceClient(get("order")),
		OrderQueryClient:                 pb.NewOrderQueryServiceClient(get("order")),
		OrderStatsClient:                 pb.NewOrderStatsServiceClient(get("stats_reader")),
		OrderItemCommandClient:           pb.NewOrderItemCommandServiceClient(get("order-item")),
		OrderItemQueryClient:             pb.NewOrderItemQueryServiceClient(get("order-item")),
		ProductCommandClient:             pb.NewProductCommandServiceClient(get("product")),
		ProductQueryClient:               pb.NewProductQueryServiceClient(get("product")),
		ReviewCommandClient:              pb.NewReviewCommandServiceClient(get("review")),
		ReviewQueryClient:                pb.NewReviewQueryServiceClient(get("review")),
		ReviewDetailCommandClient:        pb.NewReviewDetailCommandServiceClient(get("review-detail")),
		ReviewDetailQueryClient:          pb.NewReviewDetailQueryServiceClient(get("review-detail")),
		ShippingCommandClient:            pb.NewShippingCommandServiceClient(get("shipping-address")),
		ShippingQueryClient:              pb.NewShippingQueryServiceClient(get("shipping-address")),
		SliderCommandClient:              pb.NewSliderCommandServiceClient(get("slider")),
		SliderQueryClient:                pb.NewSliderQueryServiceClient(get("slider")),
		TransactionCommandClient:         pb.NewTransactionCommandServiceClient(get("transaction")),
		TransactionQueryClient:           pb.NewTransactionQueryServiceClient(get("transaction")),
		TransactionStatsClient:           pb.NewTransactionStatsServiceClient(get("stats_reader")),
		TransactionStatsByMerchantClient: pb.NewTransactionStatsByMerchantServiceClient(get("stats_reader")),
	}
}
