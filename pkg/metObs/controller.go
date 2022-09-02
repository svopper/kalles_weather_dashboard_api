package metObs

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/metObs")
	routes.GET("/", GetMetObs)
}
