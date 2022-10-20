package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/test_enigma"))
	dsn := "host=localhost user=postgres password=kukang12 dbname=aset_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Aset{}, &User{})
	DB = db

}
