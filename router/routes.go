package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGroup(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
