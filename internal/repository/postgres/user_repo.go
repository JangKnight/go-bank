package postgres

import (
	"context"

	"example.com/go-bank/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
    return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.pool.Exec(
		ctx,
		`INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)`,
		user.Name, user.Email, user.PasswordHash,
	)
	return err
}
	