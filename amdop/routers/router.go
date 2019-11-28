package routers

import (
	"amdop/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	beego.Router("/admin", &controllers.MainController{}, "get:Admin")
	beego.Router("/", &controllers.MainController{}, "get:Front")

	admin.Run()
}
