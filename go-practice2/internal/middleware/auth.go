package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "log"
)

func APIKeyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        key := c.GetHeader("X-API-Key")
        log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

        if key != "secret123" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
