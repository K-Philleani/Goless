package main

import (
	db "Goless/go_mysql/database"
)
func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":9999")
}
