package user_model

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

const (
	USER_TABLENAME = "users"
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

func (u *User) TableName() string {
	return USER_TABLENAME
}

func Create(uid, status int64, name, nickname, pwd, email,
	sex, roleId, phone string) (user User) {
	user, exist := QueryById(uid)
	if exist == true {
		return user
	} else {
		o := orm.NewOrm()
		o.Using("mine")
		newuser := new(User)
		// newuser.Id = uid
		newuser.Name = name
		newuser.Nickname = nickname
		newuser.Pwd = pwd
		newuser.Email = email
		newuser.Sex = sex
		newuser.Roleid = roleId
		newuser.Status = status
		newuser.Phone = phone
		o.Insert(newuser)

		return *newuser
	}
}

//根据ID查找
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

//删除用户
func DeleteUserById(id int64) bool {
	o := orm.NewOrm()
	o.Using("default")

	//根据ID得到用户模型
	if num, err := o.Delete(&User{Id: id}); err == nil {
		fmt.Println("删除影响的行数：", num)
		return true
	} else {
		return false
	}
}

//更新
func UpdateUserById(id int, table string, field map[string]interface{}) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(table).Filter("Id", id).Update(field)
	if err == nil {
		return true
	}
	return false
}

func QueryUserById(name string) (User, error) {
	var user User

	o := orm.NewOrm()
	qs := o.QueryTable("users")

	err := qs.Filter("Name", name).One(&user)
	fmt.Println(err)

	if err == nil {
		fmt.Println(user.Name)
		return user, nil
	}
	return user, err
}

//用户数据列表
func GetUserList() (user []User) {
	o := orm.NewOrm()
	qs := o.QueryTable("users")
	var us []User
	cnt, err := qs.Filter("id__gt", 0).OrderBy("-id").Limit(10, 0).All(&us)
	if err == nil {
		fmt.Printf("count:", cnt)
	}
	return us
}

func QueryUserBySql(sql string, qarms []string) bool {
	o := orm.NewOrm()
	o.Raw(sql, qarms)

	return true
}

//分页数据
func UserLimitList(perPage, page int) (users []User) {
	o := orm.NewOrm()
	qs := o.QueryTable("users")

	var us []User
	cnt, err := qs.Limit(perPage, (page-1)*perPage).All(&us)
	if err == nil {
		fmt.Printf("count:", cnt)
	}
	return us
}

//获取用户总个数
func GetDataNum() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("users")

	var us []User
	num, err := qs.Filter("id__gt", 0).All(&us)
	if err != nil {
		return 0
	} else {
		return num
	}
}
