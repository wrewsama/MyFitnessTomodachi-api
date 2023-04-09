package initialisers

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}