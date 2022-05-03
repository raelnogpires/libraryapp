package author

import (
	"database/sql"
	"fmt"

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

	rows, err := s.DB.Query("SELECT * FROM LibraryDB.authors;")

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

func (s *Service) GetById(Id int64) (*Author, error) {
	var a Author

	// O método Prepare verifica se a consulta é válida
	stmt, err := s.DB.Prepare("SELECT * FROM LibraryDB.authors WHERE id = ?;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(Id).Scan(&a.Id, &a.Name)
	if err != nil {
		return nil, err
	}

	// Retorna a posição de memória de a
	return &a, nil
}

func (s *Service) Create(name string) error {
	// Iniciamos uma transação
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO LibraryDB.authors VALUES (DEFAULT, ?);")
	if err != nil {
		return err
	}

	defer stmt.Close()

	// O método Exec retorna um result, porém não há interesse nele
	// sendo assim é possível ignorá-lo com _
	_, err = stmt.Exec(name)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) Update(a *Author) error {
	if a.Id == 0 {
		// Caso o id seja 0 (não presente no db), o sistema gera um erro
		// para prevenir um update/delete sem WHERE
		return fmt.Errorf("invalid id")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE LibraryDB.authors SET name = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.Name, a.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) Delete(Id int64) error {
	if Id == 0 {
		return fmt.Errorf("invalid id")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM LibraryDB.authors WHERE id = ?;", Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
