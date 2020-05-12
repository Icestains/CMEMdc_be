package models

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth UserAuth
	db.Select("user_id").Where(UserAuth{Name : username, Password : password}).First(&auth)
	if auth.UserID > 0 {
		return true
	}

	return false
}