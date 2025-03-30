package repository

import (
	"context"

	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type authRepository struct {
	DBLive *pgxpool.Pool
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{
		DBLive: db,
	}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, email, password, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	row := r.DBLive.QueryRow(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == pgx.ErrNoRows {
		logger.LogError("User not found")
		return nil, nil
	} else if err != nil {
		logger.LogError("Error while get user by email, cause: " + err.Error())
		return nil, err
	}

	return &user, nil
}
