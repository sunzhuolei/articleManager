package routers

import (
	"Test/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)
var filter = func(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil{
		ctx.Redirect(302,"/login")
	}
}
func init() {
	//beego.InsertFilter("/regist",beego.BeforeRouter,filter)
	beego.InsertFilter("/article",beego.BeforeRouter,filter)
	beego.InsertFilter("/addarticle",beego.BeforeRouter,filter)
	beego.InsertFilter("/update",beego.BeforeRouter,filter)
	beego.InsertFilter("/addtype",beego.BeforeRouter,filter)
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/logout", &controllers.LoginController{},"get:LogOut")
    beego.Router("/regist", &controllers.RegistController{})
    beego.Router("/article", &controllers.ArticleController{})
    beego.Router("/addarticle", &controllers.ArticleController{},"get:AddArticle;post:HandleAdd")
    beego.Router("/update", &controllers.ArticleController{},"get:UpdateGet")
    beego.Router("/addtype", &controllers.ArticleController{},"get:AddType;post:TypePost")
    beego.Router("/test1", &controllers.ArticleController{},"get:TestInsert")
    beego.Router("/test2", &controllers.ArticleController{},"get:TestSelect")
}
