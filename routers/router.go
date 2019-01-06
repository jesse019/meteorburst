package routers

import (
	"github.com/meteorburst/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	beego.Router("ws/join", &controllers.HomeController{}, "get:Join")
}
