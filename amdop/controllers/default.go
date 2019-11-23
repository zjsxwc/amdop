package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/syyongx/php2go"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Admin() {
	entry, err := c.getSPAConfig("admin_entry")
	if err == nil {
		c.Ctx.Redirect(302, entry)
		return
	}
	c.Data["json"] = &map[string]interface{}{"status": false, "info": "未配置单页", "code": err.Error()}
	c.ServeJSON()
}

func (c *MainController) Front() {
	entry, err := c.getSPAConfig("front_entry")
	if err == nil {
		c.Ctx.Redirect(302, entry)
		return
	}
	c.Data["json"] = &map[string]interface{}{"status": false, "info": "未配置单页", "code": err.Error()}
	c.ServeJSON()
}

func (c *MainController) getSPAConfig(key string) (string, error) {
	json, err := php2go.FileGetContents("static/spa/config.json")
	if err == nil {
		configMap := make(map[string]string)
		err := php2go.JSONDecode([]byte(json), &configMap)
		if err == nil {
			value, ok := configMap[key]
			if ok {
				return value, nil
			}
		}
	}
	return "", errors.New("NO_SPA_CONFIG")
}
