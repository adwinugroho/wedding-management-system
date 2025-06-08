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
	Login(ctx context.Context, email string, password string) (*models.User, string, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{authRepository: authRepository}
}

func (s *authService) Login(ctx context.Context, email string, password string) (*models.User, string, error) {
	user, err := s.authRepository.GetUserByEmail(ctx, email)
	if err != nil {
		logger.LogError("Error while get user by email, cause: " + err.Error())
		return nil, "", models.NewError("500-General-Error", "Internal server error")
	}

	if user == nil {
		return nil, "", models.NewError("404-User-not-found", "User not found")
	}

	if user.Password == nil {
		return nil, "", models.NewError("401-Invalid-Password", "Invalid password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		logger.LogError("Error while compare hash and password, cause: " + err.Error())
		return nil, "", models.NewError("401-Invalid-Password", "Invalid password")
	}

	token, err := s.generateJWTToken(*user)
	if err != nil {
		logger.LogError("Error while generate JWT token, cause: " + err.Error())
		return nil, "", models.NewError("500-General-Error", "Internal server error")
	}

	return user, token, nil
}

func (s *authService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.authRepository.GetUserByEmail(ctx, email)
}

func (s *authService) generateJWTToken(payload models.User) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":    payload.ID,
		"email": payload.Email,
		"name":  payload.Name,
		"role":  payload.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":   time.Now().Unix(),
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
