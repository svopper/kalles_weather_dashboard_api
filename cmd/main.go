package main

import (
	"log"
	"os"

	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/router"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/metObs"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/oceanObs"
)

func main() {
	router := router.InstantiateRouter()
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("PORT environment variable not set. Defaulting to 8080")
		port = "8080"
	}

	metObs.RegisterRoutes(router)
	oceanObs.RegisterRoutes(router)
	router.Run("localhost:8080")
}
