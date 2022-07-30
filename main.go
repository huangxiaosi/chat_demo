package main

import (
	"chat-demo/conf"
	"chat-demo/router"
)

func main() {
	conf.Init()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
