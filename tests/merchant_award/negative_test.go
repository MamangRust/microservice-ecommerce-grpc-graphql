package merchant_award_test

import (
	"context"

	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent merchant award must map to codes.NotFound (404), not Internal.
func (s *MerchantAwardGapiTestSuite) TestMerchantAwardGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindById(ctx, &pb.FindByIdMerchantAwardRequest{Id: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent merchant award must be NotFound, got %v: %s", st.Code(), st.Message())
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *MerchantAwardRepositoryTestSuite) TestMerchantAwardFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.MerchantAwardQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
