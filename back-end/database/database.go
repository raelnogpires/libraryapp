package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// the var db receives the memory pointer for the db engine
// making it's use simplier
var db *gorm.DB

func InitDB() {
	var dsn string = os.Getenv("CONNECTION_STR")

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if error isn't null, is treated
	if err != nil {
		fmt.Println("something went wrong, couldn't connect with mysql database")
		log.Fatal("error: ", err)
	}

	fmt.Println("db running")

	// https://gorm.io/docs/index.html
	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	// migrations.RunMigrations(db)
	// seeders.Seeder(db)
}

func CloseConnection() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

// exporting our db for querys
func GetDB() *gorm.DB {
	return db
}
