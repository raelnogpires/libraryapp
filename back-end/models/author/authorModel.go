package author

import (
	"database/sql"

	"github.com/raelnogpires/libraryapp/back-end/database"
)

type UseCase interface {
	GetAll() ([]*Author, error)
	GetById(Id int64) (*Author, error)
	Create(a *Author) error
	Update(a *Author) error
	Delete(Id int64) error
}

// Estrutura que recebe a conexão com o DB
type Service struct {
	DB *sql.DB
}

// Retorna um ponteiro em memória para determinada estrutura
func NewService(db *sql.DB) *Service {
	return &Service{
		DB: database.GetDB(),
	}
}

func (s *Service) GetAll() ([]*Author, error) {
	var result []*Author

	rows, err := s.DB.Query("SELECT id, name FROM librarydb.authors")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var a Author
		err = rows.Scan(&a.Id, &a.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, &a)
	}

	return result, nil
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
