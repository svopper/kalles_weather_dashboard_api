package forecast

import "github.com/gin-gonic/gin"

func GetForecast(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Forecast",
	})
}
