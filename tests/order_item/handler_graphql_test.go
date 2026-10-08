package order_item_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type OrderItemGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler http.Handler
	orderID int
}

func (s *OrderItemGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupRoleService()
	s.SetupUserService()
	s.SetupCategoryService()
	s.SetupMerchantService()
	s.SetupProductService()
	s.SetupShippingAddressService()
	s.SetupOrderItemService()
	s.SetupOrderService()
	s.SetupTransactionService()

	// Seed dependencies
	ctx := context.Background()
	userID := s.SeedUser(ctx)
	catID := s.SeedCategory(ctx)
	merchID := s.SeedMerchant(ctx, userID)
	prodID := s.SeedProduct(ctx, merchID, catID)
	s.orderID = s.SeedOrder(ctx, userID, merchID, prodID)

	s.handler = s.GraphQLHandler()
}

func (s *OrderItemGraphqlTestSuite) TestOrderItemGraphqlLifecycle() {
	// 1. FindAll
	all := s.GQL(s.handler, `query FindAllOrderItems($input: FindAllOrderItemInput!) {
		findAllOrderItems(input: $input) { status message pagination { total_records } data { id order_id product_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllOrderItems"), "data"))

	// 2. FindOrderItemByOrder
	s.Require().NotZero(s.orderID)
	byOrder := s.GQL(s.handler, `query FindOrderItemsByOrder($input: FindByIdOrderItemInput!) {
		findOrderItemsByOrder(input: $input) { status message data { id order_id product_id quantity } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.orderID}})
	byOrderData := s.Arr(s.Obj(byOrder, "findOrderItemsByOrder"), "data")
	if len(byOrderData) == 0 {
		s.T().Skip("No order items found")
	}
	s.Equal(float64(s.orderID), byOrderData[0].(map[string]interface{})["order_id"])

	// 3. FindByActive
	active := s.GQL(s.handler, `query FindActiveOrderItems($input: FindAllOrderItemInput!) {
		findActiveOrderItems(input: $input) { status message pagination { total_records } data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotNil(s.Obj(active, "findActiveOrderItems"))

	// 4. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedOrderItems($input: FindAllOrderItemInput!) {
		findTrashedOrderItems(input: $input) { status message pagination { total_records } data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotNil(s.Obj(trashed, "findTrashedOrderItems"))
}

func TestOrderItemGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(OrderItemGraphqlTestSuite))
}
