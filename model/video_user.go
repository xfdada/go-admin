package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/global/page"
	"go-admin/utils"
	"go-admin/utils/loggers"
	"go-admin/utils/uuid"
	"gopkg.in/gomail.v2"
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
		Password:    utils.EncodeMD5(param["password"].(string)),
		Description: param["desc"].(string),
		IsLive:      0,
	})
	if res.Error != nil {
		return res.Error
	}
	go sendMail(param["email"].(string), uuids)
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

func sendMail(email, uuids string) {
	mail := []string{email}
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(global.Email.User, "祥富官方影院")) //此方法可以取别名
	m.SetHeader("To", mail...)
	m.SetHeader("Subject", "账号激活邮件")
	body := `
<h1 style="text-align:center;font-size:18px">祥富影院</h1>
<p> 欢迎注册本影院，为了更好的服务您，方便您找回密码，请点击此链接激活账号</p>
<p> <a href="http://127.0.0.1:8001/api/check/` + uuids + `">激活</a></p>
`
	m.SetBody("text/html", body)
	d := gomail.NewDialer(global.Email.Host, global.Email.Port, global.Email.User, global.Email.Pwd)
	err := d.DialAndSend(m)
	if err != nil {
		loggers.Logs(fmt.Sprintf("发送邮件给%s发生错误，错误详情是：err:%v\n", email, err))
	}
}
