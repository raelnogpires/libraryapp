package services

import (
	"errors"

	"github.com/raelnogpires/libraryapp/src/database"
	"github.com/raelnogpires/libraryapp/src/models"
)

func GetAllAuthors() ([]*models.Author, error) {
	db := database.GetDB()
	var a []*models.Author

	// search an array/slice of authors
	err := db.Find(&a).Error

	if err != nil {
		// returns error for treatment at controller layter
		return nil, errors.New("internal server error")
	}

	// returns all the authors
	return a, nil
}

func GetAuthorById(ID int) (*models.Author, error) {
	db := database.GetDB()
	var a *models.Author

	// search the first row from db that matches the model passed as pointer and
	// parameters (in this case, id)
	err := db.First(&a, ID).Error
	if err != nil {
		return nil, errors.New("author not found")
	}

	// returns one author
	return a, nil
}

func CreateAuthor(a *models.Author) error {
	db := database.GetDB()

	err := db.Create(&a).Error
	if err != nil {
		return errors.New("couldn't register author")
	}

	return nil
}

func EditAuthor(a *models.Author) error {
	db := database.GetDB()
	var check models.Author

	// checks if author exist
	err := db.First(&check, a.ID).Error
	if err != nil {
		return errors.New("author not found")
	}

	// updates author
	// https://pkg.go.dev/gorm.io/gorm#DB.Save
	err = db.Save(&a).Error
	if err != nil {
		return errors.New("couldn't edit author")
	}

	return nil
}

func DeleteAuthor(ID int) error {
	db := database.GetDB()
	var a *models.Author

	// checks if author exist
	err := db.First(&a, ID).Error
	if err != nil {
		return errors.New("author not found")
	}

	// deletes author
	err = db.Delete(&models.Author{}, ID).Error
	if err != nil {
		return errors.New("couldn't delete author")
	}

	return nil
}
