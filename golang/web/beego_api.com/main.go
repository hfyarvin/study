package main

import (
	"fmt"

	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("db")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn)
	fmt.Printf("数据库链接成功！%s\n", conn)
}

func main() {
	o := orm.NewOrm()
	o.Using("default")
	beego.Run()
}
