package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func Init() {
	dsn := "root:00000000@tcp(172.0.0.1:3306)/ginhello"
	DB, err := sql.Open("mysql", dsn)
	defer DB.Close()
	if err != nil {
		log.Println("err: ", err)
		return
	}
	log.Println("数据库连接成功")
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
