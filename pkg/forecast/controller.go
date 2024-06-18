package forecast

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/forecast")
	routes.GET("", GetForecast)
}
