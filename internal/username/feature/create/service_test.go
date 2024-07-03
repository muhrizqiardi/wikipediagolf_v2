package create

import (
	"context"
	"errors"
	"testing"

	"github.com/lib/pq"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestCreate_InvalidUsername(t *testing.T) {
	var (
		mr      = &mockRepository{insertErr: nil}
		payload = CreateUsernameRequest{
			UID:      "mockID",
			Username: "invalid username",
		}
		s   = NewService(context.TODO(), mr)
		err = s.Create(payload)
	)

	testutil.AssertError(t, err)
	testutil.CompareError(t, ErrInvalidUsername, err)
}

func TestCreate_RepositoryErrorCode23505(t *testing.T) {
	var (
		mr      = &mockRepository{insertErr: &pq.Error{Code: "23505"}}
		payload = CreateUsernameRequest{
			UID:      "mockID",
			Username: "validUsername",
		}
		s   = NewService(context.TODO(), mr)
		err = s.Create(payload)
	)

	testutil.AssertError(t, err)
	testutil.CompareError(t, ErrDuplicateUsername, err)
}
func TestCreate_RepositoryError(t *testing.T) {
	var (
		mr      = &mockRepository{insertErr: errors.New("")}
		payload = CreateUsernameRequest{
			UID:      "mockID",
			Username: "validUsername",
		}
		s   = NewService(context.TODO(), mr)
		err = s.Create(payload)
	)

	testutil.AssertError(t, err)
	testutil.CompareError(t, ErrCreateUsername, err)
}

func TestCreate_NoError(t *testing.T) {
	var (
		mr      = &mockRepository{insertErr: nil}
		payload = CreateUsernameRequest{
			UID:      "mockID",
			Username: "validUsername",
		}
		s = NewService(context.TODO(), mr)
	)

	testutil.AssertNoError(t, s.Create(payload))
}
