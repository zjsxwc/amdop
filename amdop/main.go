package main

import (
	_ "amdop/models"
	_ "amdop/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/beego/admin/src/models"
	"os"
)

func main() {
	beego.Run()
}

func init()  {
	os.Setenv("TZ", "UTC")
	//基础类型 int、string等默认已经注册了的，不需要为session注册，需要注册特定类型，在文件`sess_utils.go`里init注册了的也不用再注册
	gob.Register(models.User{})
	gob.Register(map[string]bool{})
	gob.Register(map[string]string{})
}


