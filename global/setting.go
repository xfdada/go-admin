package global

import "go-admin/config"

var (
	Server  *config.Server
	Mysql   *config.Mysql
	Captcha *config.Captcha
	JWT     *config.JWT
	Upload  *config.Upload
)
