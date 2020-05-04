package Database

import (
	"CMEMdc_be/Database/handlers"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Userinfo *handlers.UserinfoHandler
	User     *handlers.UserHandler
	EmqxJS   *handlers.EmqxJSHandler
}

//创建一个新的服务？
func NewService(db *gorm.DB) Service {
	return Service{
		Userinfo: handlers.NewUserinfoHandler(db),
		User:     handlers.NewUserHandler(db),
		EmqxJS:   handlers.NewEmqxJSHandler(db),
	}
}
