package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func main() {
	dsn := "root:00000000@tcp(127.0.0.1:3306)/gocase"
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		log.Println("数据库连接失败， err: ", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println("数据库连接失败， err: ", err)
		return
	}
	log.Println("数据库连接成功")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	router := gin.Default()
	router.POST("/person", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")

		sqlStr := "insert into person (first_name, last_name) values(?, ?)"
		rs, err := db.Exec(sqlStr, firstName, lastName)
		if err != nil {
			log.Println("数据插入失败, err: ", err)
			return
		}
		id, err := rs.LastInsertId()
		if err != nil {
			log.Println("数据获取结果失败, err: ", err)
			return
		}
		fmt.Println("insert person id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(200, gin.H{
			"message": msg,
			"id": id,
			"firstName": firstName,
			"lastName": lastName,
		})
	})

	router.Run(":8888")
}
