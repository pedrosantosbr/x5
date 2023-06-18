package services

import (
	"context"

	"github.com/pedrosantosbr/x5/internal"
)

type UserRepository interface {
	Create(ctx context.Context, user *UserCreateParams) (internal.User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Check(password, hash string) error
}

type User struct {
	repo   UserRepository
	hasher PasswordHasher
}

// NewUser Service...
func NewUser(repo UserRepository, hasher PasswordHasher) *User {
	return &User{
		repo:   repo,
		hasher: hasher,
	}
}

func Create(u *User, ctx context.Context, params *UserCreateParams) (internal.User, error) {
	if err := params.Validate(); err != nil {
		return internal.User{}, err
	}

	hashedPassword, err := u.hasher.Hash(params.Password)
	if err != nil {
		return internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "hasher.Hash")
	}

	params.Password = hashedPassword

	user, err := u.repo.Create(ctx, params)
	if err != nil {
		return internal.User{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "repo.Create")
	}

	return user, nil
}