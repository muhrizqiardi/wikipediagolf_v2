package repository

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
)

type Repository interface {
	InsertRoom(roomCode, state string) (*model.Room, error)
	InsertRoomMember(roomID uuid.UUID, userUID string, isOwner bool) (*model.RoomMember, error)
	DeleteRoomMember(roomID uuid.UUID, userUID string) error
	GetRoomByCode(roomCode string) (*model.Room, error)
	GetRoomByID(roomID uuid.UUID) (*model.Room, error)
	GetRoomMembers(roomID uuid.UUID) ([]model.RoomMember, error)
	GetRoomBelongToMember(userUID uuid.UUID) (*model.Room, error)
	UpdateRoomState(roomID uuid.UUID, newState string) (*model.Room, error)
	Delete(roomID uuid.UUID) (*model.Room, error)
}
