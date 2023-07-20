package routes

import (
	"etcd-vision/middlewares"
	v1 "etcd-vision/routes/v1"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"os"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//Panic 处理
	r.Use(gin.RecoveryWithWriter(os.Stdout))
	// gzip
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{".*"})))
	// gin中间件注册
	registerMiddleware(r)
	// 路由注册
	setUpRouter(r)
	return r
}

// 注册gin中间件
func registerMiddleware(r *gin.Engine) {
	r.Use(gin.Recovery())
	// Cors
	r.Use(middlewares.Cors)
	// swagger
	middlewares.RegisterSwaggerMiddleware(r)
	// 404
	r.NoRoute(middlewares.NoRouteHandler)
	r.Use(middlewares.Trace)
}

// 设置路由
func setUpRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1.RegisterRouter(api)
	}
}
