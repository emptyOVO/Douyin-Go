package controller

import (
	"Douyin/cache"
	"Douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
)

func SendMail(c *gin.Context) {
	email := c.Query("email")
	// 生成验证码
	code := utils.GenerateCode()
	// 发送验证码
	err := utils.SendVerificationCode(email, code)
	if err != nil {
		fmt.Println("Failed to send verification code:", err)
		return
	}
	//redis缓存验证码
	conn := cache.RedisPool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)
	//往集合中加验证码
	_, err = conn.Do("ECODE", email, code)
	if err != nil {
		return
	}
	if err != nil {
		log.Println(err)
	}
}
