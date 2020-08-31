package main
import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type student struct {
	id int
	name string
	class string
	sex string
}

// 连接数据库
func initDB() (err error){
	dsn := "root:00000000@tcp(127.0.0.1:3306)/gouse"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

// 单行数据查询
func queryRow(id int) {
	var stu student
	sqlStr := `select id, name, class, sex from student where id=?`
	err := db.QueryRow(sqlStr, id).Scan(&stu.id, &stu.name, &stu.class, &stu.sex)
	if err != nil {
		log.Println("查询失败, err:", err)
		return
	}
	fmt.Printf("id:%d, name: %s, class:%s, sex:%s\n", stu.id, stu.name, stu.class, stu.sex)
}

// 多行数据查询
func queryMultiRows(id int) {
	sqlStr := `select id, name, class, sex from student where id> ?`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		log.Println("数据查询失败, err:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var stu student
		err := rows.Scan(&stu.id, &stu.name, &stu.class, &stu.sex)
		if err != nil {
			log.Println("结果返回失败,err:", err)
			return
		}
		fmt.Printf("id:%d, name:%s, class:%s, sex:%s\n", stu.id, stu.name, stu.class, stu.sex)
	}
}

// 插入数据
func insertRow(name, class, sex string) {
	sqlStr := `insert into student (name, class, sex) values(?, ?, ?)`
	row, err := db.Exec(sqlStr, name, class, sex)
	if err != nil {
		log.Println("数据插入失败, err:", err)
		return
	}
	count, err := row.LastInsertId()
	if err != nil {
		log.Println("数据返回结果失败, err", err)
		return
	}
	fmt.Printf("插入数据id:%d\n", count)
}

func main() {
	err := initDB()
	if err != nil {
		log.Println("数据库连接失败")
		return
	}
	fmt.Println("数据库连接成功")
	//queryRow(1)
	insertRow("刘环宇", "计算机一班", "男")
	queryMultiRows(0)
}
