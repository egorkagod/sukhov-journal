package auth

import "context"

type UserCreateRepoDTO struct {
	Login        string
	PasswordHash string
	Role         string
}

type UserRepo interface {
	GetbyID(ctx context.Context, id int64) (*User, error)
	GetByLogin(ctx context.Context, login string) (*User, error)
	Create(ctx context.Context, data UserCreateRepoDTO) error
}

type userInMemoryRepo struct {
	data []User
}

func NewRepo() UserRepo {
	return &userInMemoryRepo{data: make([]User, 0)}
}

func (r *userInMemoryRepo) GetbyID(ctx context.Context, id int64) (*User, error) {
	if id >= 0 && id < int64(len(r.data)) {
		return &r.data[id], nil
	}
	return nil, UserNotFoundErr
}

func (r *userInMemoryRepo) GetByLogin(ctx context.Context, login string) (*User, error) {
	for _, user := range r.data {
		if user.Login == login {
			return &user, nil
		}
	}
	return nil, UserNotFoundErr
}

func (r *userInMemoryRepo) Create(ctx context.Context, data UserCreateRepoDTO) error {
	user, err := r.GetByLogin(ctx, data.Login) // делаем проверку из-за InMemory репозитория
	switch {
	case err == UserNotFoundErr:
	case user != nil:
		return UserAlreadyExistsErr
	default:
		return err
	}

	newUser := User{ID: int64(len(r.data)), Login: data.Login, PasswordHash: data.PasswordHash, Role: data.Role}
	r.data = append(r.data, newUser)
	return nil
}
