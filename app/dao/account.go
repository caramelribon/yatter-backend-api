package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// Create a new account
func (r *account) CreateUser(ctx context.Context, account *object.Account) error {
    _, err := r.db.ExecContext(ctx, "INSERT INTO account (username, password_hash) VALUES (?, ?)", account.Username, account.PasswordHash)
    if err != nil {
        return fmt.Errorf("failed to insert account into db: %w", err)
    }
    return nil
}

// Find account by username
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

// Find account by ID
func (r *account) FindById(ctx context.Context, id int64) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}
