package transaction_test

import (
	"context"
	"net/http"
	"testing"

	gt "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/graphtest"
	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type TransactionGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler       http.Handler
	transactionID int
	userID        int
	merchID       int
	orderID       int
}

func (s *TransactionGraphqlTestSuite) SetupSuite() {
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
	merchID := s.SeedMerchant(ctx, userID)
	catID := s.SeedCategory(ctx)
	prodID := s.SeedProduct(ctx, merchID, catID)
	orderID := s.SeedOrder(ctx, userID, merchID, prodID)
	s.SeedShippingAddress(ctx, orderID)
	s.SeedOrderItem(ctx, orderID, prodID)

	s.userID = userID
	s.merchID = merchID
	s.orderID = orderID

	// createTransaction derives the caller from the request context (populated
	// by AuthMiddleware in the real gateway), so emulate an authenticated user.
	s.handler = gt.WithUser(s.GraphQLHandler(), s.userID)
}

func (s *TransactionGraphqlTestSuite) TestTransactionGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateTransaction($input: CreateTransactionRequest!) {
		createTransaction(input: $input) { status message data { id order_id merchant_id payment_method amount } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"order_id":       s.orderID,
			"merchant_id":    s.merchID,
			"payment_method": "credit_card",
			"amount":         100000,
		},
	})
	createData := s.Obj(s.Obj(create, "createTransaction"), "data")
	s.transactionID = int(createData["id"].(float64))
	s.Require().NotZero(s.transactionID)

	// 2. FindAll
	all := s.GQL(s.handler, `query FindAllTransaction($input: FindAllTransactionRequest!) {
		findAllTransaction(input: $input) { status message pagination { total_records } data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllTransaction"), "data"))

	// 3. FindById
	byID := s.GQL(s.handler, `query FindTransactionById($input: FindByIdTransactionRequest!) {
		findTransactionById(input: $input) { status message data { id order_id merchant_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.transactionID}})
	s.Equal(float64(s.transactionID), s.Obj(s.Obj(byID, "findTransactionById"), "data")["id"])

	// 4. FindByMerchant
	byMerchant := s.GQL(s.handler, `query FindTransactionByMerchant($input: FindAllTransactionMerchantRequest!) {
		findTransactionByMerchant(input: $input) { status message pagination { total_records } data { id merchant_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"merchant_id": s.merchID, "page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(byMerchant, "findTransactionByMerchant"), "data"))

	// 5. FindByActive
	active := s.GQL(s.handler, `query FindByActive($input: FindAllTransactionRequest!) {
		findByActive(input: $input) { status message pagination { total_records } data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findByActive"), "data"))

	// 6. Update
	updated := s.GQL(s.handler, `mutation UpdateTransaction($input: UpdateTransactionRequest!) {
		updateTransaction(input: $input) { status message data { id payment_method amount } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"transaction_id": s.transactionID,
			"order_id":       s.orderID,
			"merchant_id":    s.merchID,
			"payment_method": "bank_transfer",
			"amount":         120000,
		},
	})
	s.Equal("bank_transfer", s.Obj(s.Obj(updated, "updateTransaction"), "data")["payment_method"])

	// 7. Trash
	s.GQL(s.handler, `mutation TrashedTransaction($input: FindByIdTransactionRequest!) {
		trashedTransaction(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.transactionID}})

	// 8. FindByTrashed
	trashed := s.GQL(s.handler, `query FindByTrashed($input: FindAllTransactionRequest!) {
		findByTrashed(input: $input) { status message pagination { total_records } data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findByTrashed"), "data"))

	// 9. Restore
	s.GQL(s.handler, `mutation RestoreTransaction($input: FindByIdTransactionRequest!) {
		restoreTransaction(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.transactionID}})

	// 10. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashedTransaction($input: FindByIdTransactionRequest!) {
		trashedTransaction(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.transactionID}})
	s.GQL(s.handler, `mutation DeleteTransactionPermanent($input: FindByIdTransactionRequest!) {
		deleteTransactionPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.transactionID}})

	// 11. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllTransaction { status message } }`, nil)

	// 12. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllTransactionPermanent { status message } }`, nil)
}

func TestTransactionGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransactionGraphqlTestSuite))
}
