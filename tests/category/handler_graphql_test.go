package category_test

import (
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type CategoryGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler    http.Handler
	categoryID int
}

func (s *CategoryGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupCategoryService()

	s.handler = s.GraphQLHandler()
}

func (s *CategoryGraphqlTestSuite) TestCategoryGraphqlLifecycle() {
	// 1. Create
	create := s.GQLMultipart(s.handler, `mutation CreateCategory($input: CreateCategoryInput!) {
		createCategory(input: $input) { status message data { id name slug_category } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"name":           "Test Category",
			"description":    "Test Description",
			"slug_category":  "test-category",
			"image_category": nil,
		},
	}, []tests.UploadFile{
		{VariablePath: "input.image_category", Filename: "test.jpg", Content: []byte("dummy image content")},
	})
	createData := s.Obj(s.Obj(create, "createCategory"), "data")
	s.categoryID = int(createData["id"].(float64))
	s.Equal("Test Category", createData["name"])

	// 2. FindById
	byID := s.GQL(s.handler, `query FindCategoryById($input: FindByIdCategoryInput!) {
		findCategoryById(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.categoryID}})
	s.Equal(float64(s.categoryID), s.Obj(s.Obj(byID, "findCategoryById"), "data")["id"])

	// 3. FindAll
	all := s.GQL(s.handler, `query FindAllCategories($input: FindAllCategoryInput!) {
		findAllCategories(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllCategories"), "data"))

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindActiveCategories($input: FindAllCategoryInput!) {
		findActiveCategories(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findActiveCategories"), "data"))

	// 5. Update
	updated := s.GQLMultipart(s.handler, `mutation UpdateCategory($input: UpdateCategoryInput!) {
		updateCategory(input: $input) { status message data { id name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"category_id":    s.categoryID,
			"name":           "Updated Category",
			"description":    "Updated Description",
			"slug_category":  "updated-category",
			"image_category": nil,
		},
	}, []tests.UploadFile{
		{VariablePath: "input.image_category", Filename: "updated.jpg", Content: []byte("dummy image content")},
	})
	s.Equal("Updated Category", s.Obj(s.Obj(updated, "updateCategory"), "data")["name"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashCategory($input: FindByIdCategoryInput!) {
		trashCategory(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.categoryID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindTrashedCategories($input: FindAllCategoryInput!) {
		findTrashedCategories(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findTrashedCategories"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreCategory($input: FindByIdCategoryInput!) {
		restoreCategory(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.categoryID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashCategory($input: FindByIdCategoryInput!) {
		trashCategory(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.categoryID}})
	s.GQL(s.handler, `mutation DeleteCategoryPermanent($input: FindByIdCategoryInput!) {
		deleteCategoryPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.categoryID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllCategories { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllCategoriesPermanent { status message } }`, nil)
}

func TestCategoryGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CategoryGraphqlTestSuite))
}
