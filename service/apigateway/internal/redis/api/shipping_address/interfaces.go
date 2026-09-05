package shippingaddress_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type ShippingAddressQueryCache interface {
	GetShippingAddressAllCache(ctx context.Context, req *model.FindAllShippingRequest) (*model.APIResponsePaginationShipping, bool)
	SetShippingAddressAllCache(ctx context.Context, req *model.FindAllShippingRequest, data *model.APIResponsePaginationShipping)

	GetShippingAddressActiveCache(ctx context.Context, req *model.FindAllShippingRequest) (*model.APIResponsePaginationShippingDeleteAt, bool)
	SetShippingAddressActiveCache(ctx context.Context, req *model.FindAllShippingRequest, data *model.APIResponsePaginationShippingDeleteAt)

	GetShippingAddressTrashedCache(ctx context.Context, req *model.FindAllShippingRequest) (*model.APIResponsePaginationShippingDeleteAt, bool)
	SetShippingAddressTrashedCache(ctx context.Context, req *model.FindAllShippingRequest, data *model.APIResponsePaginationShippingDeleteAt)

	GetCachedShippingAddressCache(ctx context.Context, shipping_id int) (*model.APIResponseShipping, bool)
	SetCachedShippingAddressCache(ctx context.Context, data *model.APIResponseShipping)

	GetCachedShippingAddressByOrderCache(ctx context.Context, order_id int) (*model.APIResponseShipping, bool)
	SetCachedShippingAddressByOrderCache(ctx context.Context, data *model.APIResponseShipping)
}

type ShippingAddressCommandCache interface {
	DeleteShippingAddressCache(ctx context.Context, shipping_id int)
}
