package main

import (
	_ "amdop/src/models"
	_ "amdop/src/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}



