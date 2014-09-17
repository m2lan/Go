package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["myproject/controllers:HomeController"] = append(beego.GlobalControllerRouter["myproject/controllers:HomeController"],
		beego.ControllerComments{
			"Get",
			"/home",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["myproject/controllers:MainController"] = append(beego.GlobalControllerRouter["myproject/controllers:MainController"],
		beego.ControllerComments{
			"Get",
			"/",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["myproject/controllers:MainController"] = append(beego.GlobalControllerRouter["myproject/controllers:MainController"],
		beego.ControllerComments{
			"AddUser",
			"/addUser",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["myproject/controllers:MainController"] = append(beego.GlobalControllerRouter["myproject/controllers:MainController"],
		beego.ControllerComments{
			"NewList",
			"/new",
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["myproject/controllers:MainController"] = append(beego.GlobalControllerRouter["myproject/controllers:MainController"],
		beego.ControllerComments{
			"Post",
			"/addUser",
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["myproject/controllers:MainController"] = append(beego.GlobalControllerRouter["myproject/controllers:MainController"],
		beego.ControllerComments{
			"UpdateUser",
			"/updateUser",
			[]string{"get"},
			nil})

}
