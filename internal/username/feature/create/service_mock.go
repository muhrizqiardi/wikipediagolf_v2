package create

type mockService struct {
	createV   *CreateUsernameResponse
	createErr error
}

func (ms *mockService) Create(payload CreateUsernameRequest) (*CreateUsernameResponse, error) {
	return ms.createV, ms.createErr
}
