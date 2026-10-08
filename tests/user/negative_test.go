package user_test

import (
	"context"

	pb_user "github.com/MamangRust/microservice-ecommerce-grpc-pb/user"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent user must map to codes.NotFound (404), not Internal.
func (s *UserGapiTestSuite) TestUserGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindById(ctx, &pb_user.FindByIdUserRequest{Id: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent user must be NotFound, got %v: %s", st.Code(), st.Message())
}

// graphql: non-existent user must surface a Not Found (404) error, invalid id a
// GraphQL validation error. The user resolver maps gRPC codes to an HTTP-style
// "Not Found" message, unlike resolvers that wrap the raw gRPC error.
func (s *UserGraphqlTestSuite) TestUserGraphqlNotFound() {
	errs := s.GQLExpectError(s.handler, `query FindByIdUser($input: FindByIdUserInput!) {
		findByIdUser(input: $input) { status message data { id email } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": 999999}})
	s.Contains(errs[0].Message, "Not Found", "non-existent user must surface a Not Found error")
}

func (s *UserGraphqlTestSuite) TestUserGraphqlInvalidID() {
	// "abc" cannot satisfy Int!, so the request is rejected before the resolver runs.
	errs := s.GQLExpectError(s.handler, `query FindByIdUser($input: FindByIdUserInput!) {
		findByIdUser(input: $input) { status message data { id email } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": "abc"}})
	s.NotEmpty(errs)
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *UserRepositoryTestSuite) TestUserFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.UserQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
