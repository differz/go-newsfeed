package gin

import "github.com/gin-gonic/gin"

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) != 0 {
			var errors []error
			for _, e := range c.Errors {
				errors = append(errors, e)
			}
			responseError(c, errors...)
		}
	}
}
