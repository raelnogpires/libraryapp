package services

import (
	"errors"

	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func GetAllCategories() ([]*models.Category, error) {
	db := database.GetDB()
	var cat []*models.Category

	err := db.Find(&cat).Error
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return cat, nil
}

func GetCategoryById(ID int) (*models.Category, error) {
	db := database.GetDB()
	var cat *models.Category

	err := db.First(&cat, ID).Error
	if err != nil {
		return nil, errors.New("category not found")
	}

	return cat, nil
}

func CreateCategory(cat *models.Category) error {
	db := database.GetDB()

	err := db.Create(&cat).Error
	if err != nil {
		return errors.New("couldn't register category")
	}

	return nil
}

func EditCategory(cat *models.Category) error {
	db := database.GetDB()
	var check models.Category

	err := db.First(&check, cat.ID).Error
	if err != nil {
		return errors.New("category not found")
	}

	err = db.Save(&cat).Error
	if err != nil {
		return errors.New("couldn't update category")
	}

	return nil
}

func DeleteCategory(ID int) error {
	db := database.GetDB()
	var cat *models.Category

	err := db.First(&cat, ID).Error
	if err != nil {
		return errors.New("category not found")
	}

	err = db.Delete(&models.Category{}, ID).Error
	if err != nil {
		return errors.New("couldn't delete category")
	}

	return nil
}
