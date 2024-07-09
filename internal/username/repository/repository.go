package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/username/query"
)

type Repository struct {
	context context.Context
	db      *sql.DB
}

func NewRepository(ctx context.Context, db *sql.DB) *Repository {
	return &Repository{
		context: ctx,
		db:      db,
	}
}

type FindByUIDResult struct {
	UID       string    `db:"uid"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (r *Repository) FindByUID(uid string) (*FindByUIDResult, error) {
	var (
		query  = query.QueryFindUsernameByUID
		args   = []any{uid}
		result FindByUIDResult
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, query, args...); err != nil {
		return nil, err
	}

	return &result, nil
}
