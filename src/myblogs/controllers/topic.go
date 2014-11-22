package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.TplNames = "topic.html"
}

func (this *TopicController) Post() {
	title := this.Input().Get("title")
	content := this.Input().Get("content")

	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 302)
}

func (this *TopicController) Add() {
	this.Data["IsTopicAdd"] = true
	this.TplNames = "topicAdd.html"
}

func (this *TopicController) View() {
	this.TplNames = "topicView.html"
	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
}
