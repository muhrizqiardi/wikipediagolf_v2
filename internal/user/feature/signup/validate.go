package signup

import "github.com/go-playground/validator/v10"

var Validate = validator.New(func(v *validator.Validate) {
	v.RegisterAlias("ispassword", "min=8")
	v.RegisterAlias("isconfirm", "min=8,eqfield=Password")
	v.RegisterValidation("isusername", func(fl validator.FieldLevel) bool {
		return UsernamePattern.MatchString(fl.Field().String())
	})
})
