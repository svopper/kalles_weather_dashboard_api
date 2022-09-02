package ping

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/ping")
	routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
