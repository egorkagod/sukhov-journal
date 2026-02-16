package auth

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"journal/internal/db"
)

type UserCreateRepoDTO struct {
	Login        string
	PasswordHash string
	Role         string
}

type UserRepo interface {
	GetbyID(ctx context.Context, id uint64) (*User, error)
	GetByLogin(ctx context.Context, login string) (*User, error)
	Create(ctx context.Context, data UserCreateRepoDTO) error
}

type userRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (UserRepo, error) {
	return &userRepo{db: db}, nil
}

func (r *userRepo) GetbyID(ctx context.Context, id uint64) (*User, error) {
	var user User
	query := r.db.WithContext(ctx).First(&user, id)
	err := query.Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ErrUserNotFound
	case err != nil:
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetByLogin(ctx context.Context, login string) (*User, error) {
	var user User
	query := r.db.WithContext(ctx).Where("Login = ?", login).Find(&user)
	err := query.Error
	if err != nil {
		return nil, err
	}
	if query.RowsAffected == 0 {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

func (r *userRepo) Create(ctx context.Context, data UserCreateRepoDTO) error {
	newUser := &User{Login: data.Login, PasswordHash: data.PasswordHash, Role: data.Role}
	err := r.db.WithContext(ctx).Create(newUser).Error

	if db.IsUniqueErr(err) {
		return ErrUserAlreadyExists
	}
	return err
}
