package services

import (
	"errors"

	"github.com/raelnogpires/libraryapp/back-end/auth"
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

func Login(u *models.User) error {
	db := database.GetDB()
	var check models.User

	err := db.Where("email = ?", u.Email).First(&check).Error
	if err != nil {
		return errors.New("user not found")
	}

	if check.Password != auth.SHA256Encoder(u.Password) {
		return errors.New("invalid credentials")
	}

	return nil
}

func DeleteMe(ID int) error {
	db := database.GetDB()
	var check models.User

	err := db.Where("id = ?", ID).First(&check).Error
	if err != nil {
		return errors.New("user not found")
	}

	err = db.Delete(&models.User{}, ID).Error
	if err != nil {
		return errors.New("couldn't delete user")
	}

	return nil
}

func GetUser(email string) (*models.User, error) {
	db := database.GetDB()
	var user models.User

	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
