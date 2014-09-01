package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "index.html"

	topics, err := models.GetAllTopic(true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}
