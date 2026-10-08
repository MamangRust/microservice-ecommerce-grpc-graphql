package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"

	gt "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/graphtest"
)

// GraphQLHandler builds the apigateway GraphQL HTTP handler backed by the gRPC
// service connections the suite has registered in Conns. Suites use this
// instead of the old REST handlers now that the gateway only speaks GraphQL.
func (s *BaseTestSuite) GraphQLHandler() http.Handler {
	return gt.NewTestHandler(s.Conns, s.GetCacheStore(), s.Log)
}

// GQL executes a GraphQL operation against handler and returns the top-level
// "data" object. It fails the test on transport errors or GraphQL errors.
func (s *BaseTestSuite) GQL(handler http.Handler, query string, variables map[string]interface{}) map[string]interface{} {
	resp, err := gt.ExecuteGraphQL(handler, query, variables, "")
	s.Require().NoError(err)
	s.Require().Empty(resp.Errors, "unexpected GraphQL errors: %+v", resp.Errors)
	s.Require().NotNil(resp.Data, "GraphQL response contained no data")
	return resp.Data
}

// GQLRaw executes a GraphQL operation and returns the raw response without
// asserting that it succeeded. Use it for negative tests.
func (s *BaseTestSuite) GQLRaw(handler http.Handler, query string, variables map[string]interface{}) *gt.GraphQLResponse {
	resp, err := gt.ExecuteGraphQL(handler, query, variables, "")
	s.Require().NoError(err)
	return resp
}

// GQLExpectError executes a GraphQL operation that is expected to fail and
// returns the resulting GraphQL errors.
func (s *BaseTestSuite) GQLExpectError(handler http.Handler, query string, variables map[string]interface{}) []gt.GraphQLError {
	resp := s.GQLRaw(handler, query, variables)
	s.Require().NotEmpty(resp.Errors, "expected GraphQL errors, got data: %+v", resp.Data)
	return resp.Errors
}

// UploadFile describes one file part of a GraphQL multipart request.
type UploadFile struct {
	// VariablePath is the dot path into the variables object, e.g. "input.image_category".
	VariablePath string
	Filename     string
	Content      []byte
}

// GQLMultipart executes a GraphQL operation that carries file uploads using the
// GraphQL multipart request spec, which is what the apigateway's Upload scalar
// requires. The corresponding variable (see UploadFile.VariablePath) must be
// present and nil in variables.
func (s *BaseTestSuite) GQLMultipart(handler http.Handler, query string, variables map[string]interface{}, uploads []UploadFile) map[string]interface{} {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	operations := map[string]interface{}{"query": query}
	if variables != nil {
		operations["variables"] = variables
	}
	opsJSON, err := json.Marshal(operations)
	s.Require().NoError(err)
	s.Require().NoError(writer.WriteField("operations", string(opsJSON)))

	fileMap := map[string][]string{}
	for i, up := range uploads {
		fileMap[strconv.Itoa(i)] = []string{"variables." + up.VariablePath}
	}
	mapJSON, err := json.Marshal(fileMap)
	s.Require().NoError(err)
	s.Require().NoError(writer.WriteField("map", string(mapJSON)))

	for i, up := range uploads {
		part, err := writer.CreateFormFile(strconv.Itoa(i), up.Filename)
		s.Require().NoError(err)
		_, err = part.Write(up.Content)
		s.Require().NoError(err)
	}
	s.Require().NoError(writer.Close())

	req := httptest.NewRequest(http.MethodPost, "/query", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	s.Require().NoError(err)

	var resp gt.GraphQLResponse
	s.Require().NoError(json.Unmarshal(body, &resp), "body: %s", string(body))
	s.Require().Empty(resp.Errors, "unexpected GraphQL errors: %+v", resp.Errors)
	s.Require().NotNil(resp.Data, "GraphQL response contained no data")
	return resp.Data
}

// Obj asserts that key in m holds a JSON object and returns it.
func (s *BaseTestSuite) Obj(m map[string]interface{}, key string) map[string]interface{} {
	v, ok := m[key].(map[string]interface{})
	s.Require().True(ok, "expected %q to be an object, got %T", key, m[key])
	return v
}

// Arr asserts that key in m holds a JSON array and returns it.
func (s *BaseTestSuite) Arr(m map[string]interface{}, key string) []interface{} {
	v, ok := m[key].([]interface{})
	s.Require().True(ok, "expected %q to be an array, got %T", key, m[key])
	return v
}
