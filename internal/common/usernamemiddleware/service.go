package usernamemiddleware

type Service interface {
	FindByUID(uid string) (*FindByUIDResponse, error)
}
