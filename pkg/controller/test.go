package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TestController struct {
	beego.Controller
}

func (self *TestController) Get() {
	path := self.GetString("path")
	logs.Info(path)
	self.Ctx.WriteString(path)
}
