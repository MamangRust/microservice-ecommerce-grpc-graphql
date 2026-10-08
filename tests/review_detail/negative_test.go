package review_detail_test

import (
	"context"

	pb_review_detail "github.com/MamangRust/microservice-ecommerce-grpc-pb/review_detail"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent review detail must map to codes.NotFound (404), not Internal.
func (s *ReviewDetailGapiTestSuite) TestReviewDetailGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindById(ctx, &pb_review_detail.FindByIdReviewDetailRequest{Id: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent review detail must be NotFound, got %v: %s", st.Code(), st.Message())
}

// graphql: non-existent review detail must surface a NotFound error, invalid
// id a GraphQL validation error.
func (s *ReviewDetailGraphqlTestSuite) TestReviewDetailGraphqlNotFound() {
	errs := s.GQLExpectError(s.handler, `query FindReviewDetailById($input: FindByIdReviewDetailInput!) {
		findReviewDetailById(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": 999999}})
	s.Contains(errs[0].Message, "NotFound", "non-existent review detail must surface a NotFound error")
}

func (s *ReviewDetailGraphqlTestSuite) TestReviewDetailGraphqlInvalidID() {
	// "abc" cannot satisfy Int!, so the request is rejected before the resolver runs.
	errs := s.GQLExpectError(s.handler, `query FindReviewDetailById($input: FindByIdReviewDetailInput!) {
		findReviewDetailById(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": "abc"}})
	s.NotEmpty(errs)
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *ReviewDetailRepositoryTestSuite) TestReviewDetailFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.ReviewDetailQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
