package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id int64
	Title string `orm:"size(100)"`
	Content string `orm:"type(text)"`
}

func init() {
	orm.RegisterModel(new(Article))
}
