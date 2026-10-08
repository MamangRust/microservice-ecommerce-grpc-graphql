package review_detail_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type ReviewDetailGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler        http.Handler
	reviewDetailID int
	reviewID       int
}

func (s *ReviewDetailGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()

	// Setup dependencies
	s.SetupUserService()
	s.SetupCategoryService()
	s.SetupMerchantService()
	s.SetupProductService()
	s.SetupOrderItemService()
	s.SetupShippingAddressService()
	s.SetupOrderService()
	s.SetupReviewService()
	s.SetupReviewDetailService()

	// Seed dependencies
	ctx := context.Background()
	userID := s.SeedUser(ctx)
	catID := s.SeedCategory(ctx)
	merchID := s.SeedMerchant(ctx, userID)
	prodID := s.SeedProduct(ctx, merchID, catID)
	s.SeedOrder(ctx, userID, merchID, prodID)
	s.reviewID = s.SeedReview(ctx, userID, prodID)

	s.handler = s.GraphQLHandler()
}

func (s *ReviewDetailGraphqlTestSuite) TestReviewDetailGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateReviewDetail($input: CreateReviewDetailInput!) {
		createReviewDetail(input: $input) { status message data { id reviewId type url caption } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"reviewId": s.reviewID,
			"type":     "photo",
			"url":      "review_photo.jpg",
			"caption":  "API Test Caption",
		},
	})
	createData := s.Obj(s.Obj(create, "createReviewDetail"), "data")
	s.reviewDetailID = int(createData["id"].(float64))
	s.Equal("API Test Caption", createData["caption"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindReviewDetailById($input: FindByIdReviewDetailInput!) {
		findReviewDetailById(input: $input) { status message data { id reviewId caption } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewDetailID}})
	s.Equal(float64(s.reviewDetailID), s.Obj(s.Obj(byID, "findReviewDetailById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllReviewDetails($input: FindAllReviewDetailInput!) {
		findAllReviewDetails(input: $input) { status message pagination { total_records } data { id caption } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "pageSize": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllReviewDetails"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveReviewDetails($input: FindAllReviewDetailInput!) {
		findActiveReviewDetails(input: $input) { status message pagination { total_records } data { id caption } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "pageSize": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveReviewDetails"), "data"))

	// 5. Update
	updated := s.GQL(s.handler, `mutation UpdateReviewDetail($input: UpdateReviewDetailInput!) {
		updateReviewDetail(input: $input) { status message data { id type caption } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"reviewDetailId": s.reviewDetailID,
			"type":           "video",
			"url":            "review_video.mp4",
			"caption":        "Updated Caption",
		},
	})
	s.Equal("Updated Caption", s.Obj(s.Obj(updated, "updateReviewDetail"), "data")["caption"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashedReviewDetail($input: FindByIdReviewDetailInput!) {
		trashedReviewDetail(input: $input) { status message data { id caption } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewDetailID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedReviewDetails($input: FindAllReviewDetailInput!) {
		findTrashedReviewDetails(input: $input) { status message pagination { total_records } data { id caption } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "pageSize": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedReviewDetails"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreReviewDetail($input: FindByIdReviewDetailInput!) {
		restoreReviewDetail(input: $input) { status message data { id caption } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewDetailID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashedReviewDetail($input: FindByIdReviewDetailInput!) {
		trashedReviewDetail(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewDetailID}})
	s.GQL(s.handler, `mutation DeleteReviewDetailPermanent($input: FindByIdReviewDetailInput!) {
		deleteReviewDetailPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewDetailID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllReviewDetails { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllReviewDetailsPermanent { status message } }`, nil)
}

func TestReviewDetailGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ReviewDetailGraphqlTestSuite))
}
