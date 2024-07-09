package createmodal

import "github.com/muhrizqiardi/wikipediagolf_v2/internal/username/repository"

type mockRepository struct {
	v   *repository.FindByUIDResult
	err error
}

func (mr *mockRepository) FindByUID(uid string) (*repository.FindByUIDResult, error) {
	return mr.v, mr.err
}
