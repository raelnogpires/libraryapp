package user

type UseCase interface {
	GetById(Id int64) (*User, error)
	Create(u *User) error
	Delete(Id int64) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetById() (*User, error) {
	return nil, nil
}

func (s *Service) Create() error {
	return nil
}

func (s *Service) Delete() error {
	return nil
}
