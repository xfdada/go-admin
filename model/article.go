package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/global"
	"go-admin/global/page"
	"go-admin/utils/loggers"
	"math"
	"strconv"
)

var (
	db = global.DB
)

type Article struct {
	*Model
	Title      string `json:"title" gorm:"type:varchar(255)"`      //标题
	Desc       string `json:"desc" gorm:"type:varchar(500)"`       //简述
	Url        string `json:"url" gorm:"type:varchar(100)"`        //封面图片地址
	Content    string `json:"content"`                             // 文章内容
	CreatedBy  string `json:"created_by" gorm:"type:varchar(20)"`  // 创建者
	ModifiedBy string `json:"modified_by" gorm:"type:varchar(20)"` //修改着
}
type ArticleList struct {
	*page.Page
	Data []Article `json:"data"`
}

func NewArticle() *Article {
	ok := db.Migrator().HasTable(&Article{})
	if !ok {
		_ = db.AutoMigrate(&Article{})
		_ = db.Set("ENGINE", "InnoDB").AutoMigrate(&User{})
	}
	return &Article{}
}

func (a *Article) TableName() string {
	return "blog_articles"
}

func (a *Article) Get(id int) (*Article, error) {
	db := global.DB
	err := db.Model(a).Find(a, id).Error
	if err != nil {
		loggers.Logs(fmt.Sprintf("文章%v查询失败 detaile：%v\n", id, err))
		return nil, err
	}
	return a, nil
}

func (a *Article) List(c *gin.Context) (*ArticleList, error) {
	var article []Article
	nowPage, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	res := db.Scopes(Paginate(nowPage, pageSize)).Find(&article)
	var total int64
	db.Model(a).Count(&total)
	if res.Error != nil {
		loggers.Logs(fmt.Sprint("查询失败", "Details:", res.Error))
		return &ArticleList{}, res.Error
	}
	data := ArticleList{
		Page: &page.Page{
			PageSize: int64(pageSize),
			Total:    total,
			NowPage:  nowPage,
			Pages:    math.Ceil(float64(total) / float64(pageSize)),
		},
		Data: article,
	}
	return &data, nil
}

func (a *Article) Create(title, desc, url, Content, CreatedBy string) error {
	err := db.Model(a).Create(&Article{
		Title:     title,
		Desc:      desc,
		Url:       url,
		Content:   Content,
		CreatedBy: CreatedBy,
	}).Error
	if err != nil {
		loggers.Logs(fmt.Sprintf("文章插入失败 错误详情：%v\n", err))
		return err
	}
	return nil
}
func (a *Article) Update(id int, title, desc, url, Content, ModifiedBy string) error {
	err := db.Model(a).Where("id = ?", id).Updates(&Article{
		Title:      title,
		Desc:       desc,
		Url:        url,
		Content:    Content,
		ModifiedBy: ModifiedBy,
	}).Error
	if err != nil {
		loggers.Logs(fmt.Sprintf("文章更新失败 错误详情：%v\n", err))
		return err
	}
	return nil
}
func (a *Article) Delete(id int) error {
	db := global.DB
	err := db.Model(a).Delete(a, id).Error
	if err != nil {
		loggers.Logs(fmt.Sprintf("删除文章失败 原因是err:%v\n", err))
		return err
	}
	return nil
}
