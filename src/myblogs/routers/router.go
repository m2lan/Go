package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"myblog/controllers"
	"strings"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/about", &controllers.AboutController{})
	langs := strings.Split(beego.AppConfig.String("types"), "|")

	for _, lang := range langs {
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error(err.Error())
		}
	}
	beego.AddFuncMap("i18n", i18n.Tr)
}
