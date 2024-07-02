package create

type mockRepository struct {
	insertErr error
}

func (mr *mockRepository) Insert(uid, username string) error {
	return mr.insertErr
}
