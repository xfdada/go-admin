package sendmail

import (
	"fmt"
	"go-admin/global"
	"go-admin/utils/loggers"
	"go-admin/utils/redis"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func SendMail(email, uuids string) {
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

func SendCaptcha(email string) error {
	rand.Seed(time.Now().UnixNano())
	capt := fmt.Sprintf("%06v", rand.Intn(1000000))
	err := redis.Set(email, capt, 600)
	if err != nil {
		loggers.Logs(fmt.Sprintf("发送邮件给%s 生成验证码失败，错误详情是：err:%v\n", email, err))
	}
	mail := []string{email}
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(global.Email.User, "祥富官方影院")) //此方法可以取别名
	m.SetHeader("To", mail...)
	m.SetHeader("Subject", "登录验证码")
	body := `
	<h1 style="text-align:center;font-size:18px">祥富影院</h1>
<p> 欢迎登录本影院，您的验证码是：<span style="color:lightblue;">` + capt + `</span></p>
`

	m.SetBody("text/html", body)
	d := gomail.NewDialer(global.Email.Host, global.Email.Port, global.Email.User, global.Email.Pwd)
	err = d.DialAndSend(m)
	if err != nil {
		loggers.Logs(fmt.Sprintf("发送邮件给%s发生错误，错误详情是：err:%v\n", email, err))
		return err
	}
	return nil
}
