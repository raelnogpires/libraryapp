package author

type UseCase interface {
	GetAll() ([]*Author, error)
	GetById(Id int64) (*Author, error)
	Create(a *Author) error
	Update(a *Author) error
	Delete(Id int64) error
}

// Estrutura que recebe a conexão com o DB
type Service struct{}

// Retorna um ponteiro em memória para determinada estrutura
func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAll() ([]*Author, error) {
	return nil, nil
}

func (s *Service) GetById() (*Author, error) {
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
