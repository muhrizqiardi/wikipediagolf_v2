package signup

type mockRepository struct {
	createV   *CreateUserResponse
	createErr error
}

func (mr *mockRepository) Create(email, password string) (*CreateUserResponse, error) {
	return mr.createV, mr.createErr
}
