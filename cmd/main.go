package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/db"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/envs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/router"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/metObs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/oceanObs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/stations"
)

func main() {
	envs.ConfigureViper()
	router := router.InstantiateRouter()
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("PORT environment variable not set. Defaulting to 8080")
		port = "8080"
	}
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive!",
		})
	})

	db_client := db.Init()

	metObs.RegisterRoutes(router, db_client)
	oceanObs.RegisterRoutes(router, db_client)
	stations.RegisterRoutes(router, db_client)
	router.Run("localhost:" + port)
}
