package gin

import (
	"bytes"
	"net/http"

	"github.com/VitaliiHurin/go-newsfeed/api"
	"github.com/gin-gonic/gin"
)

func responseSuccess(c *gin.Context, data interface{}) {
	res := gin.H{
		"status": "success",
	}
	if data != nil {
		res["data"] = data
	}
	c.JSON(
		http.StatusOK,
		res,
	)
}

func responseError(c *gin.Context, errors ...error) {
	var messages string
	var code int

	if len(errors) == 1 {
		code = api.HTTPStatusCodeByError(errors[0])
	} else {
		code = http.StatusInternalServerError
	}

	var buf bytes.Buffer
	for _, err := range errors {
		buf.WriteString(err.Error())
		buf.WriteRune('\n')
	}
	messages = buf.String()

	c.JSON(
		code,
		gin.H{
			"status":   "error",
			"messages": messages,
		},
	)
}
