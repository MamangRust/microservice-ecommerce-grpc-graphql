package auth_test

import (
	"context"
	"net/http"
	"testing"

	gt "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/graphtest"
	pb_role "github.com/MamangRust/microservice-ecommerce-grpc-pb/role"
	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type AuthGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler     http.Handler
	email       string
	password    string
	accessToken string
	userID      int
}

func (s *AuthGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()

	// Order matters: the user service depends on the role service, and the auth
	// service depends on both the user and role services.
	s.SetupRoleService()
	s.SetupUserService()
	s.SetupAuthService()

	// Registration assigns the default ROLE_ADMIN role, so it must exist first.
	_, err := pb_role.NewRoleCommandServiceClient(s.Conns["role"]).CreateRole(
		context.Background(), &pb_role.CreateRoleRequest{Name: "ROLE_ADMIN"})
	s.Require().NoError(err)

	s.handler = s.GraphQLHandler()

	s.email = "auth.handler.graphql.test@example.com"
	s.password = "password123"
}

func (s *AuthGraphqlTestSuite) TestAuthGraphqlLifecycle() {
	// 1. Register
	register := s.GQL(s.handler, `mutation RegisterUser($input: RegisterInput!) {
		registerUser(input: $input) { status message data { id firstname lastname email } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"firstname":        "Auth",
			"lastname":         "GraphQL",
			"email":            s.email,
			"password":         s.password,
			"confirm_password": s.password,
		},
	})
	registerData := s.Obj(s.Obj(register, "registerUser"), "data")
	s.userID = int(registerData["id"].(float64))
	s.Require().NotZero(s.userID)
	s.Equal(s.email, registerData["email"])

	// 2. Login
	login := s.GQL(s.handler, `mutation LoginUser($input: LoginInput!) {
		loginUser(input: $input) { status message data { access_token refresh_token } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"email":    s.email,
			"password": s.password,
		},
	})
	loginData := s.Obj(s.Obj(login, "loginUser"), "data")
	s.accessToken = loginData["access_token"].(string)
	s.Require().NotEmpty(s.accessToken)

	// 3. GetMe — the resolver reads the authenticated user from the request
	// context (populated by AuthMiddleware in the real gateway), so wrap the
	// handler with WithUser to emulate an authenticated caller.
	getMe := s.GQL(gt.WithUser(s.handler, s.userID), `query GetMe($input: GetMeInput!) {
		getMe(input: $input) { status message data { id firstname lastname email } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"access_token": s.accessToken,
		},
	})
	getMeData := s.Obj(s.Obj(getMe, "getMe"), "data")
	s.Equal(s.email, getMeData["email"])
	s.Equal(float64(s.userID), getMeData["id"])
}

func TestAuthGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(AuthGraphqlTestSuite))
}
