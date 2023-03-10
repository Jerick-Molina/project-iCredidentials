package server

import (
	"net/http"
	"project/iCredidentials/internal/security"

	"github.com/gin-gonic/gin"
)

func (server *Server) AuthorizeToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, "No token found")
			ctx.Abort()
		}

		claims, err := security.TokenReader(token)

		if err != nil {

			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()

		}

		//Sets users claims
		ctx.Set("claims", claims)

		ctx.Next()
	}
}
