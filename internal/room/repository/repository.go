package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/query"
)

var (
	ErrNoRowsAffected = errors.New("no rows are affected")
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

func (r *Repository) InsertRoom(roomCode, status string) (*model.Room, error) {
	var (
		q      = query.QueryInsertRoom
		args   = []any{roomCode, status}
		result model.Room
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, q, args); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) InsertRoomMember(roomID uuid.UUID, userUID string, isOwner bool) (*model.RoomMember, error) {
	var (
		q      = query.QueryInsertRoomMember
		args   = []any{roomID, userUID, isOwner}
		result model.RoomMember
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, q, args); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) DeleteRoomMember(roomID uuid.UUID, userUID string) error {
	var (
		q    = query.QueryGetRoomBelongToMember
		args = []any{roomID}
	)

	tx, err := r.db.BeginTx(r.context, &sql.TxOptions{})
	if err != nil {
		return err
	}
	sqlResult, err := tx.Exec(q, args)
	if err != nil {
		return err
	}
	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected < 1 {
		return ErrNoRowsAffected
	}
	if err := tx.Commit(); err != nil {
		return nil
	}

	return nil
}

func (r *Repository) GetRoomByCode(roomCode string) (*model.Room, error) {
	var (
		q      = query.QueryGetRoomByCode
		args   = []any{roomCode}
		result model.Room
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, q, args); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) GetRoomByID(roomID uuid.UUID) (*model.Room, error) {
	var (
		q      = query.QueryGetRoomByID
		args   = []any{roomID}
		result model.Room
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, q, args); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) GetRoomMembers(roomID uuid.UUID) ([]model.RoomMember, error) {
	var (
		q      = query.QueryGetRoomMembers
		args   = []any{roomID}
		result []model.RoomMember
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Select(&result, q, args); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) GetRoomBelongToMember(userUID string) (*model.Room, error) {
	var (
		q      = query.QueryGetRoomBelongToMember
		args   = []any{userUID}
		result model.Room
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, q, args); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) UpdateRoomState(roomID uuid.UUID, newStatus string) (*model.Room, error) {
	var (
		q      = query.QueryGetRoomBelongToMember
		args   = []any{roomID, newStatus}
		result model.Room
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, q, args); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) Delete(roomID uuid.UUID) error {
	var (
		q    = query.QueryGetRoomBelongToMember
		args = []any{roomID}
	)

	tx, err := r.db.BeginTx(r.context, &sql.TxOptions{})
	if err != nil {
		return err
	}
	sqlResult, err := tx.Exec(q, args)
	if err != nil {
		return err
	}
	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected < 1 {
		return ErrNoRowsAffected
	}
	if err := tx.Commit(); err != nil {
		return nil
	}

	return nil
}
