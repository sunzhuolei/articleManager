package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"

	"fmt"
	"path"
	//"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"Test/models"
	"time"
	//"strconv"
	"math"
	//"os/user"
	//"github.com/mattn/go-gtk/gtk"
	//"os/user"
	"github.com/astaxie/beego/logs"
)

type ArticleController struct {
	beego.Controller
}


func (this *ArticleController) Get(){
	//数据查询

	o := orm.NewOrm()
	perpage := 10
	//curpage := 1
	page,_:= this.GetInt("page",1)
	fmt.Println("dsdssdsd",page)
	var articles []models.Article
	//获取总条数
	qs := o.QueryTable("article")
	total,_ := qs.Count()
	totalpage := int(math.Ceil(float64(total)/float64(perpage)))
	o.QueryTable("article").Limit(perpage,(page-1)*perpage).All(&articles)
	this.Data["articles"] = articles
	this.Data["curpage"] = page
	this.Data["totalpage"] = totalpage
	fmt.Println(this.Data)
	if page == 1{
		this.Data["isFirstPage"] = true
	}else{
		this.Data["isFirstPage"] = false
	}
	if page == totalpage{
		this.Data["isLastPage"] = true
	}else{
		this.Data["isLastPage"] = false
	}
	pageSlice := make([]int,3)
	if page ==1{
		pageSlice[0] = page
		pageSlice[1] = page+1
		pageSlice[2] = page+2
	}else if page>1 && page < totalpage{
		pageSlice[0] = page-1
		pageSlice[1] = page
		pageSlice[2] = page+1
	}else{
		pageSlice[0] = page-2
		pageSlice[1] = page-1
		pageSlice[2] = page
	}
	this.Data["pageslice"] = pageSlice
	this.Data["username"] = 444
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["title"] = "title.html"
	this.TplName = "index.html"

}


func (this *ArticleController) AddArticle(){
	//查询文章分类
	o := orm.NewOrm()
	var list []models.ArticleType
	o.QueryTable("article_type").All(&list)
	this.Data["list"] = list
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["title"] = "title2.html"
	this.TplName = "addarticle.html"
}


func (this *ArticleController) HandleAdd(){
	//接受参数
	title := this.GetString("title")
	content := this.GetString("content")
	typeid,_:= this.GetInt("type")
	file,header,_ := this.GetFile("photo")
	//格式和大小限制
	defer file.Close()
	ext := path.Ext(header.Filename)
	if ext !=".jpg" && ext !=".png" && ext != ".gif"{
		fmt.Println("图片格式不被支持！")
		return
	}
	if header.Size > 1024 * 1024 *10{
		fmt.Println("你在整大点试试！")
		return
	}
	err3 := this.SaveToFile("photo","./static/img/"+header.Filename)
	if err3 != nil{
		fmt.Println("上传失败！")
		return
	}
	//保存数据到数据库
	o := orm.NewOrm()
	item := new(models.Article)
	item.Title = title
	item.Content = content
	item.Addtime = int(time.Now().Unix())
	item.Type = &models.ArticleType{TypeId:typeid}
	item.Photo = "/static/img/"+header.Filename
	_,err4:=o.Insert(item)
	if err4 != nil{
		fmt.Println("插入失败！")
		return
	}
	//跳转到首页
	this.Redirect("/article",302)
}

func (this *ArticleController) UpdateGet(){
	id,_ := this.GetInt("id",1)
	//查询数据
	o := orm.NewOrm()
	article := new(models.Article)
	article.Id = id
	o.Read(article)
	var list []models.ArticleType
	o.QueryTable("article_type").All(&list)
	this.Data["article"] = article
	this.Data["list"] = list
	this.Layout = "layout.html"
	this.TplName = "updatearticle.html"
}


func (this *ArticleController) AddType(){
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["title"] = "title3.html"
	this.TplName = "addtype.html"
}

func (this *ArticleController) TypePost(){
	typename := this.GetString("type")
	o := orm.NewOrm()
	typeinfo := new(models.ArticleType)
	typeinfo.TypeName = typename
	o.Insert(typeinfo)
	this.Redirect("/article",302)
}


func (this *ArticleController) TestInsert(){
	//插入相关数据
	o := orm.NewOrm()
	username := this.GetSession("username")

	user := new(models.User)
	user.Username = username.(string)
	logs.Info(user)
	articleid,_ := this.GetInt("id",1)
	article := new(models.Article)
	article.Id = articleid
	m2m := o.QueryM2M(article,"Users")
	o.Read(user,"Username")
	//进行插入操作
	m2m.Add(user)
	this.Ctx.WriteString("suc")
}

func (this *ArticleController) TestSelect(){
	//插入相关数据
	o := orm.NewOrm()
	username := this.GetSession("username")

	user := new(models.User)
	user.Username = username.(string)
	logs.Info(user)
	articleid,_ := this.GetInt("id",1)
	article := new(models.Article)
	article.Id = articleid
	o.LoadRelated(article,"Users")
	for _,val := range article.Users{
		logs.Info(val.Username,val.Uid)
	}

	var test []models.User
	o.QueryTable("user").Filter("Articles__Article__Id",articleid).Distinct().All(&test)
	for _,val2 := range test{
		logs.Info(val2.Username,val2.Uid)
	}
	this.Ctx.WriteString("suc")
}