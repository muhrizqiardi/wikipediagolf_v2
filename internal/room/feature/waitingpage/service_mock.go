package waitingpage

type mockService struct {
	getRoomV   *GetRoomResponse
	getRoomErr error
}

func (ms *mockService) GetRoom(userUID string) (*GetRoomResponse, error) {
	return ms.getRoomV, ms.getRoomErr
}
