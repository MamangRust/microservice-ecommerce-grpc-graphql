package graph

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	mycontext "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/context"
	"github.com/MamangRust/microservice-ecommerce-grpc-apigateway/internal/middlewares"
	"github.com/vektah/gqlparser/v2/ast"
)

// adminOnly is the RBAC policy applied to role administration, mirroring the
// REST gateway (RequireRoles("Admin", "ROLE_ADMIN") on /api/role-query and
// /api/role-command).
var adminOnly = []string{"Admin", "ROLE_ADMIN"}

func expectedHasRoleFields() []string {
	return []string{
		"findAllRole",
		"findByIdRole",
		"findByActiveRole",
		"findByTrashedRole",
		"findByUserIdRole",
		"createRole",
		"updateRole",
		"trashedRole",
		"restoreRole",
		"deleteRolePermanent",
		"restoreAllRole",
		"deleteAllRolePermanent",
	}
}

func declaredRoles(field *ast.FieldDefinition) []string {
	directive := field.Directives.ForName("hasRole")
	if directive == nil {
		return nil
	}

	arg := directive.Arguments.ForName("roles")
	if arg == nil {
		return []string{}
	}

	roles := make([]string, 0, len(arg.Value.Children))
	for _, child := range arg.Value.Children {
		roles = append(roles, child.Value.Raw)
	}
	sort.Strings(roles)
	return roles
}

// TestHasRoleDirectiveCoversRoleAdministration locks the RBAC policy to role
// administration: every role query/mutation must carry @hasRole and no other
// root field may be annotated by accident.
func TestHasRoleDirectiveCoversRoleAdministration(t *testing.T) {
	schema := NewExecutableSchema(Config{Resolvers: &Resolver{}}).Schema()

	want := expectedHasRoleFields()
	sortedWant := append([]string(nil), adminOnly...)
	sort.Strings(sortedWant)

	got := map[string][]string{}
	for _, root := range []*ast.Definition{schema.Query, schema.Mutation} {
		if root == nil {
			continue
		}
		for _, field := range root.Fields {
			if roles := declaredRoles(field); len(roles) > 0 {
				got[field.Name] = roles
			}
		}
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d fields annotated with @hasRole, found %d: %v",
			len(want), len(got), sortedKeys(got))
	}

	for _, name := range want {
		roles, ok := got[name]
		if !ok {
			t.Errorf("expected %s to be annotated with @hasRole", name)
			continue
		}
		if strings.Join(roles, ",") != strings.Join(sortedWant, ",") {
			t.Errorf("expected @hasRole(roles: %v) on %s, got %v", adminOnly, name, roles)
		}
	}
}

func sortedKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

type roleCheckerFunc func(ctx context.Context, userID int, requiredRoles ...string) error

func (f roleCheckerFunc) CheckRole(ctx context.Context, userID int, requiredRoles ...string) error {
	return f(ctx, userID, requiredRoles...)
}

func withUser(next http.Handler, userID int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(mycontext.WithUserID(r.Context(), userID)))
	})
}

func rbacHandler(checker middlewares.RoleChecker) http.Handler {
	srv := handler.New(NewExecutableSchema(Config{
		Resolvers:  &Resolver{},
		Directives: DirectiveRoot{HasRole: middlewares.HasRole(checker)},
	}))
	srv.AddTransport(transport.POST{})
	return srv
}

func executeGraphQL(t *testing.T, h http.Handler, query string) map[string]interface{} {
	t.Helper()

	body, err := json.Marshal(map[string]interface{}{"query": query})
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/query", strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	var resp map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response %q: %v", rec.Body.String(), err)
	}
	return resp
}

// TestHasRoleDirectiveDeniesNonAdminThroughExecutor proves the schema
// annotation is enforced by the executor: an authenticated user without the
// required role gets a forbidden error and the resolver never runs.
func TestHasRoleDirectiveDeniesNonAdminThroughExecutor(t *testing.T) {
	var checkedUser int
	var checkedRoles []string

	checker := roleCheckerFunc(func(_ context.Context, userID int, requiredRoles ...string) error {
		checkedUser = userID
		checkedRoles = requiredRoles
		return errors.New("role not permitted")
	})

	h := withUser(rbacHandler(checker), 5)

	resp := executeGraphQL(t, h, `mutation { createRole(input: {name: "Sneaky"}) { status message } }`)

	if data, ok := resp["data"].(map[string]interface{}); ok && data["createRole"] != nil {
		t.Fatalf("expected no data for an unauthorised user, got %v", data["createRole"])
	}

	errs, ok := resp["errors"].([]interface{})
	if !ok || len(errs) == 0 {
		t.Fatalf("expected a forbidden error, got %v", resp)
	}

	message, _ := errs[0].(map[string]interface{})["message"].(string)
	if !strings.HasPrefix(message, "forbidden:") {
		t.Fatalf("expected a 'forbidden:' error, got %q", message)
	}

	if checkedUser != 5 {
		t.Fatalf("expected the checker to receive user 5, got %d", checkedUser)
	}
	if strings.Join(checkedRoles, ",") != "Admin,ROLE_ADMIN" {
		t.Fatalf("expected the schema roles to be enforced, got %v", checkedRoles)
	}
}

// TestHasRoleDirectiveSkipsUnannotatedField makes sure the directive only
// guards the annotated role administration fields.
func TestHasRoleDirectiveSkipsUnannotatedField(t *testing.T) {
	calls := 0
	checker := roleCheckerFunc(func(context.Context, int, ...string) error {
		calls++
		return nil
	})

	h := withUser(rbacHandler(checker), 5)

	// findUserById is not annotated: the directive must not run. The request is
	// expected to fail on the bare resolver's missing dependencies, which is
	// enough — what matters is that RBAC was not consulted.
	resp := executeGraphQL(t, h, `query { findUserById(input: {user_id: 1}) { status message } }`)
	if calls != 0 {
		t.Fatalf("expected the role checker not to run for an unannotated field, ran %d times", calls)
	}
	if _, ok := resp["data"]; !ok && resp["errors"] == nil {
		t.Fatalf("expected a GraphQL response, got %v", resp)
	}
}
