package controllers

import (
	"fmt"
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
type User struct {
	Id   int    `form:"_"`
	Name string `form:"username"`
	Pass string `form:"password"`
}

// @router / [get]
func (this *MainController) Get() {
	// this.Data["Website"] = "beego.me"
	// this.Data["Email"] = "astaxie@gmail.com"
	// this.TplNames = "index.tpl"
	// this.Ctx.WriteString("Hello world!!!")
	Users := struct {
		Name string
		Age  int
	}{
		Name: "zhangsan",
		Age:  10,
	}

	array := []int{1, 2, 3, 4, 5, 6, 7}

	this.Data["User"] = Users
	this.Data["Array"] = array
	this.TplNames = "index.tpl"
}

// @router /addUser [get]
func (this *MainController) AddUser() {
	flash := beego.ReadFromRequest(&this.Controller)
	if _, ok := flash.Data["notice"]; ok {
		//显示设置成功
		this.TplNames = "addUser.html"
	} else if _, ok = flash.Data["error"]; ok {
		//显示错误
		this.TplNames = "addUser.html"
	} else {
		// 不然默认显示设置页面
		this.TplNames = "addUser.html"
	}

}

// @router /new [get]
func (this *MainController) NewList() {
	tx := this.GetSession("new")
	if tx == nil {
		this.SetSession("new", "aaaa")
	}
	this.Ctx.WriteString(fmt.Sprintln(this.GetSession("new")))
}

// @router /addUser [post]
func (this *MainController) Post() {
	flash := beego.NewFlash()
	u := User{}

	if error := this.ParseForm(&u); error != nil {
		return
	}

	if u.Name == "" {
		flash.Error("username not null")
		flash.Store(&this.Controller)
		this.Redirect("/addUser", 302)
		return
	}
	flash.Notice("Settings saved!")
	flash.Store(&this.Controller)
	// this.Ctx.WriteString(u.Name)
	this.Redirect("/addUser", 302)
}

// @router /updateUser [get]
func (this *MainController) UpdateUser() {
	this.Ctx.Output.Body([]byte("updateuser"))
}
