package role_test

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

type RoleGraphQLTestSuite struct {
	tests.BaseTestSuite
	graphqlH http.Handler
	roleID   int
}

func (s *RoleGraphQLTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupRoleService()

	cacheStore := s.GetCacheStore()
	s.graphqlH = graphtest.NewTestHandler(s.Conns, cacheStore, s.Log)
}

func (s *RoleGraphQLTestSuite) TearDownSuite() {
	s.BaseTestSuite.TearDownSuite()
}

func gqlRole(h http.Handler, query string, variables map[string]interface{}) map[string]interface{} {
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

func (s *RoleGraphQLTestSuite) Test1_CreateRole() {
	result := gqlRole(s.graphqlH, `mutation { createRole(input: {name:"GraphQL Role"}) { status message data { id name } } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["createRole"].(map[string]interface{})["status"])
	data := result["data"].(map[string]interface{})["createRole"].(map[string]interface{})["data"].(map[string]interface{})
	s.Equal("GraphQL Role", data["name"])
	s.roleID = int(data["id"].(float64))
}

func (s *RoleGraphQLTestSuite) Test2_FindAllRole() {
	vars := map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}}
	result := gqlRole(s.graphqlH, `query FindAllRole($input: FindAllRoleInput) { findAllRole(input: $input) { status message data { id name } pagination { current_page page_size total_pages total_records } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["findAllRole"].(map[string]interface{})["status"])
}

func (s *RoleGraphQLTestSuite) Test3_FindByIdRole() {
	s.Require().NotZero(s.roleID)
	vars := map[string]interface{}{"input": map[string]interface{}{"role_id": s.roleID}}
	result := gqlRole(s.graphqlH, `query FindByIdRole($input: FindByIdRoleInput!) { findByIdRole(input: $input) { status message data { id name } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["findByIdRole"].(map[string]interface{})["status"])
}

func (s *RoleGraphQLTestSuite) Test4_UpdateRole() {
	s.Require().NotZero(s.roleID)
	vars := map[string]interface{}{"input": map[string]interface{}{"id": s.roleID, "name": "Updated GraphQL Role"}}
	result := gqlRole(s.graphqlH, `mutation UpdateRole($input: UpdateRoleInput!) { updateRole(input: $input) { status message data { id name } } }`, vars)
	s.Equal("success", result["data"].(map[string]interface{})["updateRole"].(map[string]interface{})["status"])
}

func (s *RoleGraphQLTestSuite) Test5_RestoreAllAndDeleteAll() {
	result := gqlRole(s.graphqlH, `mutation { restoreAllRole { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["restoreAllRole"].(map[string]interface{})["status"])

	result = gqlRole(s.graphqlH, `mutation { deleteAllRolePermanent { status message } }`, nil)
	s.Equal("success", result["data"].(map[string]interface{})["deleteAllRolePermanent"].(map[string]interface{})["status"])
}

func TestRoleGraphQLSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleGraphQLTestSuite))
}
