package controller

import (
	initializers "basic/Initializers"
	"basic/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var data models.Person
	if err := c.BindJSON(&data); err != nil {
		log.Println(err.Error())
	}
	var user models.Person
	initializers.DB.First(&user, "username = ?", data.Username)
	fmt.Println(user.ID)
	if user.ID != 0 {
		c.Status(http.StatusContinue)
		fmt.Println(http.StatusContinue)
		return
	}
	hashedpass,err:=bcrypt.GenerateFromPassword([]byte(data.Password),bcrypt.DefaultCost)
	if err!=nil{
		log.Fatal("Error in hashing password")
	}
	data.Password=string(hashedpass)
	result := initializers.DB.Create(&data)
	fmt.Println(result)
	fmt.Println(http.StatusAccepted)
	c.Status(http.StatusAccepted)
}
