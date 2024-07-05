package create

import (
	"testing"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
	"github.com/muhrizqiardi/wikipediagolf_v2/test/testutil"
)

func TestService_RepositoryConflictError(t *testing.T) {
	var (
		mr = &mockRepository{
			insertRoomV:         nil,
			insertRoomErr:       &pq.Error{Code: "23505"},
			insertRoomMemberV:   nil,
			insertRoomMemberErr: nil,
		}
		c            = NewCodeGenerator()
		s            = NewService(c, mr)
		mockOwnerUID = "mockOwnerUID"
	)
	_, err := s.Create(mockOwnerUID)
	testutil.AssertError(t, err)
}

func TestService_NoError(t *testing.T) {
	var (
		mockOwnerUID = "mockOwnerUID"
		mr           = &mockRepository{
			insertRoomV: &model.Room{
				ID:     uuid.New(),
				Code:   "123456",
				Status: "open",
			},
			insertRoomErr: nil,
			insertRoomMemberV: &model.RoomMember{
				ID:      uuid.New(),
				IsOwner: true,
				RoomID:  uuid.New(),
				UserUID: mockOwnerUID,
			},
			insertRoomMemberErr: nil,
		}
		c = NewCodeGenerator()
		s = NewService(c, mr)
	)
	_, err := s.Create(mockOwnerUID)
	testutil.AssertNoError(t, err)
}
