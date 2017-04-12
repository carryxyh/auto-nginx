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

var defaultAlias = "default"

func init() {

	db = config.GetDBconfig()

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase(db.Alias, db.DriverName, db.DataSource, db.MaxIdleConns, db.MaxOpenConns)

	orm.RegisterDataBase(defaultAlias, db.DriverName, db.DataSource, db.MaxIdleConns, db.MaxOpenConns)

	orm.RegisterModel(new(Template))

	if db.CreateTable {
		orm.RunSyncdb(db.Alias, false, true)
	}

	ormer = orm.NewOrm()
}

/**
	根据id查询
 */
func QueryById(id int) Template {
	t := Template{Id: id}
	err := ormer.Read(&t)
	if err != nil {
		panic(err)
	}
	return t
}

/**
	查询所有的记录
 */
func QueryAll() []*Template {
	var t []*Template
	qs := ormer.QueryTable("template")
	qs.All(&t)
	return t
}

/**
	插入一条模板
 */
func Insert(tmpl Template) int64 {
	id, err := ormer.Insert(&tmpl)
	if err != nil {
		panic(err)
	}
	return id
}

/**
	根据id删除
 */
func DeleteById(id int) bool {
	t := Template{Id: id}
	num, err := ormer.Delete(&t)
	if err != nil {
		panic(err)
	}
	return num == 1
}

/**
	更新一条数据
 */
func Update(tmpl Template) bool {
	num, err := ormer.Update(&tmpl)
	if err != nil {
		panic(err)
	}
	return num == 1
}