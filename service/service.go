package service

import (
    "github.com/go-martini/martini"// 用martini框架
)

func NewServer(port string) {
    m := martini.Classic()
    m.Get("/", func(params martini.Params) string {
        return "Hello world! " 
    })
    m.RunOnAddr(":"+port)
}
