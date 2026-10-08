package order_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type OrderGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler http.Handler
	orderID int
	userID  int
	merchID int
	prodID  int
}

func (s *OrderGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupMerchantService()
	s.SetupCategoryService()
	s.SetupProductService()
	s.SetupOrderItemService()
	s.SetupShippingAddressService()
	s.SetupTransactionService()
	s.SetupOrderService()

	// Seed dependencies
	ctx := context.Background()
	userID := s.SeedUser(ctx)
	catID := s.SeedCategory(ctx)
	merchID := s.SeedMerchant(ctx, userID)
	prodID := s.SeedProduct(ctx, merchID, catID)

	s.userID = userID
	s.merchID = merchID
	s.prodID = prodID

	s.handler = s.GraphQLHandler()
}

func (s *OrderGraphqlTestSuite) TestOrderGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateOrder($input: CreateOrderInput!) {
		createOrder(input: $input) { status message data { id merchant_id user_id total_price } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_id": s.merchID,
			"user_id":     s.userID,
			"total_price": 1000,
			"items": []interface{}{
				map[string]interface{}{"product_id": s.prodID, "quantity": 1, "price": 1000},
			},
			"shipping": map[string]interface{}{
				"alamat":          "Test Alamat",
				"provinsi":        "Test Provinsi",
				"kota":            "Test Kota",
				"courier":         "Test Courier",
				"shipping_method": "Test Method",
				"shipping_cost":   100,
				"negara":          "Test Negara",
			},
		},
	})
	createData := s.Obj(s.Obj(create, "createOrder"), "data")
	s.orderID = int(createData["id"].(float64))
	s.Require().NotZero(s.orderID)

	// 2. FindById
	byID := s.GQL(s.handler, `query FindOrderById($input: FindByIdOrderInput!) {
		findOrderById(input: $input) { status message data { id merchant_id total_price } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.orderID}})
	s.Equal(float64(s.orderID), s.Obj(s.Obj(byID, "findOrderById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllOrders($input: FindAllOrderInput!) {
		findAllOrders(input: $input) { status message pagination { total_records } data { id total_price } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllOrders"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveOrders($input: FindAllOrderInput!) {
		findActiveOrders(input: $input) { status message pagination { total_records } data { id total_price } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveOrders"), "data"))

	// 5. Update (order_item_id/shipping_id 0 mirror the old REST body, which
	// omitted them: the service creates a new item and falls back to the
	// order's existing shipping address).
	updated := s.GQL(s.handler, `mutation UpdateOrder($input: UpdateOrderInput!) {
		updateOrder(input: $input) { status message data { id total_price } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"order_id":    s.orderID,
			"user_id":     s.userID,
			"total_price": 1500,
			"items": []interface{}{
				map[string]interface{}{"order_item_id": 0, "product_id": s.prodID, "quantity": 1, "price": 1500},
			},
			"shipping": map[string]interface{}{
				"shipping_id":     0,
				"alamat":          "Updated Alamat",
				"provinsi":        "Updated Provinsi",
				"kota":            "Updated Kota",
				"courier":         "Updated Courier",
				"shipping_method": "Updated Method",
				"shipping_cost":   200,
				"negara":          "Updated Negara",
			},
		},
	})
	s.Obj(s.Obj(updated, "updateOrder"), "data")

	// 6. Trash
	s.GQL(s.handler, `mutation TrashOrder($input: FindByIdOrderInput!) {
		trashOrder(input: $input) { status message data { id total_price } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.orderID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedOrders($input: FindAllOrderInput!) {
		findTrashedOrders(input: $input) { status message pagination { total_records } data { id total_price } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedOrders"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreOrder($input: FindByIdOrderInput!) {
		restoreOrder(input: $input) { status message data { id total_price } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.orderID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashOrder($input: FindByIdOrderInput!) {
		trashOrder(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.orderID}})
	s.GQL(s.handler, `mutation DeleteOrderPermanent($input: FindByIdOrderInput!) {
		deleteOrderPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.orderID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllOrders { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllOrdersPermanent { status message } }`, nil)
}

func TestOrderGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(OrderGraphqlTestSuite))
}
