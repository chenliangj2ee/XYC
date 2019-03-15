package routers

import (
	"XYC/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Index")
	beego.Router("/index", &controllers.MainController{},"get:Index")
	beego.Router("/shop", &controllers.MainController{},"get:Shop")
	beego.Router("/about-us", &controllers.MainController{},"get:About")
	beego.Router("/blog", &controllers.MainController{},"get:Blog")
	beego.Router("/contact-us", &controllers.MainController{},"get:Contact")
}
