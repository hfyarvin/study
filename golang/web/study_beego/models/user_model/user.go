package user_model

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	Users_TableName = "users"
)

type Users struct {
	Id           int64     `orm:"id" json:"id"`
	Uuid         string    `orm:"-" json:"-"`
	FirstName    string    `orm:"first_name VARCHAR(50) NOTNULL"         json:"first_name"`
	LastName     string    `orm:"last_name VARCHAR(50) NOTNULL"          json:"last_name"`
	Password     string    `orm:"-"                                      json:"password"`
	PasswordHash string    `orm:"password_hash VARCHAR(200) NOTNULL"     json:"password_hash"`
	Email        string    `orm:"email VARCHAR(100) NOTNULL"             json:"email"`
	Phone        string    `orm:"phone VARCHAR(30) NOTNULL"              json:"phone"`
	Company      string    `orm:"company VARCHAR(200)"                   json:"company"`
	Address      string    `orm:"address VARCHAR(400)"                   json:"address"`
	City         string    `orm:"city VARCHAR(100)"                      json:"city"`
	Regin        string    `orm:"regin VARCHAR(100)"                     json:"regin"`
	ZipCode      string    `orm:"zip_code VARCHAR(30)"                   json:"zip_code,omitempty"`
	Country      string    `orm:"country VARCHAR(100)"                   json:"country"`
	IdentityCard string    `orm:"identity_card VARCHAR(50)"              json:"identity_card"`
	IsActive     int       `orm:"is_active BIT(1)"                       json:"is_active"`
	Created      time.Time `orm:"created DATETIME(0) NOTNULL"            json:"created"`
	Updated      time.Time `orm:"updated DATETIME(0) NOTNULL"            json:"updated"`
}

func (this *Users) TableName() string {
	return Users_TableName
}
