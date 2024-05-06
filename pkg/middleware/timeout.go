package middleware

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func timeoutResponse(c *gin.Context) {
	c.Status(http.StatusRequestTimeout)
}

func Timeout() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(1000*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse))
}
