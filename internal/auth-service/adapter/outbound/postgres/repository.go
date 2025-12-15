package postgres

import (
	"context"
	/* 	"errors"
	 */
	"github.com/anfastk/MERGESPACE/internal/auth-service/application/port/outbound"
	"github.com/anfastk/MERGESPACE/internal/auth-service/domain/entity"
	/* 	"github.com/jackc/pgx/v5"
	 */"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

var _ outbound.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *pgxpool.Pool) outbound.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	/* m := fromDomain(user)

	query := `
		INSERT INTO users (
			id, email, username, password_hash, status, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		m.ID,
		m.Email,
		m.Username,
		m.PasswordHash,
		m.Status,
		m.CreatedAt,
		m.UpdatedAt,
	)
	*/
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	/* query := `
		SELECT id, email, username, password_hash, status, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var m UserModel
	err := r.db.QueryRow(ctx, query, email).Scan(
		&m.UserID,
		&m.Email,
		&m.Username,
		&m.PasswordHash,
		&m.Status,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	*/
	return nil, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	/* query := `
		SELECT id, email, username, password_hash, status, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var m UserModel
	err := r.db.QueryRow(ctx, query, id).Scan(
		&m.UserID,
		&m.Email,
		&m.Username,
		&m.PasswordHash,
		&m.Status,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	*/
	return nil, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {

	/* query := `
		SELECT id, email, username, password_hash, status, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	var m userModel
	err := r.db.QueryRow(ctx, query, username).Scan(
		&m.ID,
		&m.Email,
		&m.Username,
		&m.PasswordHash,
		&m.Status,
		&m.CreatedAt,
		&m.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	} */

	return nil, nil
}
