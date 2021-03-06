package migrations

import (
	"github.com/raelnogpires/libraryapp/src/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Author{})
	db.AutoMigrate(models.Category{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Book{})
}
