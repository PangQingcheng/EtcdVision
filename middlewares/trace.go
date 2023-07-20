package middlewares

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Trace(c *gin.Context) {
	log.Debugf("request url: %s \n", c.Request.URL.String())
	c.Next()
}
