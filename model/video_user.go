package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/global/page"
	"go-admin/utils"
	"go-admin/utils/loggers"
	"go-admin/utils/sendmail"
	"go-admin/utils/uuid"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type VideoUser struct {
	ID          int64          `json:"id" gorm:"primaryKey;type:bigint;autoIncrement"`
	Uuid        string         `json:"uuid" gorm:"type:varchar(20);unique"`
	NickName    string         `json:"nick_name" gorm:"type:varchar(30)"`
	HeadIcon    string         `json:"head_icon" gorm:"type:varchar(100)"`
	Phone       string         `json:"phone" gorm:"type:varchar(15)"`
	Email       string         `json:"email" gorm:"type:varchar(50)"`
	Password    string         `json:"-" gorm:"type:varchar(200)"`
	Description string         `json:"description" gorm:"type:varchar(200)"`
	SingIp      string         `json:"sing_ip" gorm:"type:varchar(20)"`
	CreatedAt   *LocalTime     `json:"created_at" gorm:"type:datetime;column:sing_time;"`
	IsLive      uint8          `json:"is_live" gorm:"type:tinyint(1)"`
	UpdateAt    *LocalTime     `json:"update_at" gorm:"type:datetime;column:sing_up;"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
type VideoUserList struct {
	page.Page
	Data []VideoUser
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
	uuids := uuid.GetUuid()
	res := db.Model(v).Create(&VideoUser{
		Uuid:        uuids,
		NickName:    param["nickName"].(string),
		HeadIcon:    param["headIcon"].(string),
		Phone:       param["phone"].(string),
		Email:       param["email"].(string),
		Password:    utils.HashAndSalt(param["password"].(string)),
		Description: param["desc"].(string),
		IsLive:      0,
	})
	if res.Error != nil {
		return res.Error
	}
	go sendmail.SendMail(param["email"].(string), uuids)
	return nil
}

func (v *VideoUser) List(c *gin.Context) (*VideoUserList, error) {
	var list []VideoUser
	nowPage, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	res := db.Scopes(Paginate(nowPage, pageSize)).Find(&list)
	var total int64
	db.Model(v).Count(&total)
	if res.Error != nil {
		return nil, res.Error
	}
	return &VideoUserList{
		Page: page.Page{
			PageSize: int64(pageSize),
			Total:    total,
			NowPage:  nowPage,
			Pages:    math.Ceil(float64(total) / float64(pageSize)),
		},
		Data: list,
	}, nil
}
func (v *VideoUser) OnLive(uuids string) error {
	res := db.Model(v).Where("uuid = ?", uuids).Update("is_live", 1)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (v *VideoUser) Login(userName, password string) bool {
	res := db.Model(v).Where("email = ? and is_live = 1", userName).First(v)
	if res.Error != nil {
		loggers.Logs(fmt.Sprintf("查询用户失败 错误详情是%v", res.Error))
		return false
	}
	ok := utils.ComparePasswords(v.Password, password)
	if !ok {
		return false
	}
	return true
}
