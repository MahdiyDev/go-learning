package main

import (
	"tm/api/routes"
	"tm/config"
	"tm/lib"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	env := config.Env{}
	err := env.Load()
	if err != nil {
		panic(err)
	}
	conn := lib.DatabaseConnection{DATABASE_URL: env.DATABASE_URL}
	err = conn.Setup()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes := routes.Routes{}
	routes.Setup(router, &conn)

	router.Run(":" + env.PORT)
}
