package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Create a new status
	CreateStatus(ctx context.Context, status *object.Status, accountId int64) error
	// Find a status by ID
	FindById(ctx context.Context, id int64) (*object.Status, error)

	// Get statuses
	GetPublicStatuses(ctx context.Context, query *object.QueryParams) ([]*object.Status, error)
	GetHomeStatuses(ctx context.Context, query *object.QueryParams, accountId int64) ([]*object.Status, error)

	// Delete a status
	DeleteStatus(ctx context.Context, statusId int64) error
}
