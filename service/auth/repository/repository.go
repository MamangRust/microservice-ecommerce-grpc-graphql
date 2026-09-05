package repository

import (
	db "github.com/MamangRust/microservice-ecommerce-auth/database/schema"
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
)

type Repositories struct {
	User         UserRepository
	RefreshToken RefreshTokenRepository
	UserRole     UserRoleRepository
	Role         RoleRepository
	ResetToken   ResetTokenRepository
}

func NewRepositories(DB *db.Queries,
	userQuery pb.UserQueryServiceClient,
	userCommand pb.UserCommandServiceClient,
	roleQuery pb.RoleQueryServiceClient,
	roleCommand pb.RoleCommandServiceClient,
) *Repositories {
	return &Repositories{
		User:         NewUserRepository(userQuery, userCommand),
		RefreshToken: NewRefreshTokenRepository(DB),
		UserRole:     NewUserRoleRepository(roleCommand),
		Role:         NewRoleRepository(roleQuery),
		ResetToken:   NewResetTokenRepository(DB),
	}
}
