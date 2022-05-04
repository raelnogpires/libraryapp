package migrations

import (
	"github.com/raelnogpires/libraryapp/back-end/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Author{})
	db.AutoMigrate(models.Category{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Book{})

	// https://www.w3schools.com/sql/sql_foreignkey.asp
	// apparently GORM doesn't create FK's in AutoMigrate even if its tagged in the Model
	// so this is the way to go
	db.Exec(
		"ALTER TABLE books ADD FOREIGN KEY (category_id) REFERENCES categories(id)",
	)
	db.Exec(
		"ALTER TABLE books ADD FOREIGN KEY (author_id) REFERENCES authors(id)",
	)
}
