package main

import (
    "cloudgo/service"
    flag "github.com/spf13/pflag"
    "os"
)

const (
    PORT string = "8080" // 默认端口8080
)

func main() {
    port := os.Getenv("PORT") //检查端口
    if len(port) == 0 {
        port = PORT //没有端口则使用默认端口
    }
    //对端口号进行解析
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
    service.NewServer(port) //启动server端口服务
}
