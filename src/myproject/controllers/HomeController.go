package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
)

type HomeController struct {
	beego.Controller
}

type user struct {
	Age int
}

// @router /home [get]
func (this *HomeController) Get() {
	// this.Ctx.WriteString("HomeController")
	this.Abort("404")
	valid := validation.Validation{}
	u := user{20}

	valid.MaxSize(u.Age, 15, "ageMax")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	if v := valid.Max(u.Age, 12, "age"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
	}

	this.TplNames = "addUser.html"
}
