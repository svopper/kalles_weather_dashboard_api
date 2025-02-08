package forecast

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// Implement DMI forecast EDR service

type handler struct {
	DB *redis.Client
}

func RegisterRoutes(r *gin.Engine, db *redis.Client) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/forecast-edr")
	routes.GET("/temperature", h.GetTemperatureForecast)
}
