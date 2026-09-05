package category_test

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/graphtest"
	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type CategoryGraphQLTestSuite struct {
	tests.BaseTestSuite
	graphqlH    http.Handler
	categoryID  int
}

func (s *CategoryGraphQLTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupCategoryService()

	cacheStore := s.GetCacheStore()
	s.graphqlH = graphtest.NewTestHandler(s.Conns, cacheStore, s.Log)
}

func (s *CategoryGraphQLTestSuite) TearDownSuite() {
	s.BaseTestSuite.TearDownSuite()
}

func gqlCategory(h http.Handler, query string, variables map[string]interface{}) map[string]interface{} {
	body := map[string]interface{}{"query": query, "variables": variables}
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	var result map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &result)
	return result
}

// gqlCategoryMultipart issues a GraphQL multipart request so the required
// image_category Upload! field can carry a real file part.
func gqlCategoryMultipart(h http.Handler, query string, variables map[string]interface{}) map[string]interface{} {
	// The Upload! field must be present (as null) in the operations payload so
	// the multipart transport can inject the file part into it.
	input, _ := variables["input"].(map[string]interface{})
	input["image_category"] = nil
	operations, _ := json.Marshal(map[string]interface{}{"query": query, "variables": variables})
	fileMap, _ := json.Marshal(map[string][]string{"0": {"variables.input.image_category"}})

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	_ = w.WriteField("operations", string(operations))
	_ = w.WriteField("map", string(fileMap))
	part, _ := w.CreateFormFile("0", "category.jpg")
	_, _ = part.Write([]byte("fake image content"))
	_ = w.Close()

	req, _ := http.NewRequest(http.MethodPost, "/query", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	var result map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &result)
	return result
}

func (s *CategoryGraphQLTestSuite) Test1_CreateCategory() {
	vars := map[string]interface{}{"input": map[string]interface{}{
		"name":          "GraphQL Category",
		"description":   "Test Description",
		"slug_category": "graphql-category",
	}}
	result := gqlCategoryMultipart(s.graphqlH, `mutation CreateCategory($input: CreateCategoryInput!) { createCategory(input: $input) { status message data { id name slug_category } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["createCategory"].(map[string]interface{})["status"])
}

func (s *CategoryGraphQLTestSuite) Test2_FindAllCategory() {
	vars := map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}}
	result := gqlCategory(s.graphqlH, `query FindAllCategory($input: FindAllCategoryInput!) { findAllCategories(input: $input) { status message data { id name } pagination { current_page page_size total_pages total_records } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["findAllCategories"].(map[string]interface{})["status"])
}

func (s *CategoryGraphQLTestSuite) Test3_RestoreAllAndDeleteAll() {
	result := gqlCategory(s.graphqlH, `mutation { restoreAllCategories { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["restoreAllCategories"].(map[string]interface{})["status"])

	result = gqlCategory(s.graphqlH, `mutation { deleteAllCategoriesPermanent { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["deleteAllCategoriesPermanent"].(map[string]interface{})["status"])
}

func TestCategoryGraphQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CategoryGraphQLTestSuite))
}
