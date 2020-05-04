package model

import (
	"database/sql/driver"
	"encoding/json"
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
	UserID   uint
	Name     string `json:"name" binding:"required" gorm:"size:20;not null;unique"`
	Password string `json:"password" binding:"required" gorm:"size:16;not null"`
}

type Userinfo struct {
	Uid        int64  `gorm:"primary_key"`
	Username   string `json:"username"`
	Department string `json:"department"`
	Created    string `json:"created"`
}

type UserRegister struct {
	Username   string `json:"username"`
	Department string `json:"department"`
	Created    string `json:"created"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
