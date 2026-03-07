package service

import (
	"context"
	"time"

	"example.com/go-bank/internal/crypto"
	"example.com/go-bank/internal/domain"
)

// UserRepo defines the interface for user data persistence.
type UserRepo interface {
	Create(ctx context.Context, u *domain.User) error
}

// UserService provides user-related operations.
type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, name *string, email, password string) error {

	// Hash PW
	hashedPassword, err := crypto.HashPassword(password) 
	if err != nil {
		return err
	}
	// Validation
	user := &domain.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}
	

	return s.repo.Create(ctx, user)
}