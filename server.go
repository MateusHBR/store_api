package main

import (
	"log"

	"github.com/MateusHBR/mallus_api/adapter/database"
	"github.com/MateusHBR/mallus_api/router"
	"github.com/MateusHBR/mallus_api/server"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	db := database.PQDatabase{}
	dbConnct, err := db.OpenConnection()
	if err != nil {
		log.Fatal("Failed to open Database Connection")
	}
	defer dbConnct.Close()

	s := server.New().WithDatabaseConnection(dbConnct)

	router.RegisterRoutes(&s, engine, router.PingGroup)

	engine.Run()
}
