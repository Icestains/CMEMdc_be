package models

import (
	"CMEMdc_be/utils/logging"
	"time"
)

type MqttClient struct {
	ID        int
	Clientid  string `json:"ClientId"`
	State     int
	Node      string
	OnlineAt  time.Time
	OfflineAt time.Time
	CreateAt  time.Time
}

func FindAllEmqxClients() (res []MqttClient) {

	if err := db.Find(&res).Error; err != nil {
		logging.Error(err.Error())
	}
	return
}
