package auth

import (
	"github.com/MateusHBR/mallus_api/server"
	"github.com/gin-gonic/gin"
)

const BREARER_SCHEMA = "Bearer "

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BREARER_SCHEMA):]
		token, err := AuthService{
			jwtKey: "asdasdsadasasd-key",
		}.validateToken(tokenString)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(server.MakeUnauthorizedResponse())
			return
		}

		c.Next()
	}
}
