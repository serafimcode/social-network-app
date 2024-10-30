package repositories

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"social-network-app/internal/domain"
)

type UserRepository struct {
	db *sqlx.DB
}

func BuildUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	builder := sq.Select("*").From("users")

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	err = ur.db.SelectContext(ctx, &users, sql, args...)

	if err != nil {
		return nil, err
	}

	return users, nil
}
