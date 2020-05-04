package Postgresql

import (
	"CMEMdc_be/Database/model"
	"CMEMdc_be/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//		"user=postgres password=0852 dbname=postgres sslmode=disable",
// ConnectSQL ...
//链接数据库
func NewDatabase(cfg *config.DatabaseConfig) (*gorm.DB, error) {

	dbSource := "user=" + cfg.User + " password=" + cfg.Password + " dbname=" + cfg.DbName + " sslmode=disable"
	DB, err := gorm.Open("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	if err := InitDatabase(DB); err != nil {
		return nil, err
	}
	return DB, nil

}

//初始化数据库
func InitDatabase(db *gorm.DB) error {

	//db.LogMode(config.Debug) // auto migrate
	models := []interface{}{
		&model.Userinfo{},
		&model.User{},
	}
	if err := db.AutoMigrate(models...).Error; err != nil {
		return err
	}

	return nil

}
