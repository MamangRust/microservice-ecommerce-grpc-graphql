package cart_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/graphtest"
	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type CartGraphQLTestSuite struct {
	tests.BaseTestSuite
	graphqlH http.Handler
}

func (s *CartGraphQLTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupUserService()
	s.SetupCategoryService()
	s.SetupMerchantService()
	s.SetupProductService()
	cacheStore := s.GetCacheStore()
	s.graphqlH = graphtest.NewTestHandler(s.Conns, cacheStore, s.Log)
}

func (s *CartGraphQLTestSuite) TearDownSuite() {
	s.BaseTestSuite.TearDownSuite()
}

func gqlCart(h http.Handler, query string, variables map[string]interface{}) map[string]interface{} {
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

func (s *CartGraphQLTestSuite) Test1_SmokeQuery() {
	// Basic smoke test - just verify GraphQL handler works
	result := gqlCart(s.graphqlH, "{ __typename }", nil)
	s.NotNil(result)
}

func TestCartGraphQLSuite(t *testing.T) {
	if testing.Short() { t.Skip("skipping integration test") }
	suite.Run(t, new(CartGraphQLTestSuite))
}
