package main

import "github.com/liyuanwu2020/msgo"

func main() {
	engine := msgo.Default()

	engine.Run(":80")
}
