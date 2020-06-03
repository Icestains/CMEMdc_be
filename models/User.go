package models

import (
	"CMEMdc_be/utils/logging"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
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
	Num        int    `gorm:"AUTO_INCREMENT"` // 自增
	Name       string `json:"name" binding:"required" gorm:"size:20;not null;unique"`
	Password   string `json:"password" binding:"required" gorm:"size:16;not null"`
	Email      string `json:"email" gorm:"size:50;"`
	Permission string `json:"permission"`
}

type UserAuth struct {
	UserID   uint   `gorm:"primary_key"`
	Name     string `json:"name" binding:"required" gorm:"size:20;not null;unique"`
	Password string `json:"password" binding:"required" gorm:"size:16;not null"`
}

//注册
func Create(regInfo *User) error {
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	db.NewRecord(regInfo)
	if !db.HasTable(&UserAuth{}) {
		db.CreateTable(&UserAuth{})
	}
	var err error
	err = db.Create(&regInfo).Error
	if err != nil {
		logging.Error(err)
		return err
	}
	err = db.Create(&UserAuth{Name: regInfo.Name, Password: regInfo.Password}).Error
	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
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

// 根据用户名查找用户信息
func FindUserInfo(name string) (User, error) {
	var user User
	err := db.Where("name = ?", name).Find(&user).Error
	return user, err
}

// 验证用户信息
func CheckAuth(username, password string) (bool, error) {
	var auth UserAuth
	err := db.Select("user_id").Where(UserAuth{Name: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if auth.UserID > 0 {
		return true, nil
	}
	return false, nil
}
