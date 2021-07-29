package model

import (
	"fmt"
	"go-admin/global"
	"go-admin/utils/loggers"
	"strconv"
)

type User struct {
	*Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func NewUser() User {
	return User{}
}
func (u User) TableName() string {
	return "blog_user"
}

func (u User) Get(id string) (User, int) {
	db := global.DB
	res := db.Where("id = ?", id).First(&u)
	if res.Error != nil {
		loggers.Logs(fmt.Sprint("查询失败", id, "Details:", res.Error))
		return User{}, 0
	}
	if res.RowsAffected <= 0 {
		return User{}, 0
	}
	return u, 1
}

func (u User) List() ([]User, error) {
	var user []User
	db := global.DB
	res := db.Find(&user)
	if res.Error != nil {
		loggers.Logs(fmt.Sprint("查询失败", "Details:", res.Error))
		return []User{}, res.Error
	}
	return user, nil
}

func (u User) Add(name, address, age string) error {
	db := global.DB
	ages, _ := strconv.Atoi(age)
	err := db.Create(&User{Name: name, Age: ages, Address: address}).Error
	if err != nil {
		loggers.Logs(fmt.Sprint("插入数据失败", "Details:", err))
		return err
	}
	return nil
}

func (u User) Update(id, name, address, age string) error {
	db := global.DB
	ages, _ := strconv.Atoi(age)
	err := db.Model(&u).Where("id = ?", id).Updates(User{Name: name, Age: ages, Address: address}).Error
	if err != nil {
		loggers.Logs(fmt.Sprint("更新数据失败", "Details:", err))
		return err
	}
	return nil
}
func (u User) Delete(id string) (int64, error) {
	db := global.DB
	res := db.Delete(&u, id)
	if res.Error != nil {
		loggers.Logs(fmt.Sprint("删除数据失败", "Details:", res.Error))
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
