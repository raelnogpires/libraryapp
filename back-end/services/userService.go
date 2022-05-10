package services

import (
	"errors"

	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/models"
)

func RegisterUser(u *models.User) error {
	db := database.GetDB()

	db.Raw("SELECT * FROM users WHERE email = ?", u.Email).Scan(&u)
	if u.ID != uint(0) {
		return errors.New("email already registered")
	}

	err := db.Create(&u).Error
	if err != nil {
		return errors.New("couldn't register user")
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
