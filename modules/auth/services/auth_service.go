package services

import (
	"context"
	"time"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/adwinugroho/wedding-management-system/modules/auth/repository"
	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	GenerateJWTToken(ctx context.Context, user models.User) (*models.User, string, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	BlacklistToken(ctx context.Context, token string, userID string, expiresAt time.Time) error
	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
	CleanupExpiredBlacklistedTokens(ctx context.Context) error
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{authRepository: authRepository}
}

func (s *authService) GenerateJWTToken(ctx context.Context, user models.User) (*models.User, string, error) {
	token, err := s.generateJWTToken(user)
	if err != nil {
		logger.LogError("Error while generate JWT token, cause: " + err.Error())
		return nil, "", models.NewError("500-General-Error", "Internal Server Error.")
	}

	return &user, token, nil
}

func (s *authService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.authRepository.GetUserByEmail(ctx, email)
}

func (s *authService) BlacklistToken(ctx context.Context, token string, userID string, expiresAt time.Time) error {
	return s.authRepository.BlacklistToken(ctx, token, userID, expiresAt)
}

func (s *authService) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	return s.authRepository.IsTokenBlacklisted(ctx, token)
}

func (s *authService) CleanupExpiredBlacklistedTokens(ctx context.Context) error {
	return s.authRepository.CleanupExpiredBlacklistedTokens(ctx)
}

func (s *authService) generateJWTToken(payload models.User) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":   payload.ID,
		"role": payload.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":  time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
