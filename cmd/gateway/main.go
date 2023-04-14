package main

import (
	"github.com/liyuanwu2020/msgo/engine"
	"github.com/liyuanwu2020/msgo/engine/gateway"
	"github.com/liyuanwu2020/msgo/register"
	"github.com/liyuanwu2020/order/service"
)

func main() {
	e := engine.Default()
	e.Use(engine.Logging, engine.Recovery)
	e.SetRegister(register.MsNacosDefault())
	e.Get("/user2/*", service.Route)

	var configs []gateway.GWConfig
	configs = append(configs, gateway.GWConfig{
		Name:        "user",
		Path:        "/user/**",
		Host:        "127.0.0.1",
		Port:        9003,
		ServiceName: "user",
	}, gateway.GWConfig{
		Name: "goods",
		Path: "/goods/**",
		Host: "127.0.0.1",
		Port: 9002,
	})
	e.SetGateConfigs(configs)
	e.Run(":80")

}
