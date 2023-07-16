package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestErr(c *gin.Context, err error) {
	message := ""
	if err != nil {
		message = err.Error()
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
	})
}
