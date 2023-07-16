package main

import (
	"etcd-vision/routes"
	"fmt"
)

func main() {
	// 创建一个默认的 Gin 引擎
	r := routes.InitRouter()

	fmt.Println("listening on :8080")
	// 启动服务器，监听指定端口
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
