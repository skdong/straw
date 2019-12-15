package controller

import(
	"github.com/astaxie/beego"
)

type MainController struct{
	beego.Controller
}

func (this * MainController) Get(){
	this.Redirect("/static/index.html", 302)
}