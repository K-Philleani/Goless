package model

import db "Goless/GMP/database"

type Person struct {
	Id    		int 	`json:"id" form:"id"`
	FirstName 	string 	`json:"first_name" form:"first_name"`
	LastName 	string 	`json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	sqlStr := "insert into person (first_name, last_name) values(?, ?)"
	rs, err := db.SqlDB.Exec(sqlStr, p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	if err != nil {
		return
	}
	return
}

