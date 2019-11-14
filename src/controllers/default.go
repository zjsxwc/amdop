package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Xx struct {
	Vv int64
}

func x(px interface{})  {
	fmt.Println(px)
}


func (c *MainController) Get() {
	c.Data["Website"] = "GetGetGetGet"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	//x(new(Xx))

	c.SetSession("xxccc", "teststring")
	fmt.Println(c.GetSession("xxccc"))
}


func (c *MainController) Foo() {
	c.Data["Website"] = "FooFooFooFoo"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	x(new(Xx))


}
