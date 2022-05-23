package seeders

import (
	"fmt"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	db.Begin()

	queries := [3]string{CreateTables, PopulateAuthorsAndCategories, PopulateBooks}

	db.Exec("DROP DATABASE IF EXISTS LibraryDB;")
	err := db.Exec("CREATE DATABASE LibraryDB;").Error
	if err != nil {
		fmt.Println("right here")
		db.Rollback()
	}

	for _, v := range queries {
		err := db.Exec(v).Error
		if err != nil {
			fmt.Println("error in exec query loop")
			db.Rollback()
			break
		}
	}

	db.Commit()
}
