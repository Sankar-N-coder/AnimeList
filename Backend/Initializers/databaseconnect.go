package initializers

import (
	"basic/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	Loadenv()
}

var DB *gorm.DB

func Connectdatabase() {
	var err error
	dsn := os.Getenv("DB_CON")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error in Connecting")
	}
	fmt.Println(DB,"Success")
	DB.AutoMigrate(&models.Person{})
}
