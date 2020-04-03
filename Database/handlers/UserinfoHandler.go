package handlers

import (
	"CMEMdc_be/Database/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserinfoHandler struct {
	db *gorm.DB
}

func NewUserinfoHandler(db *gorm.DB) *UserinfoHandler {
	return &UserinfoHandler{db}
}

func (h *UserinfoHandler) FindAll() (*[]model.Userinfo, error) {
	var res []model.Userinfo
	if err := h.db.Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (h *UserinfoHandler) Create(regInfo *model.Userinfo) error {
	return h.db.Create(regInfo).Error
}

func (h *UserinfoHandler) Delete(uid int64) error {
	var users model.Userinfo
	h.db.Where("uid = ?", uid).Find(&users)
	fmt.Println(users)
	//return err

	return h.db.Delete(&users).Error
	//return
}

func (h *UserinfoHandler) Update(newUserInfo *model.Userinfo) error {
	var userinfo model.Userinfo
	h.db.Where("uid = ?", newUserInfo.Uid).Find(&userinfo)
	return h.db.Model(&userinfo).Updates(model.Userinfo{
		Username:   newUserInfo.Username,
		Department: newUserInfo.Department,
	}).Error
}
