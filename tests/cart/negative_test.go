package cart_test

import (
	"context"

	pb_cart "github.com/MamangRust/microservice-ecommerce-grpc-pb/cart"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// gapi: cart has no FindById query; the valid negative path is validation of
// create payload — quantity 0 must be rejected as InvalidArgument.
func (s *CartGapiTestSuite) TestCartGapiInvalidQuantity() {
	ctx := context.Background()
	_, err := s.commandClient.Create(ctx, &pb_cart.CreateCartRequest{
		UserId:    1,
		ProductId: 1,
		Quantity:  0,
	})
	s.Require().Error(err)
	st, ok := status.FromError(err)
	s.Require().True(ok, "expected a gRPC status error")
	s.Equal(codes.InvalidArgument, st.Code(), "cart create with quantity 0 must be InvalidArgument, got %v: %s", st.Code(), st.Message())
}

// graphql: a malformed create payload (a non-integer where Int! is required)
// must be rejected by GraphQL validation before the resolver runs.
func (s *CartGraphqlTestSuite) TestCartGraphqlInvalidInput() {
	errs := s.GQLExpectError(s.handler, `mutation CreateCart($input: CreateCartInput!) {
		createCart(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{
		"user_id":    s.userID,
		"product_id": 1,
		"quantity":   "not-a-number",
	}})
	s.NotEmpty(errs)
}
