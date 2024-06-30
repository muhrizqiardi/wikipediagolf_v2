package signup

import "github.com/go-playground/validator/v10"

type ValidationErrors struct {
	validator.ValidationErrors
}

func (errs *ValidationErrors) Fields() map[string]string {
	result := make(map[string]string)
	for _, err := range errs.ValidationErrors {
		result[err.StructField()] = err.Error()
	}

	return result
}

var Validate = validator.New(func(v *validator.Validate) {
	v.RegisterAlias("ispassword", "min=8")
	v.RegisterAlias("isconfirm", "min=8,eqfield=Password")
	v.RegisterValidation("isusername", func(fl validator.FieldLevel) bool {
		return UsernamePattern.MatchString(fl.Field().String())
	})
})
