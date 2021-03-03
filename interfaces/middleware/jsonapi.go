package middleware

import (
	"github.com/gin-gonic/gin"
	"golang-ddd-starter/helpers"
)

/**
* jsonapi header
 */
func Jsonapi() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Accept") != "application/vnd.api+json" {
			helpers.ReturnStatusUnsupportedMediaType(c)
			c.Abort()
			return
		}
		c.Writer.Header().Set("Accept", "application/vnd.api+json")
		c.Writer.Header().Set("Content-Type", "application/vnd.api+json")
		c.Next()
	}
}
