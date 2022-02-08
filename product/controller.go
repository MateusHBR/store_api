package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MateusHBR/mallus_api/product/repository"
	"github.com/MateusHBR/mallus_api/product/service"
	"github.com/MateusHBR/mallus_api/server"
)

func ListProducts(s *server.Server, engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {

		productService := service.ProductService{
			Repository: repository.ProductRepository{
				DB: s.DB,
			},
		}

		productService.GetAllProducts()

		products, _ := productService.GetAllProducts()

		c.JSON(http.StatusOK, products)
	}
}
