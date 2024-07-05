package create

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
)

type mockRepository struct {
	insertRoomV   *model.Room
	insertRoomErr error

	insertRoomMemberV   *model.RoomMember
	insertRoomMemberErr error
}

func (mr *mockRepository) InsertRoom(roomCode, status string) (*model.Room, error) {
	return mr.insertRoomV, mr.insertRoomErr
}

func (mr *mockRepository) InsertRoomMember(roomID uuid.UUID, userUID string, isOwner bool) (*model.RoomMember, error) {
	return mr.insertRoomMemberV, mr.insertRoomErr
}
