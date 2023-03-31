package server

import (
	"net/http"
	"project/iCredidentials/internal/mongodb"

	"github.com/gin-gonic/gin"
)

func (server *Server) RegisterWebsite(ctx *gin.Context) {

	var params mongodb.RegisterWebsiteParams

	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := server.database.RegisterWebsiteTx(ctx, params); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Valid")
}

func (server *Server) GetAllRegisteredWebsites(ctx *gin.Context) {

	var userId string

	//bind

	result, err := server.database.GetRegisteredWebsites(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}
