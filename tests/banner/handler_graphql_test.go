package banner_test

import (
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type BannerGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler  http.Handler
	bannerID int
}

func (s *BannerGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupBannerService()

	s.handler = s.GraphQLHandler()
}

func (s *BannerGraphqlTestSuite) TestBannerGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateBanner($input: CreateBannerInput!) {
		createBanner(input: $input) { status message data { banner_id name is_active } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"name":       "Test Banner",
			"start_date": "2024-01-01",
			"end_date":   "2024-12-31",
			"start_time": "00:00:00",
			"end_time":   "23:59:59",
			"is_active":  true,
		},
	})
	createData := s.Obj(s.Obj(create, "createBanner"), "data")
	s.bannerID = int(createData["banner_id"].(float64))
	s.Equal("Test Banner", createData["name"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindBannerById($input: FindByIdBannerInput!) {
		findBannerById(input: $input) { status message data { banner_id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.bannerID}})
	s.Equal(float64(s.bannerID), s.Obj(s.Obj(byID, "findBannerById"), "data")["banner_id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllBanners($input: FindAllBannerInput!) {
		findAllBanners(input: $input) { status message pagination { total_records } data { banner_id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllBanners"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveBanners($input: FindAllBannerInput!) {
		findActiveBanners(input: $input) { status message pagination { total_records } data { banner_id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveBanners"), "data"))

	// 5. Update
	updated := s.GQL(s.handler, `mutation UpdateBanner($input: UpdateBannerInput!) {
		updateBanner(input: $input) { status message data { banner_id name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"banner_id":  s.bannerID,
			"name":       "Updated Banner",
			"start_date": "2024-01-01",
			"end_date":   "2024-12-31",
			"start_time": "00:00:00",
			"end_time":   "23:59:59",
			"is_active":  false,
		},
	})
	s.Equal("Updated Banner", s.Obj(s.Obj(updated, "updateBanner"), "data")["name"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashBanner($input: FindByIdBannerInput!) {
		trashBanner(input: $input) { status message data { banner_id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.bannerID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedBanners($input: FindAllBannerInput!) {
		findTrashedBanners(input: $input) { status message pagination { total_records } data { banner_id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedBanners"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreBanner($input: FindByIdBannerInput!) {
		restoreBanner(input: $input) { status message data { banner_id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.bannerID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashBanner($input: FindByIdBannerInput!) {
		trashBanner(input: $input) { status message data { banner_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.bannerID}})
	s.GQL(s.handler, `mutation DeleteBannerPermanent($input: FindByIdBannerInput!) {
		deleteBannerPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.bannerID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllBanners { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllBannersPermanent { status message } }`, nil)
}

func TestBannerGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(BannerGraphqlTestSuite))
}
