package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("./"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&Article{}); err != nil {
		log.Println(err)
	}

	return db, err
}
