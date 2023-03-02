package server

import (
	"project/iCredidentials/internal/mongodb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	database *mongodb.Database
}

func InitServer(db *mongodb.Database) *Server {
	server := &Server{database: db}

	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowAllOrigins: true, //To be changed
		AllowMethods:    []string{"GET", "POST"},
	}))

	server.router = route

	return server

}

func (server *Server) Start(host string) error {
	return server.router.Run(host)
}

func (server *Server) InitRoutes() {

	api := server.router.Group("/api")
	{
		api.POST("/signin")
		api.POST("/signup")
	}
}
