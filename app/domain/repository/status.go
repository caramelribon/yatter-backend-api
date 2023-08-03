package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Create a new status
	CreateStatus(ctx context.Context, status *object.Status) error
}
