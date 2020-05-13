package models

import (
	"CMEMdc_be/utils/logging"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"CMEMdc_be/utils/setting"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                        error
		dbType, dbName, user, password, host, port string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		logging.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()

	//dbSource := "host=" + sec.Host + " port=" + sec.Port + " user=" + user + " password=" + password + " dbname=" + dbName + " sslmode=disable"
	//connStr := "postgres://postgres:0852@postgresql/5434?sslmode=disable"
	connStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		dbName)
	db, err = gorm.Open(dbType, connStr)

	if err != nil {
		log.Println(err)
	}

	//db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&User{}, &UserAuth{})
}

func CloseDB() {
	defer db.Close()
}
