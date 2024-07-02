package create

type mockService struct {
	createErr error
}

func (ms *mockService) Create(payload CreateUsernameRequest) error {
	return ms.createErr
}
