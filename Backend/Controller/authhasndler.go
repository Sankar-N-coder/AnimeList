package controller

import (
	initializers "basic/Initializers"
	"basic/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var data  models.Person
	if err := c.BindJSON(&data); err != nil {
		log.Println(err.Error())
	}
	result:=initializers.DB.Create(&data)
	fmt.Println(result)
}
