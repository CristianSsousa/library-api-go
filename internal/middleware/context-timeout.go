package middleware

import (
	"context"
	"library-api-go/pkg/apperros"
	"library-api-go/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeoutMiddleware (timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		
		done := make(chan struct{})

		go func() {
			c.Next()
			close(done)
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			response.ErrorResponse(c, http.StatusRequestTimeout, apperros.ErrTimeOut)
			c.Abort()
			return
		}

	}
}
	