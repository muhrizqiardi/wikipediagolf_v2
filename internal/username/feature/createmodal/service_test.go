package createmodal

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/internal/username/repository"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestService_UsernameNotFound(t *testing.T) {
	var (
		mr      = &mockRepository{err: sql.ErrNoRows}
		s       = newService(mr)
		mockUID = "mockUID"
		_, err  = s.FindByUID(mockUID)
	)

	testutil.AssertError(t, err)
	testutil.CompareError(t, ErrUsernameNotFound, err)
}

func TestService_RepositoryError(t *testing.T) {
	var (
		mr      = &mockRepository{err: errors.New("")}
		s       = newService(mr)
		mockUID = "mockUID"
		_, err  = s.FindByUID(mockUID)
	)

	testutil.AssertError(t, err)
}

func TestService_NoError(t *testing.T) {
	var (
		mr = &mockRepository{
			v: &repository.FindByUIDResult{
				UID:      "mockUID",
				Username: "mockUsername",
			},
			err: nil,
		}
		s       = newService(mr)
		mockUID = "mockUID"
		_, err  = s.FindByUID(mockUID)
	)

	testutil.AssertNoError(t, err)
}
