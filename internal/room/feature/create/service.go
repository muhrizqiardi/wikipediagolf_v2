package create

import "github.com/google/uuid"

type CodeGenerator interface {
	Generate() string
}

type codeGenerator struct {
}

func NewCodeGenerator() *codeGenerator {
	return &codeGenerator{}
}

func (c *codeGenerator) Generate() string {
	return uuid.NewString()[:6]
}

type Service interface {
	Create(ownerUID string) (*CreateRoomResponse, error)
}

type service struct {
	codeGenerator CodeGenerator
	repository    Repository
}

func NewService(c CodeGenerator, repository Repository) *service {
	return &service{
		codeGenerator: c,
		repository:    repository,
	}
}

func (s *service) Create(ownerUID string) (*CreateRoomResponse, error) {
	code := s.codeGenerator.Generate()
	room, err := s.repository.InsertRoom(code, "open")
	if err != nil {
		return nil, err
	}
	member, err := s.repository.InsertRoomMember(room.ID, ownerUID, true)
	if err != nil {
		return nil, err
	}

	return &CreateRoomResponse{*room, *member}, nil
}
