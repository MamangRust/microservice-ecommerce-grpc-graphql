package review_test

import (
	"context"

	pb_review "github.com/MamangRust/microservice-ecommerce-grpc-pb/review"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: review has no FindById query RPC; Update on a non-existent review must
// map to codes.NotFound (404), not Internal.
func (s *ReviewGapiTestSuite) TestReviewGapiNotFound() {
	ctx := context.Background()
	_, err := s.commandClient.Update(ctx, &pb_review.UpdateReviewRequest{
		ReviewId: 999999,
		Rating:   5,
		Comment:  "not found",
	})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "update on non-existent review must be NotFound, got %v: %s", st.Code(), st.Message())
}

// graphql: update on a non-existent review must surface a NotFound error,
// invalid id a GraphQL validation error.
func (s *ReviewGraphqlTestSuite) TestReviewGraphqlNotFound() {
	// Valid input required — otherwise gateway validation fires before the NotFound lookup.
	errs := s.GQLExpectError(s.handler, `mutation UpdateReview($input: UpdateReviewRequest!) {
		updateReview(input: $input) { status message data { id comment } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"review_id": 999999,
			"name":      "not found",
			"comment":   "not found",
			"rating":    5,
		},
	})
	s.Contains(errs[0].Message, "NotFound", "update on non-existent review must surface a NotFound error")
}

func (s *ReviewGraphqlTestSuite) TestReviewGraphqlInvalidID() {
	// "abc" cannot satisfy Int!, so the request is rejected before the resolver runs.
	errs := s.GQLExpectError(s.handler, `mutation UpdateReview($input: UpdateReviewRequest!) {
		updateReview(input: $input) { status message data { id comment } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"review_id": "abc",
			"name":      "not found",
			"comment":   "not found",
			"rating":    5,
		},
	})
	s.NotEmpty(errs)
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *ReviewRepositoryTestSuite) TestReviewFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.ReviewQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
