package services

import (
	"context"
	"time"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/adwinugroho/wedding-management-system/modules/auth/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GenerateJWTToken(ctx context.Context, user models.User) (*models.User, string, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	BlacklistToken(ctx context.Context, token string, userID string, expiresAt time.Time) error
	IsTokenBlacklisted(ctx context.Context, token string) (bool, error)
	CleanupExpiredBlacklistedTokens(ctx context.Context) error
	RegisterUser(ctx context.Context, user models.User) (*models.User, error)
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

func (s *authService) RegisterUser(ctx context.Context, user models.User) (*models.User, error) {
	// Check if user already exists
	existingUser, err := s.authRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		logger.LogError("Error while checking existing user, cause: " + err.Error())
		return nil, models.NewError("500-General-Error", "Internal Server Error.")
	}
	if existingUser != nil {
		return nil, models.NewError("400-User-Exists", "User with this email already exists.")
	}

	// Hash the password if provided
	if user.Password != nil {
		hashedPassword, err := s.hashPassword(*user.Password)
		if err != nil {
			logger.LogError("Error while hashing password, cause: " + err.Error())
			return nil, models.NewError("500-General-Error", "Internal Server Error.")
		}
		user.Password = &hashedPassword
	}

	// Set default values
	if user.Provider == "" {
		user.Provider = "local"
	}
	if user.Role == "" {
		user.Role = "USER"
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Create the user
	err = s.authRepository.CreateUser(ctx, &user)
	if err != nil {
		logger.LogError("Error while creating user, cause: " + err.Error())
		return nil, models.NewError("500-General-Error", "Internal Server Error.")
	}

	return &user, nil
}

func (s *authService) hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.LogError("Error hashing password, cause: " + err.Error())
		return "", models.NewError("500-General-Error", "Internal Server Error.")
	}

	return string(hashPassword), nil
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
