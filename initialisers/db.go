package initialisers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	sqluser := os.Getenv("SQLUSER")
	password := os.Getenv("PASSWORD")
	hostname := os.Getenv("HOSTNAME")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")

	dsn :=  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
                        sqluser,
						password,
						hostname,
						dbport,
						dbname)
	log.Printf("Connecting to %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	DB = db
}
