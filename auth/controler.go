package auth

import (
	"fmt"
	"net/http"

	"github.com/MateusHBR/mallus_api/server"
	"github.com/gin-gonic/gin"
)

func SignUp(s *server.Server, _ *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: get from os.Getenv("ACCESS_SECRET")
		var authService Service = AuthService{
			jwtKey: "asdasdsadasasd-key",
			repository: authRepository{
				DB: s.DB,
			},
		}

		var createAccountDTO CreateAccountDTO
		if err := c.ShouldBindJSON(&createAccountDTO); err != nil {
			c.AbortWithStatusJSON(server.MakeBadRequestResponse(""))
			return
		}

		res, err := authService.createAccount(createAccountDTO.toEntity())

		if err != nil {
			c.AbortWithStatusJSON(server.MakeInternalServerErrorResponse())
			return
		}

		c.JSON(http.StatusCreated, AuthResponseDTO{Token: res})
	}
}

func Refresh(s *server.Server, _ *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exist := c.Get("claims")

		if !exist {
			c.AbortWithStatusJSON(server.MakeInternalServerErrorResponse())
			return
		}

		fmt.Println(claims)
		//TODO: get from os.Getenv("ACCESS_SECRET")
		var authService Service = AuthService{
			jwtKey: "asdasdsadasasd-key",
		}

		//TODO: implement refresh token
		res, err := authService.refreshToken()

		if err != nil {
			c.AbortWithStatusJSON(server.MakeInternalServerErrorResponse())
			return
		}

		c.JSON(http.StatusCreated, AuthResponseDTO{Token: res})
	}
}
