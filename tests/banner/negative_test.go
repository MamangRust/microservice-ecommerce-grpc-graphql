package banner_test

import (
	"context"

	pb_banner "github.com/MamangRust/microservice-ecommerce-grpc-pb/banner"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent banner must map to codes.NotFound (404), not Internal.
func (s *BannerGapiTestSuite) TestBannerGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindById(ctx, &pb_banner.FindByIdBannerRequest{Id: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent banner must be NotFound, got %v: %s", st.Code(), st.Message())
}

// graphql: non-existent banner must surface a NotFound error, invalid id a
// GraphQL validation error.
func (s *BannerGraphqlTestSuite) TestBannerGraphqlNotFound() {
	errs := s.GQLExpectError(s.handler, `query FindBannerById($input: FindByIdBannerInput!) {
		findBannerById(input: $input) { status message data { banner_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": 999999}})
	s.Contains(errs[0].Message, "NotFound", "non-existent banner must surface a NotFound error")
}

func (s *BannerGraphqlTestSuite) TestBannerGraphqlInvalidID() {
	// "abc" cannot satisfy Int!, so the request is rejected before the resolver runs.
	errs := s.GQLExpectError(s.handler, `query FindBannerById($input: FindByIdBannerInput!) {
		findBannerById(input: $input) { status message data { banner_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": "abc"}})
	s.NotEmpty(errs)
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *BannerRepositoryTestSuite) TestBannerFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.BannerQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
