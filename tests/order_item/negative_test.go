package order_item_test

import (
	"context"

	pb_order_item "github.com/MamangRust/microservice-ecommerce-grpc-pb/order_item"
)

// order_item has no single-record lookup: FindOrderItemByOrder returns an empty
// list (success) for a non-existent order, so there is no NotFound path.
//
// gapi: a non-existent order must return an empty result, not an error.
func (s *OrderItemGapiTestSuite) TestOrderItemGapiEmptyResult() {
	ctx := context.Background()
	res, err := s.queryClient.FindOrderItemByOrder(ctx, &pb_order_item.FindByIdOrderItemRequest{Id: 999999})
	s.NoError(err)
	s.NotNil(res)
	s.Empty(res.Data)
}

// graphql: a non-existent order returns an empty data list, not an error.
func (s *OrderItemGraphqlTestSuite) TestOrderItemGraphqlEmptyResult() {
	data := s.GQL(s.handler, `query FindOrderItemsByOrder($input: FindByIdOrderItemInput!) {
		findOrderItemsByOrder(input: $input) { status message data { id order_id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": 999999}})
	s.Empty(s.Arr(s.Obj(data, "findOrderItemsByOrder"), "data"))
}

// graphql: "abc" cannot satisfy Int!, so the request is rejected before the
// resolver runs.
func (s *OrderItemGraphqlTestSuite) TestOrderItemGraphqlInvalidID() {
	errs := s.GQLExpectError(s.handler, `query FindOrderItemsByOrder($input: FindByIdOrderItemInput!) {
		findOrderItemsByOrder(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": "abc"}})
	s.NotEmpty(errs)
}

// repository: FindOrderItemByOrder on a non-existent order must return an empty
// result without error.
func (s *OrderItemRepositoryTestSuite) TestOrderItemFindByOrderEmpty() {
	ctx := context.Background()
	items, err := s.repo.OrderItemQuery.FindOrderItemByOrder(ctx, 999999)
	s.NoError(err)
	s.Empty(items)
}
