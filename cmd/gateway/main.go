package main

import (
	"github.com/liyuanwu2020/msgo/engine"
	"log"
)

func main() {
	e := engine.Default()
	e.Get("/user", func(ctx *engine.Context) {
		log.Println("hello")
	})
	e.Run(":80")

	//engine := msgo.Default()
	//engine.OpenGateWay = true
	//var configs []gateway.GWConfig
	//configs = append(configs, gateway.GWConfig{
	//	Name: "order",
	//	Path: "/order/**",
	//	Host: "127.0.0.1",
	//	Port: 9003,
	//}, gateway.GWConfig{
	//	Name: "goods",
	//	Path: "/goods/**",
	//	Host: "127.0.0.1",
	//	Port: 9002,
	//})
	//engine.SetGateConfigs(configs)
	//engine.Run(":80")
}
