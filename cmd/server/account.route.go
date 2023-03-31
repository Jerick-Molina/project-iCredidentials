package server

import (
	"fmt"
	"net/http"
	"project/iCredidentials/internal/mongodb"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateAccount(ctx *gin.Context) {
	var params mongodb.CreateAccountParams
	//Has special code or is using default ?
	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	websiteId, _ := ctx.GetQuery("id")
	data, err := server.database.CreateAccountTx(ctx, params, websiteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)

}

func (server *Server) SignIn(ctx *gin.Context) {

	var params mongodb.SignInParams

	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	websiteId, _ := ctx.GetQuery("id")
	fmt.Println(params.WebsiteId)
	results, err := server.database.SignInTx(ctx, params, websiteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, results)
}
