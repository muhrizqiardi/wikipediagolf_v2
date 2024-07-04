package waitingpage

type Service interface {
	GetRoom(userUID string) (*GetRoomResponse, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) GetRoom(userUID string) (*GetRoomResponse, error) {
	room, err := s.repository.GetRoomBelongToMember(userUID)
	if err != nil {
		return nil, err
	}

	members, err := s.repository.GetRoomMembers(room.ID)
	if err != nil {
		return nil, err
	}

	resultMembers := make([]GetRoomResponseMember, 0, len(members))
	for _, member := range members {
		displayName, err := s.repository.GetRoomMemberDisplayName(member.UserUID)
		if err != nil {
			continue
		}
		resultMembers = append(resultMembers, GetRoomResponseMember{
			ID:          member.ID,
			IsOwner:     member.IsOwner,
			RoomID:      member.RoomID,
			UserUID:     member.UserUID,
			Username:    member.Username,
			DisplayName: displayName,
			CreatedAt:   member.CreatedAt,
			UpdatedAt:   member.UpdatedAt,
		})
	}
	result := GetRoomResponse{
		Room:    *room,
		Members: resultMembers,
	}

	return &result, nil
}
