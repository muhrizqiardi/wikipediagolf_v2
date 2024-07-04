package create

type mockService struct {
	v   *CreateRoomResponse
	err error
}

func (ms *mockService) Create(ownerUID string) (*CreateRoomResponse, error) {
	return ms.v, ms.err
}
