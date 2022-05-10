package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/raelnogpires/libraryapp/back-end/auth"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const Bearer = "Bearer "
		header := ctx.GetHeader("authorization")
		if header == "" {
			ctx.JSON(401, gin.H{
				"error": "token not found or unauthorized",
			})
			return
		}

		token := header[len(Bearer):]

		if !auth.NewJWTService().ValidateToken(token) {
			ctx.JSON(401, gin.H{
				"error": "unauthorized",
			})
			return
		}
	}
}
