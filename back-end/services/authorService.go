package services

import (
	"strconv"

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

func GetAuthorById(ID string) (*models.Author, error) {
	intid, err := strconv.Atoi(ID)
	if err != nil {
		return nil, err
	}

	db := database.GetDB()
	var a *models.Author

	err = db.First(&a, intid).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}

func CreateAuthor(a models.Author) error {
	db := database.GetDB()

	err := db.Create(&a).Error
	if err != nil {
		return err
	}

	return nil
}

func EditAuthor(a models.Author, ID string) (*models.Author, error) {
	n, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		return nil, err
	}

	db := database.GetDB()
	a.ID = uint(n)

	err = db.Save(&a).Error
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func DeleteAuthor(ID string) error {
	intid, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	db := database.GetDB()
	err = db.Delete(&models.Author{}, intid).Error
	if err != nil {
		return err
	}

	return nil
}
