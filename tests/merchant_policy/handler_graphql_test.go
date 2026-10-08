package merchant_policy_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type MerchantPolicyGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler    http.Handler
	policyID   int
	merchantID int
	userID     int
}

func (s *MerchantPolicyGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupMerchantService()
	s.SetupMerchantPolicyService()

	ctx := context.Background()
	s.userID = s.SeedUser(ctx)
	s.merchantID = s.SeedMerchant(ctx, s.userID)

	s.handler = s.GraphQLHandler()
}

func (s *MerchantPolicyGraphqlTestSuite) TestMerchantPolicyGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateMerchantPolicy($input: CreateMerchantPoliciesInput!) {
		createMerchantPolicy(input: $input) { status message data { id merchant_id title } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_id": s.merchantID,
			"policy_type": "Return",
			"title":       "Return Policy",
			"description": "30-day returns accepted",
		},
	})
	createData := s.Obj(s.Obj(create, "createMerchantPolicy"), "data")
	s.policyID = int(createData["id"].(float64))
	s.Equal("Return Policy", createData["title"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindMerchantPolicyById($input: FindByIdMerchantPoliciesInput!) {
		findMerchantPolicyById(input: $input) { status message data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.policyID}})
	s.Equal(float64(s.policyID), s.Obj(s.Obj(byID, "findMerchantPolicyById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllMerchantPolicies($input: FindAllMerchantPoliciesInput!) {
		findAllMerchantPolicies(input: $input) { status message pagination { total_records } data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllMerchantPolicies"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveMerchantPolicies($input: FindAllMerchantPoliciesInput!) {
		findActiveMerchantPolicies(input: $input) { status message pagination { total_records } data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveMerchantPolicies"), "data"))

	// 5. Update
	updated := s.GQL(s.handler, `mutation UpdateMerchantPolicy($input: UpdateMerchantPoliciesInput!) {
		updateMerchantPolicy(input: $input) { status message data { id title } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_policy_id": s.policyID,
			"policy_type":        "Shipping",
			"title":              "Shipping Policy",
			"description":        "Free shipping over $50",
		},
	})
	s.Equal("Shipping Policy", s.Obj(s.Obj(updated, "updateMerchantPolicy"), "data")["title"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashMerchantPolicy($input: FindByIdMerchantPoliciesInput!) {
		trashMerchantPolicy(input: $input) { status message data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.policyID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedMerchantPolicies($input: FindAllMerchantPoliciesInput!) {
		findTrashedMerchantPolicies(input: $input) { status message pagination { total_records } data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedMerchantPolicies"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreMerchantPolicy($input: FindByIdMerchantPoliciesInput!) {
		restoreMerchantPolicy(input: $input) { status message data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.policyID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashMerchantPolicy($input: FindByIdMerchantPoliciesInput!) {
		trashMerchantPolicy(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.policyID}})
	s.GQL(s.handler, `mutation DeleteMerchantPolicyPermanent($input: FindByIdMerchantPoliciesInput!) {
		deleteMerchantPolicyPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.policyID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllMerchantPolicies { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllMerchantPoliciesPermanent { status message } }`, nil)
}

func TestMerchantPolicyGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantPolicyGraphqlTestSuite))
}
