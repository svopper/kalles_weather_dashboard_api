package stations

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type handler struct {
	DB *redis.Client
}

func RegisterRoutes(r *gin.Engine, db *redis.Client) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/stations")
	routes.GET("metObs", h.GetMetObsStations)
	routes.GET("metObs/:id", h.GetMetObsStation)
	routes.GET("oceanObs", h.GetOceanObsStations)
	routes.GET("oceanObs/:id", h.GetOceanObsStation)
}
