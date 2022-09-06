package oceanObs

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/oceanObs")
	routes.GET("", GetOceanObs)
}
