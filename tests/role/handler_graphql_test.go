package role_test

import (
	"net/http"
	"testing"

	tests "github.com/MamangRust/microservice-ecommerce-test"
	"github.com/stretchr/testify/suite"
)

type RoleGraphqlTestSuite struct {
	tests.BaseTestSuite
	handler http.Handler
	roleID  int
}

func (s *RoleGraphqlTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	s.SetupRoleService()

	s.handler = s.GraphQLHandler()
}

func (s *RoleGraphqlTestSuite) TestRoleGraphqlLifecycle() {
	// 1. Create
	create := s.GQL(s.handler, `mutation CreateRole($input: CreateRoleInput!) {
		createRole(input: $input) { status message data { id name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{"name": "API Role"},
	})
	createData := s.Obj(s.Obj(create, "createRole"), "data")
	s.roleID = int(createData["id"].(float64))
	s.Require().NotZero(s.roleID)
	s.Equal("API Role", createData["name"])

	// 2. FindAll
	all := s.GQL(s.handler, `query FindAllRole($input: FindAllRoleInput) {
		findAllRole(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(all, "findAllRole"), "data"))

	// 3. FindById
	byID := s.GQL(s.handler, `query FindByIdRole($input: FindByIdRoleInput!) {
		findByIdRole(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"role_id": s.roleID}})
	s.Equal(float64(s.roleID), s.Obj(s.Obj(byID, "findByIdRole"), "data")["id"])

	// 4. FindByActive
	active := s.GQL(s.handler, `query FindByActiveRole($input: FindAllRoleInput) {
		findByActiveRole(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(active, "findByActiveRole"), "data"))

	// 5. FindByUserId
	s.GQL(s.handler, `query FindByUserIdRole($input: FindByIdUserRoleInput!) {
		findByUserIdRole(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"user_id": 1}})

	// 6. Update
	updated := s.GQL(s.handler, `mutation UpdateRole($input: UpdateRoleInput!) {
		updateRole(input: $input) { status message data { id name } }
	}`, map[string]interface{}{
		"input": map[string]interface{}{"id": s.roleID, "name": "Updated API Role"},
	})
	s.Equal("Updated API Role", s.Obj(s.Obj(updated, "updateRole"), "data")["name"])

	// 7. Trash
	s.GQL(s.handler, `mutation TrashedRole($input: FindByIdRoleInput!) {
		trashedRole(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"role_id": s.roleID}})

	// 8. FindByTrashed
	trashed := s.GQL(s.handler, `query FindByTrashedRole($input: FindAllRoleInput) {
		findByTrashedRole(input: $input) { status message pagination { total_records } data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"page": 1, "page_size": 10}})
	s.NotEmpty(s.Arr(s.Obj(trashed, "findByTrashedRole"), "data"))

	// 9. Restore
	s.GQL(s.handler, `mutation RestoreRole($input: FindByIdRoleInput!) {
		restoreRole(input: $input) { status message data { id name } }
	}`, map[string]interface{}{"input": map[string]interface{}{"role_id": s.roleID}})

	// 10. DeletePermanent (trash first so the row is eligible)
	s.GQL(s.handler, `mutation TrashedRole($input: FindByIdRoleInput!) {
		trashedRole(input: $input) { status message data { id } }
	}`, map[string]interface{}{"input": map[string]interface{}{"role_id": s.roleID}})
	s.GQL(s.handler, `mutation DeleteRolePermanent($input: FindByIdRoleInput!) {
		deleteRolePermanent(input: $input) { status message }
	}`, map[string]interface{}{"input": map[string]interface{}{"role_id": s.roleID}})

	// 11. RestoreAll
	s.GQL(s.handler, `mutation { restoreAllRole { status message } }`, nil)

	// 12. DeleteAll
	s.GQL(s.handler, `mutation { deleteAllRolePermanent { status message } }`, nil)
}

func TestRoleGraphqlSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleGraphqlTestSuite))
}
