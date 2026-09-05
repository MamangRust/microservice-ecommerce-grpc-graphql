package cart_test

import (
	"context"

	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: cart has no FindById query; the valid negative path is validation of
// create payload — quantity 0 must be rejected as InvalidArgument.
func (s *CartGapiTestSuite) TestCartGapiInvalidQuantity() {
	ctx := context.Background()
	_, err := s.commandClient.Create(ctx, &pb.CreateCartRequest{
		UserId:    1,
		ProductId: 1,
		Quantity:  0,
	})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.InvalidArgument, st.Code(), "cart create with quantity 0 must be InvalidArgument, got %v: %s", st.Code(), st.Message())
}
