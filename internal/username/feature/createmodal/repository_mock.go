package createmodal

type mockRepository struct {
	findByUIDV   *FindByUIDResponse
	findByUIDErr error
}

func (mr *mockRepository) FindByUID(uid string) (*FindByUIDResponse, error) {
	return mr.findByUIDV, mr.findByUIDErr
}
