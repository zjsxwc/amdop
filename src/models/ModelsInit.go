package models

import (
	"os"
)

func init() {
	args := os.Args
	for _, v := range args {
		//初始化Src下的业务Model数据库
		if v == "-syncdb" {
			_ = new(Article)
			_ = new(Comment)
		}
	}
}
