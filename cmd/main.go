package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/svopper/kalles_weather_dashboard_v2/docs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/db"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/envs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/router"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/forecast"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/metObs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/oceanObs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/stations"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Title = "Kalle's Weather API"
	docs.SwaggerInfo.Description = "This is a weather API for Kalle's Weather Dashboard App"
	docs.SwaggerInfo.Version = "1.0"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func main() {
	envs.ConfigureViper()
	router := router.InstantiateRouter()
	setupSwagger(router)
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

	dbClient := db.Init()

	metObs.RegisterRoutes(router, dbClient)
	oceanObs.RegisterRoutes(router, dbClient)
	stations.RegisterRoutes(router, dbClient)
	forecast.RegisterRoutes(router, dbClient)
	router.Run("localhost:" + port)
}
