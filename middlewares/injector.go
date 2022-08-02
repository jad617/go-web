package middlewares

import "github.com/gin-gonic/gin"

func InjectContext(ctxKey string, ctxValue string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ctxKey, ctxValue)
		c.Next()
	}
}
