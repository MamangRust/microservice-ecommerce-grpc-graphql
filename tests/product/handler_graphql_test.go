package product_test

import (
	"context"
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type ProductGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler      http.Handler
	productID    int
	merchantID   int
	categoryID   int
	categoryName string
}

func (s *ProductGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupCategoryService()
	s.SetupMerchantService()
	s.SetupProductService()

	// Seed dependencies
	ctx := context.Background()
	userID := s.SeedUser(ctx)
	s.categoryID = s.SeedCategory(ctx)
	s.merchantID = s.SeedMerchant(ctx, userID)
	s.categoryName = "Seed Category" // Default from SeedCategory if not specified

	s.handler = s.GraphQLHandler()
}

func (s *ProductGraphqlTestSuite) TestProductGraphqlLifecycle() {
	// 1. Create
	create := s.GQLMultipart(s.handler, `mutation CreateProduct($input: CreateProductInput!) {
		createProduct(input: $input) { status message data { id merchantId categoryId name countInStock slugProduct } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"merchantId":   s.merchantID,
			"categoryId":   s.categoryID,
			"name":         "Test Product",
			"description":  "Test Description",
			"price":        1000,
			"countInStock": 10,
			"brand":        "Test Brand",
			"weight":       1,
			"rating":       5,
			"slugProduct":  "test-product",
			"imageProduct": nil,
		},
	}, []tests.UploadFile{
		{VariablePath: "input.imageProduct", Filename: "product.jpg", Content: []byte("dummy image content")},
	})
	createData := s.Obj(s.Obj(create, "createProduct"), "data")
	s.productID = int(createData["id"].(float64))
	s.Equal("Test Product", createData["name"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindProductById($input: FindByIdProductInput!) {
		findProductById(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.productID}})
	s.Equal(float64(s.productID), s.Obj(s.Obj(byID, "findProductById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllProducts($input: FindAllProductInput!) {
		findAllProducts(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "pageSize": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllProducts"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveProducts($input: FindAllProductInput!) {
		findActiveProducts(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "pageSize": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveProducts"), "data"))

	// 5. FindByMerchant
	byMerchant := s.GQL(s.handler, `query FindProductsByMerchant($input: FindAllProductMerchantInput!) {
		findProductsByMerchant(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"merchantId": s.merchantID, "page": 1, "pageSize": 10}})
	s.NotEmpty(s.Arr(s.Obj(byMerchant, "findProductsByMerchant"), "data"))

	// 6. FindByCategory
	byCategory := s.GQL(s.handler, `query FindProductsByCategory($input: FindAllProductCategoryInput!) {
		findProductsByCategory(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"categoryName": s.categoryName, "page": 1, "pageSize": 10}})
	s.Obj(byCategory, "findProductsByCategory")

	// 7. Update
	updated := s.GQLMultipart(s.handler, `mutation UpdateProduct($input: UpdateProductInput!) {
		updateProduct(input: $input) { status message data { id name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"productId":    s.productID,
			"merchantId":   s.merchantID,
			"categoryId":   s.categoryID,
			"name":         "Updated Product",
			"description":  "Updated Description",
			"price":        2000,
			"countInStock": 20,
			"brand":        "Updated Brand",
			"weight":       2,
			"rating":       4,
			"slugProduct":  "updated-product",
			"imageProduct": nil,
		},
	}, []tests.UploadFile{
		{VariablePath: "input.imageProduct", Filename: "updated.jpg", Content: []byte("dummy image content")},
	})
	s.Equal("Updated Product", s.Obj(s.Obj(updated, "updateProduct"), "data")["name"])

	// 8. Trash
	s.GQL(s.handler, `mutation TrashedProduct($input: FindByIdProductInput!) {
		trashedProduct(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.productID}})

	// 9. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedProducts($input: FindAllProductInput!) {
		findTrashedProducts(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "pageSize": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedProducts"), "data"))

	// 10. Restore
	s.GQL(s.handler, `mutation RestoreProduct($input: FindByIdProductInput!) {
		restoreProduct(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.productID}})

	// 11. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashedProduct($input: FindByIdProductInput!) {
		trashedProduct(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.productID}})
	s.GQL(s.handler, `mutation DeleteProductPermanent($input: FindByIdProductInput!) {
		deleteProductPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.productID}})

	// 12. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllProducts { status message } }`, nil)

	// 13. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllProductsPermanent { status message } }`, nil)
}

// graphql: a non-existent product must surface a NotFound error.
func (s *ProductGraphqlTestSuite) TestProductGraphqlNotFound() {
	errs := s.GQLExpectError(s.handler, `query FindProductById($input: FindByIdProductInput!) {
		findProductById(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": 999999}})
	s.Contains(errs[0].Message, "NotFound", "non-existent product must surface a NotFound error")
}

// graphql: a non-numeric id must be rejected as a GraphQL validation error.
func (s *ProductGraphqlTestSuite) TestProductGraphqlInvalidID() {
	errs := s.GQLExpectError(s.handler, `query FindProductById($input: FindByIdProductInput!) {
		findProductById(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": "abc"}})
	s.NotEmpty(errs)
}

func TestProductGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ProductGraphqlTestSuite))
}
