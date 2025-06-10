package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

func ConfigureContentType() func(c *gin.Context) {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Header("Content-Type", "application/json; charset=utf-8")
		} else {
			c.Header("Content-Type", "text/html; charset=utf-8")
		}
		c.Next()
	}
}

func ConfigureLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
