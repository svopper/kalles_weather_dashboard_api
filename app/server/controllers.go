package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svopper/kalles_weather_dashboard_v2/app/server/ocean"
	"github.com/svopper/kalles_weather_dashboard_v2/app/server/temperature"
)

func InstantiateControllers() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	api := router.Group("/api")
	{
		api.GET("/metObs", temperature.GetIndex)
		api.GET("/ocean", ocean.GetOcean)
		// api.GET("/map", station_map.GetMap)
	}

	return router
}
