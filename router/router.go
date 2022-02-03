package router

import "github.com/gin-gonic/gin"

type mallusRoute func(*gin.Engine)

func RegisterRoutes(engine *gin.Engine, routes ...mallusRoute) {
	for _, route := range routes {
		route(engine)
	}
}
