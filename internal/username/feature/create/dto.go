package create

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateUsernameRequest struct {
	UID      string `schema:"uid"`
	Username string `schema:"username" validate:"isusername"`
}

func (c *CreateUsernameRequest) Valid(ctx context.Context) error {
	if err := Validate.StructCtx(ctx, c); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			return &ValidationErrors{err}
		}
	}

	return nil
}

type CreateUsernameResponse struct {
	UID       string    `db:"uid"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
