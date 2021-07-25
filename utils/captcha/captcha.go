package captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"go-admin/global"
)

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
