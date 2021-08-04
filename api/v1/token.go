package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/global/errcode"
	"go-admin/global/response"
	"go-admin/utils/captcha"
	"go-admin/utils/jwts"
	"go-admin/utils/sendmail"
)

//@Tags 固定接口
//@Summary 获取token
//@Produce json
//@Success 200   "成功"
//@Failure 400 {object} errcode.Error "未找到"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /api/get_token [get]
func GetToken(c *gin.Context) {
	token, err := jwts.GenerateToken("xfdada", "go-admin") //这里可以自定义从数据库中取用户的账号和密码生成
	if err != nil {
		c.JSON(200, gin.H{"msg": "failed", "err": fmt.Sprintf("GenerateToken err:%v", err)})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "token": token})
}

func GetCapt(c *gin.Context) {
	id, capt := captcha.GetCaptcha()
	if id != "" {
		c.JSON(200, gin.H{"code": 200, "id": id, "capt": capt})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code": 110, "msg": "获取验证码失败"})
}
func Verify(c *gin.Context) {
	res := response.NewResponse(c)
	id := c.PostForm("cid")
	val := c.PostForm("captcha")
	if id == "" || val == "" {
		res.ToError(errcode.ParamsError)
		return
	}
	// 同时在内存清理掉这个图片
	ok := captcha.Verify(id, val)
	if !ok {
		res.ToError(errcode.CodeError)
		return
	}
	res.ToResponse(errcode.Success)
}

//发送验证码

func SendEmail(c *gin.Context) {
	res := response.NewResponse(c)
	username := c.PostForm("email")
	ok := sendmail.SendCaptcha(username)
	if ok != nil {
		res.ToError(errcode.ServerError)
		return
	}
	res.ToResponse(errcode.Success)
}
