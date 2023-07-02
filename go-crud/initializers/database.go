package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	//postgres://dnntippe:8x3-f6Cb5KFEWFScdLC056xGfMexhxRx@lallah.db.elephantsql.com/dnntippe
	// 5432 PostgreSQL default db portu
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

}
