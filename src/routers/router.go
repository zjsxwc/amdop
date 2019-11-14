package routers

import (
	"amdop/src/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/foo", &controllers.MainController{}, "Get:Foo")

	admin.Run()
}
