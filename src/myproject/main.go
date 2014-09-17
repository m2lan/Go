package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myproject/models"
	_ "myproject/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
