---
name: audit-and-fill-test-gaps
description: Systematic audit of test coverage gaps by inspecting project conventions, then filling missing test files following the established patterns.
source: auto-skill
extracted_at: '2026-06-28T05:43:05.200Z'
---

# Audit and Fill Test Coverage Gaps

Use this skill when the user asks to "continue writing tests", "add missing tests", "complete test coverage", or similar. It systematically discovers the project's test conventions, identifies gaps, and fills missing test files.

## Step 1: Discover established test patterns

Read representative examples of each test layer to understand the project's conventions:

```
tests/<service>/repository_test.go
tests/<service>/service_test.go
tests/<service>/handler_gapi_test.go    # gRPC
tests/<service>/handler_api_test.go     # HTTP / Echo
```

Key traits to note:
- Is it `testify/suite` or `testing.T`? 
- Does it use `testing.Short()` guard?
- Internal vs external test package (`package xxx` vs `package xxx_test`)?
- Mock vs real infrastructure (testcontainers, real DB)?
- Setup/teardown pattern (shared `BaseTestSuite`?)
- Orchestration helpers for dependency services

**Why:** Every project has its own conventions. Following them verbatim avoids style-review friction.

## Step 2: Read shared test infrastructure

```
tests/base.go          # BaseTestSuite
tests/test_setup.go    # Container lifecycle
tests/orchestration_helpers.go  # Service bootstrapping
tests/go.mod           # Dependencies + replace directives
```

Understand:
- What `BaseTestSuite` provides (DBPool, RedisClient, GetCacheStore, RegisterServer, GetConnection, Conns map, SeedXXX helpers)
- How services are wired (each `SetupXxxService()` registers a gRPC server + stores conn in `s.Conns["xxx"]`)
- The `GOWORK=off` requirement (because replace directives conflict with go.work)

## Step 3: Audit all services for gaps

Glob all test files and all service modules:

```
tests/**/*_test.go
service/*/go.mod
```

Create a matrix of service → test files to identify which layer(s) are missing.

For each service, expect:
- `repository_test.go`
- `service_test.go`
- `handler_gapi_test.go`
- `handler_api_test.go`

Some services also have stats sub-domains (4 extra files with `stats_` prefix).

Services like `migrate`, `seeder`, `email` (Kafka consumer) may not need the standard 4-layer pattern — verify this before filing them as gaps.

## Step 4: Read the missing layer's source code

For the missing test file, read:
- The HTTP handler source (routes, method signatures, request/response types)
- Request DTOs (`shared/domain/requests/`)
- Response DTOs (`shared/domain/response/`)
- Mapper (`shared/mapper/<service>/`)
- gRPC proto-generated types (`shared/pb/`)
- Any existing tests in the same service for reference

## Step 5: Write the missing test file

Follow the exact patterns from Step 1. For HTTP handler tests (`handler_api_test.go`):

- **If handlers use `c.Get("user_id")` (JWT auth)**: add an Echo middleware in SetupSuite that sets the value:
  ```go
  s.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
      return func(c echo.Context) error {
          c.Set("user_id", s.userID)
          return next(c)
      }
  })
  ```
- **Seed a user** for auth context and pass its ID as the "authenticated" user
- **Seed all dependent entities** needed for request payloads (category, merchant, product, etc.)
- **Test the full lifecycle** in one method (Create → FindAll → Update → Delete) following the project's pattern

## Step 6: Verify

```bash
cd tests && GOWORK=off go build ./<service>/...
cd tests && GOWORK=off go vet ./<service>/...
```

Fix any unused imports or compilation errors found by vet.

## Upgrade Repository Tests to Full Lifecycle Coverage

Use this when the user asks to "complete repository tests", "add lifecycle tests", or existing repo tests only cover Create/FindByID.

### Step 1: Research repository interfaces

For each service, read the repository interface definition:

```
service/<name>/repository/interfaces.go
```

Build a matrix of available methods:

```
Command: Create, Update, Trash, Restore, DeletePermanent, RestoreAll, DeleteAll
Query:   FindByID, FindAll, FindActive, FindTrashed
```

Also check shared request structs in `shared/domain/requests/<service>.go` for Update types and pagination/filter types.

### Step 2: Identify lifecycle outliers

Not all services support the full soft-delete lifecycle:

| Outlier     | What's missing | Lifecycle to use |
|-------------|---------------|------------------|
| **cart**    | No Update, Trash, Restore, FindByID, FindActive, FindTrashed — only CreateCart, FindCarts, DeletePermanent, DeleteAllPermanently | Create → Find → DeletePermanent → DeleteAllPermanently → Verify empty |
| **order_item** | No FindByID | Use FindOrderItemByOrder / FindAll instead of FindByID |
| **user**    | Extra methods: UpdatePassword, UpdateIsVerified, FindByEmail, FindByEmailWithPassword | Standard lifecycle + extra numbered tests |

### Step 3: Write the lifecycle test pattern

Use a single `TestXxxLifecycle()` method (preferred) or numbered methods. The standard lifecycle for most services is:

```go
pageReq := &requests.FindAllXxx{Search: "test", Page: 1, PageSize: 10}

// 1. Create
created, err := s.repo.XxxCommand.Create(ctx, req)
s.NoError(err)
s.NotNil(created)
entityID := int(created.XxxID)

// 2. FindByID
found, err := s.repo.XxxQuery.FindByID(ctx, entityID)
s.NoError(err)

// 3. FindAll
all, err := s.repo.XxxQuery.FindAll(ctx, pageReq)
s.NoError(err)
s.NotEmpty(all)

// 4. FindActive (before trash — should appear)
active, err := s.repo.XxxQuery.FindActive(ctx, pageReq)
s.NoError(err)
s.NotEmpty(active)

// 5. Update
updateReq := &requests.UpdateXxxRequest{XxxID: &entityID, ...}
updated, err := s.repo.XxxCommand.Update(ctx, updateReq)
s.NoError(err)
s.NotNil(updated)

// 6. Trash
trashed, err := s.repo.XxxCommand.Trash(ctx, entityID)
s.NoError(err)
s.NotNil(trashed)

// 7. FindTrashed
trashedItems, err := s.repo.XxxQuery.FindTrashed(ctx, pageReq)
s.NoError(err)
s.NotEmpty(trashedItems)

// 8. FindActive after trash — should NOT include
activeAfterTrash, err := s.repo.XxxQuery.FindActive(ctx, pageReq)
s.NoError(err)
for _, item := range activeAfterTrash {
    s.NotEqual(entityID, int(item.XxxID))
}

// 9. Restore
restored, err := s.repo.XxxCommand.Restore(ctx, entityID)
s.NoError(err)
s.NotNil(restored)

// 10. Trash again → DeletePermanent
_, err = s.repo.XxxCommand.Trash(ctx, entityID)
s.Require().NoError(err)
success, err := s.repo.XxxCommand.DeletePermanent(ctx, entityID)
s.NoError(err)
s.True(success)

// 11-12. RestoreAll + DeleteAll (use a second entity)
```

### Step 4: For paginated queries, always provide `Search`

Pagination structs like `FindAllXxx` have `Search string` with `validate:"required"`. At the repository level, the validation tag isn't enforced (it's a service-layer concern), but always provide `Search: "test"` to be safe.

### Step 5: Handle `TearDownSuite`

Every upgraded test must include a `TearDownSuite()` override:

```go
func (s *XxxRepositoryTestSuite) TearDownSuite() {
    s.BaseTestSuite.TearDownSuite()
}
```

### Step 6: Verify

```bash
# Single service
cd tests && GOWORK=off go vet ./<service>/

# All at once
cd tests && GOWORK=off go vet ./slider/ ./merchant_detail/ ./merchant_policy/ ./merchant_award/ ./merchant_business/ ./review_detail/ ./order/ ./shipping_address/ ./review/ ./product/ ./transaction/ ./cart/ ./order_item/ ./user/
```

Common vet issues to fix:
- **Unused variable**: `s.SeedProduct()` returns a value you don't need — use `_` or just call without capturing
- **Wrong return count**: Some repository methods return 2 values (rows, error), not 3 with a total. Check the actual interface
- **Non-existent field**: Generated sqlc types may not have the fields you expect from the service layer. Check `db.*Row` types in `pkg/database/schema/`

## Common pitfalls in this project

| Pitfall | Fix |
|---------|-----|
| `GOWORK=off` not set | Always use `GOWORK=off` when running from `tests/` |
| Cache import ambiguity | Cart uses `github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache/cart` for API handler, but `github.com/MamangRust/microservice-ecommerce-grpc-cart/cache` for service/repo |
| Auth middleware missing | Cart, merchant, review handlers extract `c.Get("user_id").(int)` — must inject via Echo middleware in test |
| Route naming quirk | Cart has inverted route groups: command endpoints under `/api/cart-query/`, query endpoint under `/api/cart-command/` |
| `s.Require()` vs `s.Assert()` | Use `s.Require()` for setup/preconditions, `s.*` (soft) for test assertions |
| `db.*Row` types differ from service DTOs | Repository methods return sqlc-generated types from `pkg/database/schema/`. Don't assume they have the same fields as service-layer response types. Grep the actual `.sqlc.gen.go` file to verify field names |
