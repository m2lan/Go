package main

import (
	"github.com/astaxie/beego"
<<<<<<< HEAD
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
=======
	_ "myproject/routers"
)

func main() {
>>>>>>> 11e05332a743e61629d966e998bc9ee56732bdf7
	beego.Run()
}
