package models

import "github.com/jinzhu/gorm"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

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
