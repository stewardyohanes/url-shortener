package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stewardyohanes/url-shortener/handlers"
	"github.com/stewardyohanes/url-shortener/middleware"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/shorten", middleware.RedisRateLimit(), handlers.Shorten)
	r.GET("/:short_code", handlers.Redirect)

	return r
}
