package config

import (
	"log"
	"github.com/jinzhu/gorm"
	"emad.com/models"
	_"github.com/jinzhu/gorm/dialects/mysql"
        )

var DB *gorm.DB

func ConnectDB()  {
	db, err := gorm.Open("mysql","root:@/go_fiber?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		log.Fatalf("Could not connect to database: %s", err)
	}

	db.AutoMigrate(
		&models.Users{}, 
		models.Auths{},
	)
	DB = db
	
}