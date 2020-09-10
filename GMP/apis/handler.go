package apis

import (
	"Goless/GMP/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
