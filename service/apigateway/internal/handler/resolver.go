package graph

import (
	errorstd "errors"
	"fmt"
	"time"

	apicache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache"
	graphql "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/mapper"
	rolepermission "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/permission/role"
	auth_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/auth"
	banner_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/banner"
	cart_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/cart"
	category_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/category"
	merchant_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/merchant"
	merchantawards_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/merchant_awards"
	merchantbusiness_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/merchant_business"
	merchantdetail_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/merchant_detail"
	merchantpolicies_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/merchant_policies"
	order_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/order"
	orderitem_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/order_item"
	product_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/product"
	review_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/review"
	reviewdetail_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/review_detail"
	role_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/role"
	shippingaddress_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/shipping_address"
	slider_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/slider"
	transaction_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/transaction"
	user_cache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/redis/api/user"
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
	"github.com/MamangRust/microservice-ecommerce-pkg/kafka"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-pkg/upload_image"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type ServiceConnections struct {
	AuthClient               *grpc.ClientConn
	RoleClient               *grpc.ClientConn
	UserClient               *grpc.ClientConn
	CategoryClient           *grpc.ClientConn
	MerchantClient           *grpc.ClientConn
	OrderItemClient          *grpc.ClientConn
	OrderClient              *grpc.ClientConn
	ProductClient            *grpc.ClientConn
	TransactionClient        *grpc.ClientConn
	CartClient               *grpc.ClientConn
	ReviewClient             *grpc.ClientConn
	SliderClient             *grpc.ClientConn
	ShippingClient           *grpc.ClientConn
	BannerClient             *grpc.ClientConn
	MerchantAwardClient      *grpc.ClientConn
	MerchantBusinessClient   *grpc.ClientConn
	MerchantDetailClient     *grpc.ClientConn
	MerchantPolicyClient     *grpc.ClientConn
	ReviewDetailClient       *grpc.ClientConn
	MerchantSocialLinkClient *grpc.ClientConn
	StatsReaderClient        *grpc.ClientConn
}

type Resolver struct {
	AuthGraphql               *AuthHandleGraphql
	RoleGraphql               *RoleHandleGraphql
	UserGraphql               *UserHandleGraphql
	CartGraphql               *CartHandleGraphql
	BannerGraphql             *BannerHandleGraphql
	CategoryGraphql           *CategoryHandleGraphql
	MerchantGraphql           *MerchantHandleGraphql
	MerchantAwardGraphql      *MerchantAwardHandleGraphql
	MerchantBusinessGraphql   *MerchantBusinessHandleGraphql
	MerchantDetailGraphql     *MerchantDetailHandleGraphql
	MerchantPolicyGraphql     *MerchantPolicyHandleGraphql
	MerchantSocialLinkGraphql *MerchantSocialLinkHandleGraphql
	OrderGraphql              *OrderHandleGraphql
	OrderItemGraphql          *OrderItemHandleGraphql
	ProductGraphql            *ProductHandleGraphql
	ReviewGraphql             *ReviewHandleGraphql
	ReviewDetailGraphql       *ReviewDetailHandleGraphql
	ShippingAddressGraphql    *ShippingAddressHandleGraphql
	SliderGraphql             *SliderHandleGraphql
	TransactionGraphql        *TransactionHandleGraphql
	StatsRead                 *StatsReadHandleGraphql
	ResolverHandle            *resolverHandler
}

type GRPCClients struct {
	AuthClient                       pb_auth.AuthServiceClient
	RoleCommandClient                pb_role.RoleCommandServiceClient
	RoleQueryClient                  pb_role.RoleQueryServiceClient
	UserRoleQueryClient              pb_user_role.UserRoleQueryServiceClient
	UserCommandClient                pb_user.UserCommandServiceClient
	UserQueryClient                  pb_user.UserQueryServiceClient
	BannerCommandClient              pb_banner.BannerCommandServiceClient
	BannerQueryClient                pb_banner.BannerQueryServiceClient
	CartCommandClient                pb_cart.CartCommandServiceClient
	CartQueryClient                  pb_cart.CartQueryServiceClient
	CategoryCommandClient            pb_category.CategoryCommandServiceClient
	CategoryQueryClient              pb_category.CategoryQueryServiceClient
	CategoryStatsClient              pb_category.CategoryStatsServiceClient
	CategoryStatsByMerchantClient    pb_category.CategoryStatsByMerchantServiceClient
	CategoryStatsByIdClient          pb_category.CategoryStatsByIdServiceClient
	MerchantCommandClient            pb_merchant.MerchantCommandServiceClient
	MerchantQueryClient              pb_merchant.MerchantQueryServiceClient
	MerchantAwardCommandClient       pb_merchant_award.MerchantAwardCommandServiceClient
	MerchantAwardQueryClient         pb_merchant_award.MerchantAwardQueryServiceClient
	MerchantBusinessCommandClient    pb_merchant_business.MerchantBusinessCommandServiceClient
	MerchantBusinessQueryClient      pb_merchant_business.MerchantBusinessQueryServiceClient
	MerchantDetailCommandClient      pb_merchant_detail.MerchantDetailCommandServiceClient
	MerchantDetailQueryClient        pb_merchant_detail.MerchantDetailQueryServiceClient
	MerchantPolicyCommandClient      pb_merchant_policy.MerchantPolicyCommandServiceClient
	MerchantPolicyQueryClient        pb_merchant_policy.MerchantPolicyQueryServiceClient
	MerchantSocialLinkClient         pb_merchant_social_link.MerchantSocialCommandServiceClient
	OrderCommandClient               pb_order.OrderCommandServiceClient
	OrderQueryClient                 pb_order.OrderQueryServiceClient
	OrderStatsClient                 pb_order.OrderStatsServiceClient
	OrderStatsByMerchantClient       pb_order.OrderStatsByMerchantServiceClient
	OrderItemCommandClient           pb_order_item.OrderItemCommandServiceClient
	OrderItemQueryClient             pb_order_item.OrderItemQueryServiceClient
	ProductCommandClient             pb_product.ProductCommandServiceClient
	ProductQueryClient               pb_product.ProductQueryServiceClient
	ReviewCommandClient              pb_review.ReviewCommandServiceClient
	ReviewQueryClient                pb_review.ReviewQueryServiceClient
	ReviewDetailCommandClient        pb_review_detail.ReviewDetailCommandServiceClient
	ReviewDetailQueryClient          pb_review_detail.ReviewDetailQueryServiceClient
	ShippingCommandClient            pb_shipping_address.ShippingCommandServiceClient
	ShippingQueryClient              pb_shipping_address.ShippingQueryServiceClient
	SliderCommandClient              pb_slider.SliderCommandServiceClient
	SliderQueryClient                pb_slider.SliderQueryServiceClient
	TransactionCommandClient         pb_transaction.TransactionCommandServiceClient
	TransactionQueryClient           pb_transaction.TransactionQueryServiceClient
	TransactionStatsClient           pb_transaction.TransactionStatsServiceClient
	TransactionStatsByMerchantClient pb_transaction.TransactionStatsByMerchantServiceClient
}

type Deps struct {
	Clients     *GRPCClients
	Logger      logger.LoggerInterface
	Mapping     *graphql.GraphqlMapper
	Cache       *cache.CacheStore
	ImageUpload upload_image.ImageUploads
	Kafka       *kafka.Kafka
}

func NewResolver(deps *Deps) *Resolver {
	obs, _ := observability.NewObservability(
		"graphql-client",
		deps.Logger,
	)

	resolver := NewResolverHandler(obs, deps.Logger)

	// RBAC for the @hasRole directive: roles are resolved through the Kafka
	// request-role/response-role handshake against the role service's consumer
	// and cached in Redis, matching the payment gateway's RolePermission.
	rolePermission := rolepermission.NewRolePermission(
		deps.Kafka,
		"request-role",
		"response-role",
		5*time.Second,
		deps.Logger,
		apicache.NewRoleCache(deps.Cache),
	)

	return &Resolver{
		AuthGraphql: &AuthHandleGraphql{
			AuthClient: deps.Clients.AuthClient,
			Mapping:    deps.Mapping.AuthGraphqlMapper,
			Logger:     deps.Logger,
			Cache:      auth_cache.NewMencache(deps.Cache),
		},
		RoleGraphql: &RoleHandleGraphql{
			RoleCommandClient:   deps.Clients.RoleCommandClient,
			RoleQueryClient:     deps.Clients.RoleQueryClient,
			UserRoleQueryClient: deps.Clients.UserRoleQueryClient,
			Mapping:             deps.Mapping.RoleGraphqlMapper,
			Logger:              deps.Logger,
			Cache:               role_cache.NewRoleMencache(deps.Cache),
			Permission:          rolePermission,
		},
		UserGraphql: &UserHandleGraphql{
			UserCommandClient: deps.Clients.UserCommandClient,
			UserQueryClient:   deps.Clients.UserQueryClient,
			Mapping:           deps.Mapping.UserGraphqlMapper,
			Logger:            deps.Logger,
			Cache:             user_cache.NewUserMencache(deps.Cache),
		},
		BannerGraphql: &BannerHandleGraphql{
			BannerCommandClient: deps.Clients.BannerCommandClient,
			BannerQueryClient:   deps.Clients.BannerQueryClient,
			Mapping:             deps.Mapping.BannerGraphqlMapper,
			Logger:              deps.Logger,
			Cache:               banner_cache.NewBannerMencache(deps.Cache),
		},
		CartGraphql: &CartHandleGraphql{
			CartCommandClient: deps.Clients.CartCommandClient,
			CartQueryClient:   deps.Clients.CartQueryClient,
			Mapping:           deps.Mapping.CartGraphqlMapper,
			Logger:            deps.Logger,
			Cache:             cart_cache.NewCartMencache(deps.Cache),
		},
		CategoryGraphql: &CategoryHandleGraphql{
			CategoryCommandClient: deps.Clients.CategoryCommandClient,
			CategoryQueryClient:   deps.Clients.CategoryQueryClient,
			Mapping:               deps.Mapping.CategoryGraphqlMapper,
			Logger:                deps.Logger,
			Cache:                 category_cache.NewCategoryMencache(deps.Cache),
			UploadImage:           deps.ImageUpload,
		},
		MerchantGraphql: &MerchantHandleGraphql{
			MerchantCommandClient: deps.Clients.MerchantCommandClient,
			MerchantQueryClient:   deps.Clients.MerchantQueryClient,
			Mapping:               deps.Mapping.MerchantGraphqlMapper,
			Logger:                deps.Logger,
			Cache:                 merchant_cache.NewMerchantMencache(deps.Cache),
		},
		MerchantAwardGraphql: &MerchantAwardHandleGraphql{
			MerchantAwardCommandClient: deps.Clients.MerchantAwardCommandClient,
			MerchantAwardQueryClient:   deps.Clients.MerchantAwardQueryClient,
			Mapping:                    deps.Mapping.MerchantAwardGraphqlMapper,
			Logger:                     deps.Logger,
			Cache:                      merchantawards_cache.NewMerchantAward(deps.Cache),
		},
		MerchantBusinessGraphql: &MerchantBusinessHandleGraphql{
			MerchantBusinessCommandClient: deps.Clients.MerchantBusinessCommandClient,
			MerchantBusinessQueryClient:   deps.Clients.MerchantBusinessQueryClient,
			Mapping:                       deps.Mapping.MerchantBusinessGraphqlMapper,
			Logger:                        deps.Logger,
			Cache:                         merchantbusiness_cache.NewMerchantBusinessMencache(deps.Cache),
		},
		MerchantDetailGraphql: &MerchantDetailHandleGraphql{
			MerchantDetailCommandClient: deps.Clients.MerchantDetailCommandClient,
			MerchantDetailQueryClient:   deps.Clients.MerchantDetailQueryClient,
			Mapping:                     deps.Mapping.MerchantDetailGraphqlMapper,
			UploadImage:                 deps.ImageUpload,
			Logger:                      deps.Logger,
			Cache:                       merchantdetail_cache.NewMerchantDetailMencache(deps.Cache),
		},
		MerchantPolicyGraphql: &MerchantPolicyHandleGraphql{
			MerchantPolicyCommandClient: deps.Clients.MerchantPolicyCommandClient,
			MerchantPolicyQueryClient:   deps.Clients.MerchantPolicyQueryClient,
			Mapping:                     deps.Mapping.MerchantPolicyGraphqlMapper,
			Logger:                      deps.Logger,
			Cache:                       merchantpolicies_cache.NewMerchantPoliciesMencache(deps.Cache),
		},
		MerchantSocialLinkGraphql: &MerchantSocialLinkHandleGraphql{
			MerchantSocialLinkClient: deps.Clients.MerchantSocialLinkClient,
			Mapping:                  deps.Mapping.MerchantSocialLinkGraphqlMapper,
			Logger:                   deps.Logger,
		},
		OrderGraphql: &OrderHandleGraphql{
			OrderCommandClient: deps.Clients.OrderCommandClient,
			OrderQueryClient:   deps.Clients.OrderQueryClient,
			Mapping:            deps.Mapping.OrderGraphqlMapper,
			Logger:             deps.Logger,
			Cache:              order_cache.OrderNewMencache(deps.Cache),
		},
		OrderItemGraphql: &OrderItemHandleGraphql{
			OrderItemCommandClient: deps.Clients.OrderItemCommandClient,
			OrderItemQueryClient:   deps.Clients.OrderItemQueryClient,
			Mapping:                deps.Mapping.OrderItemGraphqlMapper,
			Logger:                 deps.Logger,
			Cache:                  orderitem_cache.NewOrderItemMencache(deps.Cache),
		},
		ProductGraphql: &ProductHandleGraphql{
			ProductCommandClient: deps.Clients.ProductCommandClient,
			ProductQueryClient:   deps.Clients.ProductQueryClient,
			Mapping:              deps.Mapping.ProductGraphqlMapper,
			UploadImage:          deps.ImageUpload,
			Logger:               deps.Logger,
			Cache:                product_cache.NewProductMencache(deps.Cache),
		},
		ReviewGraphql: &ReviewHandleGraphql{
			ReviewCommandClient: deps.Clients.ReviewCommandClient,
			ReviewQueryClient:   deps.Clients.ReviewQueryClient,
			Mapping:             deps.Mapping.ReviewGraphqlMapper,
			Logger:              deps.Logger,
			Cache:               review_cache.NewReviewMencache(deps.Cache),
		},
		ReviewDetailGraphql: &ReviewDetailHandleGraphql{
			ReviewDetailCommandClient: deps.Clients.ReviewDetailCommandClient,
			ReviewDetailQueryClient:   deps.Clients.ReviewDetailQueryClient,
			Mapping:                   deps.Mapping.ReviewDetailGraphqlMapper,
			Logger:                    deps.Logger,
			Cache:                     reviewdetail_cache.NewReviewDetailMencache(deps.Cache),
		},
		ShippingAddressGraphql: &ShippingAddressHandleGraphql{
			ShippingCommandClient: deps.Clients.ShippingCommandClient,
			ShippingQueryClient:   deps.Clients.ShippingQueryClient,
			Mapping:               deps.Mapping.ShippingAddresGraphqlMapper,
			Logger:                deps.Logger,
			Cache:                 shippingaddress_cache.NewShippingAddressMencache(deps.Cache),
		},
		SliderGraphql: &SliderHandleGraphql{
			SliderCommandClient: deps.Clients.SliderCommandClient,
			SliderQueryClient:   deps.Clients.SliderQueryClient,
			Mapping:             deps.Mapping.SliderGraphqlMapper,
			UploadImage:         deps.ImageUpload,
			Logger:              deps.Logger,
			Cache:               slider_cache.NewSliderMencache(deps.Cache),
		},
		TransactionGraphql: &TransactionHandleGraphql{
			TransactionCommandClient: deps.Clients.TransactionCommandClient,
			TransactionQueryClient:   deps.Clients.TransactionQueryClient,
			Mapping:                  deps.Mapping.TransactionGraphqlMapper,
			Logger:                   deps.Logger,
			Cache:                    transaction_cache.NewTransactionMencache(deps.Cache),
		},
		StatsRead: &StatsReadHandleGraphql{
			CategoryStats:              deps.Clients.CategoryStatsClient,
			CategoryStatsById:          deps.Clients.CategoryStatsByIdClient,
			CategoryStatsByMerchant:    deps.Clients.CategoryStatsByMerchantClient,
			OrderStats:                 deps.Clients.OrderStatsClient,
			OrderStatsByMerchant:       deps.Clients.OrderStatsByMerchantClient,
			TransactionStats:           deps.Clients.TransactionStatsClient,
			TransactionStatsByMerchant: deps.Clients.TransactionStatsByMerchantClient,
		},
		ResolverHandle: resolver,
	}
}

type AuthHandleGraphql struct {
	AuthClient pb_auth.AuthServiceClient
	Mapping    graphql.AuthGraphqlMapper
	Logger     logger.LoggerInterface
	Cache      auth_cache.AuthMencache
}

type RoleHandleGraphql struct {
	RoleCommandClient   pb_role.RoleCommandServiceClient
	RoleQueryClient     pb_role.RoleQueryServiceClient
	UserRoleQueryClient pb_user_role.UserRoleQueryServiceClient
	Mapping             graphql.RoleGraphqlMapper
	Logger              logger.LoggerInterface
	Cache               role_cache.RoleMencache

	// Permission validates a user's roles over the role service's gRPC
	// FindByUserId, caching the result in Redis. The @hasRole directive (and,
	// for HTTP routes, middlewares.RequireRoles) use it to enforce role based
	// access.
	Permission rolepermission.RolePermission
}

type UserHandleGraphql struct {
	UserCommandClient pb_user.UserCommandServiceClient
	UserQueryClient   pb_user.UserQueryServiceClient
	Mapping           graphql.UserGraphqlMapper
	Logger            logger.LoggerInterface
	Cache             user_cache.UserMencache
}

type BannerHandleGraphql struct {
	BannerCommandClient pb_banner.BannerCommandServiceClient
	BannerQueryClient   pb_banner.BannerQueryServiceClient
	Mapping             graphql.BannerGraphqlMapper
	Logger              logger.LoggerInterface
	Cache               banner_cache.BannerMencache
}

type CartHandleGraphql struct {
	CartCommandClient pb_cart.CartCommandServiceClient
	CartQueryClient   pb_cart.CartQueryServiceClient
	Mapping           graphql.CartGraphqlMapper
	Logger            logger.LoggerInterface
	Cache             cart_cache.CartMencache
}

type CategoryHandleGraphql struct {
	CategoryCommandClient pb_category.CategoryCommandServiceClient
	CategoryQueryClient   pb_category.CategoryQueryServiceClient
	Mapping               graphql.CategoryGraphqlMapper
	UploadImage           upload_image.ImageUploads
	Logger                logger.LoggerInterface
	Cache                 category_cache.CategoryMencache
}

type MerchantHandleGraphql struct {
	MerchantCommandClient pb_merchant.MerchantCommandServiceClient
	MerchantQueryClient   pb_merchant.MerchantQueryServiceClient
	Mapping               graphql.MerchantGraphqlMapper
	Logger                logger.LoggerInterface
	Cache                 merchant_cache.MerchantMencache
}

type MerchantAwardHandleGraphql struct {
	MerchantAwardCommandClient pb_merchant_award.MerchantAwardCommandServiceClient
	MerchantAwardQueryClient   pb_merchant_award.MerchantAwardQueryServiceClient
	Mapping                    graphql.MerchantAwardGraphqlMapper
	Logger                     logger.LoggerInterface
	Cache                      merchantawards_cache.MerchantAwardMencache
}

type MerchantBusinessHandleGraphql struct {
	MerchantBusinessCommandClient pb_merchant_business.MerchantBusinessCommandServiceClient
	MerchantBusinessQueryClient   pb_merchant_business.MerchantBusinessQueryServiceClient
	Mapping                       graphql.MerchantBusinessGraphqlMapper
	Logger                        logger.LoggerInterface
	Cache                         merchantbusiness_cache.MerchantBusinessMencache
}

type MerchantDetailHandleGraphql struct {
	MerchantDetailCommandClient pb_merchant_detail.MerchantDetailCommandServiceClient
	MerchantDetailQueryClient   pb_merchant_detail.MerchantDetailQueryServiceClient
	Mapping                     graphql.MerchantDetailGraphqlMapper
	UploadImage                 upload_image.ImageUploads
	Logger                      logger.LoggerInterface
	Cache                       merchantdetail_cache.MerchantDetailMencache
}

type MerchantPolicyHandleGraphql struct {
	MerchantPolicyCommandClient pb_merchant_policy.MerchantPolicyCommandServiceClient
	MerchantPolicyQueryClient   pb_merchant_policy.MerchantPolicyQueryServiceClient
	Mapping                     graphql.MerchantPolicyGraphqlMapper
	Logger                      logger.LoggerInterface
	Cache                       merchantpolicies_cache.MerchantPoliciesMencache
}

type MerchantSocialLinkHandleGraphql struct {
	MerchantSocialLinkClient pb_merchant_social_link.MerchantSocialCommandServiceClient
	Mapping                  graphql.MerchantSocialLinkGraphqlMapper
	Logger                   logger.LoggerInterface
}

type OrderHandleGraphql struct {
	OrderCommandClient pb_order.OrderCommandServiceClient
	OrderQueryClient   pb_order.OrderQueryServiceClient
	Mapping            graphql.OrderGraphqlMapper
	Logger             logger.LoggerInterface
	Cache              order_cache.OrderMencache
}

type OrderItemHandleGraphql struct {
	OrderItemCommandClient pb_order_item.OrderItemCommandServiceClient
	OrderItemQueryClient   pb_order_item.OrderItemQueryServiceClient
	Mapping                graphql.OrderItemGraphqlMapper
	Logger                 logger.LoggerInterface
	Cache                  orderitem_cache.OrderItemMencache
}

type ProductHandleGraphql struct {
	ProductCommandClient pb_product.ProductCommandServiceClient
	ProductQueryClient   pb_product.ProductQueryServiceClient
	Mapping              graphql.ProductGraphqlMapper
	UploadImage          upload_image.ImageUploads
	Logger               logger.LoggerInterface
	Cache                product_cache.ProductMencache
}

type ReviewHandleGraphql struct {
	ReviewCommandClient pb_review.ReviewCommandServiceClient
	ReviewQueryClient   pb_review.ReviewQueryServiceClient
	Mapping             graphql.ReviewGraphqlMapper
	Logger              logger.LoggerInterface
	Cache               review_cache.ReviewMencache
}

type ReviewDetailHandleGraphql struct {
	ReviewDetailCommandClient pb_review_detail.ReviewDetailCommandServiceClient
	ReviewDetailQueryClient   pb_review_detail.ReviewDetailQueryServiceClient
	Mapping                   graphql.ReviewDetailGraphqlMapper
	Logger                    logger.LoggerInterface
	Cache                     reviewdetail_cache.ReviewDetailMencache
}

type ShippingAddressHandleGraphql struct {
	ShippingCommandClient pb_shipping_address.ShippingCommandServiceClient
	ShippingQueryClient   pb_shipping_address.ShippingQueryServiceClient
	Mapping               graphql.ShippingAddresGraphqlMapper
	Logger                logger.LoggerInterface
	Cache                 shippingaddress_cache.ShippingAddressMencache
}

type SliderHandleGraphql struct {
	SliderCommandClient pb_slider.SliderCommandServiceClient
	SliderQueryClient   pb_slider.SliderQueryServiceClient
	Mapping             graphql.SliderGraphqlMapper
	UploadImage         upload_image.ImageUploads
	Logger              logger.LoggerInterface
	Cache               slider_cache.SliderMencache
}

type TransactionHandleGraphql struct {
	TransactionCommandClient pb_transaction.TransactionCommandServiceClient
	TransactionQueryClient   pb_transaction.TransactionQueryServiceClient
	Mapping                  graphql.TransactionGraphqlMapper
	Logger                   logger.LoggerInterface
	Cache                    transaction_cache.TransactionMencache
}

type StatsReadHandleGraphql struct {
	CategoryStats              pb_category.CategoryStatsServiceClient
	CategoryStatsById          pb_category.CategoryStatsByIdServiceClient
	CategoryStatsByMerchant    pb_category.CategoryStatsByMerchantServiceClient
	OrderStats                 pb_order.OrderStatsServiceClient
	OrderStatsByMerchant       pb_order.OrderStatsByMerchantServiceClient
	TransactionStats           pb_transaction.TransactionStatsServiceClient
	TransactionStatsByMerchant pb_transaction.TransactionStatsByMerchantServiceClient
}

func (h *Resolver) handleGraphQLError(err error, operation string) *errors.AppError {
	if err == nil {
		return nil
	}

	var appErr *errors.AppError
	if errorstd.As(err, &appErr) {
		return appErr
	}

	return errors.NewInternalError(err).WithMessage("Failed to " + operation)
}

func (h *Resolver) parseValidationErrors(err error) []errors.ValidationError {
	var validationErrs []errors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, errors.ValidationError{
				Field:   fe.Field(),
				Message: h.getValidationMessage(fe),
			})
		}
		return validationErrs
	}

	return []errors.ValidationError{
		{
			Field:   "general",
			Message: err.Error(),
		},
	}
}

func (h *Resolver) getValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s", fe.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s", fe.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	default:
		return fmt.Sprintf("Validation failed on '%s' tag", fe.Tag())
	}
}
