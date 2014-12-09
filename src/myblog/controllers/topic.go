package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"strconv"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopics"] = true
	this.TplNames = "admin/topic.html"
}

func (this *TopicController) Post() {
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	classID, _ := strconv.Atoi(this.Input().Get("classID"))

	_, err := models.AddTopic(title, content, int64(classID))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin/topics", 302)
}

func (this *TopicController) Add() {
	this.Data["IsTopicAdd"] = true
	this.TplNames = "topicAdd.html"
	classList, err := models.FindClassify(0)
	if err != nil {
		beego.Error(err)
	}
	this.Data["ClassList"] = classList
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

type TopicAddController struct {
	beego.Controller
}

func (this *TopicAddController) Get() {
	this.Data["IsTopicAdd"] = true
	this.TplNames = "admin/topic_add.html"
	classList, err := models.FindClassify(0)
	if err != nil {
		beego.Error(err)
	}
	this.Data["ClassList"] = classList
}
