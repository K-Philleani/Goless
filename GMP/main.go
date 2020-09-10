package main

import (
	db "Goless/GMP/database"
)
func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":9999")
}
