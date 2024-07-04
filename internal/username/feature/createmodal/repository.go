package createmodal

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByUID(uid string) (*FindByUIDResponse, error)
}

type repository struct {
	context context.Context
	db      *sql.DB
}

func newRepository(ctx context.Context, db *sql.DB) *repository {
	return &repository{
		context: ctx,
		db:      db,
	}
}

func (r *repository) FindByUID(uid string) (*FindByUIDResponse, error) {
	var (
		query  = QueryFindUsernameByUID
		args   = []any{uid}
		result FindByUIDResponse
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
