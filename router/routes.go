package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MateusHBR/mallus_api/product"
	"github.com/MateusHBR/mallus_api/server"
)

func PingGroup(s *server.Server, engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.GET("/product", product.ListProducts(s, engine))
}
