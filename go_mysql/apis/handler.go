package apis

import (
	"Goless/go_mysql/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func AddPerson(c *gin.Context) {
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")

	p := model.Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Println(err)
		return
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(200, gin.H{
		"msg": msg,
	})
}

func GetPerson(c *gin.Context) {
	var p model.Person
	persons, err := p.GetPerson()
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"persons": persons,
	})
}

func GetPersonOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println(id)
		return
	}

	p := model.Person{Id: id}
	person, err := p.GetPersonOne()
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"person": person,
	})
}