package signup

type mockService struct {
	v   *CreateUserResponse
	err error
}

func (ms *mockService) SignUp(payload CreateUserRequest) (*CreateUserResponse, error) {
	return ms.v, ms.err
}
