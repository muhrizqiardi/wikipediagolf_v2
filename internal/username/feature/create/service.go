package create

type Service interface {
	Create(payload CreateUsernameRequest) (*CreateUsernameResponse, error)
}
