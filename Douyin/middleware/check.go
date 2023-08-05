package middleware

import (
	"Douyin/cache"
	"Douyin/common"
	"Douyin/controller"
	"Douyin/dao"
	"Douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func Check() gin.HandlerFunc {

	return func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("username")
		//正则表示5～16字节，允许字母、数字、下划线，以字母开头
		matchString := "^[a-zA-Z][a-zA-Z0-9_]{4,15}$"
		usernameMatch, _ := regexp.MatchString(matchString, username)
		passwordMatch, _ := regexp.MatchString(matchString, password)

		if usernameMatch && passwordMatch != true {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1, StatusMsg: "Account or password is illegal",
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Set("password", password)
		//挂起
		c.Next()
	}
}
func TokenAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		//得到token字段
		//1.get请求
		token := c.Query("token")
		if token == "" {
			//2.post请求
			token = c.PostForm("token")
		}
		// 两种情况下来，判断是否有token
		if token == "" {
			c.JSON(http.StatusOK, controller.UserResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "token 不存在"},
			})
			//终止
			c.Abort()
		}
		t, claim, err := utils.ParseToken(token)
		//判断是否有效
		if !t.Valid || err != nil {
			c.JSON(http.StatusOK, controller.UserResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "token有效期过了或者" + err.Error()},
			})
			c.Abort()
			return
		}
		//1.首先到redis中查找，没有的话再去mysql中查找，减少读写压力
		var isExists = true

		err = cache.UserIsExists(claim.UserId)
		if err != nil {
			fmt.Println(err)
			//在redis中不存在
			isExists = false
		}
		if !isExists {
			//进行db查找
			var user *dao.User
			user, err = dao.GetUserInstance().QueryUserByID(claim.UserId)
			if err != nil {
				c.JSON(http.StatusOK, controller.UserResponse{
					Response: common.Response{StatusCode: 1, StatusMsg: "token find failed"},
				})
				c.Abort()
				return
			}
			if user.ID == 0 {
				c.JSON(http.StatusOK, controller.UserResponse{
					Response: common.Response{StatusCode: 1, StatusMsg: "id is not exists"},
				})
				c.Abort()
				return
			}
		}
		c.Set("userid", claim.UserId)
		c.Next()
	}
}
