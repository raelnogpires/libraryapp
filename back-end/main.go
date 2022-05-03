package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/raelnogpires/libraryapp/back-end/database"
	"github.com/raelnogpires/libraryapp/back-end/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
	s := server.NewServer()

	s.Run()
}
