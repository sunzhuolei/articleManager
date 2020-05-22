package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/astaxie/beego/orm"
	"Test/models"
	"crypto/md5"
	"encoding/hex"
	"time"
)

type RegistController struct {
	beego.Controller
}


func (this *RegistController) Get(){
	this.TplName = "regist.html"
}

func (this *RegistController) Post(){
	//获取用户名
	username := this.Input().Get("username")
	fmt.Println(username)
	//获取密码
	password := this.Input().Get("password")
	if username == "" || password == ""{
		fmt.Println("账号密码不能为空！")
		this.TplName = "regist.html"
		return
	}
	//存储账号密码
	o := orm.NewOrm()
	user := new(models.User)
	user.Username = username
	hash := md5.New()
	hash.Write([]byte(password))
	//密码加密处理
	user.Password = hex.EncodeToString(hash.Sum(nil))
	user.Addtime = int(time.Now().Unix())
	//进行插入操作
	n,_ := o.Insert(user)
	if n>0{
		//插入成功
		fmt.Println(user)
		this.Ctx.WriteString("suc")
		fmt.Println("注册成功！")
		this.Redirect("/login",302)
	}else{
		fmt.Println("注册失败")
	}

	//this.TplName = "login.html"
}