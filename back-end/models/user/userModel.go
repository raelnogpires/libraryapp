package user

import (
	"database/sql"
	"fmt"

	"github.com/raelnogpires/libraryapp/back-end/database"
)

type UseCase interface {
	GetById(Id int64) (*User, error)
	Create(u *User) error
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

func (s *Service) GetById(Id int64) (*User, error) {
	var u User

	stmt, err := s.DB.Prepare("SELECT id, username, email FROM LibraryDB.users WHERE id = ?;")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(Id).Scan(&u.Id, &u.Username, &u.Email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *Service) Create(username string, email string, password string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO LibraryDB.users VALUES (DEFAULT, ?, ?, ?);")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(username, email, password)
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

	_, err = tx.Exec("DELETE FROM LibraryDB.users WHERE id = ?;", Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
