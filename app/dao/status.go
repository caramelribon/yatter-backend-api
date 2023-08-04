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
// CreateStatus : ステータスを作成
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
// FindById : IDからステータスを取得
func (r *status) FindById(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)
	query := `
		SELECT s.id, s.content, s.create_at, a.id AS "account.id", a.username AS "account.username", a.create_at AS "account.create_at"
		FROM status s
		JOIN account a ON s.account_id = a.id
		WHERE s.id = ?`
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}

// Get statuses
// GetStatuses : ステータスを取得
func (r *status) GetStatuses(ctx context.Context, query *object.QueryParams) (*[]object.Status, error) {
	var statuses *[]object.Status
	var args []interface{}

	// クエリの生成
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
		// 空のスライスを返す
		return statuses, nil
	} else {
		if query.SinceId != 0 && query.MaxId != 0 {
    q += " WHERE s.id > ? AND s.id < ?"
    args = append(args, query.MaxId, query.SinceId)
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
	}

	// クエリの実行
	err := r.db.SelectContext(ctx, &statuses, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find statuses from db: %w", err)
	}

	return statuses, nil
}
