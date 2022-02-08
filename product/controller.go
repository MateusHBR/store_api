package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MateusHBR/mallus_api/adapter/database"
	repository "github.com/MateusHBR/mallus_api/product/repository"
	service "github.com/MateusHBR/mallus_api/product/service"
)

func ListProducts(engine *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {

		productService := service.ProductService{
			Repository: repository.ProductRepository{
				DatabaseAdapter: database.PQDatabase{},
			},
		}

		productService.GetAllProducts()

		products, _ := productService.GetAllProducts()

		c.JSON(http.StatusOK, products)
	}
}
