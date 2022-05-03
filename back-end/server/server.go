package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/server/routes"
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

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("server running at port: %s", s.port)
	log.Fatal(router.Run(":" + s.port))
}
