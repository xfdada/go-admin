package model

type VideoUser struct {
	ID          int64  `json:"id" gorm:"primaryKey;type:bigint;autoIncrement"`
	Uuid        string `json:"uuid" gorm:"type:varchar(12);unique"`
	NickName    string `json:"nick_name" gorm:"type:varchar(30)"`
	HeadIcon    string `json:"head_icon" gorm:"type:varchar(100)"`
	Phone       string `json:"phone" gorm:"type:varchar(15)"`
	Email       string `json:"email" gorm:"type:varchar(50)"`
	Password    string `json:"-" gorm:"type:varchar(200)"`
	Description string `json:"description" gorm:"type:varchar(200)"`
}
