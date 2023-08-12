package utils

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type Claims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"admin"`
	jwt.StandardClaims
}

const expire_time = 3 * time.Hour

// GenerateToken 生成token的函数
func GenerateToken(username string, userId int64) (string, error) {

	// 加载RSA私钥
	keys, err := loadKeys()
	if err != nil {
		log.Println("Failed to load keys")
		return "err", err
	}
	// 创建一个token
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
	// 使用私钥签署token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := tokenClaims.SignedString(keys.PrivateKey)
	log.Println(token)
	return token, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	// 加载RSA公钥
	keys, err := loadKeys()
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return keys.PublicKey, nil
	})
	claimsF := token.Claims.(*Claims)
	return token, claimsF, err
}
