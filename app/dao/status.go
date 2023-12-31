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
	// Implementation for repository.Satus
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// Create a new status
func (r *status) CreateStatus(ctx context.Context, status *object.Status, accountId int64) error {
	result, err := r.db.ExecContext(ctx, "INSERT INTO status (account_id, content) VALUES (?, ?)", accountId, status.Content)
	if err != nil {
		return fmt.Errorf("failed to insert status into db: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get status ID from db: %w", err)
	}
	status.ID = id

	return nil
}

// Find a status by ID
func (r *status) FindById(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)

	// create query
	query := `
		SELECT
		  s.id,
			s.content,
			s.create_at,
			a.id AS "account.id",
			a.username AS "account.username",
			a.create_at AS "account.create_at"
		FROM status s
		JOIN account a ON s.account_id = a.id
		WHERE s.id = ?`

	// execute query
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}

// Get statuses (public)
func (r *status) GetPublicStatuses(ctx context.Context, query *object.QueryParams) ([]*object.Status, error) {
	var statuses []*object.Status
	var args []interface{}

	// create query
	q := `
		SELECT
		  s.id,
			s.content,
			s.create_at,
			a.id AS "account.id",
			a.username AS "account.username",
			a.create_at AS "account.create_at"
		FROM status s
		LEFT JOIN account a ON s.account_id = a.id
		`
	if query.OnlyMedia {
		// return empty slice
		return statuses, nil
	} else {
		if query.SinceId != 0 && query.MaxId != 0 {
			q += " WHERE s.id > ? AND s.id < ?"
			args = append(args, query.SinceId, query.MaxId)
		} else {
			if query.SinceId != 0 {
				q += " WHERE s.id > ?"
				args = append(args, query.SinceId)
			}
			if query.MaxId != 0 {
				q += " WHERE s.id < ?"
				args = append(args, query.MaxId)
			}
		}
		if(query.Limit != 0) {
			q += " LIMIT ?"
			args = append(args, query.Limit)
		}
	}

	// execute query
	err := r.db.SelectContext(ctx, &statuses, q, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find statuses from db: %w", err)
	}

	return statuses, nil
}

// Get statuses (home)
func (r *status) GetHomeStatuses(ctx context.Context, query *object.QueryParams, accountId int64) ([]*object.Status, error) {
	var statuses []*object.Status
	var args []interface{}

	// create query
	q := `
		SELECT
		  s.id,
			s.content,
			s.create_at,
			a.id AS "account.id",
			a.username AS "account.username",
			a.create_at AS "account.create_at"
		FROM status s
		LEFT JOIN account a ON s.account_id = a.id
		`

	q += " WHERE s.account_id = ?"
	args = append(args, accountId)

	if query.OnlyMedia {
		// return empty slice
		return statuses, nil
	} else {
		if query.SinceId != 0 && query.MaxId != 0 {
			q += " AND s.id > ? AND s.id < ?"
			args = append(args, query.SinceId, query.MaxId)
		} else {
			if query.SinceId != 0 {
				q += " AND s.id > ?"
				args = append(args, query.SinceId)
			}
			if query.MaxId != 0 {
				q += " AND s.id < ?"
				args = append(args, query.MaxId)
			}
		}
		if(query.Limit != 0) {
			q += " LIMIT ?"
			args = append(args, query.Limit)
		}
	}

	// execute query
	err := r.db.SelectContext(ctx, &statuses, q, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find statuses from db: %w", err)
	}

	return statuses, nil
}

// Delete a status
func (r *status) DeleteStatus(ctx context.Context, statusId int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM status WHERE id = ?", statusId)
	if err != nil {
		return fmt.Errorf("failed to delete status from db: %w", err)
	}

	return nil
}

