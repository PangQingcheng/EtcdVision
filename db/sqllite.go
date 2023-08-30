package db

import (
	"etcd-vision/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.ETCDDataSource{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
