package category

type UseCase interface {
	GetAll() ([]*Category, error)
	GetById(Id int64) (*Category, error)
	Create(c *Category) error
	Update(c *Category) error
	Delete(Id int64) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAll() ([]*Category, error) {
	return nil, nil
}

func (s *Service) GetById() (*Category, error) {
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
