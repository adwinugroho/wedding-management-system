package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

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
	// TODO: to secure token when user already logout
	BlacklistToken(ctx context.Context, token string, userID string, expiresAt time.Time) error
	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
	CleanupExpiredBlacklistedTokens(ctx context.Context) error
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{
		DBLive: db,
	}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, email, password, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	row := r.DBLive.QueryRow(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err == pgx.ErrNoRows {
		logger.LogError("User not found")
		return nil, nil
	} else if err != nil {
		logger.LogError("Error while get user by email, cause: " + err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) BlacklistToken(ctx context.Context, token string, userID string, expiresAt time.Time) error {
	// Hash the token for storage
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	query := `
		INSERT INTO blacklisted_tokens (token_hash, user_id, expires_at, reason)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (token_hash) DO NOTHING
	`

	_, err := r.DBLive.Exec(ctx, query, tokenHash, userID, expiresAt, "logout")
	if err != nil {
		logger.LogError("Error while blacklisting token, cause: " + err.Error())
		return err
	}

	return nil
}

func (r *authRepository) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	// Hash the token for lookup
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	query := `
		SELECT EXISTS(
			SELECT 1 FROM blacklisted_tokens 
			WHERE token_hash = $1 AND expires_at > now()
		)
	`

	var exists bool
	err := r.DBLive.QueryRow(ctx, query, tokenHash).Scan(&exists)
	if err != nil {
		logger.LogError("Error while checking if token is blacklisted, cause: " + err.Error())
		return false, err
	}

	return exists, nil
}

func (r *authRepository) CleanupExpiredBlacklistedTokens(ctx context.Context) error {
	query := `DELETE FROM blacklisted_tokens WHERE expires_at <= now()`

	_, err := r.DBLive.Exec(ctx, query)
	if err != nil {
		logger.LogError("Error while cleaning up expired blacklisted tokens, cause: " + err.Error())
		return err
	}

	return nil
}
