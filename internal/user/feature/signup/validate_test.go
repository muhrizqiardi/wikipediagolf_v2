package signup

import (
	"context"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestValidate(t *testing.T) {
	t.Run("should register validation", func(t *testing.T) {
		{
			type testType struct {
				Password        string `validate:"ispassword"`
				ConfirmPassword string `validate:"isconfirm"`
				ExpErr          bool
			}
			tests := []testType{
				{Password: "invld", ConfirmPassword: "", ExpErr: true},
				{Password: "validPassword-123", ConfirmPassword: "invld", ExpErr: true},
				{Password: "validPassword-123", ConfirmPassword: "validInequal-123", ExpErr: true},
				{Password: "validPassword-123", ConfirmPassword: "validPassword-123", ExpErr: false},
			}
			for _, test := range tests {
				t.Run("Password:"+test.Password+"; ConfirmPassword:"+test.ConfirmPassword, func(t *testing.T) {
					if test.ExpErr {
						testutil.AssertError(t, Validate.StructCtx(context.TODO(), &test))
					} else {
						testutil.AssertNoError(t, Validate.StructCtx(context.TODO(), &test))
					}
				})
			}
		}
		{
			type testType struct {
				Username string `validate:"isusername"`
				ExpErr   bool
			}
			tests := []testType{
				{Username: "invalid username", ExpErr: true},
				{Username: "invalid username/?", ExpErr: true},

				{Username: "validUsername123", ExpErr: false},
				{Username: "validusername", ExpErr: false},
				{Username: "valid_username", ExpErr: false},
				{Username: "valid-username", ExpErr: false},
			}
			for _, test := range tests {
				t.Run("Username:"+test.Username, func(t *testing.T) {
					if test.ExpErr {
						testutil.AssertError(t, Validate.StructCtx(context.TODO(), &test))
					} else {
						testutil.AssertNoError(t, Validate.StructCtx(context.TODO(), &test))
					}
				})
			}
		}
	})
}
