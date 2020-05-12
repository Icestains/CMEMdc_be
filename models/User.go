package models

import (
	"errors"
	"fmt"
	"time"
)

// 基本模型的定义
type BasicInfo struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	BasicInfo
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增
	Name     string `json:"name" binding:"required" gorm:"size:20;not null;unique"`
	Password string `json:"password" binding:"required" gorm:"size:16;not null"`
	Email    string `json:"email" gorm:"size:50;"`

	Permission string
}

type UserAuth struct {
	UserID   uint   `gorm:"primary_key"`
	Name     string `json:"name" binding:"required" gorm:"size:20;not null;unique"`
	Password string `json:"password" binding:"required" gorm:"size:16;not null"`
}

//查找用户名
func FindUserByName(name, password string) (bool, error) {
	var res User
	if err := db.Where("name = ?", name).Find(&res).Error; err != nil {
		return false, err
	} else {
		if password == "" {
			return true, errors.New("password required")
		} else if res.Password == password {
			return true, nil
		} else {
			return true, errors.New("wrong password")
		}
	}
}

//注册
func Create(regInfo *User) error {
	db.NewRecord(regInfo) // => 主键为空返回`true`
	if !db.HasTable(&UserAuth{}) {
		db.CreateTable(&UserAuth{})
	}
	err := db.Create(&regInfo).Error
	fmt.Println(db.Create(&UserAuth{Name: regInfo.Name, Password: regInfo.Password}).Value)
	return err
}

//
////注销？
//func Delete(id uint) error {
//	var user model.User
//	return h.db.Where("ID = ?", id).Delete(&user).Error
//}
//
////修改用户信息 用户名 密码
//func Update(newUser *model.User) error {
//	var userinfo model.User
//	h.db.Where("uid = ?", newUser.ID).Find(&userinfo)
//	return h.db.Model(&userinfo).Updates(model.User{
//		Name:     newUser.Name,
//		Password: newUser.Password,
//	}).Error
//}
//


func FindUserInfo(name string) (User, error) {
	var user User
	//err := db.Where("name = ?", name).Find(&user).Error
	err := db.Raw("select * from users where name = ?", name).Scan(&user).Error
	return user, err
}
