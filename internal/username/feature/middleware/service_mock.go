package middleware

type mockService struct {
	findByUIDV   *FindByUIDResponse
	findByUIDErr error
}

func (mr *mockService) FindByUID(uid string) (*FindByUIDResponse, error) {
	return mr.findByUIDV, mr.findByUIDErr
}
