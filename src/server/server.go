package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/src/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	var p string = os.Getenv("SERVER_PORT")
	return Server{
		port:   p,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("server running at port: %s", s.port)
	log.Fatal(router.Run(":" + s.port))
}
