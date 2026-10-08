package slider_test

import (
	"context"

	pb_slider "github.com/MamangRust/microservice-ecommerce-grpc-pb/slider"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent slider must map to codes.NotFound (404), not Internal.
func (s *SliderGapiTestSuite) TestSliderGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindById(ctx, &pb_slider.FindByIdSliderRequest{Id: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent slider must be NotFound, got %v: %s", st.Code(), st.Message())
}

// graphql: non-existent slider must surface a NotFound error, invalid id a
// GraphQL validation error. Slider has no findById query, so trashedSlider is
// used to exercise the same lookup path.
func (s *SliderGraphqlTestSuite) TestSliderGraphqlNotFound() {
	errs := s.GQLExpectError(s.handler, `mutation TrashedSlider($input: FindByIdSliderRequest!) {
		trashedSlider(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": 999999}})
	s.Contains(errs[0].Message, "NotFound", "non-existent slider must surface a NotFound error")
}

func (s *SliderGraphqlTestSuite) TestSliderGraphqlInvalidID() {
	// "abc" cannot satisfy Int!, so the request is rejected before the resolver runs.
	errs := s.GQLExpectError(s.handler, `mutation TrashedSlider($input: FindByIdSliderRequest!) {
		trashedSlider(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": "abc"}})
	s.NotEmpty(errs)
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *SliderRepositoryTestSuite) TestSliderFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.SliderQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
