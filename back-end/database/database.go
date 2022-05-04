package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/raelnogpires/libraryapp/back-end/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// the var db receives the memory pointer for the db engine
// making it's use simplier
var db *gorm.DB

func InitDB() {
	var db_user string = os.Getenv("DB_USER")
	var db_password string = os.Getenv("DB_PASSWORD")
	var db_host string = os.Getenv("DB_HOST")
	var db_name string = os.Getenv("DB_NAME")

	// connection string is created using fmt.Sprintf + os.Getenv methods
	// for protection of delicate information
	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_user,
		db_password,
		db_host,
		db_name,
	)

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

	migrations.RunMigrations(db)
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
