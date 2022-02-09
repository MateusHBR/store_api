package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MateusHBR/mallus_api/server"
)

func ListProducts(s *server.Server, _ *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productService service = productService{
			repository: productRepository{
				DB: s.DB,
			},
		}

		products, _ := productService.getAllProducts()

		c.JSON(http.StatusOK, ProductListDTO{}.fromEntity(products))
	}
}

func AddProduct(s *server.Server, _ *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productService service = productService{
			repository: productRepository{
				DB: s.DB,
			},
		}

		var productDto ProductDTO
		if err := c.BindJSON(&productDto); err != nil {
			fmt.Println("Failed to read body")
		}

		product, _ := productService.addProduct(productDto.toEntity())

		c.JSON(http.StatusOK, ProductDTO{}.fromEntity(product))
	}
}

func UpdateProduct(s *server.Server, _ *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productService service = productService{
			repository: productRepository{
				DB: s.DB,
			},
		}

		var productDto ProductDTO
		if err := c.BindJSON(&productDto); err != nil {
			fmt.Println("Failed to read body")
		}

		productDto.ID = c.Param("id")

		product, _ := productService.updateProduct(productDto.toEntity())

		c.JSON(http.StatusOK, ProductDTO{}.fromEntity(product))
	}
}
