package services

import (
	"context"

	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/adwinugroho/wedding-management-system/modules/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, email string, password string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{authRepository: authRepository}
}

func (s *authService) Login(ctx context.Context, email string, password string) (*models.User, error) {
	user, err := s.authRepository.GetUserByEmail(ctx, email)
	if err != nil {
		logger.LogError("Error while get user by email, cause: " + err.Error())
		return nil, models.NewError("500-General-Error", "Internal server error")
	}

	if user == nil {
		return nil, models.NewError("404-User-not-found", "User not found")
	}

	if user.Password == nil {
		return nil, models.NewError("401-Invalid-Password", "Invalid password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		logger.LogError("Error while compare hash and password, cause: " + err.Error())
		return nil, models.NewError("401-Invalid-Password", "Invalid password")
	}

	// TODO: Implement generate token JWT

	return user, nil
}

func (s *authService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.authRepository.GetUserByEmail(ctx, email)
}
