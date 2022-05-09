package services

import (
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func RegisterUser(u *models.User) error {
	db := database.GetDB()
	var check models.User

	err := db.First(&check, u.Email).Error
	if err != nil {
		return err
	}

	err = db.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func Login(u *models.UserLogin) error {
	db := database.GetDB()
	var check models.User

	err := db.First(&check, u.Username).Error
	if err != nil {
		return err
	}

	return nil
}

func GetUserById(ID int) (*models.User, error) {
	db := database.GetDB()
	var user *models.User

	err := db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
