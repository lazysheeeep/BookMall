package main

import (
	"BookMall/config"
	"BookMall/route"
)

func main() {
	config.Init()
	r := route.NewRouter()
	_ = r.Run(config.HttpPort)
}
