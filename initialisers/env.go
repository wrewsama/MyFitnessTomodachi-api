package initialisers

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}