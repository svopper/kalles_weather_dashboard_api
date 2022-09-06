package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InstantiateRouter() *gin.Engine {
	router := gin.New()

	config := cors.Default()
	router.Use(config)
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	return router
}
