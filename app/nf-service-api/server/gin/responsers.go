package gin

import (
	"net/http"

	"github.com/VitaliiHurin/go-newsfeed/app/nf-service-api/api"
	"github.com/gin-gonic/gin"
)

func responseSuccess(c *gin.Context, data interface{}) {
	res := gin.H{
		"status": "success",
	}
	if data != nil {
		res["data"] = data
	}
	c.JSON(http.StatusOK, res)
}

func responseError(c *gin.Context, errors ...error) {
	res := gin.H{
		"status": "error",
	}
	var code int
	if len(errors) == 1 {
		code = api.HTTPStatusCodeByError(errors[0])
		res["message"] = errors[0].Error()
	} else {
		code = http.StatusInternalServerError
		var messages []string
		for _, err := range errors {
			messages = append(messages, err.Error())
		}
		res["messages"] = messages
	}
	c.JSON(code, res)
}
