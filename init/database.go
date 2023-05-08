package Init

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	//postgres://ianwnkgi:vq7io-P4E43aW9XcfItkWq3ujlD2LEgv@tiny.db.elephantsql.com/ianwnkgi
	dsn := "host=tiny.db.elephantsql.com user=ianwnkgi password=vq7io-P4E43aW9XcfItkWq3ujlD2LEgv dbname=ianwnkgi port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail. DataBase")
	}
}
