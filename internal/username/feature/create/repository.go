package create

import (
	"context"
	"database/sql"
)

type Repository interface {
	Insert(uid, username string) error
}

type repository struct {
	context context.Context
	db      *sql.DB
}

func NewRepository(ctx context.Context, db *sql.DB) *repository {
	return &repository{
		context: ctx,
		db:      db,
	}
}

func (r *repository) Insert(uid, username string) error {
	var (
		query = QueryInsertUsername
		args  = []any{uid, username}
	)

	tx, err := r.db.BeginTx(r.context, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Commit()

	if _, err := tx.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
