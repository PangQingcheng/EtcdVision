package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// 404 handler
func NoRouteHandler(c *gin.Context) {
	log.Infof("Not Found Path: %s", c.Request.RequestURI)
	c.String(http.StatusNotFound, "%s", "Page Not Found")
}
