package config

import (
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/testutil"
)

func TestGetConfig(t *testing.T) {
	type testCase struct {
		args   []string
		getenv func(string) string
		exp    Config
	}

	tests := []testCase{
		{
			args: []string{"-host=localhost", "-port=3001", "-migrate", "-seed"},
			getenv: func(s string) string {
				switch s {
				case "DATABASE_URL":
					return "postgres://user:mockpw@localhost:5432/dbname"
				}
				return ""
			},
			exp: Config{
				Host:        "localhost",
				Port:        3001,
				IsMigrate:   true,
				IsSeed:      true,
				DatabaseURL: "postgres://user:mockpw@localhost:5432/dbname",
			},
		},
		{
			args: []string{},
			getenv: func(s string) string {
				switch s {
				case "DATABASE_URL":
					return "postgres://user:mockpw@localhost:5432/dbname"
				}
				return ""
			},
			exp: Config{
				Host:        "0.0.0.0",
				Port:        3000,
				IsMigrate:   false,
				IsSeed:      false,
				DatabaseURL: "postgres://user:mockpw@localhost:5432/dbname",
			},
		},
	}

	for _, test := range tests {
		var (
			exp = test.exp
			got = GetConfig(test.args, test.getenv)
		)

		testutil.AssertEqualCMP(t, exp, got)
	}
}
