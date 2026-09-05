package user_test

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

type UserGraphQLTestSuite struct {
	tests.BaseTestSuite
	graphqlH http.Handler
	userID   int
}

func (s *UserGraphQLTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupRoleService()
	s.SetupUserService()

	cacheStore := s.GetCacheStore()
	s.graphqlH = graphtest.NewTestHandler(s.Conns, cacheStore, s.Log)
}

func (s *UserGraphQLTestSuite) TearDownSuite() {
	s.BaseTestSuite.TearDownSuite()
}

func gqlUser(h http.Handler, query string, variables map[string]interface{}) map[string]interface{} {
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

func (s *UserGraphQLTestSuite) Test1_CreateUser() {
	result := gqlUser(s.graphqlH, `mutation { createUser(input: {firstname:"GQL",lastname:"User",email:"gql.user@test.com",password:"pass123",confirm_password:"pass123"}) { status data { id email } } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["createUser"].(map[string]interface{})["status"])
}

func (s *UserGraphQLTestSuite) Test2_FindAllUser() {
	vars := map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}}
	result := gqlUser(s.graphqlH, `query FindAllUser($input: FindAllUserInput) { findAllUsers(input: $input) { status message data { id email } pagination { current_page page_size total_pages total_records } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["findAllUsers"].(map[string]interface{})["status"])
}

func (s *UserGraphQLTestSuite) Test3_RestoreAllAndDeleteAll() {
	result := gqlUser(s.graphqlH, `mutation { restoreAllUser { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["restoreAllUser"].(map[string]interface{})["status"])

	result = gqlUser(s.graphqlH, `mutation { deleteAllUserPermanent { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["deleteAllUserPermanent"].(map[string]interface{})["status"])
}

func TestUserGraphQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(UserGraphQLTestSuite))
}
