package handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requeridToken := os.Getenv("TOKEN")
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("error: token not found"))
			return
		}

		if requeridToken != token {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("error: invalid token"))
			return
		}
		ctx.Next()
	}
}
