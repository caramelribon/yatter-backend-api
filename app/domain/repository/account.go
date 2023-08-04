package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)

	// Create a new account
	CreateUser(ctx context.Context, account *object.Account) error
	
	// Fetch account which has specified ID
	FindById(ctx context.Context, id int64) (*object.Account, error)
}
