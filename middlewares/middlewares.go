package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func ConfigureContentType() func(c *gin.Context) {
	return func(c *gin.Context) {
		// verify if the request is starting with /api
		// /api/*
		if c.Request.URL.Path == "/api" || c.Request.URL.Path == "/api/" {
			// Set the Content-Type header to application/json for API requests
			c.Header("Content-Type", "application/json; charset=utf-8")
		} else {
			// Set the Content-Type header to text/html for non-API requests
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
