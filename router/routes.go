package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MateusHBR/mallus_api/auth"
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
	engine.POST("/product", product.AddProduct(s, engine))
	engine.PUT("/product/:id", product.UpdateProduct(s, engine))
	engine.DELETE("/product/:id", product.DeleteProduct(s, engine))
}

func AuthGroup(s *server.Server, engine *gin.Engine) {
	engine.POST("/signup", auth.SignUp(s, engine))
	engine.POST("/refresh", auth.AuthorizeJWT(), auth.Refresh(s, engine))
}
