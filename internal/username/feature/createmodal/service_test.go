package createmodal

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestService_UsernameNotFound(t *testing.T) {
	var (
		mr      = &mockRepository{findByUIDErr: sql.ErrNoRows}
		s       = NewService(mr)
		mockUID = "mockUID"
		_, err  = s.FindByUID(mockUID)
	)

	testutil.AssertError(t, err)
	testutil.CompareError(t, ErrUsernameNotFound, err)
}

func TestService_RepositoryError(t *testing.T) {
	var (
		mr      = &mockRepository{findByUIDErr: errors.New("")}
		s       = NewService(mr)
		mockUID = "mockUID"
		_, err  = s.FindByUID(mockUID)
	)

	testutil.AssertError(t, err)
}

func TestService_NoError(t *testing.T) {
	var (
		mr = &mockRepository{
			findByUIDV: &FindByUIDResponse{
				UID:      "mockUID",
				Username: "mockUsername",
			},
			findByUIDErr: nil,
		}
		s       = NewService(mr)
		mockUID = "mockUID"
		_, err  = s.FindByUID(mockUID)
	)

	testutil.AssertNoError(t, err)
}
