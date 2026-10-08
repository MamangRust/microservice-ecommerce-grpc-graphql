package merchant_detail_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type MerchantDetailGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler    http.Handler
	detailID   int
	merchantID int
	userID     int
}

func (s *MerchantDetailGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupMerchantService()
	s.SetupMerchantDetailService()

	// Seed dependencies
	ctx := context.Background()
	s.userID = s.SeedUser(ctx)
	s.merchantID = s.SeedMerchant(ctx, s.userID)

	s.handler = s.GraphQLHandler()
}

func (s *MerchantDetailGraphqlTestSuite) TestMerchantDetailGraphqlLifecycle() {
	// 1. Create
	create := s.GQLMultipart(s.handler, `mutation CreateMerchantDetail($input: CreateMerchantDetailInput!) {
		createMerchantDetail(input: $input) { status message data { id merchant_id display_name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_id":       s.merchantID,
			"display_name":      "Test Merchant",
			"cover_image":       nil,
			"logo":              nil,
			"short_description": "A test merchant",
			"website_url":       "http://example.com",
		},
	}, []tests.UploadFile{
		{VariablePath: "input.cover_image", Filename: "cover.jpg", Content: []byte("dummy image content")},
		{VariablePath: "input.logo", Filename: "logo.jpg", Content: []byte("dummy image content")},
	})
	createData := s.Obj(s.Obj(create, "createMerchantDetail"), "data")
	s.detailID = int(createData["id"].(float64))
	s.Equal("Test Merchant", createData["display_name"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindMerchantDetailById($input: FindByIdMerchantDetailInput!) {
		findMerchantDetailById(input: $input) { status message data { id display_name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.detailID}})
	s.Equal(float64(s.detailID), s.Obj(s.Obj(byID, "findMerchantDetailById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllMerchantDetails($input: FindAllMerchantDetailInput!) {
		findAllMerchantDetails(input: $input) { status message pagination { total_records } data { id display_name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllMerchantDetails"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveMerchantDetails($input: FindAllMerchantDetailInput!) {
		findActiveMerchantDetails(input: $input) { status message pagination { total_records } data { id display_name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveMerchantDetails"), "data"))

	// 5. Update
	updated := s.GQLMultipart(s.handler, `mutation UpdateMerchantDetail($input: UpdateMerchantDetailInput!) {
		updateMerchantDetail(input: $input) { status message data { id display_name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_detail_id": s.detailID,
			"display_name":       "Updated Merchant",
			"cover_image":        nil,
			"logo":               nil,
			"short_description":  "Updated short description",
			"website_url":        "http://updated.com",
		},
	}, []tests.UploadFile{
		{VariablePath: "input.cover_image", Filename: "updated_cover.jpg", Content: []byte("dummy image content")},
		{VariablePath: "input.logo", Filename: "updated_logo.jpg", Content: []byte("dummy image content")},
	})
	s.Equal("Updated Merchant", s.Obj(s.Obj(updated, "updateMerchantDetail"), "data")["display_name"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashMerchantDetail($input: FindByIdMerchantDetailInput!) {
		trashMerchantDetail(input: $input) { status message data { id display_name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.detailID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedMerchantDetails($input: FindAllMerchantDetailInput!) {
		findTrashedMerchantDetails(input: $input) { status message pagination { total_records } data { id display_name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedMerchantDetails"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreMerchantDetail($input: FindByIdMerchantDetailInput!) {
		restoreMerchantDetail(input: $input) { status message data { id display_name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.detailID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashMerchantDetail($input: FindByIdMerchantDetailInput!) {
		trashMerchantDetail(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.detailID}})
	s.GQL(s.handler, `mutation DeleteMerchantDetailPermanent($input: FindByIdMerchantDetailInput!) {
		deleteMerchantDetailPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.detailID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllMerchantDetails { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllMerchantDetailsPermanent { status message } }`, nil)
}

func TestMerchantDetailGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantDetailGraphqlTestSuite))
}
