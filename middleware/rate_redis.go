package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stewardyohanes/url-shortener/config"
)

const (
	RateLimitWindows = time.Minute
	MaxRequests = 10
)

func RedisRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", ip)

		count, err := config.RedisClient.Incr(config.RedisCtx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Terjadi kesalahan saat memproses permintaan.",
			})
			return
		}

		if count == 1 {
			config.RedisClient.Expire(config.RedisCtx, key, RateLimitWindows)
		}

		fmt.Printf("🔢 %s sudah request %dx\n", ip, count)

		if count > int64(MaxRequests) {
			ttl, _ := config.RedisClient.TTL(config.RedisCtx, key).Result()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Terlalu banyak permintaan, coba lagi nanti.",
				"retry_after": ttl.Seconds(),
			})
			return
		}

		c.Next()
	}
}
