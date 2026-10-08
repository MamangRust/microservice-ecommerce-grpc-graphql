package apps

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
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
	"github.com/MamangRust/microservice-ecommerce-grpc-pb/user_role"
	"github.com/MamangRust/microservice-ecommerce-pkg/auth"
	"github.com/MamangRust/microservice-ecommerce-pkg/dotenv"
	"github.com/MamangRust/microservice-ecommerce-pkg/kafka"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	otel_pkg "github.com/MamangRust/microservice-ecommerce-pkg/otel"
	redisclient "github.com/MamangRust/microservice-ecommerce-pkg/redis"
	"github.com/MamangRust/microservice-ecommerce-pkg/upload_image"
	sharedcache "github.com/MamangRust/microservice-ecommerce-shared/cache"
	sharedobservability "github.com/MamangRust/microservice-ecommerce-shared/observability"
	"github.com/go-chi/chi/v5"
	"github.com/grafana/pyroscope-go"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceAddresses struct {
	Auth             string
	Role             string
	User             string
	Category         string
	Merchant         string
	OrderItem        string
	Order            string
	Product          string
	Transaction      string
	Cart             string
	Review           string
	Slider           string
	Shipping         string
	Banner           string
	MerchantAward    string
	MerchantBusiness string
	MerchantDetail   string
	MerchantPolicy   string
	ReviewDetail     string
}

func loadServiceAddresses() *ServiceAddresses {
	return &ServiceAddresses{
		Auth:             getEnvOrDefault("GRPC_AUTH_ADDR", "localhost:50051"),
		Role:             getEnvOrDefault("GRPC_ROLE_ADDR", "localhost:50052"),
		User:             getEnvOrDefault("GRPC_USER_ADDR", "localhost:50053"),
		Category:         getEnvOrDefault("GRPC_CATEGORY_ADDR", "localhost:50054"),
		Merchant:         getEnvOrDefault("GRPC_MERCHANT_ADDR", "localhost:50055"),
		OrderItem:        getEnvOrDefault("GRPC_ORDER_ITEM_ADDR", "localhost:50056"),
		Order:            getEnvOrDefault("GRPC_ORDER_ADDR", "localhost:50057"),
		Product:          getEnvOrDefault("GRPC_PRODUCT_ADDR", "localhost:50058"),
		Transaction:      getEnvOrDefault("GRPC_TRANSACTION_ADDR", "localhost:50059"),
		Cart:             getEnvOrDefault("GRPC_CART_ADDR", "localhost:50060"),
		Review:           getEnvOrDefault("GRPC_REVIEW_ADDR", "localhost:50061"),
		Slider:           getEnvOrDefault("GRPC_SLIDER_ADDR", "localhost:50062"),
		Shipping:         getEnvOrDefault("GRPC_SHIPPING_ADDRESS_ADDR", "localhost:50063"),
		Banner:           getEnvOrDefault("GRPC_BANNER_ADDR", "localhost:50064"),
		MerchantAward:    getEnvOrDefault("GRPC_MERCHANT_AWARD_ADDR", "localhost:50065"),
		MerchantBusiness: getEnvOrDefault("GRPC_MERCHANT_BUSINESS_ADDR", "localhost:50066"),
		MerchantDetail:   getEnvOrDefault("GRPC_MERCHANT_DETAIL_ADDR", "localhost:50067"),
		MerchantPolicy:   getEnvOrDefault("GRPC_MERCHANT_POLICY_ADDR", "localhost:50068"),
		ReviewDetail:     getEnvOrDefault("GRPC_REVIEW_DETAIL_ADDR", "localhost:50069"),
	}
}

func createServiceConnections(addresses *ServiceAddresses, logger logger.LoggerInterface) (*graph.ServiceConnections, error) {
	var connections graph.ServiceConnections

	conns := map[string]*string{
		"Auth":             &addresses.Auth,
		"Role":             &addresses.Role,
		"User":             &addresses.User,
		"Category":         &addresses.Category,
		"Merchant":         &addresses.Merchant,
		"OrderItem":        &addresses.OrderItem,
		"Order":            &addresses.Order,
		"Product":          &addresses.Product,
		"Transaction":      &addresses.Transaction,
		"Cart":             &addresses.Cart,
		"Review":           &addresses.Review,
		"Slider":           &addresses.Slider,
		"Shipping":         &addresses.Shipping,
		"Banner":           &addresses.Banner,
		"MerchantAward":    &addresses.MerchantAward,
		"MerchantBusiness": &addresses.MerchantBusiness,
		"MerchantDetail":   &addresses.MerchantDetail,
		"MerchantPolicy":   &addresses.MerchantPolicy,
		"ReviewDetail":     &addresses.ReviewDetail,
	}

	for name, addr := range conns {
		conn, err := createConnection(*addr, name, logger)
		if err != nil {
			return nil, err
		}

		switch name {
		case "Auth":
			connections.AuthClient = conn
		case "Role":
			connections.RoleClient = conn
		case "User":
			connections.UserClient = conn
		case "Category":
			connections.CategoryClient = conn
		case "Merchant":
			connections.MerchantClient = conn
		case "OrderItem":
			connections.OrderItemClient = conn
		case "Order":
			connections.OrderClient = conn
		case "Product":
			connections.ProductClient = conn
		case "Transaction":
			connections.TransactionClient = conn
		case "Cart":
			connections.CartClient = conn
		case "Review":
			connections.ReviewClient = conn
		case "Slider":
			connections.SliderClient = conn
		case "Shipping":
			connections.ShippingClient = conn
		case "Banner":
			connections.BannerClient = conn
		case "MerchantAward":
			connections.MerchantAwardClient = conn
		case "MerchantBusiness":
			connections.MerchantBusinessClient = conn
		case "MerchantDetail":
			connections.MerchantDetailClient = conn
			connections.MerchantSocialLinkClient = conn
		case "MerchantPolicy":
			connections.MerchantPolicyClient = conn
		case "ReviewDetail":
			connections.ReviewDetailClient = conn
		case "StatsReader":
			connections.StatsReaderClient = conn
		}
	}

	return &connections, nil
}

func createConnection(address, serviceName string, logger logger.LoggerInterface) (*grpc.ClientConn, error) {
	logger.Info(fmt.Sprintf("Connecting to %s service at %s", serviceName, address))
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to %s service", serviceName), zap.Error(err))
		return nil, err
	}
	return conn, nil
}

func closeConnections(conns *graph.ServiceConnections, log logger.LoggerInterface) {
	for name, conn := range map[string]*grpc.ClientConn{
		"Auth":               conns.AuthClient,
		"Role":               conns.RoleClient,
		"User":               conns.UserClient,
		"Category":           conns.CategoryClient,
		"Merchant":           conns.MerchantClient,
		"OrderItem":          conns.OrderItemClient,
		"Order":              conns.OrderClient,
		"Product":            conns.ProductClient,
		"Transaction":        conns.TransactionClient,
		"Cart":               conns.CartClient,
		"Review":             conns.ReviewClient,
		"Slider":             conns.SliderClient,
		"Shipping":           conns.ShippingClient,
		"Banner":             conns.BannerClient,
		"MerchantAward":      conns.MerchantAwardClient,
		"MerchantBusiness":   conns.MerchantBusinessClient,
		"MerchantDetail":     conns.MerchantDetailClient,
		"MerchantPolicy":     conns.MerchantPolicyClient,
		"ReviewDetail":       conns.ReviewDetailClient,
		"MerchantSocialLink": conns.MerchantSocialLinkClient,
	} {
		if conn != nil {
			if err := conn.Close(); err != nil {
				log.Error(fmt.Sprintf("Failed to close %s connection", name), zap.Error(err))
			}
		}
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvIntOrDefault(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsed
}

// initPyroscope wires continuous profiling when PYROSCOPE_SERVER is configured.
// Without the env var the gateway keeps running unprofiled instead of failing
// to boot.
func initPyroscope() error {
	server := os.Getenv("PYROSCOPE_SERVER")
	if server == "" {
		return nil
	}

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "apigateway",
		ServerAddress:   server,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
		Tags: map[string]string{
			"service": "apigateway",
			"env":     os.Getenv("ENV"),
			"version": os.Getenv("VERSION"),
		},
	})

	return err
}

type Client struct {
	Logger logger.LoggerInterface
	Server *http.Server
}

func RunClient() (*Client, func(), error) {
	flag.Parse()

	addresses := loadServiceAddresses()

	if err := dotenv.Viper(); err != nil {
		fmt.Printf("Warning: Failed to load .env file: %v\n", err)
	}

	ctx := context.Background()

	telemetry := otel_pkg.NewTelemetry(otel_pkg.Config{
		ServiceName:          "apigateway",
		ServiceVersion:       "1.0.0",
		Environment:          "development",
		Endpoint:             getEnvOrDefault("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317"),
		Insecure:             true,
		EnableRuntimeMetrics: true,
	})
	if err := telemetry.Init(ctx); err != nil {
		fmt.Printf("Warning: Failed to initialize telemetry: %v\n", err)
	}

	log, err := logger.NewLogger("apigateway", telemetry.GetLogger())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create logger: %w", err)
	}

	if err := initPyroscope(); err != nil {
		log.Warn("Failed to initialize pyroscope", zap.Error(err))
	}

	log.Debug("Creating gRPC connections...")
	conns, err := createServiceConnections(addresses, log)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect services: %w", err)
	}

	tokenManager, err := auth.NewManager(viper.GetString("SECRET_KEY"))
	if err != nil {
		log.Fatal("Failed to create token manager", zap.Error(err))
	}

	myredis := redisclient.NewRedisClient(&redisclient.Config{
		Host:         viper.GetString("REDIS_HOST"),
		Port:         viper.GetString("REDIS_PORT"),
		Password:     viper.GetString("REDIS_PASSWORD"),
		DB:           viper.GetInt("REDIS_DB_APIGATEWAY"),
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 3,
	})

	if err := myredis.Client.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to ping redis", zap.Error(err))
	}

	cacheMetrics, err := sharedobservability.NewCacheMetrics("apigateway")
	if err != nil {
		log.Error("Failed to initialize cache metrics for apigateway cache store", zap.Error(err))
	}
	store := sharedcache.NewCacheStore(myredis.Client, log, cacheMetrics)

	imageUpload := upload_image.NewImageUpload(log)

	graphqlMapper := graphqlmapper.NewGraphqlMapper()

	grpcClients := &graph.GRPCClients{
		AuthClient:                       pb_auth.NewAuthServiceClient(conns.AuthClient),
		RoleCommandClient:                pb_role.NewRoleCommandServiceClient(conns.RoleClient),
		RoleQueryClient:                  pb_role.NewRoleQueryServiceClient(conns.RoleClient),
		UserRoleQueryClient:              pb_user_role.NewUserRoleQueryServiceClient(conns.RoleClient),
		UserCommandClient:                pb_user.NewUserCommandServiceClient(conns.UserClient),
		UserQueryClient:                  pb_user.NewUserQueryServiceClient(conns.UserClient),
		BannerCommandClient:              pb_banner.NewBannerCommandServiceClient(conns.BannerClient),
		BannerQueryClient:                pb_banner.NewBannerQueryServiceClient(conns.BannerClient),
		CartCommandClient:                pb_cart.NewCartCommandServiceClient(conns.CartClient),
		CartQueryClient:                  pb_cart.NewCartQueryServiceClient(conns.CartClient),
		CategoryCommandClient:            pb_category.NewCategoryCommandServiceClient(conns.CategoryClient),
		CategoryQueryClient:              pb_category.NewCategoryQueryServiceClient(conns.CategoryClient),
		CategoryStatsClient:              pb_category.NewCategoryStatsServiceClient(conns.StatsReaderClient),
		CategoryStatsByMerchantClient:    pb_category.NewCategoryStatsByMerchantServiceClient(conns.StatsReaderClient),
		CategoryStatsByIdClient:          pb_category.NewCategoryStatsByIdServiceClient(conns.StatsReaderClient),
		MerchantCommandClient:            pb_merchant.NewMerchantCommandServiceClient(conns.MerchantClient),
		MerchantQueryClient:              pb_merchant.NewMerchantQueryServiceClient(conns.MerchantClient),
		MerchantAwardCommandClient:       pb_merchant_award.NewMerchantAwardCommandServiceClient(conns.MerchantAwardClient),
		MerchantAwardQueryClient:         pb_merchant_award.NewMerchantAwardQueryServiceClient(conns.MerchantAwardClient),
		MerchantBusinessCommandClient:    pb_merchant_business.NewMerchantBusinessCommandServiceClient(conns.MerchantBusinessClient),
		MerchantBusinessQueryClient:      pb_merchant_business.NewMerchantBusinessQueryServiceClient(conns.MerchantBusinessClient),
		MerchantDetailCommandClient:      pb_merchant_detail.NewMerchantDetailCommandServiceClient(conns.MerchantDetailClient),
		MerchantDetailQueryClient:        pb_merchant_detail.NewMerchantDetailQueryServiceClient(conns.MerchantDetailClient),
		MerchantPolicyCommandClient:      pb_merchant_policy.NewMerchantPolicyCommandServiceClient(conns.MerchantPolicyClient),
		MerchantPolicyQueryClient:        pb_merchant_policy.NewMerchantPolicyQueryServiceClient(conns.MerchantPolicyClient),
		MerchantSocialLinkClient:         pb_merchant_social_link.NewMerchantSocialCommandServiceClient(conns.MerchantSocialLinkClient),
		OrderCommandClient:               pb_order.NewOrderCommandServiceClient(conns.OrderClient),
		OrderQueryClient:                 pb_order.NewOrderQueryServiceClient(conns.OrderClient),
		OrderStatsClient:                 pb_order.NewOrderStatsServiceClient(conns.StatsReaderClient),
		OrderStatsByMerchantClient:       pb_order.NewOrderStatsByMerchantServiceClient(conns.StatsReaderClient),
		OrderItemCommandClient:           pb_order_item.NewOrderItemCommandServiceClient(conns.OrderItemClient),
		OrderItemQueryClient:             pb_order_item.NewOrderItemQueryServiceClient(conns.OrderItemClient),
		ProductCommandClient:             pb_product.NewProductCommandServiceClient(conns.ProductClient),
		ProductQueryClient:               pb_product.NewProductQueryServiceClient(conns.ProductClient),
		ReviewCommandClient:              pb_review.NewReviewCommandServiceClient(conns.ReviewClient),
		ReviewQueryClient:                pb_review.NewReviewQueryServiceClient(conns.ReviewClient),
		ReviewDetailCommandClient:        pb_review_detail.NewReviewDetailCommandServiceClient(conns.ReviewDetailClient),
		ReviewDetailQueryClient:          pb_review_detail.NewReviewDetailQueryServiceClient(conns.ReviewDetailClient),
		ShippingCommandClient:            pb_shipping_address.NewShippingCommandServiceClient(conns.ShippingClient),
		ShippingQueryClient:              pb_shipping_address.NewShippingQueryServiceClient(conns.ShippingClient),
		SliderCommandClient:              pb_slider.NewSliderCommandServiceClient(conns.SliderClient),
		SliderQueryClient:                pb_slider.NewSliderQueryServiceClient(conns.SliderClient),
		TransactionCommandClient:         pb_transaction.NewTransactionCommandServiceClient(conns.TransactionClient),
		TransactionQueryClient:           pb_transaction.NewTransactionQueryServiceClient(conns.TransactionClient),
		TransactionStatsClient:           pb_transaction.NewTransactionStatsServiceClient(conns.StatsReaderClient),
		TransactionStatsByMerchantClient: pb_transaction.NewTransactionStatsByMerchantServiceClient(conns.StatsReaderClient),
	}

	resolver := graph.NewResolver(&graph.Deps{
		Clients:     grpcClients,
		Logger:      log,
		Mapping:     graphqlMapper,
		Cache:       store,
		ImageUpload: imageUpload,
		Kafka:       kafka.NewKafka(log, []string{viper.GetString("KAFKA_BROKERS")}),
	})

	srv := setupGraphql(tokenManager, resolver, log)

	go func() {
		log.Info(fmt.Sprintf("🚀 Starting GraphQL server on %s", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("GraphQL server error", zap.Error(err))
		}
	}()

	shutdown := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		log.Info("Shutting down GraphQL API Gateway...")
		if err := srv.Shutdown(ctx); err != nil {
			log.Error("HTTP server shutdown error", zap.Error(err))
		}
		closeConnections(conns, log)

		if err := telemetry.Shutdown(context.Background()); err != nil {
			log.Error("Telemetry shutdown failed", zap.Error(err))
		}

		log.Info("Shutdown complete ✅")
	}

	return &Client{
		Logger: log,
		Server: srv,
	}, shutdown, nil
}

func setupGraphql(token auth.TokenManager, resolver *graph.Resolver, logger logger.LoggerInterface) *http.Server {
	port := getEnvOrDefault("CLIENT_PORT", "5000")

	logger.Debug("Starting GraphQL server", zap.String("port", getEnvOrDefault("CLIENT_PORT", "5000")))

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
		// RBAC is enforced per operation by the @hasRole directive, because
		// every GraphQL operation shares the single POST /query endpoint.
		Directives: graph.DirectiveRoot{
			HasRole: middlewares.HasRole(resolver.RoleGraphql.Permission),
		},
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

	// Zero means "use the package defaults".
	rateLimiter := middlewares.NewRateLimiter(
		getEnvIntOrDefault("GATEWAY_RATE_LIMIT_RPS", 0),
		getEnvIntOrDefault("GATEWAY_RATE_LIMIT_BURST", 0),
	)

	r := chi.NewRouter()

	r.Get("/", playground.Handler("GraphQL Playground", "/query"))

	// Outermost first: profile labels, then flood protection, then auth.
	r.Handle("/query", middlewares.PyroscopeMiddleware()(
		rateLimiter.Middleware(
			middlewares.AuthMiddleware(token, logger)(srv),
		),
	))

	logger.Info("GraphQL Playground running",
		zap.String("url", "http://localhost:"+port),
		zap.String("endpoint", "/query"),
	)

	return &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}
}
