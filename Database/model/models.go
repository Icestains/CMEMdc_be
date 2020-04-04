package model

import "time"

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

//type Email struct {
//	UserID     int    `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
//	Email      string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
//	Subscribed bool
//}

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
