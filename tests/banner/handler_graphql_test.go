package banner_test

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

type BannerGraphQLTestSuite struct {
	tests.BaseTestSuite
	graphqlH http.Handler
	bannerID int
}

func (s *BannerGraphQLTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupBannerService()

	cacheStore := s.GetCacheStore()
	s.graphqlH = graphtest.NewTestHandler(s.Conns, cacheStore, s.Log)
}

func (s *BannerGraphQLTestSuite) TearDownSuite() {
	s.BaseTestSuite.TearDownSuite()
}

func gqlBanner(h http.Handler, query string, variables map[string]interface{}) map[string]interface{} {
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

func (s *BannerGraphQLTestSuite) Test1_CreateBanner() {
	vars := map[string]interface{}{"input": map[string]interface{}{
		"name":       "GraphQL Banner",
		"start_date": "2026-01-01",
		"end_date":   "2026-12-31",
		"start_time": "08:00:00",
		"end_time":   "17:00:00",
		"is_active":  true,
	}}
	result := gqlBanner(s.graphqlH, `mutation CreateBanner($input: CreateBannerInput!) { createBanner(input: $input) { status message data { banner_id name } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["createBanner"].(map[string]interface{})["status"])
}

func (s *BannerGraphQLTestSuite) Test2_FindAllBanner() {
	vars := map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}}
	result := gqlBanner(s.graphqlH, `query FindAllBanner($input: FindAllBannerInput!) { findAllBanners(input: $input) { status message data { banner_id name } pagination { current_page page_size total_pages total_records } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["findAllBanners"].(map[string]interface{})["status"])
}

func (s *BannerGraphQLTestSuite) Test3_RestoreAllAndDeleteAll() {
	result := gqlBanner(s.graphqlH, `mutation { restoreAllBanners { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["restoreAllBanners"].(map[string]interface{})["status"])

	result = gqlBanner(s.graphqlH, `mutation { deleteAllBannersPermanent { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["deleteAllBannersPermanent"].(map[string]interface{})["status"])
}

func TestBannerGraphQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(BannerGraphQLTestSuite))
}
