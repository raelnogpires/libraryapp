package services

import (
	"errors"

	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

// ugly? kinda, but returns everything's necessary
var allQuery string = `SELECT
	b.id, b.name, b.description,
	c.id AS category_id, c.name AS category_name,
	a.id AS author_id, a.name AS author_name, b.img_url
	FROM books AS b
	INNER JOIN categories AS c
		ON b.category_id = c.id
	INNER JOIN authors AS a
		ON b.author_id = a.id
	ORDER BY b.id;`

func GetAllBooks() ([]*models.FullBook, error) {
	db := database.GetDB()
	var books []*models.FullBook

	err := db.Raw(allQuery).Scan(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

// also ugly
var idQuery string = `SELECT
	b.id, b.name, b.description,
	c.id AS category_id, c.name AS category_name,
	a.id AS author_id, a.name AS author_name, b.img_url
	FROM books AS b
	INNER JOIN categories AS c
		ON b.category_id = c.id
	INNER JOIN authors AS a
		ON b.author_id = a.id
	WHERE b.id = ?
	ORDER BY b.id;`

func GetBookById(ID int) (*models.FullBook, error) {
	db := database.GetDB()
	var book *models.FullBook

	db.Raw(idQuery, ID).Scan(&book)
	if book == nil {
		return nil, errors.New("not found")
	}

	return book, nil
}

func CreateBook(b *models.Book) error {
	db := database.GetDB()

	err := db.Create(&b).Error
	if err != nil {
		return err
	}

	return nil
}

func EditBook(b *models.Book) error {
	db := database.GetDB()
	var check models.Book

	err := db.First(&check, b.ID).Error
	if err != nil {
		return err
	}

	err = db.Save(&b).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(ID int) error {
	db := database.GetDB()
	var book *models.Book

	err := db.First(&book, ID).Error
	if err != nil {
		return err
	}

	err = db.Delete(&models.Book{}, ID).Error
	if err != nil {
		return err
	}

	return nil
}
