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

// 更新数据
func updateRow(sex string, id int) {
	sqlStr := `update student set sex =? where id = ?`
	row, err := db.Exec(sqlStr, sex, id)
	if err != nil {
		log.Println("更新数据失败, err:", err)
		return
	}

	count, err := row.RowsAffected()
	if err != nil {
		log.Println("数据结果失败, err:", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", count)
}

// 删除数据

func deleteRow(id int) {
	sqlStr := `delete from student where id =?`
	row, err := db.Exec(sqlStr, id)
	if err != nil {
		log.Println("删除数据失败, err:", err)
		return
	}
	count, err := row.RowsAffected()
	if err != nil {
		log.Println("数据结果失败, err:", err)
		return
	}
	fmt.Printf("删除了:%d行数据\n", count)
}

// 预处理
func prepareRow() {
	sqlStr := `insert into student (name, class, sex) values(?, ?, ?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println("预处理失败, err:", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("王木", "计算机一班", "男")
	if err != nil {
		log.Println("数据插入失败,err:", err)
		return
	}
	_, err = stmt.Exec("孙天宇", "计算机一班", "nan")
	if err != nil {
		log.Println("数据插入失败, err:", err)
		return
	}
}


func main() {
	err := initDB()
	if err != nil {
		log.Println("数据库连接失败")
		return
	}
	fmt.Println("数据库连接成功")
	//queryRow(1)
	//insertRow("刘环宇", "计算机一班", "男")
	updateRow("nan", 7)
	//deleteRow(2)
	//prepareRow()
	queryMultiRows(0)
}
