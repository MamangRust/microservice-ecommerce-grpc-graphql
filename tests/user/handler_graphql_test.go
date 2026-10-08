package user_test

import (
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type UserGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler   http.Handler
	userID    int
	userEmail string
}

func (s *UserGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()

	// The user service resolves roles through the role service.
	s.SetupRoleService()
	s.SetupUserService()

	s.handler = s.GraphQLHandler()
}

func (s *UserGraphqlTestSuite) TestUserGraphqlLifecycle() {
	// 1. Create
	s.userEmail = "handler.user.graphql@example.com"
	create := s.GQL(s.handler, `mutation CreateUser($input: CreateUserInput!) {
		createUser(input: $input) { status message data { id firstname lastname email } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"firstname":        "Handler",
			"lastname":         "User",
			"email":            s.userEmail,
			"password":         "password123",
			"confirm_password": "password123",
		},
	})
	createData := s.Obj(s.Obj(create, "createUser"), "data")
	s.userID = int(createData["id"].(float64))
	s.Require().NotZero(s.userID)
	s.Equal(s.userEmail, createData["email"])

	// 2. FindAll
	all := s.GQL(s.handler, `query FindAllUsers($input: FindAllUserInput) {
		findAllUsers(input: $input) { status message pagination { total_records } data { id email } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllUsers"), "data"))

	// 3. FindById
	byID := s.GQL(s.handler, `query FindByIdUser($input: FindByIdUserInput!) {
		findByIdUser(input: $input) { status message data { id email } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.userID}})
	s.Equal(float64(s.userID), s.Obj(s.Obj(byID, "findByIdUser"), "data")["id"])

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindByActiveUsers($input: FindAllUserInput) {
		findByActiveUsers(input: $input) { status message pagination { total_records } data { id email } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findByActiveUsers"), "data"))

	// 5. Update
	updated := s.GQL(s.handler, `mutation UpdateUser($input: UpdateUserInput!) {
		updateUser(input: $input) { status message data { id firstname lastname email } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{
			"id":               s.userID,
			"firstname":        "Updated",
			"lastname":         "User",
			"email":            s.userEmail,
			"password":         "password123",
			"confirm_password": "password123",
		},
	})
	s.Equal("Updated", s.Obj(s.Obj(updated, "updateUser"), "data")["firstname"])

	// 6. Trash
	s.GQL(s.handler, `mutation TrashedUser($input: FindByIdUserInput!) {
		trashedUser(input: $input) { status message data { id email } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.userID}})

	// 7. FindByTrashed
	trashed := s.GQL(s.handler, `query FindByTrashedUsers($input: FindAllUserInput) {
		findByTrashedUsers(input: $input) { status message pagination { total_records } data { id email } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findByTrashedUsers"), "data"))

	// 8. Restore
	s.GQL(s.handler, `mutation RestoreUser($input: FindByIdUserInput!) {
		restoreUser(input: $input) { status message data { id email } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.userID}})

	// 9. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashedUser($input: FindByIdUserInput!) {
		trashedUser(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.userID}})
	s.GQL(s.handler, `mutation DeleteUserPermanent($input: FindByIdUserInput!) {
		deleteUserPermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"id": s.userID}})

	// 10. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllUser { status message } }`, nil)

	// 11. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllUserPermanent { status message } }`, nil)
}

func TestUserGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(UserGraphqlTestSuite))
}
