package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserCredentialsServiceDTO struct {
	Login    string
	Password string
}

type UserService interface {
	Register(ctx context.Context, data UserCredentialsServiceDTO) error
	Login(ctx context.Context, data UserCredentialsServiceDTO) (*User, error)
}

type userService struct {
	repo UserRepo
}

func NewService(repo UserRepo) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, data UserCredentialsServiceDTO) error {
	hash, err := hashPassword(data.Password)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, UserCreateRepoDTO{Login: data.Login, PasswordHash: hash, Role: "USER"})
}

func (s *userService) Login(ctx context.Context, data UserCredentialsServiceDTO) (*User, error) {
	user, err := s.repo.GetByLogin(ctx, data.Login)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.Password)); err != nil {
		return nil, CredentialsIsIncorrect
	}
	return user, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
