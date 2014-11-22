package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
	"myproject/controllers"
)

func init() {

	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/addUser", &controllers.MainController{}, "get:AddUser")
	// var FilterUser = func(ctx *context.Context) {
	// 	_, ok := ctx.Input.Session("uid").(int)
	// 	fmt.Println(ctx.Request.RequestURI)
	// 	if !ok && ctx.Request.RequestURI != "/addUser" {
	// 		ctx.Redirect(302, "/addUser")
	// 	}
	// }
	// beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.HomeController{})
}
