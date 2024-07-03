package signup

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Email           string `schema:"email" validate:"email"`
	Password        string `schema:"password" validate:"ispassword"`
	ConfirmPassword string `schema:"confirmPassword" validate:"isconfirm"`
}

func (c *CreateUserRequest) Valid(ctx context.Context) error {
	if err := Validate.StructCtx(ctx, c); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			return &ValidationErrors{err}
		}
	}

	return nil
}

type CreateUserResponse struct {
	UID           string
	Email         string
	EmailVerified bool
	PhoneNumber   string
	DisplayName   string
	PhotoURL      string
	Disabled      bool
	Token         string
}
