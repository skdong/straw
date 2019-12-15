package controller

import (
	"path"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/skdong/straw/pkg/context"
	"github.com/skdong/straw/pkg/content"
)

type ContentController struct{
	beego.Controller
}

func (self * ContentController) Get(){
	relativePath := self.GetString("path")
	absolutePath := path.Join(context.WorkSpace, relativePath)
	info, err := content.GetInfo(absolutePath)
	if err != nil {
		self.Ctx.ResponseWriter.Status = http.StatusNotFound
		self.Ctx.WriteString(fmt.Sprint(err))
	}else{
		buf, err := json.Marshal(info)
		if err != nil{
			self.Ctx.ResponseWriter.Status = http.StatusInternalServerError
			self.Ctx.WriteString(fmt.Sprint(err))
		}
		self.Ctx.WriteString(string(buf))
	}

}