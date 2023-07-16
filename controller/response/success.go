package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
}
