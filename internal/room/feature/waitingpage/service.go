package waitingpage

type Service interface {
	GetRoom(userUID string) (*GetRoomResponse, error)
}
