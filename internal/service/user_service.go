package service

import (
	"context"
	"time"

	"example.com/go-bank/internal/crypto"
	"example.com/go-bank/internal/domain"
)

type UserRepo interface {
	Save(ctx context.Context, u *domain.User) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, name, email, password string) error {

	hashedPassword, err := crypto.HashPassword(password) 
	if err != nil {
		return err
	}

	user := &domain.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}

	return s.repo.Save(ctx, user)
}