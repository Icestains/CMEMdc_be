package handlers

import (
	"CMEMdc_be/Database/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

//查找用户名
func (h *UserHandler) FindUserByName(name, password string) (bool, error) {
	var res model.User
	if err := h.db.Where("name = ?", name).Find(&res).Error; err != nil {
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
func (h *UserHandler) Create(regInfo *model.User) error {
	h.db.NewRecord(regInfo) // => 主键为空返回`true`

	return h.db.Create(&regInfo).Error
}

//注销？
func (h *UserHandler) Delete(id uint) error {
	var user model.User
	return h.db.Where("ID = ?", id).Delete(&user).Error
}

//修改用户信息 用户名 密码
func (h *UserHandler) Update(newUser *model.User) error {
	var userinfo model.User
	h.db.Where("uid = ?", newUser.ID).Find(&userinfo)
	return h.db.Model(&userinfo).Updates(model.User{
		Name:     newUser.Name,
		Password: newUser.Password,
	}).Error
}

func (h *UserHandler) FindUserEmail(name string) (model.User, error) {
	var user model.User
	err := h.db.Where("name = ?", name).Find(&user).Error
	return user, err
}

////登录
//func (h *UserHandler) Login(User *model.User) {
//	var userinfo model.User
//
//
//}
