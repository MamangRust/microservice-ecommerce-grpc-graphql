package merchant_business_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type MerchantBusinessGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler    http.Handler
	businessID int
	merchantID int
	userID     int
}

func (s *MerchantBusinessGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupMerchantService()
	s.SetupMerchantBusinessService()

	ctx := context.Background()
	s.userID = s.SeedUser(ctx)
	s.merchantID = s.SeedMerchant(ctx, s.userID)

	s.handler = s.GraphQLHandler()
}

func (s *MerchantBusinessGraphqlTestSuite) TestMerchantBusinessGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateMerchantBusiness($input: CreateMerchantBusinessInput!) {
		createMerchantBusiness(input: $input) { status message data { id merchant_id business_type } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_id":         s.merchantID,
			"business_type":       "Retail",
			"tax_id":              "123-456-789",
			"established_year":    2020,
			"number_of_employees": 10,
			"website_url":         "http://example.com",
		},
	})
	createData := s.Obj(s.Obj(create, "createMerchantBusiness"), "data")
	s.businessID = int(createData["id"].(float64))
	s.Equal("Retail", createData["business_type"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindMerchantBusinessById($input: FindByIdMerchantBusinessInput!) {
		findMerchantBusinessById(input: $input) { status message data { id business_type } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.businessID}})
	s.Equal(float64(s.businessID), s.Obj(s.Obj(byID, "findMerchantBusinessById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllMerchantBusinesses($input: FindAllMerchantBusinessInput!) {
		findAllMerchantBusinesses(input: $input) { status message pagination { total_records } data { id business_type } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllMerchantBusinesses"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveMerchantBusinesses($input: FindAllMerchantBusinessInput!) {
		findActiveMerchantBusinesses(input: $input) { status message pagination { total_records } data { id business_type } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveMerchantBusinesses"), "data"))

	// 5. Update
	updated := s.GQL(s.handler, `mutation UpdateMerchantBusiness($input: UpdateMerchantBusinessInput!) {
		updateMerchantBusiness(input: $input) { status message data { id business_type } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchant_business_info_id": s.businessID,
			"business_type":             "Wholesale",
			"tax_id":                    "987-654-321",
			"established_year":          2021,
			"number_of_employees":       20,
			"website_url":               "http://updated.com",
		},
	})
	s.Equal("Wholesale", s.Obj(s.Obj(updated, "updateMerchantBusiness"), "data")["business_type"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashMerchantBusiness($input: FindByIdMerchantBusinessInput!) {
		trashMerchantBusiness(input: $input) { status message data { id business_type } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.businessID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedMerchantBusinesses($input: FindAllMerchantBusinessInput!) {
		findTrashedMerchantBusinesses(input: $input) { status message pagination { total_records } data { id business_type } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedMerchantBusinesses"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreMerchantBusiness($input: FindByIdMerchantBusinessInput!) {
		restoreMerchantBusiness(input: $input) { status message data { id business_type } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.businessID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashMerchantBusiness($input: FindByIdMerchantBusinessInput!) {
		trashMerchantBusiness(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.businessID}})
	s.GQL(s.handler, `mutation DeleteMerchantBusinessPermanent($input: FindByIdMerchantBusinessInput!) {
		deleteMerchantBusinessPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.businessID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllMerchantBusinesses { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllMerchantBusinessesPermanent { status message } }`, nil)
}

func TestMerchantBusinessGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantBusinessGraphqlTestSuite))
}
