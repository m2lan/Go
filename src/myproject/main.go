package main

import (
	"github.com/astaxie/beego"
	_ "myproject/routers"
)

func main() {
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}
