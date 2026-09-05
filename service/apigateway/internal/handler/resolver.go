package graph

import (
	errorstd "errors"
	"fmt"

	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"github.com/MamangRust/microservice-ecommerce-pkg/upload_image"
	"github.com/MamangRust/microservice-ecommerce-shared/cache"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"github.com/MamangRust/microservice-ecommerce-shared/observability"
	pb "github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphql "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper"
	auth_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/auth"
	banner_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/banner"
	cart_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/cart"
	category_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/category"
	merchant_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/merchant"
	merchantawards_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/merchant_awards"
	merchantbusiness_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/merchant_business"
	merchantdetail_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/merchant_detail"
	merchantpolicies_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/merchant_policies"
	order_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/order"
	orderitem_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/order_item"
	product_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/product"
	review_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/review"
	reviewdetail_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/review_detail"
	role_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/role"
	shippingaddress_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/shipping_address"
	slider_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/slider"
	transaction_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/transaction"
	user_cache "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/redis/api/user"
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
	StatsReaderClient         *grpc.ClientConn
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
	ResolverHandle            *resolverHandler
}

type GRPCClients struct {
	AuthClient                       pb.AuthServiceClient
	RoleCommandClient                pb.RoleCommandServiceClient
	RoleQueryClient                  pb.RoleQueryServiceClient
	UserCommandClient                pb.UserCommandServiceClient
	UserQueryClient                  pb.UserQueryServiceClient
	BannerCommandClient              pb.BannerCommandServiceClient
	BannerQueryClient                pb.BannerQueryServiceClient
	CartCommandClient                pb.CartCommandServiceClient
	CartQueryClient                  pb.CartQueryServiceClient
	CategoryCommandClient            pb.CategoryCommandServiceClient
	CategoryQueryClient              pb.CategoryQueryServiceClient
	CategoryStatsClient              pb.CategoryStatsServiceClient
	CategoryStatsByMerchantClient    pb.CategoryStatsByMerchantServiceClient
	CategoryStatsByIdClient          pb.CategoryStatsByIdServiceClient
	MerchantCommandClient            pb.MerchantCommandServiceClient
	MerchantQueryClient              pb.MerchantQueryServiceClient
	MerchantAwardCommandClient       pb.MerchantAwardCommandServiceClient
	MerchantAwardQueryClient         pb.MerchantAwardQueryServiceClient
	MerchantBusinessCommandClient    pb.MerchantBusinessCommandServiceClient
	MerchantBusinessQueryClient      pb.MerchantBusinessQueryServiceClient
	MerchantDetailCommandClient      pb.MerchantDetailCommandServiceClient
	MerchantDetailQueryClient        pb.MerchantDetailQueryServiceClient
	MerchantPolicyCommandClient      pb.MerchantPolicyCommandServiceClient
	MerchantPolicyQueryClient        pb.MerchantPolicyQueryServiceClient
	MerchantSocialLinkClient         pb.MerchantSocialCommandServiceClient
	OrderCommandClient               pb.OrderCommandServiceClient
	OrderQueryClient                 pb.OrderQueryServiceClient
	OrderStatsClient                 pb.OrderStatsServiceClient
	OrderItemCommandClient           pb.OrderItemCommandServiceClient
	OrderItemQueryClient             pb.OrderItemQueryServiceClient
	ProductCommandClient             pb.ProductCommandServiceClient
	ProductQueryClient               pb.ProductQueryServiceClient
	ReviewCommandClient              pb.ReviewCommandServiceClient
	ReviewQueryClient                pb.ReviewQueryServiceClient
	ReviewDetailCommandClient        pb.ReviewDetailCommandServiceClient
	ReviewDetailQueryClient          pb.ReviewDetailQueryServiceClient
	ShippingCommandClient            pb.ShippingCommandServiceClient
	ShippingQueryClient              pb.ShippingQueryServiceClient
	SliderCommandClient              pb.SliderCommandServiceClient
	SliderQueryClient                pb.SliderQueryServiceClient
	TransactionCommandClient         pb.TransactionCommandServiceClient
	TransactionQueryClient           pb.TransactionQueryServiceClient
	TransactionStatsClient           pb.TransactionStatsServiceClient
	TransactionStatsByMerchantClient pb.TransactionStatsByMerchantServiceClient
}

type Deps struct {
	Clients     *GRPCClients
	Logger      logger.LoggerInterface
	Mapping     *graphql.GraphqlMapper
	Cache       *cache.CacheStore
	ImageUpload upload_image.ImageUploads
}

func NewResolver(deps *Deps) *Resolver {
	obs, _ := observability.NewObservability(
		"graphql-client",
		deps.Logger,
	)

	resolver := NewResolverHandler(obs, deps.Logger)

	return &Resolver{
		AuthGraphql: &AuthHandleGraphql{
			AuthClient: deps.Clients.AuthClient,
			Mapping:    deps.Mapping.AuthGraphqlMapper,
			Logger:     deps.Logger,
			Cache:      auth_cache.NewMencache(deps.Cache),
		},
		RoleGraphql: &RoleHandleGraphql{
			RoleCommandClient: deps.Clients.RoleCommandClient,
			RoleQueryClient:   deps.Clients.RoleQueryClient,
			Mapping:           deps.Mapping.RoleGraphqlMapper,
			Logger:            deps.Logger,
			Cache:             role_cache.NewRoleMencache(deps.Cache),
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
			CategoryCommandClient:         deps.Clients.CategoryCommandClient,
			CategoryQueryClient:           deps.Clients.CategoryQueryClient,
			CategoryStatsClient:           deps.Clients.CategoryStatsClient,
			CategoryStatsByMerchantClient: deps.Clients.CategoryStatsByMerchantClient,
			CategoryStatsByIdClient:       deps.Clients.CategoryStatsByIdClient,
			Mapping:                       deps.Mapping.CategoryGraphqlMapper,
			Logger:                        deps.Logger,
			Cache:                         category_cache.NewCategoryMencache(deps.Cache),
			UploadImage:                   deps.ImageUpload,
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
			OrderStatsClient:   deps.Clients.OrderStatsClient,
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
			TransactionCommandClient:         deps.Clients.TransactionCommandClient,
			TransactionQueryClient:           deps.Clients.TransactionQueryClient,
			TransactionStatsClient:           deps.Clients.TransactionStatsClient,
			TransactionStatsByMerchantClient: deps.Clients.TransactionStatsByMerchantClient,
			Mapping:                          deps.Mapping.TransactionGraphqlMapper,
			Logger:                           deps.Logger,
			Cache:                            transaction_cache.NewTransactionMencache(deps.Cache),
		},
		ResolverHandle: resolver,
	}
}

type AuthHandleGraphql struct {
	AuthClient pb.AuthServiceClient
	Mapping    graphql.AuthGraphqlMapper
	Logger     logger.LoggerInterface
	Cache      auth_cache.AuthMencache
}

type RoleHandleGraphql struct {
	RoleCommandClient pb.RoleCommandServiceClient
	RoleQueryClient   pb.RoleQueryServiceClient
	Mapping           graphql.RoleGraphqlMapper
	Logger            logger.LoggerInterface
	Cache             role_cache.RoleMencache
}

type UserHandleGraphql struct {
	UserCommandClient pb.UserCommandServiceClient
	UserQueryClient   pb.UserQueryServiceClient
	Mapping           graphql.UserGraphqlMapper
	Logger            logger.LoggerInterface
	Cache             user_cache.UserMencache
}

type BannerHandleGraphql struct {
	BannerCommandClient pb.BannerCommandServiceClient
	BannerQueryClient   pb.BannerQueryServiceClient
	Mapping             graphql.BannerGraphqlMapper
	Logger              logger.LoggerInterface
	Cache               banner_cache.BannerMencache
}

type CartHandleGraphql struct {
	CartCommandClient pb.CartCommandServiceClient
	CartQueryClient   pb.CartQueryServiceClient
	Mapping           graphql.CartGraphqlMapper
	Logger            logger.LoggerInterface
	Cache             cart_cache.CartMencache
}

type CategoryHandleGraphql struct {
	CategoryCommandClient         pb.CategoryCommandServiceClient
	CategoryQueryClient           pb.CategoryQueryServiceClient
	CategoryStatsClient           pb.CategoryStatsServiceClient
	CategoryStatsByMerchantClient pb.CategoryStatsByMerchantServiceClient
	CategoryStatsByIdClient       pb.CategoryStatsByIdServiceClient
	Mapping                       graphql.CategoryGraphqlMapper
	UploadImage                   upload_image.ImageUploads
	Logger                        logger.LoggerInterface
	Cache                         category_cache.CategoryMencache
}

type MerchantHandleGraphql struct {
	MerchantCommandClient pb.MerchantCommandServiceClient
	MerchantQueryClient   pb.MerchantQueryServiceClient
	Mapping               graphql.MerchantGraphqlMapper
	Logger                logger.LoggerInterface
	Cache                 merchant_cache.MerchantMencache
}

type MerchantAwardHandleGraphql struct {
	MerchantAwardCommandClient pb.MerchantAwardCommandServiceClient
	MerchantAwardQueryClient   pb.MerchantAwardQueryServiceClient
	Mapping                    graphql.MerchantAwardGraphqlMapper
	Logger                     logger.LoggerInterface
	Cache                      merchantawards_cache.MerchantAwardMencache
}

type MerchantBusinessHandleGraphql struct {
	MerchantBusinessCommandClient pb.MerchantBusinessCommandServiceClient
	MerchantBusinessQueryClient   pb.MerchantBusinessQueryServiceClient
	Mapping                       graphql.MerchantBusinessGraphqlMapper
	Logger                        logger.LoggerInterface
	Cache                         merchantbusiness_cache.MerchantBusinessMencache
}

type MerchantDetailHandleGraphql struct {
	MerchantDetailCommandClient pb.MerchantDetailCommandServiceClient
	MerchantDetailQueryClient   pb.MerchantDetailQueryServiceClient
	Mapping                     graphql.MerchantDetailGraphqlMapper
	UploadImage                 upload_image.ImageUploads
	Logger                      logger.LoggerInterface
	Cache                       merchantdetail_cache.MerchantDetailMencache
}

type MerchantPolicyHandleGraphql struct {
	MerchantPolicyCommandClient pb.MerchantPolicyCommandServiceClient
	MerchantPolicyQueryClient   pb.MerchantPolicyQueryServiceClient
	Mapping                     graphql.MerchantPolicyGraphqlMapper
	Logger                      logger.LoggerInterface
	Cache                       merchantpolicies_cache.MerchantPoliciesMencache
}

type MerchantSocialLinkHandleGraphql struct {
	MerchantSocialLinkClient pb.MerchantSocialCommandServiceClient
	Mapping                  graphql.MerchantSocialLinkGraphqlMapper
	Logger                   logger.LoggerInterface
}

type OrderHandleGraphql struct {
	OrderCommandClient pb.OrderCommandServiceClient
	OrderQueryClient   pb.OrderQueryServiceClient
	OrderStatsClient   pb.OrderStatsServiceClient
	Mapping            graphql.OrderGraphqlMapper
	Logger             logger.LoggerInterface
	Cache              order_cache.OrderMencache
}

type OrderItemHandleGraphql struct {
	OrderItemCommandClient pb.OrderItemCommandServiceClient
	OrderItemQueryClient   pb.OrderItemQueryServiceClient
	Mapping                graphql.OrderItemGraphqlMapper
	Logger                 logger.LoggerInterface
	Cache                  orderitem_cache.OrderItemMencache
}

type ProductHandleGraphql struct {
	ProductCommandClient pb.ProductCommandServiceClient
	ProductQueryClient   pb.ProductQueryServiceClient
	Mapping              graphql.ProductGraphqlMapper
	UploadImage          upload_image.ImageUploads
	Logger               logger.LoggerInterface
	Cache                product_cache.ProductMencache
}

type ReviewHandleGraphql struct {
	ReviewCommandClient pb.ReviewCommandServiceClient
	ReviewQueryClient   pb.ReviewQueryServiceClient
	Mapping             graphql.ReviewGraphqlMapper
	Logger              logger.LoggerInterface
	Cache               review_cache.ReviewMencache
}

type ReviewDetailHandleGraphql struct {
	ReviewDetailCommandClient pb.ReviewDetailCommandServiceClient
	ReviewDetailQueryClient   pb.ReviewDetailQueryServiceClient
	Mapping                   graphql.ReviewDetailGraphqlMapper
	Logger                    logger.LoggerInterface
	Cache                     reviewdetail_cache.ReviewDetailMencache
}

type ShippingAddressHandleGraphql struct {
	ShippingCommandClient pb.ShippingCommandServiceClient
	ShippingQueryClient   pb.ShippingQueryServiceClient
	Mapping               graphql.ShippingAddresGraphqlMapper
	Logger                logger.LoggerInterface
	Cache                 shippingaddress_cache.ShippingAddressMencache
}

type SliderHandleGraphql struct {
	SliderCommandClient pb.SliderCommandServiceClient
	SliderQueryClient   pb.SliderQueryServiceClient
	Mapping             graphql.SliderGraphqlMapper
	UploadImage         upload_image.ImageUploads
	Logger              logger.LoggerInterface
	Cache               slider_cache.SliderMencache
}

type TransactionHandleGraphql struct {
	TransactionCommandClient         pb.TransactionCommandServiceClient
	TransactionQueryClient           pb.TransactionQueryServiceClient
	TransactionStatsClient           pb.TransactionStatsServiceClient
	TransactionStatsByMerchantClient pb.TransactionStatsByMerchantServiceClient
	Mapping                          graphql.TransactionGraphqlMapper
	Logger                           logger.LoggerInterface
	Cache                            transaction_cache.TransactionMencache
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
