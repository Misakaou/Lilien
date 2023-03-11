package middlewares

import (
	"strings"
	"lilien-server/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		config := config.GetConfig()
		reqKey := ctx.Request.Header.Get("X-Auth-Key")
		reqSecret := ctx.Request.Header.Get("X-Auth-Secret")

		var key string
		var secret string
		if key = config.GetString("http.auth.key"); len(strings.TrimSpace(key)) == 0 {
			ctx.AbortWithStatus(500)
		}
		if secret = config.GetString("http.auth.secret"); len(strings.TrimSpace(secret)) == 0 {
			ctx.AbortWithStatus(401)
		}
		if key != reqKey || secret != reqSecret {
			ctx.AbortWithStatus(401)
			return
		}
		ctx.Next()
	}
}
