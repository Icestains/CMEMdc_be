package handlers

import (
	"CMEMdc_be/Database/model"
	"github.com/jinzhu/gorm"
)

type EmqxJSHandler struct {
	db *gorm.DB
}

func NewEmqxJSHandler(db *gorm.DB) *EmqxJSHandler {
	return &EmqxJSHandler{db}
}

func (h *EmqxJSHandler) FindAll() (*[]model.EmqxJS, error) {
	var res []model.EmqxJS

	if err := h.db.Raw("SELECT msgid, topic, payload FROM mqtt_msg ").Scan(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
