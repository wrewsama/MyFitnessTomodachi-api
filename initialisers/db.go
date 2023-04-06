package initialisers 

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	DB = db
}
