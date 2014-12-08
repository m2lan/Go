package controllers

import (
	"github.com/astaxie/beego"
	. "myblog/models"
	"strconv"
)

type ClassifyController struct {
	beego.Controller
}

func (this *ClassifyController) Get() {
	username := this.GetSession("username")
	if username == nil {
		this.Redirect("/admin/login", 302)
	} else {
		user, err := GetUserByName(username.(string))
		if err != nil {
			beego.Error(err)
		}
		classify, err := FindClassify(0)
		if err != nil {
			beego.Error(err)
		}
		this.Data["User"] = user
		this.Data["IsClassify"] = true
		this.Data["Classify"] = classify
		maxid, err := GetClassifyMaxID()
		if err != nil {
			beego.Error(err)
		}
		this.Data["MaxID"] = maxid + 1
		this.TplNames = "admin/classify.html"
	}
}

func (this *ClassifyController) Post() {
	title := this.GetString("title")
	err := AddClassify(title)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin/classify", 302)
}

type ClassifyOptController struct {
	beego.Controller
}

func (this *ClassifyOptController) Get() {
	id := this.Ctx.Input.Param(":id")
	convID, _ := strconv.Atoi(id)
	err := DeleteClassify(int64(convID))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin/classify", 302)
}

type ClassifyUpdateController struct {
	beego.Controller
}

func (this *ClassifyUpdateController) Get() {
	id := this.Ctx.Input.Param(":id")
	convID, _ := strconv.Atoi(id)
	classify, err := FindClassify(int64(convID))
	if err != nil {
		beego.Error(err)
	}
	this.Data["Classify"] = classify
	this.TplNames = "admin/classify_update.html"
}

func (this *ClassifyUpdateController) Post() {
	id := this.GetString("id")
	convID, _ := strconv.Atoi(id)
	title := this.GetString("title")
	err := UpdateClassify(int64(convID), title)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin/classify", 302)
}
