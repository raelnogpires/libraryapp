package category

import (
	"database/sql"
	"fmt"

	"github.com/raelnogpires/libraryapp/back-end/database"
)

type UseCase interface {
	GetAll() ([]*Category, error)
	GetById(Id int64) (*Category, error)
	Create(c *Category) error
	Update(c *Category) error
	Delete(Id int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: database.GetDB(),
	}
}

func (s *Service) GetAll() ([]*Category, error) {
	var result []*Category

	rows, err := s.DB.Query("SELECT * FROM LibraryDB.categories;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c Category
		err = rows.Scan(&c.Id, &c.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, &c)
	}

	return result, nil
}

func (s *Service) GetById(Id int64) (*Category, error) {
	var c Category

	stmt, err := s.DB.Prepare("SELECT * FROM LibraryDB.categories WHERE id = ?;")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(Id).Scan(&c.Id, &c.Name)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (s *Service) Create(name string) error {
	tx, err := s.DB.Begin()

	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO LibraryDB.categories VALUES (DEFAULT, ?);")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(name)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) Update(c *Category) error {
	if c.Id == 0 {
		return fmt.Errorf("invalid id")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE LibraryDB.categories SET name = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(c.Name, c.Id)
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

	_, err = tx.Exec("DELETE FROM LibraryDB.categories WHERE id = ?", Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
