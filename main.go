package main

import (
    "etcd-vision/routes"
    "fmt"
    "os/exec"
)

func main() {
    // 创建一个默认的 Gin 引擎
    r := routes.InitRouter()

    fmt.Println("listening on :8080")
    ctx := make(chan bool)
    go func() {
        // 启动服务器，监听指定端口
        err := r.Run(":8080")
        if err != nil {
            panic(err)
        }
        ctx <- false
    }()

    url := "http://localhost:8080/console/index.html "
    fmt.Println("打开浏览器并访问： " + url + " 开始使用！")
    fmt.Println("正常尝试自动打开浏览器...")
    err := exec.Command("cmd", "-c", "start", url).Start()
    if err != nil {
        fmt.Println(err)
        fmt.Println("打开浏览器失败，请复制链接到浏览器访问...")
    }
    <-ctx
}
