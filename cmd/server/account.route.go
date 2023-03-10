package server

import (
	"fmt"
	"net/http"
	"project/iCredidentials/internal/mongodb"

	"github.com/gin-gonic/gin"
)

func (server *Server) CreateAccount(ctx *gin.Context) {
	var acc mongodb.AccountCreateAccountParams
	//Has special code or is using default ?
	if err := ctx.BindJSON(&acc); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(acc)

	data, err := server.database.CreateAccountTx(ctx, acc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)

}

type SignInParams struct {
	Username string `json:"Username" bson:"Username"`
	Password string `json:"Password" bson:"Password"`
}

func (server *Server) SignIn(ctx *gin.Context) {

	var params SignInParams

	if err := ctx.BindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	results, err := server.database.SignInTx(ctx, params.Username, params.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, results)
}
