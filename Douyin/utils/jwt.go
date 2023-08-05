package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 设置jwt密钥
var jwtSecret = []byte("feige")

type Claims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"admin"`
	jwt.StandardClaims
}

const expire_time = 3 * time.Hour

// GenerateToken 生成token的函数
func GenerateToken(username string, userId int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire_time)

	claims := Claims{
		userId, // 自行添加的信息
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 设置token过期时间
			IssuedAt:  time.Now().Unix(),
			Subject:   "userToken",
		},
	}
	// 生成token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	claimsF := token.Claims.(*Claims)
	return token, claimsF, err
}
