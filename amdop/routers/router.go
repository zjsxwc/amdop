package routers

import (
	"amdop/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	beego.Router("/admin", &controllers.MainController{}, "Get:Admin")
	beego.Router("/", &controllers.MainController{}, "Get:Front")

	admin.Run()
}
