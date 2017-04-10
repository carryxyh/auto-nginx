package dao

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"auto-ng/config"
)

type Template struct {
	Id          int
	ProjectName string `orm:"size(50)"`
	Content     string `orm:"size(200)"`
}

var db *config.DBconfig

var ormer orm.Ormer

func init() {

	db = config.GetDBconfig()

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase(db.Alias, db.DriverName, db.DataSource, db.Conns)

	orm.RegisterModel(new(Template))

	orm.RunSyncdb(db.Alias, false, true)

	ormer = orm.NewOrm()
}

func QueryById(id int) Template {
	t := Template{Id: id}
	err := ormer.Read(&t)
	if err != nil {
		panic(err)
	}
	return t
}

func Insert(tmpl *Template) int {
	id, err := ormer.Insert(&tmpl)
	if err != nil {
		panic(err)
	}
	return id
}

func DeleteById(id int) bool {
	t := Template{Id: id}
	num, err := ormer.Delete(&t)
	if err != nil {
		panic(err)
	}
	return num == 1
}

func Update(tmpl *Template) bool {
	num, err := ormer.Update(&tmpl)
	if err != nil {
		panic(err)
	}
	return num == 1
}