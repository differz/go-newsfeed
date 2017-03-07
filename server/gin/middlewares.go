package gin

import "github.com/gin-gonic/gin"

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) != 0 {
			responseError(c, c.Errors...)
		}
	}
}