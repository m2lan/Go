package controllers

import (
	"github.com/astaxie/beego"
	. "myblog/models"
	"myblog/utils"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	username := this.GetSession("username")
	if username == nil {
		this.Redirect("/admin/login", 302)
	} else {
		user, err := GetUserByName(username.(string))
		if err != nil {
			beego.Error(err)
		}
		all, used := utils.DiskUsages(beego.AppConfig.String("path"))
		this.Data["User"] = user
		this.Data["All"] = all
		this.Data["Used"] = used
		this.Data["IsAdmin"] = true
		this.TplNames = "admin/index.html"
	}
}
