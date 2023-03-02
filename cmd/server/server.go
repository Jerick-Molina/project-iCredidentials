package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func InitServer() *Server {
	server := &Server{}

	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
	}))

	server.router = route

	return server

}

func (server *Server) Start(host string) error {
	return server.router.Run(host)
}

func (server *Server) InitRoutes() {

}
