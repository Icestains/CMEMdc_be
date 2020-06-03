package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"CMEMdc_be/utils/setting"
)

var db *gorm.DB

func Setup() {

	connStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.DbName)

	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, connStr)
	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{}, &UserAuth{})
}

func CloseDB() {
	defer db.Close()
}
