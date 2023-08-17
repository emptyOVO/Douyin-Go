package controller

import (
	"Douyin/cache"
	"Douyin/common"
	"Douyin/dao"
	"Douyin/service"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"strconv"
)

type UserRegisterResponse struct {
	common.Response
	Token  string `json:"token"`   // 用户鉴权token
	UserId int64  `json:"user_id"` // 用户id
}

type UserResponse struct {
	common.Response
	User dao.User `json:"user"`
}

// Register 用户注册接口
// 已增加邮箱验证码验证，相关代码已注释
func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	//captcha := c.Query("captcha")

	conn := cache.RedisPool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)
	//if captcha==cache.getCaptcha(username){
	info, err := service.UserRegister(username, password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	} else {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: common.Response{StatusCode: 0},
			Token:    info.Token,
			UserId:   info.UserID,
		})
		return
	}
	//}

}

// Login 用户登录接口
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	info, err := service.UserLogin(username, password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	} else {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: common.Response{StatusCode: 0},
			Token:    info.Token,
			UserId:   info.UserID,
		})
		return
	}
}

func UserInfo(c *gin.Context) {
	var err error
	var user *dao.User

	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	user, err = service.GetUserInfo(int64(userid))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//返回成功
	c.JSON(http.StatusOK, UserResponse{
		Response: common.Response{StatusCode: 0, StatusMsg: "successful"},
		User:     *user,
	})
	return
}
