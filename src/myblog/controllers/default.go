package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/wetalk/modules/utils"
	"myblog/models"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Prepare() {
	this.Data["Lang"] = "en-US"
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "index-test.html"
	q := this.GetString("q")
	c, _ := strconv.Atoi(this.Input().Get("c"))
	page, _ := strconv.Atoi(this.Input().Get("p"))
	if page == 0 {
		page = 1
	}

	var num int64
	if len(q) > 0 {
		topics, err := models.SearchTopic(q)
		if err != nil {
			beego.Error(err)
		}

		count, err := models.GetTopicCount(q, int64(c))
		if err != nil {
			beego.Error(err)
		}
		num = count
		this.Data["Topics"] = topics
	} else {
		topics, err := models.GetAllTopic(page, int64(c))
		if err != nil {
			beego.Error(err)
		}

		count, err := models.GetTopicCount(q, int64(c))
		if err != nil {
			beego.Error(err)
		}
		num = count
		this.Data["Topics"] = topics
	}
	classifyCount, err := models.GetClassifyCount()
	if err != nil {
		beego.Error(err)
	}
	this.Data["classifyCount"] = classifyCount
	pageNum, _ := beego.AppConfig.Int("pageNum")
	p := utils.NewPaginator(this.Ctx.Request, pageNum, num)
	this.Data["paginator"] = p
}
