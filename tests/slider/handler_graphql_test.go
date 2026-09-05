package slider_test

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

type SliderGraphQLTestSuite struct {
	tests.BaseTestSuite
	graphqlH http.Handler
}

func (s *SliderGraphQLTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupSliderService()
	cacheStore := s.GetCacheStore()
	s.graphqlH = graphtest.NewTestHandler(s.Conns, cacheStore, s.Log)
}

func (s *SliderGraphQLTestSuite) TearDownSuite() { s.BaseTestSuite.TearDownSuite() }

func gqlSlider(h http.Handler, query string, variables map[string]interface{}) map[string]interface{} {
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

func (s *SliderGraphQLTestSuite) Test1_FindAllSlider() {
	vars := map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}}
	result := gqlSlider(s.graphqlH, `query($input: FindAllSliderRequest) { findAllSliders(input: $input) { status message pagination { current_page } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["findAllSliders"].(map[string]interface{})["status"])
}

func (s *SliderGraphQLTestSuite) Test2_RestoreAllAndDeleteAll() {
	result := gqlSlider(s.graphqlH, `mutation { restoreAllSliders { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["restoreAllSliders"].(map[string]interface{})["status"])
	result = gqlSlider(s.graphqlH, `mutation { deleteAllSlidersPermanent { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["deleteAllSlidersPermanent"].(map[string]interface{})["status"])
}

func TestSliderGraphQLSuite(t *testing.T) {
	if testing.Short() { t.Skip("skipping integration test") }
	suite.Run(t, new(SliderGraphQLTestSuite))
}
