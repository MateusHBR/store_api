package auth

import (
	"net/http"

	"github.com/MateusHBR/mallus_api/server"
	"github.com/gin-gonic/gin"
)

func SignUp(s *server.Server, _ *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		var authService Service = AuthService{
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
