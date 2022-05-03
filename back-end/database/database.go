package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Atribuímos o database para a variável db, tornando seu uso mais simples.
var db *gorm.DB

func InitDB() {
	// Criamos a string de conexão usando os métodos fmt.Sprintf + os.Getenv
	// para sigilo de informações delicadas.
	var connectionString string = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	// Caso o erro não seja nulo, tratamos ele
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

// Função para exportar nosso database e realizar querys
func GetDB() *gorm.DB {
	return db
}
