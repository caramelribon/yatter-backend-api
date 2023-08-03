package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Create a new status
	CreateStatus(ctx context.Context, status *object.Status) error
	// Find a status by ID
	FindById(ctx context.Context, id int64) (*object.Status, error)
}
