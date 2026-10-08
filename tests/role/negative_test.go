package role_test

import (
	"context"

	pb_role "github.com/MamangRust/microservice-ecommerce-grpc-pb/role"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent role must map to codes.NotFound (404), not Internal.
func (s *RoleGapiTestSuite) TestRoleGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindByIdRole(ctx, &pb_role.FindByIdRoleRequest{RoleId: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent role must be NotFound, got %v: %s", st.Code(), st.Message())
}

// graphql: non-existent role must surface a Not Found (404) error, invalid id a
// GraphQL validation error. The role resolver maps gRPC codes to an HTTP-style
// "Not Found" message, unlike resolvers that wrap the raw gRPC error.
func (s *RoleGraphqlTestSuite) TestRoleGraphqlNotFound() {
	errs := s.GQLExpectError(s.handler, `query FindByIdRole($input: FindByIdRoleInput!) {
		findByIdRole(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"role_id": 999999}})
	s.Contains(errs[0].Message, "Not Found", "non-existent role must surface a Not Found error")
}

func (s *RoleGraphqlTestSuite) TestRoleGraphqlInvalidID() {
	// "abc" cannot satisfy Int!, so the request is rejected before the resolver runs.
	errs := s.GQLExpectError(s.handler, `query FindByIdRole($input: FindByIdRoleInput!) {
		findByIdRole(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"role_id": "abc"}})
	s.NotEmpty(errs)
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *RoleRepositoryTestSuite) TestRoleFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.RoleQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
