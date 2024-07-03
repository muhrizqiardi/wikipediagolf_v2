package create

import (
	"context"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestValidate(t *testing.T) {
	t.Run("should register validation", func(t *testing.T) {
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
	})
}
