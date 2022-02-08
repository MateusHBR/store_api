package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MateusHBR/mallus_api/server"
)

func ListProducts(s *server.Server, _ *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		productService := productService{
			repository: productRepository{
				DB: s.DB,
			},
		}

		products, _ := productService.getAllProducts()

		c.JSON(http.StatusOK, ProductListDTO{}.fromEntity(products))
	}
}
