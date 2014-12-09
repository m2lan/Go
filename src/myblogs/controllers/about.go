package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io"
	"log"
	"strings"
)

type AboutController struct {
	beego.Controller
}

type History struct {
	Year, Day, Title, Detail string
}

const jsonStream = `
		{"Year":"2014","Day":"8.30 10:00","Title":"创建项目","Detail":"项目架构采用国内beego框架+orm"}
		{"Year":"2014","Day":"8.30 10:30","Title":"配置路由","Detail":"配置路由"}
		{"Year":"2014","Day":"8.30 11:00","Title":"创建模板","Detail":"项目模板使用bootstrap"}
		{"Year":"2014","Day":"8.30 12:00","Title":"添加首页","Detail":"首页显示文章列表"}
		{"Year":"2014","Day":"8.30 13:00","Title":"添加文章","Detail":"文章添加"}
		{"Year":"2014","Day":"8.31 10:00","Title":"添加国际化","Detail":"集成国际化i18n"}
		{"Year":"2014","Day":"8.31 11:00","Title":"添加编辑器","Detail":"添加文章页面集成ueditor"}
		{"Year":"2014","Day":"8.31 12:00","Title":"添加关于","Detail":"按照时间轴显示"}
		{"Year":"2014","Day":"8.31 13:00","Title":"添加右边栏","Detail":"显示浏览前十的文章"}
		{"Year":"2014","Day":"8.31 14:00","Title":"添加登陆页面","Detail":"添加登陆页面"}
   `

func (this *AboutController) Get() {
	this.Data["IsAbout"] = true
	this.TplNames = "about.html"
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	i := 0
	list := make(map[int]History)
	for {
		var h History
		if err := dec.Decode(&h); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		list[i] = h
		i++
	}
	this.Data["List"] = list
}
