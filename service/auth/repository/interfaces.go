package repository

import (
	"context"

	db "github.com/MamangRust/microservice-ecommerce-auth/database/schema"
	dto "github.com/MamangRust/microservice-ecommerce-auth/dto"
	"github.com/MamangRust/microservice-ecommerce-shared/domain/requests"
	"github.com/jackc/pgx/v5"
)

// UserRepository defines the data access layer for user-related operations.
//
//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*dto.User, error)

	FindByEmailAndVerify(ctx context.Context, email string) (*dto.GetUserByEmailAndVerifyRow, error)

	FindById(ctx context.Context, user_id int) (*dto.GetUserByIDRow, error)

	CreateUser(ctx context.Context, request *requests.RegisterRequest) (*dto.CreateUserRow, error)

	UpdateUserIsVerified(ctx context.Context, user_id int, is_verified bool) (*dto.UpdateUserIsVerifiedRow, error)

	UpdateUserPassword(ctx context.Context, user_id int, password string) (*dto.UpdateUserPasswordRow, error)

	FindByVerificationCode(ctx context.Context, verification_code string) (*dto.GetUserByVerificationCodeRow, error)
}

type ResetTokenRepository interface {
	FindByToken(ctx context.Context, code string) (*db.ResetToken, error)

	CreateResetToken(ctx context.Context, req *requests.CreateResetTokenRequest) (*db.ResetToken, error)

	CreateResetTokenInTx(ctx context.Context, tx pgx.Tx, req *requests.CreateResetTokenRequest) (*db.ResetToken, error)

	DeleteResetToken(ctx context.Context, user_id int) error
}

type RefreshTokenRepository interface {
	FindByToken(ctx context.Context, token string) (*db.RefreshToken, error)

	FindByUserId(ctx context.Context, user_id int) (*db.RefreshToken, error)

	CreateRefreshToken(ctx context.Context, req *requests.CreateRefreshToken) (*db.RefreshToken, error)

	UpdateRefreshToken(ctx context.Context, req *requests.UpdateRefreshToken) (*db.RefreshToken, error)

	DeleteRefreshToken(ctx context.Context, token string) error

	DeleteRefreshTokenByUserId(ctx context.Context, user_id int) error
}

type UserRoleRepository interface {
	AssignRoleToUser(ctx context.Context, req *requests.CreateUserRoleRequest) (*dto.UserRole, error)

	RemoveRoleFromUser(ctx context.Context, req *requests.RemoveUserRoleRequest) error
}

type RoleRepository interface {
	FindById(ctx context.Context, id int) (*dto.Role, error)

	FindByName(ctx context.Context, name string) (*dto.Role, error)
}
