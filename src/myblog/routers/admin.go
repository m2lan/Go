package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/admin/logout", &controllers.LogoutController{})
	beego.Router("/admin/classify", &controllers.ClassifyController{})
	beego.Router("/admin/classify/add", &controllers.ClassifyController{})
	beego.Router("/admin/classify/delete/:id", &controllers.ClassifyOptController{})
	beego.Router("/admin/classify/update/?:id", &controllers.ClassifyUpdateController{})
	beego.Router("/admin/topic/add", &controllers.TopicAddController{})
	beego.Router("/admin/topics", &controllers.TopicController{})
}
