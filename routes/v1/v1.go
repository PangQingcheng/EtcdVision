package v1

import (
	"etcd-vision/controller/etcd"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		v1.POST("/etcds", etcd.NewETCD)
		v1.GET("/etcds", etcd.ListETCD)
		v1.DELETE("/etcds/:name", etcd.DeleteETCD)

		v1.POST("/etcds/:name/connect", etcd.ConnectETCD)
		v1.POST("/etcds/:name/disconnect", etcd.DisConnectETCD)
		v1.GET("/etcds/:name/status", etcd.GetETCDConnectStatus)

		v1.GET("/etcds/:name/keys", etcd.GetETCDKeys)

		v1.GET("/etcds/:name/value", etcd.GetETCDValue)
		v1.POST("/etcds/:name/value", etcd.SetETCDValue)
	}
}
