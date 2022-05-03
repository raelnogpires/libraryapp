package book

import (
	"database/sql"
	"fmt"

	"github.com/raelnogpires/libraryapp/back-end/database"
)

type UseCase interface {
	GetAll() ([]*Book, error)
	GetById(Id int64) (*Book, error)
	Create(b *Book) error
	Update(b *Book) error
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

func (s *Service) GetAll() ([]*Book, error) {
	var result []*Book

	rows, err := s.DB.Query("SELECT * FROM LibraryDB.books;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var b Book
		err = rows.Scan(&b.Id, &b.Name, &b.Description, &b.CategoryId, &b.AuthorId, &b.ImgUrl)
		if err != nil {
			return nil, err
		}
		result = append(result, &b)
	}

	return result, nil
}

func (s *Service) GetById(Id int64) (*Book, error) {
	var b Book

	stmt, err := s.DB.Prepare("SELECT * FROM LibraryDB.books WHERE id = ?;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(Id).Scan(&b.Id, &b.Name, &b.Description, &b.CategoryId, &b.AuthorId, &b.ImgUrl)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (s *Service) Create(name string, description string, categoryId int64, authorId int64, imgUrl string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO LibraryDB.books VALUES (DEFAULT, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, description, categoryId, authorId, imgUrl)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) Update(Id int64, Name string, Description string, ImgUrl string) error {
	if Id == 0 {
		return fmt.Errorf("invalid id")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(
		`UPDATE LibraryDB.books
		SET name = ?,
		description = ?,
		img_url = ?
		WHERE id = ?;`,
	)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(Name, Description, ImgUrl, Id)
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

	_, err = tx.Exec("DELETE FROM LibraryDB.books WHERE id = ?;", Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
