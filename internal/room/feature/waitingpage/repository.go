package waitingpage

import (
	"github.com/google/uuid"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/model"
	"github.com/muhrizqiardi/wikipediagolf_v2/internal/room/repository"
)

type Repository interface {
	GetRoomMemberDisplayName(userUID string) (string, error)
	GetRoomMembers(roomID uuid.UUID) ([]repository.GetRoomMembersRow, error)
	GetRoomBelongToMember(userUID string) (*model.Room, error)
}
