package main

import (
	_ "./models"
	_ "./routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}



