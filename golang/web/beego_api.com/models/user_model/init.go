package user_model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//驱动
func init() {
	orm.RegisterModel(new(User))
}
