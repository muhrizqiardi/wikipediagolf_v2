package create

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Insert(uid, username string) error
	Find(username string) (*FindUsernameResponse, error)
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

func (r *repository) Find(username string) (*FindUsernameResponse, error) {
	var (
		query = QueryFindUsername
		args  = []any{username}
		dest  FindUsernameResponse
	)

	dbx := sqlx.NewDb(r.db, "postgresql")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer txx.Commit()

	if err := txx.Get(&dest, query, args...); err != nil {
		return nil, err
	}

	return &dest, nil
}
