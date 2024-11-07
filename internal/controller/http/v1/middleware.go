package v1

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		statusCode := c.Writer.Status()

		duration := time.Since(start)

		log.Printf("status_code: %d, duration: %d ns\n", statusCode, duration.Nanoseconds())
	}
}
