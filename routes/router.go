package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stewardyohanes/url-shortener/handlers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/shorten", handlers.Shorten)
	r.GET("/:short_code", handlers.Redirect)

	return r
}
