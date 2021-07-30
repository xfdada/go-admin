package model

import (
	"fmt"
	"go-admin/global"
	"go-admin/global/page"
	"go-admin/utils/loggers"
)

type Article struct {
	*Model
	Title      string `json:"title"`       //标题
	Desc       string `json:"desc"`        //简述
	Url        string `json:"url"`         //封面图片地址
	Content    string `json:"content"`     // 文章内容
	CreatedBy  string `json:"created_by"`  // 创建者
	ModifiedBy string `json:"modified_by"` //修改着
}
type ArticleList struct {
	*page.Page
	Data interface{}
}

func NewArticle() *Article {
	return &Article{}
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

func (a *Article) Create() error {

	return nil
}
