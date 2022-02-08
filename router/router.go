package router

import (
	"github.com/MateusHBR/mallus_api/server"
	"github.com/gin-gonic/gin"
)

type mallusRoute func(*server.Server, *gin.Engine)

func RegisterRoutes(server *server.Server, engine *gin.Engine, routes ...mallusRoute) {
	for _, route := range routes {
		route(server, engine)
	}
}
