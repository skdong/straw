package main

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/skdong/straw/pkg/controller"
)

func initLogs(){
	logs.SetLogger("console")
}

func initApp(){
	initLogs()
}

func loadRouters(){
	beego.Router("/", &controller.MainController{})
	/*
	beego.Router("/text", &controller.TextController{})
	beego.Router("/dir", &controller.DirController{})
	beego.Router("/test", &controller.TestController{})
	*/
	beego.Router("/content", &controller.ContentController{})
}

func main(){
	initApp()
	loadRouters()
	beego.Run()
}