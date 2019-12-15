package controller

import (
	"path"
	"github.com/astaxie/beego"
	"github.com/skdong/straw/pkg/context"
	"github.com/skdong/straw/pkg/text"
)

// Text API
type TextController struct{
	beego.Controller
}

func (self *TextController)Get(){
	relativePath := self.GetString("path")
	absolutePath := path.Join(context.WorkSpace, relativePath)
	self.Ctx.WriteString(text.GetText(absolutePath))
}
