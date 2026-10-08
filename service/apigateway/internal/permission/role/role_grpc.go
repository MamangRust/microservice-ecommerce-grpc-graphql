package rolepermission

import (
	"context"
	"errors"
	"strconv"
	"time"

	apicache "github.com/MamangRust/microservice-ecommerce-grpc-apigateway/cache"
	pb_role "github.com/MamangRust/microservice-ecommerce-grpc-pb/role"
	pb_user_role "github.com/MamangRust/microservice-ecommerce-grpc-pb/user_role"
	"github.com/MamangRust/microservice-ecommerce-pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// defaultRoleTimeout bounds a single role lookup against the role service.
const defaultRoleTimeout = 5 * time.Second

// RoleQueryClient is the slice of the role gRPC service needed to resolve a
// user's roles. pb_user_role.UserRoleQueryServiceClient satisfies it.
type RoleQueryClient interface {
	FindByUserId(ctx context.Context, in *pb_user_role.FindByIdUserRoleRequest, opts ...grpc.CallOption) (*pb_role.ApiResponsesRole, error)
}

type rolePermissionGRPC struct {
	client  RoleQueryClient
	logger  logger.LoggerInterface
	cache   apicache.RoleCache
	timeout time.Duration
}

// NewRolePermissionGRPC resolves a user's roles through the role service's gRPC
// FindByUserId and caches the result in Redis, so the gateway can authorise
// GraphQL operations with `@hasRole`.
//
// This is the transport used by the `@hasRole` directive: unlike the Kafka
// request/response handshake (NewRolePermission, topics request-role /
// response-role) it fails fast on a single gRPC call instead of holding a 5s
// timeout, and it needs no broker-side consumer to be running before the
// gateway can authorise admin operations. It is the same pattern as
// RoleValidatorGRPC in the REST gateway and as the pointofsale GraphQL gateway.
func NewRolePermissionGRPC(client RoleQueryClient, logger logger.LoggerInterface, cache apicache.RoleCache) RolePermission {
	return &rolePermissionGRPC{
		client:  client,
		logger:  logger,
		cache:   cache,
		timeout: defaultRoleTimeout,
	}
}

func (p *rolePermissionGRPC) ValidateRole(ctx context.Context, userID int) ([]string, error) {
	cacheKey := strconv.Itoa(userID)

	if p.cache != nil {
		if roles, found := p.cache.GetRoleCache(ctx, cacheKey); found {
			p.logger.Debug("Role found in cache", zap.Int("user_id", userID), zap.Strings("roles", roles))
			return roles, nil
		}
	}

	if p.client == nil {
		p.logger.Error("Role validation client is not configured", zap.Int("user_id", userID))
		return nil, errors.New("role validation failed")
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	res, err := p.client.FindByUserId(ctxTimeout, &pb_user_role.FindByIdUserRoleRequest{UserId: int32(userID)})
	if err != nil {
		p.logger.Error("Role validation via gRPC failed", zap.Int("user_id", userID), zap.Error(err))
		return nil, errors.New("role validation failed")
	}

	roles := make([]string, 0, len(res.GetData()))
	for _, role := range res.GetData() {
		roles = append(roles, role.GetName())
	}

	if len(roles) == 0 {
		p.logger.Debug("Role validation failed (no roles)", zap.Int("user_id", userID))
		return nil, errors.New("role validation failed")
	}

	if p.cache != nil {
		p.cache.SetRoleCache(ctx, cacheKey, roles)
	}

	return roles, nil
}

func (p *rolePermissionGRPC) CheckRole(ctx context.Context, userID int, requiredRoles ...string) error {
	roles, err := p.ValidateRole(ctx, userID)
	if err != nil {
		return err
	}

	for _, userRole := range roles {
		for _, required := range requiredRoles {
			if userRole == required {
				return nil
			}
		}
	}

	p.logger.Debug("User does not have required role",
		zap.Int("user_id", userID),
		zap.Strings("user_roles", roles),
		zap.Strings("required_roles", requiredRoles))

	return errors.New("role not permitted")
}
