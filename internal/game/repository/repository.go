package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/model"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/game/query"
)

var (
	ErrNoRowsAffected = errors.New("no rows are affected")
	ErrRequest        = errors.New("request returns error")
)

type Repository struct {
	context    context.Context
	httpClient *http.Client
	db         *sql.DB
}

func NewRepository(ctx context.Context, httpClient *http.Client, db *sql.DB) *Repository {
	return &Repository{
		context:    ctx,
		httpClient: httpClient,
		db:         db,
	}
}

type GetRandomTitle struct {
}

func (r *Repository) GetRandomSummary(language string) (*model.Summary, error) {
	url := fmt.Sprintf("https://%s.wikipedia.org/api/rest_v1/page/random/summary", language)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= http.StatusOK && http.StatusBadRequest > res.StatusCode {
		var summary model.Summary
		if err := json.NewDecoder(res.Body).Decode(&summary); err != nil {
			return nil, err
		}
		return &summary, nil
	} else {
		return nil, ErrRequest
	}
}

func (r *Repository) CreateGame(roomID uuid.UUID, index int, language, fromTitle, toTitle string) (*model.Game, error) {
	var (
		q      = query.QueryCreateGame
		args   = []any{roomID, index, language, fromTitle, toTitle}
		result model.Game
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, q, args...); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) GetLatestGame(roomID uuid.UUID) (*model.Game, error) {
	var (
		q      = query.QueryCreateGame
		args   = []any{roomID}
		result model.Game
	)

	dbx := sqlx.NewDb(r.db, "postgres")
	txx, err := dbx.BeginTxx(r.context, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	if err := txx.Get(&result, q, args...); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) UpdateGame(id, roomID uuid.UUID, isFinished bool) error {
	var (
		q    = query.QueryUpdateGame
		args = []any{id, roomID, isFinished}
	)

	tx, err := r.db.BeginTx(r.context, &sql.TxOptions{})
	if err != nil {
		return err
	}
	sqlResult, err := tx.Exec(q, args...)
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
	if err := txx.Get(&result, q, args...); err != nil {
		return nil, err
	}
	if err := txx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}
