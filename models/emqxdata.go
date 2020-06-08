package models

import (
	"CMEMdc_be/utils/logging"
	"database/sql/driver"
	"encoding/json"
	"sort"
)

type MqttMsg struct {
	ID      int          `json:"id"`
	Msgid   string       `json:"msgid"`
	Sender  string       `json:"sender"`
	Topic   string       `json:"topic"`
	Qos     int          `json:"qos"`
	Retain  int          `json:"retain"`
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
	//string 不能默认看成 []byte ，只好先声明成 string 再转化成 json了。。。
	if e := json.Unmarshal([]byte(value.(string)), &t); e != nil {
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

func FindAllEmqxData() (res []MqttMsg) {

	err := db.Find(&res).Error
	//err := db.Raw("SELECT msgid, topic, payload FROM mqtt_msg").Scan(&res).Error
	if err != nil {
		logging.Error(err.Error())
	}
	return
}

type MagSlice []MqttMsg

func (p MagSlice) Len() int           { return len(p) }
func (p MagSlice) Less(i, j int) bool { return p[i].ID < p[j].ID }
func (p MagSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func FindEmqxDataBySender(Sender string) (res MagSlice) {
	err := db.Order("id desc").Where("sender = ?", Sender).Find(&res).Error
	if err != nil {
		logging.Error(err.Error())
	}
	sort.Sort(&res)
	return
}
func FindEmqxDataByTopic(topic string) (res []MqttMsg) {
	err := db.Where("topic = ?", topic).Find(&res).Error
	if err != nil {
		logging.Error(err.Error())
	}
	return
}
