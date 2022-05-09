package services

import (
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func GetAllAuthors() ([]*models.Author, error) {
	db := database.GetDB()
	var a []*models.Author

	err := db.Find(&a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}

func GetAuthorById(ID int) (*models.Author, error) {
	db := database.GetDB()
	var a *models.Author

	err := db.First(&a, ID).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}

func CreateAuthor(a *models.Author) error {
	db := database.GetDB()

	err := db.Create(&a).Error
	if err != nil {
		return err
	}

	return nil
}

func EditAuthor(a *models.Author) error {
	db := database.GetDB()

	// error isn't being returned
	err := db.Model(&models.Author{}).Where("id = ?", a.ID).Update("name", a.Name).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteAuthor(ID int) error {
	db := database.GetDB()

	// error isn't being returned
	err := db.Delete(&models.Author{}, ID).Error
	if err != nil {
		return err
	}

	return nil
}
