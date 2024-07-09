package signup

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/repository"

type mockRepository struct {
	createV   *repository.CreateUserResult
	createErr error
}

func (mr *mockRepository) Create(email, password string) (*repository.CreateUserResult, error) {
	return mr.createV, mr.createErr
}
