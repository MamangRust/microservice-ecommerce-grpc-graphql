package order_item_test

import (
	"context"

	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

// order_item has no single-record lookup: FindOrderItemByOrder returns an empty
// list (success) for a non-existent order, so there is no NotFound path.
//
// gapi: a non-existent order must return an empty result, not an error.
func (s *OrderItemGapiTestSuite) TestOrderItemGapiEmptyResult() {
	ctx := context.Background()
	res, err := s.queryClient.FindOrderItemByOrder(ctx, &pb.FindByIdOrderItemRequest{Id: 999999})
	s.NoError(err)
	s.NotNil(res)
	s.Empty(res.Data)
}

// repository: FindOrderItemByOrder on a non-existent order must return an empty
// result without error.
func (s *OrderItemRepositoryTestSuite) TestOrderItemFindByOrderEmpty() {
	ctx := context.Background()
	items, err := s.repo.OrderItemQuery.FindOrderItemByOrder(ctx, 999999)
	s.NoError(err)
	s.Empty(items)
}
