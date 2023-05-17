package oceanObs

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

	routes := r.Group("/oceanObs")
	routes.GET("", h.GetOceanObs)
	routes.GET(":stationId", h.GetOceanObsByStationId)
}
