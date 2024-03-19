package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load()
	if err!=nil{
		log.Fatal("Env is not loaded")
	}
}