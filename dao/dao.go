package dao

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"auto-ng/config"
)

type  Template struct {
	Id          int
	ProjectName string `orm:"size(50)"`
	Content     string `orm:"size(200)"`
}

var db *config.DBconfig

func init() {

	db = config.GetDBconfig()

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase(db.Alias, db.DriverName, db.DataSource, db.Conns)

	orm.RegisterModel(new(Template))

	orm.RunSyncdb(db.Alias, false, true)
}

func QueryById() {

}

func Insert() {

}

func DeleteById() {

}

func Update() {

}