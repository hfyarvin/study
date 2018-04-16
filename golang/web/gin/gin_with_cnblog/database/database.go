package database

import (
	"log"

	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:arvin171017@tcp(119.23.246.245:3306)/arvinDB?parseTime=true")

	if err != nil {
		log.Fatal(err.Error())
	}

	//测试连接
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
