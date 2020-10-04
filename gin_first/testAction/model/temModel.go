package model

import (
	"Goless/gin_first/testAction/initDB"
	"log"
)

type TemModel struct {
	Email 			string `form:"email" binding:"email"`
	Password 		string `form:"password"`
	PasswordAgain 	string `from:"password-again"`
}


func (user *TemModel) Save() int64{
	sqlStr := "insert into ginhello.user (email, password) values(?,?)"
	result, err := initDB.DB.Exec(sqlStr, user.Email, user.Password)
	if err != nil {
		log.Println("user insert error: ", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("user insert error: ", err)
	}
	return id
}