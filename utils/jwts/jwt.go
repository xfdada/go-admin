package jwts

import (
	"go-admin/global"
	"go-admin/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// //获取 Secret
func GetJWTSecret() []byte {

	return []byte(global.JWT.Secret)
}

//生成token
func GenerateToken(appkey, appsecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWT.Expire * time.Second)
	claims := Claims{
		AppKey:    utils.EncodeMD5(appkey),
		AppSecret: utils.EncodeMD5(appsecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWT.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

//验证token是否有效

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
