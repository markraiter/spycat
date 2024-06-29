package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/markraiter/spycat/internal/app/storage"
	"github.com/markraiter/spycat/internal/config"
	"github.com/markraiter/spycat/internal/domain"
	"github.com/markraiter/spycat/internal/lib/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserSaver interface {
	SaveUser(ctx context.Context, user *domain.User) (int, error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (*domain.User, error)
}

type AuthService struct {
	saver    UserSaver
	provider UserProvider
}

func (s *AuthService) Register(ctx context.Context, user *domain.UserRequest) (int, error) {
	const operation = "service.RegisterUser"

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", operation, err)
	}

	userResp := domain.User{
		Username: user.Username,
		Password: string(passHash),
		Email:    user.Email,
	}

	id, err := s.saver.SaveUser(ctx, &userResp)
	if err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			return 0, fmt.Errorf("%s: %w", operation, ErrAlreadyExists)
		}

		return 0, fmt.Errorf("%s: %w", operation, err)
	}

	return id, nil
}

func (s *AuthService) Login(ctx context.Context, cfg config.Auth, email, password string) (string, error) {
	const operation = "service.Login"

	user, err := s.provider.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return "", fmt.Errorf("%s: %w", operation, ErrNotFound)
		}

		return "", fmt.Errorf("%s: %w", operation, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("%s: %w", operation, ErrInvalidCredentials)
	}

	token, err := jwt.NewToken(cfg, user, cfg.AccessTTL)
	if err != nil {
		return "", fmt.Errorf("%s: %w", operation, err)
	}

	return token, nil
}
