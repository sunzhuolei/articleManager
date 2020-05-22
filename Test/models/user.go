package models

import "github.com/astaxie/beego/orm"
import _"github.com/go-sql-driver/mysql"

type User struct {
	Uid int `orm:"pk"`
	Username string
	Password string
	Addtime int
	Lastlogin int
	Articles []*Article `orm:"rel(m2m)"`
}

type Article struct {
	Id int
	Title string
	Content string
	Photo string
	Uid int
	Type *ArticleType `orm:"rel(fk)"`
	Addtime int
	Readnum int
	Users []*User `orm:"reverse(many)"`

}

/**
文章分类
 */
type ArticleType struct {
	TypeId int `orm:"pk"`
	TypeName string
	Articles []*Article `orm:"reverse(many)"`
}

func init(){
	orm.RegisterDataBase("default","mysql","root:@(127.0.0.1:3306)/acticle?charset=utf8")
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	orm.RunSyncdb("default",false,true)
}