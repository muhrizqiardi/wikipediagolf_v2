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

	getRoomBelongToMemberV   *model.Room
	getRoomBelongToMemberErr error

	deleteErr error
}

func (mr *mockRepository) InsertRoom(roomCode, status string) (*model.Room, error) {
	return mr.insertRoomV, mr.insertRoomErr
}

func (mr *mockRepository) InsertRoomMember(roomID uuid.UUID, userUID string, isOwner bool) (*model.RoomMember, error) {
	return mr.insertRoomMemberV, mr.insertRoomErr
}

func (mr *mockRepository) GetRoomBelongToMember(userUID string) (*model.Room, error) {
	return mr.getRoomBelongToMemberV, mr.getRoomBelongToMemberErr
}

func (mr *mockRepository) Delete(roomID uuid.UUID) error {
	return mr.deleteErr
}
