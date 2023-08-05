package controller

import (
	"Douyin/common"
	"Douyin/dao"
	"Douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func FriendList(c *gin.Context) {

	var userLists []dao.User
	userid, err := strconv.Atoi(c.Query("user_id"))

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FollowResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
		return
	}

	userLists, err = service.GetFriendLists(int64(userid))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FollowResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, FollowResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "successful"},
		UserLists: userLists,
	})
	return
}
