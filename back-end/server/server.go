package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   os.Getenv("SERVER_PORT"),
		server: gin.Default(),
	}
}
