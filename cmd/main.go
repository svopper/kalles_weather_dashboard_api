package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/router"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/metObs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/oceanObs"
)

func main() {
	router := router.InstantiateRouter()
	port := os.Getenv("PORT")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("pkg/common/envs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if port == "" {
		log.Println("PORT environment variable not set. Defaulting to 8080")
		port = "8080"
	}
	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	metObs.RegisterRoutes(router)
	oceanObs.RegisterRoutes(router)
	router.Run("localhost:" + port)
}
