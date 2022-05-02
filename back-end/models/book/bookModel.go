package book

type UseCase interface {
	GetAll() ([]*Book, error)
	GetById(Id int64) (*Book, error)
	Create(b *Book) error
	Update(b *Book) error
	Delete(Id int64) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAll() ([]*Book, error) {
	return nil, nil
}

func (s *Service) GetById() (*Book, error) {
	return nil, nil
}

func (s *Service) Create() error {
	return nil
}

func (s *Service) Update() error {
	return nil
}

func (s *Service) Delete() error {
	return nil
}
