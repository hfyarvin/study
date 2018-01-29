package user_model

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(100)"`
	Nickname string `orm:"size(100)"`
	Pwd      string `orm:"size(100)"`
	Email    string `orm:"size(100)"`
	Sex      string `orm:"size(2)"`
	Roleid   string `orm:"size(100)"`
	Status   int64
	Phone    string `orm:"size(16)"`
}

func Create(uid, status int64, name, nickname, pwd, email,
	sex, roleId, phone string) (user User) {
	// user,err:=
}

func QueryById(uid int64) (User, bool) {
	o := orm.NewOrm()
	u := User{Id: uid}

	err := o.Read(&u)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return u, false
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return u, false
	} else { //存在
		fmt.Println(u.Id, u.Name)
		return u, true
	}
}
