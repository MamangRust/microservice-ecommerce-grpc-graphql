package merchant_award_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type MerchantAwardGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler    http.Handler
	awardID    int
	merchantID int
}

func (s *MerchantAwardGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupMerchantService()
	s.SetupMerchantAwardService()

	ctx := context.Background()
	userID := s.SeedUser(ctx)
	s.merchantID = s.SeedMerchant(ctx, userID)

	s.handler = s.GraphQLHandler()
}

func (s *MerchantAwardGraphqlTestSuite) TestMerchantAwardGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateMerchantAward($input: CreateMerchantAwardInput!) {
		createMerchantAward(input: $input) { status message data { id merchant_id title } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_id":     s.merchantID,
			"title":           "Best Merchant 2024",
			"description":     "Award for excellence",
			"issued_by":       "E-commerce Platform",
			"issue_date":      "2024-01-01",
			"expiry_date":     "2025-01-01",
			"certificate_url": "http://example.com/cert.pdf",
		},
	})
	createData := s.Obj(s.Obj(create, "createMerchantAward"), "data")
	s.awardID = int(createData["id"].(float64))
	s.Equal("Best Merchant 2024", createData["title"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindMerchantAwardById($input: FindByIdMerchantAwardInput!) {
		findMerchantAwardById(input: $input) { status message data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.awardID}})
	s.Equal(float64(s.awardID), s.Obj(s.Obj(byID, "findMerchantAwardById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllMerchantAwards($input: FindAllMerchantAwardInput!) {
		findAllMerchantAwards(input: $input) { status message pagination { total_records } data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllMerchantAwards"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveMerchantAwards($input: FindAllMerchantAwardInput!) {
		findActiveMerchantAwards(input: $input) { status message pagination { total_records } data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveMerchantAwards"), "data"))

	// 5. Update
	updated := s.GQL(s.handler, `mutation UpdateMerchantAward($input: UpdateMerchantAwardInput!) {
		updateMerchantAward(input: $input) { status message data { id title } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_certification_id": s.awardID,
			"title":                     "Updated Award Title",
			"description":               "Updated Description",
			"issued_by":                 "Updated Issuer",
			"issue_date":                "2024-02-01",
			"expiry_date":               "2025-02-01",
			"certificate_url":           "http://example.com/updated.pdf",
		},
	})
	s.Equal("Updated Award Title", s.Obj(s.Obj(updated, "updateMerchantAward"), "data")["title"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashMerchantAward($input: FindByIdMerchantAwardInput!) {
		trashMerchantAward(input: $input) { status message data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.awardID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedMerchantAwards($input: FindAllMerchantAwardInput!) {
		findTrashedMerchantAwards(input: $input) { status message pagination { total_records } data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedMerchantAwards"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreMerchantAward($input: FindByIdMerchantAwardInput!) {
		restoreMerchantAward(input: $input) { status message data { id title } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.awardID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashMerchantAward($input: FindByIdMerchantAwardInput!) {
		trashMerchantAward(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.awardID}})
	s.GQL(s.handler, `mutation DeleteMerchantAwardPermanent($input: FindByIdMerchantAwardInput!) {
		deleteMerchantAwardPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.awardID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllMerchantAwards { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllMerchantAwardsPermanent { status message } }`, nil)
}

func TestMerchantAwardGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantAwardGraphqlTestSuite))
}
