### 安装的库

1.gorm  
2.mysql驱动  
3.viper   配置文件读取  
```
    新增配置项需要在config文件夹中添加相对于的结构体
    初始化读取在global文件夹下的 setting.go文件中
    支持配置热更新
```

4.swag    api文档
```
    安装Swagger工具
    go get -u github.com/swaggo/swag/cmd/swag
    go get -u github.com/swaggo/files
    go get -u github.com/alecthomas/template
```

5.jwt     token鉴权     
```
    "go-admin/utils/jwts"
    生成token  GenerateToken(appkey,appsecret string) token,error
    验证token有效性  ParseToken(token string) *Claims, error
    error包含错误信息，nil时验证通过 
    error.(*jwt.ValidationError).Errors
    jwt.ValidationErrorExpired 错误为 token超过有效期
    其他为token错误
```
6.base64Captcha  验证码生成
```
    通过配置文件可选内存存储，redis存储验证码，
    "go-admin/utils/captcha"包
     生成base64字符串验证码      captcha.GetCaptcha() (id，base64 string)
     检测验证码是否正确       captcha.Verify(id, val) bool
```
7.gomail邮件发送
```
    "go-admin/utils/sendmail"
    aliasName  别名 将邮件地址取成其他名字
    title       发送邮件的主题
    content     发送邮件的具体内容
    mail        要发送的邮箱地址，可以是多个
    Send(aliasName,title,content string ,mail []string)
```
