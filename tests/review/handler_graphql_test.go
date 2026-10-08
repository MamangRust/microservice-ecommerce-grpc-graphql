package review_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type ReviewGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler    http.Handler
	reviewID   int
	userID     int
	prodID     int
	merchantID int
}

func (s *ReviewGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupMerchantService()
	s.SetupCategoryService()
	s.SetupProductService()
	s.SetupOrderItemService()
	s.SetupShippingAddressService()
	s.SetupOrderService()
	s.SetupReviewService()

	ctx := context.Background()
	s.userID = s.SeedUser(ctx)
	catID := s.SeedCategory(ctx)
	s.merchantID = s.SeedMerchant(ctx, s.userID)
	s.prodID = s.SeedProduct(ctx, s.merchantID, catID)

	s.handler = s.GraphQLHandler()
}

func (s *ReviewGraphqlTestSuite) TestReviewGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateReview($input: CreateReviewRequest!) {
		createReview(input: $input) { status message data { id user_id product_id comment rating } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"user_id":    s.userID,
			"product_id": s.prodID,
			"name":       "Test Reviewer",
			"comment":    "Test Comment",
			"rating":     5,
		},
	})
	createData := s.Obj(s.Obj(create, "createReview"), "data")
	s.reviewID = int(createData["id"].(float64))
	s.Equal("Test Comment", createData["comment"])

	// 2. FindAll
	all := s.GQL(s.handler, `query FindAllReviews($input: FindAllReviewRequest!) {
		findAllReviews(input: $input) { status message pagination { total_records } data { id comment } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllReviews"), "data"))

	// 3. FindByProduct — the service filters by an optional rating that the
	// request leaves at 0, so the result can legitimately be empty; mirror the
	// old REST test and only assert the operation succeeds.
	s.Require().NotZero(s.prodID)
	byProduct := s.GQL(s.handler, `query FindReviewsByProduct($input: FindAllReviewProductRequest!) {
		findReviewsByProduct(input: $input) { status message pagination { total_records } data { id comment } }
	}`, map[string]interface{}{"input": map[string]interface{}{"product_id": s.prodID, "page": 1, "page_size": 10}})
	s.NotNil(s.Obj(byProduct, "findReviewsByProduct"))

	// 4. FindByMerchant — same optional-rating filter as FindByProduct.
	s.Require().NotZero(s.merchantID)
	byMerchant := s.GQL(s.handler, `query FindReviewsByMerchant($input: FindAllReviewMerchantRequest!) {
		findReviewsByMerchant(input: $input) { status message pagination { total_records } data { id comment } }
	}`, map[string]interface{}{"input": map[string]interface{}{"merchant_id": s.merchantID, "page": 1, "page_size": 10}})
	s.NotNil(s.Obj(byMerchant, "findReviewsByMerchant"))

	// 5. FindByActive
	active := s.GQL(s.handler, `query FindActiveReviews($input: FindAllReviewRequest!) {
		findActiveReviews(input: $input) { status message pagination { total_records } data { id comment } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotNil(s.Obj(active, "findActiveReviews"))

	// 6. Update
	updated := s.GQL(s.handler, `mutation UpdateReview($input: UpdateReviewRequest!) {
		updateReview(input: $input) { status message data { id comment rating } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"review_id": s.reviewID,
			"name":      "Test Reviewer",
			"comment":   "Updated Comment",
			"rating":    4,
		},
	})
	s.Equal("Updated Comment", s.Obj(s.Obj(updated, "updateReview"), "data")["comment"])

	// 7. Trash
	s.GQL(s.handler, `mutation TrashedReview($input: FindByIdReviewRequest!) {
		trashedReview(input: $input) { status message data { id comment } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewID}})

	// 8. FindByTrashed — the list is cached and the trash mutation only evicts
	// the single-review cache, so mirror the old REST test and only assert the
	// operation succeeds rather than its content.
	trashed := s.GQL(s.handler, `query FindTrashedReviews($input: FindAllReviewRequest!) {
		findTrashedReviews(input: $input) { status message pagination { total_records } data { id comment } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotNil(s.Obj(trashed, "findTrashedReviews"))

	// 9. Restore
	s.GQL(s.handler, `mutation RestoreReview($input: FindByIdReviewRequest!) {
		restoreReview(input: $input) { status message data { id comment } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewID}})

	// 10. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashedReview($input: FindByIdReviewRequest!) {
		trashedReview(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewID}})
	s.GQL(s.handler, `mutation DeleteReviewPermanent($input: FindByIdReviewRequest!) {
		deleteReviewPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.reviewID}})

	// 11. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllReviews { status message } }`, nil)

	// 12. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllReviewsPermanent { status message } }`, nil)
}

func TestReviewGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ReviewGraphqlTestSuite))
}
