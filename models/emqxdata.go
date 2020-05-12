package models

import (
	"database/sql/driver"
	"encoding/json"
	"log"
)

type EmqxJS struct {
	Msgid   string       `json:"msgid"`
	Topic   string       `json:"topic"`
	Payload *EmqxPayload `json:"payload"`
}

type EmqxPayload struct {
	Time       int    `json:"time"`
	Msg        string `json:"msg"`
	RandomData int    `json:"randomData"`
}

// https://github.com/jinzhu/gorm/issues/2195
// Scan 实现 gorm Scanner 接口
func (ls *EmqxPayload) Scan(value interface{}) error {
	if value == nil {
		*ls = EmqxPayload{}
		return nil
	}
	t := EmqxPayload{}
	if e := json.Unmarshal(value.([]byte), &t); e != nil {
		return e
	}
	*ls = t
	return nil
}

// Value 实现 gorm Valuer 接口
func (ls *EmqxPayload) Value() (driver.Value, error) {
	if ls == nil {
		return nil, nil
	}
	b, e := json.Marshal(*ls)
	return b, e
}

func FindAllEmqxData() (res []EmqxJS) {
	if err := db.Raw("SELECT msgid, topic, payload FROM mqtt_msg").Scan(&res).Error; err != nil {
		log.Println(err.Error())
	}
	return
}
