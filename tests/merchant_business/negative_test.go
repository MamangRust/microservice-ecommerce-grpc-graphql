package merchant_business_test

import (
	"context"

	"github.com/MamangRust/microservice-ecommerce-shared/errors"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: non-existent merchant business must map to codes.NotFound (404), not Internal.
func (s *MerchantBusinessGapiTestSuite) TestMerchantBusinessGapiNotFound() {
	ctx := context.Background()
	_, err := s.queryClient.FindById(ctx, &pb.FindByIdMerchantBusinessRequest{Id: 999999})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.NotFound, st.Code(), "non-existent merchant business must be NotFound, got %v: %s", st.Code(), st.Message())
}

// repository: FindByID on a non-existent ID must return a typed not-found error.
func (s *MerchantBusinessRepositoryTestSuite) TestMerchantBusinessFindByIDNotFound() {
	ctx := context.Background()
	_, err := s.repo.MerchantBusinessQuery.FindByID(ctx, 999999)
	s.Require().Error(err)
	var appErr *errors.AppError
	s.Require().ErrorAs(err, &appErr)
	s.Equal(errors.ErrorTypeNotFound, appErr.Type, "expected not-found error type, got %s: %v", appErr.Type, err)
}
