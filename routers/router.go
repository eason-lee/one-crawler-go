package routers

import (
	"one-crawler-go/controllers"

	"github.com/astaxie/beego"
)

func init()  {
	beego.Router("/", &controllers.IndexController{}, "GET:Index")
	beego.Router("/ones/pre/:id", &controllers.IndexController{}, "GET:PreOneView")
	beego.Router("/ones/next/:id", &controllers.IndexController{}, "GET:NextOneView")
	// beego.Router("/pre/id([0-9A-Za-z]{24})", &controllers.IndexController{}, "GET:NextView")
}
