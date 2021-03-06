package server

import (
	"log"
	"webapi-with-go/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine //Necessário para iniciar o server
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
