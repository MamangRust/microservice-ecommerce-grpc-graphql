package cart_test

import (
	"net/http"
	"testing"

	gt "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/graphtest"
	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type CartGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler http.Handler
	userID  int
	cartIDs []int
}

func (s *CartGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()

	// Bootstrap all required downstream services
	s.SetupUserService()
	s.SetupCategoryService()
	s.SetupMerchantService()
	s.SetupProductService()
	s.SetupCartService()

	// Seed a user to act as the authenticated caller
	s.userID = s.SeedUser(s.Ctx)

	// deleteAllCarts resolves the caller from the request context (populated by
	// the production AuthMiddleware), so wrap the handler to emulate an
	// authenticated user.
	s.handler = gt.WithUser(s.GraphQLHandler(), s.userID)

	s.cartIDs = make([]int, 0)
}

func (s *CartGraphqlTestSuite) TestCartGraphqlLifecycle() {
	ctx := s.Ctx

	// Seed entities needed for cart items
	categoryID := s.SeedCategory(ctx)
	merchantID := s.SeedMerchant(ctx, s.userID)
	prod1ID := s.SeedProduct(ctx, merchantID, categoryID)
	prod2ID := s.SeedProduct(ctx, merchantID, categoryID)

	// 1. Create — add first product to cart
	create := s.GQL(s.handler, `mutation CreateCart($input: CreateCartInput!) {
		createCart(input: $input) { status message data { id user_id product_id quantity } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"user_id":    s.userID,
			"product_id": prod1ID,
			"quantity":   2,
		},
	})
	createData := s.Obj(s.Obj(create, "createCart"), "data")
	cart1ID := int(createData["id"].(float64))
	s.cartIDs = append(s.cartIDs, cart1ID)
	s.Equal(float64(prod1ID), createData["product_id"])
	s.Equal(float64(2), createData["quantity"])

	// 2. Create — add second product to cart
	create2 := s.GQL(s.handler, `mutation CreateCart($input: CreateCartInput!) {
		createCart(input: $input) { status message data { id product_id quantity } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"user_id":    s.userID,
			"product_id": prod2ID,
			"quantity":   1,
		},
	})
	create2Data := s.Obj(s.Obj(create2, "createCart"), "data")
	cart2ID := int(create2Data["id"].(float64))
	s.cartIDs = append(s.cartIDs, cart2ID)

	// 3. FindAll — list cart items for the user
	all := s.GQL(s.handler, `query FindAllCarts($input: FindAllCartInput!) {
		findAllCarts(input: $input) { status message data { id product_id quantity } pagination { total_records } }
	}`, map[string]interface{}{"input": map[string]interface{}{
		"user_id": s.userID, "page": 1, "page_size": 10,
	}})
	s.GreaterOrEqual(len(s.Arr(s.Obj(all, "findAllCarts"), "data")), 2, "should have at least 2 cart items")

	// 4. Delete — remove the first cart item
	s.GQL(s.handler, `mutation DeleteCart($input: DeleteCartInput!) {
		deleteCart(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{
		"cart_id": cart1ID, "user_id": s.userID,
	}})
	s.cartIDs = s.cartIDs[1:] // remove from tracking

	// 5. DeleteAll — remove remaining cart items
	s.GQL(s.handler, `mutation DeleteAllCarts($input: DeleteCartsInput!) {
		deleteAllCarts(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"cart_ids": s.cartIDs}})

	// 6. Verify — FindAll should now be empty. The gateway list cache is not
	// invalidated by the cart mutations, so drop it before re-reading.
	_, _ = s.GetCacheStore().InvalidateCache(s.Ctx, "cart:all:*")
	verify := s.GQL(s.handler, `query FindAllCarts($input: FindAllCartInput!) {
		findAllCarts(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{
		"user_id": s.userID, "page": 1, "page_size": 10,
	}})
	s.Empty(s.Arr(s.Obj(verify, "findAllCarts"), "data"), "cart should be empty after deleting all items")
}

func TestCartGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CartGraphqlTestSuite))
}
