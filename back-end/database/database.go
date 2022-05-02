package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Atribuímos o database para a variável db, tornando seu uso mais simples.
var db sql.DB

func InitDB() {
	// Criamos a string de conexão usando os métodos fmt.Sprintf + os.Getenv
	// para sigilo de informações delicadas.
	var connectionString string = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", connectionString)

	// Caso o erro não seja nulo, tratamos ele
	if err != nil {
		log.Fatal("something went wrong: ", err)
	}

	fmt.Println("db running")

	// https://github.com/go-sql-driver/mysql#important-settings
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
}

// Função para exportar nosso database e realizar querys
func GetDB() *sql.DB {
	return &db
}
