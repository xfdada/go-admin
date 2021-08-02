package model

import (
	"gorm.io/gorm"
)

type VideoUser struct {
	ID          int64          `json:"id" gorm:"primaryKey;type:bigint;autoIncrement"`
	Uuid        string         `json:"uuid" gorm:"type:varchar(12);unique"`
	NickName    string         `json:"nick_name" gorm:"type:varchar(30)"`
	HeadIcon    string         `json:"head_icon" gorm:"type:varchar(100)"`
	Phone       string         `json:"phone" gorm:"type:varchar(15)"`
	Email       string         `json:"email" gorm:"type:varchar(50)"`
	Password    string         `json:"-" gorm:"type:varchar(200)"`
	Description string         `json:"description" gorm:"type:varchar(200)"`
	SingIp      string         `json:"sing_ip" gorm:"type:varchar(20)"`
	SingTime    string         `json:"sing_time" gorm:"type:datetime"`
	IsLive      uint           `json:"is_live" gorm:"type:tinyint(1)"`
	SingUp      string         `json:"sing_up" gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `json:"-"gorm:"index"`
}

func NewVideoUser() *VideoUser {
	_ = db.AutoMigrate(&VideoUser{})
	return &VideoUser{}
}

func (v *VideoUser) TableName() string {
	return "video_users"
}
func (v *VideoUser) Get(id int) (*VideoUser, error) {
	res := db.Model(v).Find(v, id)
	if res.Error != nil {
		return &VideoUser{}, res.Error
	}
	return v, nil
}

func (v *VideoUser) Create(param map[string]interface{}) error {
	res := db.Model(v).Create(VideoUser{
		Uuid:        "",
		NickName:    param["nickName"].(string),
		HeadIcon:    param["nickName"].(string),
		Phone:       param["nickName"].(string),
		Email:       param["nickName"].(string),
		Password:    param["nickName"].(string),
		Description: param["nickName"].(string),
		IsLive:      0,
		SingUp:      "",
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
