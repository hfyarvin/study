package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

func main() {
}

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, "postgres", "12345678", "postgres")
	var err error
	engine, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	// engine.ShowSQL()
	err = engine.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

func addAccountTest() {
	account := new(Account)
	account.Balance = 100.99
	account.Version = 1
	account.Name = "arvin"
	account = account.Create()
	fmt.Println(account.Id)
}

type Account struct {
	Id      int64     `json:"id"`
	Name    string    `xorm:"name VARCHAR(255)"    json:"name"`
	Balance float64   `xorm:"balance not null DOUBLE(10)"    json:"balance"`
	Version int       `xorm:"version default 1 INT(11)"    json:"version"`
	Created time.Time `xorm:"created not null DATETIME"    json:"created"`
	Updated time.Time `xorm:"updated not null DATETIME"    json:"updated"`
}

func (self *Account) TableName() string {
	return "accounts"
}

//创建数据表
func (self *Account) CreateTable() error {
	err := engine.CreateTables(self)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (self *Account) Create() *Account {
	_, err := engine.InsertOne(self)
	if err != nil {
		return nil
	}

	return self
}

//通过ID获取Article
func GetArticleById(id int64) *Account {
	item := new(Account)
	// has, err := engine.Where("id = ?", id).Get(item)
	has, _ := engine.Id(id).Get(item)

	if !has {
		return nil
	}
	return item
}

//ALL
func SelectAll() []*Account {
	var accs []*Account
	engine.SQL(fmt.Sprintf("select * from %s", new(Account).TableName())).Find(&accs)
	return accs
}

//利用sql删除
func DeleteUserBySQL(name string) bool {
	result, err := engine.Exec(fmt.Sprintf("delete from %s where name = '%s'", new(Account).TableName(), name))
	if err != nil {
		log.Println(err)
		return false
	}
	rows, err := result.RowsAffected()
	if err == nil && rows > 0 {
		fmt.Println("rows:", rows)
		return true
	}
	return false
}

func (self *Account) Update() {
	engine.Id(self.Id).Update(self)
}

/**
 *
 */
