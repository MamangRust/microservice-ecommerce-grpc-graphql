package shipping_address_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type ShippingAddressGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler           http.Handler
	orderID           int
	shippingAddressID int
}

func (s *ShippingAddressGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupCategoryService()
	s.SetupMerchantService()
	s.SetupProductService()
	s.SetupOrderItemService()
	s.SetupShippingAddressService()
	s.SetupOrderService()

	// Seed dependencies. SeedOrder also creates the order's shipping address,
	// which is what this suite exercises (shipping addresses have no create
	// mutation of their own).
	ctx := context.Background()
	userID := s.SeedUser(ctx)
	catID := s.SeedCategory(ctx)
	merchID := s.SeedMerchant(ctx, userID)
	prodID := s.SeedProduct(ctx, merchID, catID)
	s.orderID = s.SeedOrder(ctx, userID, merchID, prodID)

	s.handler = s.GraphQLHandler()
}

func (s *ShippingAddressGraphqlTestSuite) TestShippingAddressGraphqlLifecycle() {
	// 1. FindAll
	all := s.GQL(s.handler, `query FindAllShipping($input: FindAllShippingRequest!) {
		findAllShipping(input: $input) { status message pagination { total_records } data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllShipping"), "data"))

	// 2. FindByOrder
	s.Require().NotZero(s.orderID)
	byOrder := s.GQL(s.handler, `query FindShippingByOrder($input: FindByIdShippingRequest!) {
		findShippingByOrder(input: $input) { status message data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.orderID}})
	byOrderData := s.Obj(s.Obj(byOrder, "findShippingByOrder"), "data")
	s.shippingAddressID = int(byOrderData["id"].(float64))
	s.Require().NotZero(s.shippingAddressID)
	s.Equal(float64(s.orderID), byOrderData["order_id"])

	// 3. FindById
	byID := s.GQL(s.handler, `query FindShippingById($input: FindByIdShippingRequest!) {
		findShippingById(input: $input) { status message data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.shippingAddressID}})
	s.Equal(float64(s.shippingAddressID), s.Obj(s.Obj(byID, "findShippingById"), "data")["id"])

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveShipping($input: FindAllShippingRequest!) {
		findActiveShipping(input: $input) { status message pagination { total_records } data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveShipping"), "data"))

	// 5. Trash
	s.GQL(s.handler, `mutation TrashedShipping($input: FindByIdShippingRequest!) {
		trashedShipping(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.shippingAddressID}})

	// 6. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedShipping($input: FindAllShippingRequest!) {
		findTrashedShipping(input: $input) { status message pagination { total_records } data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedShipping"), "data"))

	// 7. Restore
	s.GQL(s.handler, `mutation RestoreShipping($input: FindByIdShippingRequest!) {
		restoreShipping(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.shippingAddressID}})

	// 8. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashedShipping($input: FindByIdShippingRequest!) {
		trashedShipping(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.shippingAddressID}})
	s.GQL(s.handler, `mutation DeleteShippingPermanent($input: FindByIdShippingRequest!) {
		deleteShippingPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.shippingAddressID}})

	// 9. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllShipping { status message } }`, nil)

	// 10. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllShippingPermanent { status message } }`, nil)
}

func TestShippingAddressGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ShippingAddressGraphqlTestSuite))
}
