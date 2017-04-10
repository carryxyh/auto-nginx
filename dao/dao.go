package dao

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type  Template struct {
	Id int
	ProjectName string `orm:"size(50)"`
	Content string `orm:"size(200)"`

}

func init() {

}
