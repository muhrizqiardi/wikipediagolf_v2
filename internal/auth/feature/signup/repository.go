package signup

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/auth/repository"

type Repository interface {
	Create(email, password string) (*repository.CreateUserResult, error)
}
