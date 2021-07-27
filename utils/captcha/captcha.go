package captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"go-admin/global"
	"go-admin/utils/loggers"
	"go-admin/utils/redis"
	"time"
)

var Captcha *base64Captcha.Store

func init() {

}

type Store struct {
	Expiration time.Duration
	PreKey     string
}

func NewDefaultRedisStore() base64Captcha.Store {
	return &Store{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

func (rs *Store) Set(id string, value string) error {
	err := redis.Set(rs.PreKey+id, value, rs.Expiration)
	if err != nil {
		loggers.Logs(fmt.Sprintf("RedisStoreGetError! err:%v", err))
		return err
	}
	return nil

}
func (rs *Store) Get(key string, clear bool) string {
	val, err := redis.Get(key)
	if err != nil {
		loggers.Logs(fmt.Sprintf("RedisStoreGetError! err:%v", err))
		return ""
	}
	if clear {
		err := redis.Del(key)
		if err != nil {
			loggers.Logs(fmt.Sprintf("RedisStoreGetError! err:%v", err))
			return ""
		}
	}
	return val
}

func (rs *Store) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}

func GetCaptcha() (string, string) {
	driver := base64Captcha.NewDriverDigit(global.Captcha.Height, global.Captcha.Width, global.Captcha.Length, global.Captcha.MaxSkew, global.Captcha.DotCount)
	// 生成base64图片
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)

	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return "", ""
	}
	return id, b64s
}
