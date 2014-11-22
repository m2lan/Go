package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {
	this.Data["Lang"] = "en-US"
}

func (this *MainController) Get() {
	fmt.Println("get")
	this.Data["IsHome"] = true
	this.TplNames = "index-test.html"

	topics, err := models.GetAllTopic(true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}
