package main

import (
	"github.com/stewardyohanes/url-shortener/config"
	"github.com/stewardyohanes/url-shortener/routes"
)

func main() {
	config.InitDB()
	config.InitRedis()
	router := routes.SetupRoutes()

	router.Run(":8080")
}
