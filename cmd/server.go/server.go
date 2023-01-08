package server

import (
	"net/http"
	"projects/iCredidentials/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	archive *database.Archive
}

func NewServer(db *database.Archive) *Server {
	server := &Server{archive: db}
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowAllOrigins: true, AllowHeaders: []string{"Authorization", "projectId", "content-type"},
		AllowMethods: []string{"GET", "POST"},
	}))
	server.router = route

	server.router.GET("/api", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, "Hello")
	})
	return server
}

func (server *Server) Start(host string) error {
	return server.router.Run(host)
}

func (server *Server) InitRoutes() {

}
