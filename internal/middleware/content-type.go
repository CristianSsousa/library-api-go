package middleware

import (
	"github.com/gin-gonic/gin"
)

func ContentTypeMiddleware(c *gin.Context)  {
	c.Header("Content-Type", "application/json")
	c.Next()
}