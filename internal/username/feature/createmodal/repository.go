package createmodal

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/username/repository"

type Repository interface {
	FindByUID(uid string) (*repository.FindByUIDResult, error)
}
