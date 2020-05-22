package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"Test/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type LoginController struct {
	beego.Controller
}


func (this *LoginController) Get(){
	username := this.Ctx.GetCookie("username")
	if username != ""{
		this.Data["persit"] = true
	}else{
		this.Data["persit"] = false
	}
	this.Data["username"] = username
	this.TplName = "login.html"
}

func (this *LoginController) Post(){
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	persit := this.Input().Get("persit")
	//根据账号密码进行相关查询
	o := orm.NewOrm()
	user := new(models.User)
	user.Username = username
	hash := md5.New()
	hash.Write([]byte(password))
	//密码加密处理
	user.Password = hex.EncodeToString(hash.Sum(nil))
	//进行相关查询操作
	err := o.Read(user,"Username","Password")
	if err != nil{
		fmt.Println("账号密码错误")
		this.Ctx.WriteString("账号密码错误")
	}else{
		fmt.Println("登录成功！")
		//保存用户登录信息
		user.Lastlogin = int(time.Now().Unix())
		fmt.Println(int(time.Now().Unix()))
		o.Update(user,"LastLogin")
		if persit == "on"{
			this.Ctx.SetCookie("username",username,time.Second *60 *60)
		}else{
			this.Ctx.SetCookie("username",username,-1)
		}
		this.SetSession("username",username)
		//this.Ctx.WriteString("suc")
		this.Redirect("/article",302)
	}
	//this.TplName = "login.html"
}

func (this *LoginController) LogOut(){
	this.DelSession("username")
	this.Redirect("/login",302)

}