package main

import (
	"github.com/MateusHBR/mallus_api/router"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	router.RegisterRoutes(engine, router.PingGroup)

	engine.Run()
}
