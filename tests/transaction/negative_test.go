package transaction_test

import (
	"context"

	pb_transaction "github.com/MamangRust/microservice-ecommerce-grpc-pb/transaction"
	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent transaction must map to codes.NotFound (404), not Internal.
func (s *TransactionGapiTestSuite) TestTransactionGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindById(ctx, &pb_transaction.FindByIdTransactionRequest{Id: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent transaction must be NotFound, got %v: %s", st.Code(), st.Message())
}

// graphql: non-existent transaction must surface a NotFound error, invalid id a
// GraphQL validation error.
func (s *TransactionGraphqlTestSuite) TestTransactionGraphqlNotFound() {
	errs := s.GQLExpectError(s.handler, `query FindTransactionById($input: FindByIdTransactionRequest!) {
		findTransactionById(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": 999999}})
	s.Contains(errs[0].Message, "NotFound", "non-existent transaction must surface a NotFound error")
}

func (s *TransactionGraphqlTestSuite) TestTransactionGraphqlInvalidID() {
	// "abc" cannot satisfy Int!, so the request is rejected before the resolver runs.
	errs := s.GQLExpectError(s.handler, `query FindTransactionById($input: FindByIdTransactionRequest!) {
		findTransactionById(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": "abc"}})
	s.NotEmpty(errs)
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *TransactionRepositoryTestSuite) TestTransactionFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.TransactionQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
