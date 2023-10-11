package main

import (
	"BookMall/cache"
	"BookMall/config"
	"BookMall/route"
)

func main() {
	config.Init()
	cache.Init()
	r := route.NewRouter()
	_ = r.Run(config.HttpPort)
}
