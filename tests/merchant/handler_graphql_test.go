package merchant_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type MerchantGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler    http.Handler
	merchantID int
	userID     int
}

func (s *MerchantGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupMerchantService()

	s.userID = s.SeedUser(context.Background())

	s.handler = s.GraphQLHandler()
}

func (s *MerchantGraphqlTestSuite) TestMerchantGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateMerchant($input: CreateMerchantInput!) {
		createMerchant(input: $input) { status message data { id user_id name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"user_id":       s.userID,
			"name":          "Test Merchant",
			"description":   "Test Description",
			"address":       "Test Address",
			"contact_email": "merchant@example.com",
			"contact_phone": "123456789",
			"status":        "active",
		},
	})
	createData := s.Obj(s.Obj(create, "createMerchant"), "data")
	s.merchantID = int(createData["id"].(float64))
	s.Equal("Test Merchant", createData["name"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindMerchantById($input: FindByIdMerchantInput!) {
		findMerchantById(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.merchantID}})
	s.Equal(float64(s.merchantID), s.Obj(s.Obj(byID, "findMerchantById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllMerchants($input: FindAllMerchantInput!) {
		findAllMerchants(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllMerchants"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveMerchants($input: FindAllMerchantInput!) {
		findActiveMerchants(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveMerchants"), "data"))

	// 5. Update
	updated := s.GQL(s.handler, `mutation UpdateMerchant($input: UpdateMerchantInput!) {
		updateMerchant(input: $input) { status message data { id name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_id":   s.merchantID,
			"user_id":       s.userID,
			"name":          "Updated Merchant",
			"description":   "Updated Description",
			"address":       "Updated Address",
			"contact_email": "updated@example.com",
			"contact_phone": "987654321",
			"status":        "active",
		},
	})
	s.Equal("Updated Merchant", s.Obj(s.Obj(updated, "updateMerchant"), "data")["name"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashMerchant($input: FindByIdMerchantInput!) {
		trashMerchant(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.merchantID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedMerchants($input: FindAllMerchantInput!) {
		findTrashedMerchants(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedMerchants"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreMerchant($input: FindByIdMerchantInput!) {
		restoreMerchant(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.merchantID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashMerchant($input: FindByIdMerchantInput!) {
		trashMerchant(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.merchantID}})
	s.GQL(s.handler, `mutation DeleteMerchantPermanent($input: FindByIdMerchantInput!) {
		deleteMerchantPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.merchantID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllMerchants { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllMerchantsPermanent { status message } }`, nil)
}

func TestMerchantGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantGraphqlTestSuite))
}
