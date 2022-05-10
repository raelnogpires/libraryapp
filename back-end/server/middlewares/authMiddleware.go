package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/auth"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const Bearer = "Bearer "
		header := ctx.GetHeader("authorization")
		if header == "Bearer" {
			// when authorization header is empty
			// const header only returns the string "Bearer"
			ctx.AbortWithStatusJSON(401, "token not found")
			return
		}

		token := header[len(Bearer):]

		if !auth.NewJWTService().ValidateToken(token) {
			ctx.AbortWithStatusJSON(401, "unauthorized")
			return
		}
	}
}
