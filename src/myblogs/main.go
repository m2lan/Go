package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	_ "myblog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	// 自动创建表
	orm.RunSyncdb("default", false, true)
	beego.Run()

}
