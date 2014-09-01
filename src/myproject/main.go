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
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
