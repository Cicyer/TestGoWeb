package database

import (
	"../table"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:ngs911@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func InsertHome(name string) error {
	home := &table.Home{
		Name:       name,
		CreateTime: time.Now(),
	}
	if err := db.Create(home).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func SelectByName(name string) (table.Home, error) {
	var home table.Home
	err := db.Model(&table.Home{}).Where("name = ?", name).Find(&home).Error
	return home, err
}
