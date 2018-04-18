package models

import (
	"fmt"
	"log"

	// db "../database"
	db "../database"
)

// var SqlDB *sql.DB
const (
	Person_TableName = "person"
)

type Person struct {
	Id        int    `json:"id"         form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name"  form:"last_name"`
}

func (this *Person) TableName() string {
	return Person_TableName
}

func (p *Person) AddPerson() bool {
	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		log.Println("err1:", err.Error())
		return false
	}

	id, err := rs.LastInsertId()
	fmt.Println(id)

	if err != nil {
		log.Println("err2:", err.Error())
		return false
	} else {
		return true
	}
}

func (p *Person) EditPerson() bool {
	rs, err := db.SqlDB.Exec("UPDATE person set first_name=?, last_name=? where id=?", p.FirstName, p.LastName, p.Id)
	if err != nil {
		return false
	}

	//影响行数
	id, err := rs.RowsAffected()
	fmt.Println(id)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 删除
func DeletePerson(pId int) bool {
	rs, err :=db.SqlDB.Exec("Delete From person where id=? ", pId)
	if err != nil {
		return false
	}
	id, err := rs.RowsAffected()
	fmt.Println(id)
	if err != nil {
		return false
	} else {
		return true
	}
}

//得到记录列表
func GetPersonList(pageno, pagesize int, search string) (persons []Person) {
	fmt.Println("搜索参数") //姓名
	persons = make([]Person, 0)
	sql := ""
	if search != "" {
		sql = fmt.Sprintf("SELECT id, first_name, last_name FROM person where last_name like '%%%s%%' or first_name like '%%%s%%' limit %d,%d", search, search, (pageno-1)*pagesize, pagesize)
	} else {
		sql = fmt.Sprintf("SELECT id, first_name, last_name FROM person limit %d,%d", (pageno-1)*pagesize, pagesize)
	}

	rows, err := db.SqlDB.Query(sql)
	if err != nil {
		return nil
	}
	defer rows.Close()

	//数据添加到数据集中
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}

	if err = rows.Err(); err != nil {
		return nil
	}
	return persons
}

//记录数
func GetRecordNum(search string) int {
	num := 0
	sql := fmt.Sprintf("SELECT id, first_name, last_name From person where last_name like '%%%s%%' or first_name like '%%%s%%'", search, search)
	rows, err := db.SqlDB.Query(sql)
	if err != nil {
		return 0
	}
	defer rows.Close()

	for rows.Next() {
		num ++ 
	}

	return num
}

func GetPersonById(id int) (p *Person) {
	var person Person

	err := db.SqlDB.QueryRow("SELECT id, first_name, last_name FROM person WHERE id = ?", id).Scan(&person.Id, &person.FirstName, &person.LastName)

	if err != nil {
		log.Println("Get Person By Id Error: ", err.Error())
	}

	return &person
}
