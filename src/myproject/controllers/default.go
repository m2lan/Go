package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "git.m2lan.com"
	this.Data["Email"] = "zyk0925@gmail.com"
	this.TplNames = "index.tpl"
}
