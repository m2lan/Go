package functions

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func db() {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	dbname := beego.AppConfig.String("dbname")

	orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@/%s?charset=utf8", username, password, dbname), 30)
}
