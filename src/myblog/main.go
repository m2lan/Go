package main

import (
	"github.com/astaxie/beego"
	_ "myblog/functions"
	_ "myblog/routers"
)

func main() {
	beego.Run()
}
