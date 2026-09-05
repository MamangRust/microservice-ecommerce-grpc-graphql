package graphqlmapper

import (
	auth "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/auth"
	banner "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/banner"
	cart "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/cart"
	category "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/category"
	merchant "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/merchant"
	merchantaward "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/merchant_award"
	merchantbusiness "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/merchant_business"
	merchantdetail "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/merchant_detail"
	merchantpolicy "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/merchant_policy"
	merchantsociallink "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/merchant_sociallink"
	order "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/order"
	orderitem "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/order_item"
	product "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/product"
	review "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/review"
	reviewdetail "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/review_detail"
	role "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/role"
	shippingaddress "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/shipping_address"
	slider "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/slider"
	transaction "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/transaction"
	user "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/user"
)

type AuthGraphqlMapper = auth.AuthGraphqlMapper
type RoleGraphqlMapper = role.RoleGraphqlMapper
type UserGraphqlMapper = user.UserGraphqlMapper
type BannerGraphqlMapper = banner.BannerGraphqlMapper
type CartGraphqlMapper = cart.CartGraphqlMapper
type CategoryGraphqlMapper = category.CategoryGraphqlMapper
type MerchantGraphqlMapper = merchant.MerchantGraphqlMapper
type MerchantAwardGraphqlMapper = merchantaward.MerchantAwardGraphqlMapper
type MerchantBusinessGraphqlMapper = merchantbusiness.MerchantBusinessGraphqlMapper
type MerchantDetailGraphqlMapper = merchantdetail.MerchantDetailGraphqlMapper
type MerchantPolicyGraphqlMapper = merchantpolicy.MerchantPolicyGraphqlMapper
type MerchantSocialLinkGraphqlMapper = merchantsociallink.MerchantSocialLinkGraphqlMapper
type OrderGraphqlMapper = order.OrderGraphqlMapper
type OrderItemGraphqlMapper = orderitem.OrderItemGraphqlMapper
type ProductGraphqlMapper = product.ProductGraphqlMapper
type ReviewGraphqlMapper = review.ReviewGraphqlMapper
type ReviewDetailGraphqlMapper = reviewdetail.ReviewDetailGraphqlMapper
type ShippingAddresGraphqlMapper = shippingaddress.ShippingAddresGraphqlMapper
type SliderGraphqlMapper = slider.SliderGraphqlMapper
type TransactionGraphqlMapper = transaction.TransactionGraphqlMapper

type GraphqlMapper struct {
	AuthGraphqlMapper               AuthGraphqlMapper
	RoleGraphqlMapper               RoleGraphqlMapper
	UserGraphqlMapper               UserGraphqlMapper
	BannerGraphqlMapper             BannerGraphqlMapper
	CartGraphqlMapper               CartGraphqlMapper
	CategoryGraphqlMapper           CategoryGraphqlMapper
	MerchantGraphqlMapper           MerchantGraphqlMapper
	MerchantAwardGraphqlMapper      MerchantAwardGraphqlMapper
	MerchantBusinessGraphqlMapper   MerchantBusinessGraphqlMapper
	MerchantDetailGraphqlMapper     MerchantDetailGraphqlMapper
	MerchantPolicyGraphqlMapper     MerchantPolicyGraphqlMapper
	MerchantSocialLinkGraphqlMapper MerchantSocialLinkGraphqlMapper
	OrderGraphqlMapper              OrderGraphqlMapper
	OrderItemGraphqlMapper          OrderItemGraphqlMapper
	ProductGraphqlMapper            ProductGraphqlMapper
	ReviewGraphqlMapper             ReviewGraphqlMapper
	ReviewDetailGraphqlMapper       ReviewDetailGraphqlMapper
	ShippingAddresGraphqlMapper     ShippingAddresGraphqlMapper
	SliderGraphqlMapper             SliderGraphqlMapper
	TransactionGraphqlMapper        TransactionGraphqlMapper
}

func NewGraphqlMapper() *GraphqlMapper {
	return &GraphqlMapper{
		AuthGraphqlMapper:               auth.NewAuthGraphqlMapper(),
		RoleGraphqlMapper:               role.NewRoleGraphqlMapper(),
		UserGraphqlMapper:               user.NewUserGraphqlMapper(),
		BannerGraphqlMapper:             banner.NewBannerResponseMapper(),
		CartGraphqlMapper:               cart.NewCartResponseMapper(),
		CategoryGraphqlMapper:           category.NewCategoryGraphqlMapper(),
		MerchantGraphqlMapper:           merchant.NewMerchantResponseMapper(),
		MerchantAwardGraphqlMapper:      merchantaward.NewMerchantAwardResponseMapper(),
		MerchantBusinessGraphqlMapper:   merchantbusiness.NewMerchantBusinessResponseMapper(),
		MerchantDetailGraphqlMapper:     merchantdetail.NewMerchantDetailResponseMapper(),
		MerchantPolicyGraphqlMapper:     merchantpolicy.NewMerchantPolicyResponseMapper(),
		MerchantSocialLinkGraphqlMapper: merchantsociallink.NewMerchantSocialLinkResponseMapper(),
		OrderGraphqlMapper:              order.NewOrderGraphqlMapper(),
		OrderItemGraphqlMapper:          orderitem.NewOrderItemGraphqlMapper(),
		ProductGraphqlMapper:            product.NewProductGraphqlMapper(),
		ReviewGraphqlMapper:             review.NewReviewResponseMapper(),
		ReviewDetailGraphqlMapper:       reviewdetail.NewReviewDetailResponseMapper(),
		ShippingAddresGraphqlMapper:     shippingaddress.NewshippingAddresResponseMapper(),
		SliderGraphqlMapper:             slider.NewSliderResponseMapper(),
		TransactionGraphqlMapper:        transaction.NewTransactionResponseMapper(),
	}
}
