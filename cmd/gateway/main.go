package main

import (
	"github.com/liyuanwu2020/msgo/engine"
	"github.com/liyuanwu2020/order/service"
)

func main() {
	e := engine.Default()
	e.Use(engine.Logging)
	e.Get("/order/*", service.Route, func(handlerFunc engine.HandlerFunc) engine.HandlerFunc {
		return func(ctx *engine.Context) {
			ctx.Logger.Info("执行顺序 method")
			handlerFunc(ctx)
		}
	}, func(handlerFunc engine.HandlerFunc) engine.HandlerFunc {
		return func(ctx *engine.Context) {
			ctx.Logger.Info("执行顺序 method2")
			handlerFunc(ctx)
		}
	})
	e.Get("/goods/*", service.Route, func(handlerFunc engine.HandlerFunc) engine.HandlerFunc {
		return func(ctx *engine.Context) {
			ctx.Logger.Info("方法级别 MiddleHandler goods1")
			handlerFunc(ctx)
		}
	}, func(handlerFunc engine.HandlerFunc) engine.HandlerFunc {
		return func(ctx *engine.Context) {
			ctx.Logger.Info("方法级别 MiddleHandler goods2")
			handlerFunc(ctx)
		}
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
