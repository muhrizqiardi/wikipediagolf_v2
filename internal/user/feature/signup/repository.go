package signup

import (
	"context"
	"database/sql"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/errorutils"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(email, password string) (*CreateUserResponse, error)
}

type userRepository struct {
	context     context.Context
	firebaseApp *firebase.App
}

func NewUserRepository(ctx context.Context, firebaseApp *firebase.App) *userRepository {
	return &userRepository{
		context:     ctx,
		firebaseApp: firebaseApp,
	}
}

func (r *userRepository) Create(email, password string) (*CreateUserResponse, error) {
	client, err := r.firebaseApp.Auth(r.context)
	if err != nil {
		return nil, err
	}

	userToCreate := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		EmailVerified(false).
		Disabled(false)

	newUser, err := client.CreateUser(r.context, userToCreate)
	if err != nil {
		switch {
		case errorutils.IsAlreadyExists(err):
			return nil, ErrDuplicateEmail
		}
		return nil, err
	}

	return &CreateUserResponse{
		UID:           newUser.UID,
		Email:         newUser.Email,
		EmailVerified: newUser.EmailVerified,
		PhoneNumber:   newUser.PhoneNumber,
		DisplayName:   newUser.DisplayName,
		PhotoURL:      newUser.PhotoURL,
		Disabled:      newUser.Disabled,
	}, nil
}

type UsernameRepository interface {
	Insert(uid, username string) error
	Find(username string) (*FindUsernameResponse, error)
}

type usernameRepository struct {
	context context.Context
	db      *sql.DB
}

func NewUsernameRepository(ctx context.Context, db *sql.DB) *usernameRepository {
	return &usernameRepository{
		context: ctx,
		db:      db,
	}
}

func (r *usernameRepository) Insert(uid, username string) error {
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

func (r *usernameRepository) Find(username string) (*FindUsernameResponse, error) {
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
