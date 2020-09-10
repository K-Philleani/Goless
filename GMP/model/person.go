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

func (p *Person) GetPerson() (persons []Person, err error){
	sqlStr := "select id, first_name, last_name from person"
	rows, err := db.SqlDB.Query(sqlStr)
	if err != nil {
		return
	}
	defer rows.Close()
	persons = make([]Person, 0)
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

