package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handleform(c *gin.Context) {
	var user struct {
		name     string
		password string
	}
	if err := c.BindJSON(&user); err != nil {
		log.Fatal("There is some issue")
	}
	fmt.Println(c.PostForm("name"))
	c.String(http.StatusOK, "Weeldone")
}
